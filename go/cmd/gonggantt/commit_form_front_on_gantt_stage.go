package main

import (
	"fmt"
	"log"
	"os"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"

	"github.com/fullstack-lang/gonggantt/go/gantt2svg"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
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
