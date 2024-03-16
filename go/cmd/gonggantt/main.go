package main

import (
	"flag"
	"log"
	"strconv"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gonggantt_stack "github.com/fullstack-lang/gonggantt/go/stack"

	gonggantt_static "github.com/fullstack-lang/gonggantt/go/static"
	gongsvg_stack "github.com/fullstack-lang/gongsvg/go/stack"

	"github.com/fullstack-lang/gonggantt/go/gantt2svg"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("gonggantt: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gonggantt_static.ServeStaticFiles(*logGINFlag)

	// setup stackGonggantt
	stackGonggantt := gonggantt_stack.NewStack(r, gonggantt_models.GanttStackName.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stackGonggantt.Probe.Refresh()

	stacksvg := gongsvg_stack.NewStack(r, gonggantt_models.SvgStackName.ToString(), "", "", "", true, true)
	stacksvg.Probe.Refresh()

	// set up the GanttSVGMapper that will intercept
	// commits on the gonggantt stage and that will
	// generate the svg
	ganttSVGMapper := new(gantt2svg.GanttSVGMapper)
	ganttSVGMapper.GanttOuputFile = *marshallOnCommit

	commitOnGanttStage := new(CommitFromFrontOnGanttStage)
	commitOnGanttStage.gongsvgStage = stacksvg.Stage
	commitOnGanttStage.ganttSVGMapper = ganttSVGMapper

	// hook on the commit from front
	stackGonggantt.Stage.OnInitCommitFromFrontCallback = commitOnGanttStage
	stackGonggantt.Stage.OnInitCommitFromBackCallback = commitOnGanttStage

	// initial publication
	ganttSVGMapper.GenerateSvg(stackGonggantt.Stage, stacksvg.Stage)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
