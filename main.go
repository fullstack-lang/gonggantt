package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	gonggantt_controllers "github.com/fullstack-lang/gonggantt/go/controllers"
	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gonggantt_orm "github.com/fullstack-lang/gonggantt/go/orm"

	"github.com/fullstack-lang/gonggantt/go/gantt2svg"

	gongsvg_controllers "github.com/fullstack-lang/gongsvg/go/controllers"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_orm "github.com/fullstack-lang/gongsvg/go/orm"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")

	backupFlag  = flag.Bool("backup", false, "read database file, generate backup and exits")
	restoreFlag = flag.Bool("restore", false, "generate restore and exits")
)

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
	gonggantt_orm.SetupModels(*logDBFlag, "./test.db")

	//
	// gongsvg database
	//
	db_inMemory := gongsvg_orm.SetupModels(*logDBFlag, ":memory:")
	// mandatory, otherwise, bizarre errors occurs

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB, err := db_inMemory.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)

	gonggantt_controllers.RegisterControllers(r)
	gongsvg_controllers.RegisterControllers(r)

	// plug the svg generator on the OnInitCommit callback
	gonggantt_models.Stage.OnInitCommitCallback = &gantt2svg.GanttToSVGTranformerSingloton

	// put all to database
	gonggantt_models.Stage.Commit()
	gongsvg_models.Stage.Commit()

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(ng, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

//go:embed ng/dist/ng
var ng embed.FS

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
