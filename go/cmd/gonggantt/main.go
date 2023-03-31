package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	gonggantt_go "github.com/fullstack-lang/gonggantt/go"
	gonggantt_fullstack "github.com/fullstack-lang/gonggantt/go/fullstack"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gonggantt_static "github.com/fullstack-lang/gonggantt/go/static"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	"github.com/fullstack-lang/gonggantt/go/gantt2svg"

	gongsvg_go "github.com/fullstack-lang/gongsvg/go"
	gongsvg_fullstack "github.com/fullstack-lang/gongsvg/go/fullstack"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	marshallOnStartup  = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	unmarshall         = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
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

	// setup the static file server and get the controller
	r := gonggantt_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	gonganttStage := gonggantt_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gonggantt/go/models")
	gongsvgStage := gongsvg_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongsvg/go/models")

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

	gongdoc_load.Load(
		"gonggantt",
		"github.com/fullstack-lang/gonggantt/go/models",
		gonggantt_go.GoModelsDir,
		gonggantt_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&gonganttStage.Map_GongStructName_InstancesNb)

	gongdoc_load.Load(
		"gongsvg",
		"github.com/fullstack-lang/gongsvg/go/models",
		gongsvg_go.GoModelsDir,
		gongsvg_go.GoDiagramsDir,
		r,
		true,
		&gonganttStage.Map_GongStructName_InstancesNb)

	// initial publication
	gantt2svg.GanttToSVGTranformerSingloton.GenerateSvg(gonganttStage, gongsvgStage)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
