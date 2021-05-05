package gantt2svg

import (
	"log"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// type of the singloton for interception of gantt commit in order to generate
// the svg
type GanttToSVGTranformer struct {
}

var GanttToSVGTranformerSingloton GanttToSVGTranformer

// callback on the commit function
func (GanttToSVGTranformer *GanttToSVGTranformer) BeforeCommit(stage *gonggantt_models.StageStruct) {

	// remove all gongsvg stage/repo
	gongsvg_models.Stage.Checkout()
	gongsvg_models.Stage.Reset()
	gongsvg_models.Stage.Commit()

	// creates a tick lane
	svg := new(gongsvg_models.SVG).Stage()
	svg.Name = "New Gantt Chart"
	svg.Display = true

	timeLine := new(gongsvg_models.Line).Stage()
	timeLine.Name = "Time Line"
	timeLine.X1 = 10
	timeLine.Y1 = 10
	timeLine.X2 = 200
	timeLine.Y2 = 10
	timeLine.FillOpacity = 1.0

	svg.Lines = append(svg.Lines, timeLine)

	gongsvg_models.Stage.Commit()

	log.Printf("Before Commit")
}
