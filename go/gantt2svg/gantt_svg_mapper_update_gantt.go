package gantt2svg

import (
	"log"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func (ganttToSVGMapper *GanttToSVGMapper) UpdateGantt(
	gongsvgStage *gongsvg_models.StageStruct,
	gongganttStage *gonggantt_models.StageStruct) {

	log.Println("GanttToSVGMapper: UpdateGantt")
}
