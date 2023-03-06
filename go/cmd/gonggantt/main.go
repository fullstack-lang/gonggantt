package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gonggantt"
	gonggantt_fullstack "github.com/fullstack-lang/gonggantt/go/fullstack"
	"github.com/fullstack-lang/gonggantt/go/gantt2svg"
	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	// import this package in order to have the scheduler start a thread that will
	// generate a new svg diagram each time the repo has been modified
	_ "github.com/fullstack-lang/gonggantt/go/gantt2svg"

	gongsvg_fullstack "github.com/fullstack-lang/gongsvg/go/fullstack"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	backupFlag  = flag.Bool("backup", false, "read database file, generate backup and exits")
	restoreFlag = flag.Bool("restore", false, "generate restore and exits")

	marshallOnStartup  = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshall         = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")

	diagrams = flag.Bool("diagrams", true, "parse diagrams (takes a few seconds)")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
	gongsvgStage *gongsvg_models.StageStruct
}

func (impl *BeforeCommitImplementation) BeforeCommit(
	gongganttStage *gonggantt_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	gongganttStage.Checkout()
	gongganttStage.Marshall(file, "github.com/fullstack-lang/gonggantt/go/models", "main")
	gantt2svg.GanttToSVGTranformerSingloton.GenerateSvg(gongganttStage, impl.gongsvgStage)
}

func main() {

	log.SetPrefix("gonggantt: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	gonganttStage, _ := gonggantt_fullstack.NewStackInstance(r, "")
	gongsvgStage, _ := gongsvg_fullstack.NewStackInstance(r, "")

	// generate injection code from the stage
	if *marshallOnStartup != "" {

		if strings.Contains(*marshallOnStartup, " ") {
			log.Fatalln(*marshallOnStartup + " must not contains blank spaces")
		}
		if strings.ToLower(*marshallOnStartup) != *marshallOnStartup {
			log.Fatalln(*marshallOnStartup + " must be lowercases")
		}

		file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnStartup))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		gonganttStage.Checkout()
		gonganttStage.Marshall(file, "github.com/fullstack-lang/gonggantt/go/models", "main")
		os.Exit(0)
	}

	if *unmarshall != "" {
		gonganttStage.Checkout()
		gonganttStage.Reset()
		gonganttStage.Commit()
		if InjectionGateway[*unmarshall] != nil {
			InjectionGateway[*unmarshall]()
		}

		gonganttStage.Commit()
	}
	if *unmarshallFromCode != "" {
		gonganttStage.Checkout()
		gonganttStage.Reset()
		gonganttStage.Commit()
		err := gonggantt_models.ParseAstFile(gonganttStage, *unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		gonganttStage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		gonganttStage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		hook.gongsvgStage = gongsvgStage
		gonganttStage.OnInitCommitFromFrontCallback = hook
	}

	// put all to database
	gonganttStage.Commit()
	gongsvg_models.Stage.Commit()

	gongdoc_load.Load(
		"gonggantt",
		"github.com/fullstack-lang/gonggantt/go/models",
		gonggantt.GoDir,
		r,
		*embeddedDiagrams,
		&gonganttStage.Map_GongStructName_InstancesNb)

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gonggantt.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	gantt2svg.GanttToSVGTranformerSingloton.GenerateSvg(gonganttStage, gongsvgStage)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
