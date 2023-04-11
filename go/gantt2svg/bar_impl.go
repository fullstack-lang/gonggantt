package gantt2svg

import (
	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type BarImpl struct {
	bar            *gonggantt_models.Bar
	ganttToRender  *gonggantt_models.Gantt
	gongganttStage *gonggantt_models.StageStruct
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

	barImpl.gongganttStage.Commit()
}
