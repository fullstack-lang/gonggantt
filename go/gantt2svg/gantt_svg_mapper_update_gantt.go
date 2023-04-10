package gantt2svg

import (
	"log"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func (ganttSVGMapper *GanttSVGMapper) UpdateGantt(
	gongsvgStage *gongsvg_models.StageStruct,
	gongganttStage *gonggantt_models.StageStruct) {

	log.Println("GanttToSVGMapper: UpdateGantt")

	for bar, rect := range ganttSVGMapper.mapBar_BarSVG {

		// we start from the rect SVG X and Width and update the bar Start and End
		_ = bar
		_ = rect
	}
}
