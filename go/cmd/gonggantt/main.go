package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	gonggantt_go "github.com/fullstack-lang/gonggantt/go"
	gonggantt_fullstack "github.com/fullstack-lang/gonggantt/go/fullstack"
	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gonggantt_orm "github.com/fullstack-lang/gonggantt/go/orm"
	gonggantt_probe "github.com/fullstack-lang/gonggantt/go/probe"
	gonggantt_static "github.com/fullstack-lang/gonggantt/go/static"

	"github.com/fullstack-lang/gonggantt/go/gantt2svg"

	gongsvg_go "github.com/fullstack-lang/gongsvg/go"
	gongsvg_fullstack "github.com/fullstack-lang/gongsvg/go/fullstack"
	gongsvg_probe "github.com/fullstack-lang/gongsvg/go/probe"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

// hook marhalling to stage
type CommitFromFrontOnGanttStage struct {
	gongsvgStage   *gongsvg_models.StageStruct
	ganttSVGMapper *gantt2svg.GanttSVGMapper
}

// BeforeCommit meets the interface for the commit on the gantt stage
// It performs 2 tasks
// 1 - update the SVG stack
// 2 - persists the data to the gantt file
func (beforeCommitFromFrontOnGanttStage *CommitFromFrontOnGanttStage) BeforeCommit(
	gongganttStage *gonggantt_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// update the gantt stage with the back repo data that was updated from the front
	gongganttStage.Checkout()

	// marshall to the file
	gongganttStage.Marshall(file, "github.com/fullstack-lang/gonggantt/go/models", "main")
	beforeCommitFromFrontOnGanttStage.ganttSVGMapper.GenerateSvg(gongganttStage, beforeCommitFromFrontOnGanttStage.gongsvgStage)
}

type OnAfterRectUpdate struct {
	gongganttStage   *gonggantt_models.StageStruct
	ganttToSVGMapper *gantt2svg.GanttSVGMapper
}

func (OnAfterRectUpdate *OnAfterRectUpdate) OnAfterUpdate(
	gongsvgStage *gongsvg_models.StageStruct,
	old *gongsvg_models.Rect,
	new *gongsvg_models.Rect) {

	log.Println("OnAfterRectUpdate, OnAfterUpdate", old.Name)

}

func main() {

	log.SetPrefix("gonggantt: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gonggantt_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	var gongganttStage *gonggantt_models.StageStruct
	gongsvgStage, gongsvgBackRepo := gongsvg_fullstack.NewStackInstance(r, gonggantt_models.SvgStackName.ToString())

	var backRepo *gonggantt_orm.BackRepoStruct

	if *marshallOnCommit != "" {
		// persistence in a SQLite file on disk in memory
		gongganttStage, backRepo = gonggantt_fullstack.NewStackInstance(r, "gonggantt")
	} else {
		// persistence in a SQLite file on disk
		gongganttStage, backRepo = gonggantt_fullstack.NewStackInstance(r, "gonggantt", "./gonggantt.db")
	}

	if *unmarshallFromCode != "" {
		gongganttStage.Checkout()
		gongganttStage.Reset()
		gongganttStage.Commit()
		err := gonggantt_models.ParseAstFile(gongganttStage, *unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		gongganttStage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		gongganttStage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {

		ganttSVGMapper := new(gantt2svg.GanttSVGMapper)
		ganttSVGMapper.GanttOuputFile = *marshallOnCommit

		commitOnGanttStage := new(CommitFromFrontOnGanttStage)
		commitOnGanttStage.gongsvgStage = gongsvgStage
		commitOnGanttStage.ganttSVGMapper = ganttSVGMapper

		// hook on the commit from front
		gongganttStage.OnInitCommitFromFrontCallback = commitOnGanttStage
		gongganttStage.OnInitCommitFromBackCallback = commitOnGanttStage

		// onAfterRectUpdate := new(OnAfterRectUpdate)
		// onAfterRectUpdate.gongganttStage = gongganttStage
		// onAfterRectUpdate.ganttToSVGMapper = ganttSVGMapper

		// // hook on the commit from front
		// gongsvgStage.OnAfterRectUpdateCallback = onAfterRectUpdate

		// put the SVG Rect logic on it ()
		gongsvg_models.SetOrchestratorOnAfterUpdate[gongsvg_models.Rect](gongsvgStage)

		// initial publication
		ganttSVGMapper.GenerateSvg(gongganttStage, gongsvgStage)
	}

	gonggantt_probe.NewProbe(r, gonggantt_go.GoModelsDir, gonggantt_go.GoDiagramsDir,
		*embeddedDiagrams, "gonggantt", gongganttStage, backRepo)

	gongsvg_probe.NewProbe(r, gongsvg_go.GoModelsDir, gongsvg_go.GoDiagramsDir,
		false, "gongsvg", gongsvgStage, gongsvgBackRepo)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
