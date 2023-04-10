package gantt2svg

import (
	"fmt"
	"log"
	"time"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func (ganttSVGMapper *GanttSVGMapper) UpdateGantt(
	gongsvgStage *gongsvg_models.StageStruct,
	gongganttStage *gonggantt_models.StageStruct,
	oldRect, newRect *gongsvg_models.Rect) {

	log.Println("GanttToSVGMapper: UpdateGantt")

	var dateChange bool

	bar := ganttSVGMapper.mapRect4Bar_Bar[oldRect]

	// we start from the rect SVG X and Width and update the bar Start and End
	startDateIn := bar.Start
	endDateIn := bar.End

	// Print the formatted time
	fmt.Println("Start time: before", startDateIn.Format("2006-01-02 15:04:05"))

	bar.Start = XtoDate(
		ganttSVGMapper.ganttToRender.XLeftLanes,
		newRect.X,
		ganttSVGMapper.ganttToRender.XRightMargin,
		ganttSVGMapper.ganttToRender.ComputedStart,
		ganttSVGMapper.ganttToRender.ComputedEnd)

	// Print the formatted time
	fmt.Println("Start time:", bar.Start.Format("2006-01-02 15:04:05"))

	if bar.Start != startDateIn {
		dateChange = true
	}

	// Print the formatted time
	fmt.Println("End time before:", endDateIn.Format("2006-01-02 15:04:05"))

	bar.End = XtoDate(
		ganttSVGMapper.ganttToRender.XLeftLanes,
		newRect.X+newRect.Width,
		ganttSVGMapper.ganttToRender.XRightMargin,
		ganttSVGMapper.ganttToRender.ComputedStart,
		ganttSVGMapper.ganttToRender.ComputedEnd)

	// Print the formatted time
	fmt.Println("End time:", bar.End.Format("2006-01-02 15:04:05"))

	if bar.End != endDateIn {
		dateChange = true
	}

	if dateChange {
		gongganttStage.Commit()
	}
}

// convert X coordinate to date
func XtoDate(leftX, x, rightX float64, ganttStartDate, ganttEndDate time.Time) (date time.Time) {

	// compute the geometric ratio of X compared to the the distance between RightX and LeftX
	ratio := (x - leftX) / (rightX - leftX)

	ganttDuration := ganttEndDate.Sub(ganttStartDate)

	durationSinceGanttStart := time.Duration(ratio * float64(ganttDuration))

	date = ganttStartDate.Add(durationSinceGanttStart)

	return
}
