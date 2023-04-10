package gantt2svg

import (
	"fmt"
	"log"
	"os"
	"time"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

func (ganttSVGMapper *GanttSVGMapper) UpdateGantt(
	gongsvgStage *gongsvg_models.StageStruct,
	gongganttStage *gonggantt_models.StageStruct,
	oldRect, newRect *gongsvg_models.Rect) {

	log.Println("GanttToSVGMapper: UpdateGantt")

	var coordChange bool

	bar := ganttSVGMapper.mapRect4Bar_Bar[oldRect]

	bar.Start = XtoDate(
		ganttSVGMapper.ganttToRender.XLeftLanes,
		newRect.X,
		ganttSVGMapper.ganttToRender.XRightMargin,
		ganttSVGMapper.ganttToRender.ComputedStart,
		ganttSVGMapper.ganttToRender.ComputedEnd)

	if newRect.X != oldRect.X {
		fmt.Println("X change: ", oldRect.X-newRect.X)
		coordChange = true
	}

	bar.End = XtoDate(
		ganttSVGMapper.ganttToRender.XLeftLanes,
		newRect.X+newRect.Width,
		ganttSVGMapper.ganttToRender.XRightMargin,
		ganttSVGMapper.ganttToRender.ComputedStart,
		ganttSVGMapper.ganttToRender.ComputedEnd)

	// Print the formatted time
	fmt.Println("End time:", bar.End.Format("2006-01-02 15:04:05"))

	if newRect.X+newRect.Width != oldRect.X+oldRect.Width {
		fmt.Println("X right change:", (newRect.X+newRect.Width)-(oldRect.X+oldRect.Width))
		coordChange = true
	}

	if coordChange {
		bar.Commit(gongganttStage)

		file, err := os.Create(fmt.Sprintf("./%s.go", ganttSVGMapper.GanttOuputFile))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		gongganttStage.Marshall(file, "github.com/fullstack-lang/gonggantt/go/models", "main")
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
