package gantt2svg

import (
	"fmt"
	"log"
	"os"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type BarImpl struct {
	bar            *gonggantt_models.Bar
	ganttToRender  *gonggantt_models.Gantt
	gongganttStage *gonggantt_models.StageStruct
	ganttOuputFile string
}

func (barImpl *BarImpl) RectUpdated(updatedRect *gongsvg_models.Rect) {

	barImpl.bar.Start = XtoDate(
		barImpl.ganttToRender.XLeftLanes,
		updatedRect.X,
		barImpl.ganttToRender.XRightMargin,
		barImpl.ganttToRender.ComputedStart,
		barImpl.ganttToRender.ComputedEnd)

	barImpl.bar.End = XtoDate(
		barImpl.ganttToRender.XLeftLanes,
		updatedRect.X+updatedRect.Width,
		barImpl.ganttToRender.XRightMargin,
		barImpl.ganttToRender.ComputedStart,
		barImpl.ganttToRender.ComputedEnd)

	barImpl.bar.Commit(barImpl.gongganttStage)

	file, err := os.Create(fmt.Sprintf("./%s.go", barImpl.ganttOuputFile))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	barImpl.gongganttStage.Marshall(file, "github.com/fullstack-lang/gonggantt/go/models", "main")
}
