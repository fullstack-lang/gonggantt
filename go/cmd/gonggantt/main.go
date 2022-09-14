package main

import (
	"embed"
	"flag"
	"fmt"
	"go/token"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gonggantt"
	gonggantt_controllers "github.com/fullstack-lang/gonggantt/go/controllers"
	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gonggantt_orm "github.com/fullstack-lang/gonggantt/go/orm"

	// gong stack for model analysis
	gong_controllers "github.com/fullstack-lang/gong/go/controllers"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_orm "github.com/fullstack-lang/gong/go/orm"
	_ "github.com/fullstack-lang/gong/ng"

	// for diagrams
	gongdoc_controllers "github.com/fullstack-lang/gongdoc/go/controllers"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_orm "github.com/fullstack-lang/gongdoc/go/orm"
	_ "github.com/fullstack-lang/gongdoc/ng"

	// import this package in order to have the scheduler start a thread that will
	// generate a new svg diagram each time the repo has been modified
	_ "github.com/fullstack-lang/gonggantt/go/gantt2svg"

	gongsvg_controllers "github.com/fullstack-lang/gongsvg/go/controllers"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_orm "github.com/fullstack-lang/gongsvg/go/orm"
	_ "github.com/fullstack-lang/gongsvg/ng"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	backupFlag  = flag.Bool("backup", false, "read database file, generate backup and exits")
	restoreFlag = flag.Bool("restore", false, "generate restore and exits")

	marshallOnStartup = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshall        = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit  = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams = flag.Bool("diagrams", true, "parse diagrams (takes a few seconds)")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gonggantt_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	gonggantt_models.Stage.Checkout()
	gonggantt_models.Stage.Marshall(file, "github.com/fullstack-lang/gonggantt/go/models", "main")
}

func main() {

	log.SetPrefix("gonggantt: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	if *backupFlag {

		// setup GORM
		db := gonggantt_orm.SetupModels(*logDBFlag, "./test.db")
		gonggantt_orm.AutoMigrate(db)
		gonggantt_models.Stage.Checkout()
		gonggantt_models.Stage.Backup("bckp")

		return
	}
	if *restoreFlag {

		// setup GORM
		db := gonggantt_orm.SetupModels(*logDBFlag, "./test.db")
		gonggantt_orm.AutoMigrate(db)
		gonggantt_models.Stage.Restore("bckp")

		return
	}

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	//
	// gonggantt
	gonggantt_orm.SetupModels(*logDBFlag, "file:memdb2?mode=memory&cache=shared")

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

		gonggantt_models.Stage.Checkout()
		gonggantt_models.Stage.Marshall(file, "github.com/fullstack-lang/gonggantt/go/models", "main")
		os.Exit(0)
	}

	// setup the stage by injecting the code from code database
	stage := &gonggantt_models.Stage
	_ = stage
	if *unmarshall != "" {
		gonggantt_models.Stage.Checkout()
		gonggantt_models.Stage.Reset()
		gonggantt_models.Stage.Commit()
		if InjectionGateway[*unmarshall] != nil {
			InjectionGateway[*unmarshall]()
		}

		gonggantt_models.Stage.Commit()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		gonggantt_models.Stage.OnInitCommitFromFrontCallback = hook
	}

	//
	// gongsvg database
	//
	db_inMemory := gongsvg_orm.SetupModels(*logDBFlag, "file:memdb1?mode=memory&cache=shared")
	// mandatory, otherwise, bizarre errors occurs

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	// add gongdoc database
	gongdoc_orm.AutoMigrate(db_inMemory)

	// add gong database
	gong_orm.AutoMigrate(db_inMemory)

	gonggantt_controllers.RegisterControllers(r)
	gongsvg_controllers.RegisterControllers(r)
	gongdoc_controllers.RegisterControllers(r)
	gong_controllers.RegisterControllers(r)

	// put all to database
	gonggantt_models.Stage.Commit()
	gongsvg_models.Stage.Commit()
	gongdoc_models.Stage.Commit()
	gong_models.Stage.Commit()

	// from the commited stage, get the number of instances per struct
	// before unmarshalling diagrams
	if *diagrams {

		// Analyse package
		modelPkg := &gong_models.ModelPkg{}

		// since the source is embedded, one needs to
		// compute the Abstract syntax tree in a special manner
		pkgs := gong_models.ParseEmbedModel(gonggantt.GoDir, "go/models")

		gong_models.WalkParser(pkgs, modelPkg)
		modelPkg.SerializeToStage()
		gong_models.Stage.Commit()

		// create the diagrams
		// prepare the model views
		pkgelt := new(gongdoc_models.Pkgelt)

		// first, get all gong struct in the model
		for gongStruct := range gong_models.Stage.GongStructs {

			// let create the gong struct in the gongdoc models
			// and put the numbre of instances
			gongStruct_ := (&gongdoc_models.GongStruct{Name: gongStruct.Name}).Stage()
			nbInstances, ok := gonggantt_models.Stage.Map_GongStructName_InstancesNb[gongStruct.Name]
			if ok {
				gongStruct_.NbInstances = nbInstances
			}
		}

		// classdiagram can only be fully in memory when they are Unmarshalled
		// for instance, the Name of diagrams or the Name of the Link
		fset := new(token.FileSet)
		pkgsParser := gong_models.ParseEmbedModel(gonggantt.GoDir, "go/diagrams")
		if len(pkgsParser) != 1 {
			log.Panic("Unable to parser, wrong number of parsers ", len(pkgsParser))
		}
		if pkgParser, ok := pkgsParser["diagrams"]; ok {
			pkgelt.Unmarshall(modelPkg, pkgParser, fset, "go/diagrams")
		}
		pkgelt.SerializeToStage()
	}
	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gonggantt.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

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
