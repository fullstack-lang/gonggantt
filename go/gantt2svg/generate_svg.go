package gantt2svg

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// type of the singloton for interception of gantt commit in order to generate
// the svg
type GanttToSVGTranformer struct {
}

var GanttToSVGTranformerSingloton GanttToSVGTranformer

// backgroundColor.fill 	string 	'white' 	The chart fill color, as an HTML color string.
// gantt.arrow 	object 	null 	For Gantt Charts, gantt.arrow controls the various properties of the arrows connecting tasks.
// gantt.arrow.angle 	number 	45 	The angle of the head of the arrow.
// gantt.arrow.color 	string 	'#000' 	The color of the arrows.
// gantt.arrow.length 	number 	8 	The length of the head of the arrow.
// gantt.arrow.radius 	number 	15 	The radius for defining the curve of the arrow between two tasks.
// gantt.arrow.spaceAfter 	number 	4 	The amount of whitespace between the head of an arrow and the task to which it points.
// gantt.arrow.width 	number 	1.4 	The width of the arrows.
// gantt.barCornerRadius 	number 	2 	The radius for defining the curve of a bar's corners.
// gantt.barHeight 	number 	null 	The height of the bars for tasks.
// gantt.criticalPathEnabled 	boolean 	true 	If true any arrows on the critical path will be styled differently.
// gantt.criticalPathStyle 	object 	null 	An object containing the style for any critical path arrows.
// gantt.criticalPathStyle.stroke 	string 	null 	The color of any critical path arrows.
// gantt.criticalPathStyle.strokeWidth 	number 	1.4 	The thickness of any critical path arrows.
// gantt.defaultStartDate 	date/number 	null 	If the start date cannot be computed from the values in the DataTable, the start date will be set to this. Accepts a date value (new Date(YYYY, M, D)) or a number, which is the number of milliseconds to use.
// gantt.innerGridHorizLine 	object 	null 	Defines the style of the inner horizontal grid lines.
// gantt.innerGridHorizLine.stroke 	string 	null 	The color of the inner horizontal grid lines.
// gantt.innerGridHorizLine.strokeWidth 	number 	1 	The width of the inner horizontal grid lines.
// gantt.innerGridTrack.fill 	string 	null 	The fill color of the inner grid track. If no innerGridDarkTrack.fill is specified, this will be applied to every grid track.
// gantt.innerGridDarkTrack.fill 	string 	null 	The fill color of the alternate, dark inner grid track.
// gantt.labelMaxWidth 	number 	300 	The maximum amount of space allowed for each task label.
// gantt.labelStyle 	object 	null

// An object containing the styles for task labels.

// labelStyle: {
//   fontName: Roboto2,
//   fontSize: 14,
//   color: '#757575'
// },

// gantt.percentEnabled 	boolean 	true 	Fills the task bar based on the percentage completed for the task.
// gantt.percentStyle.fill 	string 	null 	The color of the percentage completed portion of a task bar.
// gantt.shadowEnabled 	boolean 	true 	If set to true, draws a shadow under each task bar which has dependencies.
// gantt.shadowColor 	string 	'#000' 	Defines the color of the shadows under any task bar which has dependencies.
// gantt.shadowOffset 	number 	1 	Defines the offset, in pixels, of the shadows under any task bar which has dependencies.
// gantt.sortTasks 	boolean 	true 	Specifies that the tasks should sorted topologically, if true; otherwise use the same order as the corresponding rows of the DataTable.
// gantt.trackHeight 	number 	null 	The height of the tracks.
// width 	number 	width of the containing element 	Width of the chart, in pixels.
// height 	number 	height of the containing element 	height of the chart, in pixels.

// callback on the commit function
func (GanttToSVGTranformer *GanttToSVGTranformer) BeforeCommit(stage *gonggantt_models.StageStruct) {

	// remove all gongsvg stage/repo
	gongsvg_models.Stage.Checkout()
	gongsvg_models.Stage.Reset()
	gongsvg_models.Stage.Commit()
	gonggantt_models.Stage.Checkout()

	if len(gonggantt_models.Stage.Gantts) != 1 {
		log.Printf("It is supposed to have only one gantt chart")
		return
	}

	var ganttToRender *gonggantt_models.Gantt
	for gantt := range gonggantt_models.Stage.Gantts {
		ganttToRender = gantt
	}
	ganttToRender.ComputeStartAndEndDate()

	//
	// SVG
	//
	svg := new(gongsvg_models.SVG).Stage()
	svg.Name = "New Gantt Chart"
	svg.Display = true

	//
	// Time Line
	//
	// creates a tick lane
	LaneHeight := ganttToRender.LaneHeight
	RatioBarToLaneHeight := ganttToRender.RatioBarToLaneHeight
	barHeigth := LaneHeight * RatioBarToLaneHeight
	YTopMargin := ganttToRender.YTopMargin
	yTimeLine := LaneHeight*float64(len(stage.Lanes)) + YTopMargin

	XLeftText := ganttToRender.XLeftText
	TextHeight := ganttToRender.TextHeight

	XLeftLanes := ganttToRender.XLeftLanes
	XRightMargin := ganttToRender.XRightMargin

	timeLine := new(gongsvg_models.Line).Stage()
	timeLine.Name = "Time Line"
	timeLine.X1 = XLeftLanes
	timeLine.Y1 = yTimeLine
	timeLine.X2 = XRightMargin
	timeLine.Y2 = yTimeLine

	timeLine.Color = ganttToRender.TimeLine_Color
	timeLine.FillOpacity = ganttToRender.TimeLine_FillOpacity
	timeLine.Stroke = ganttToRender.TimeLine_Stroke
	timeLine.StrokeWidth = ganttToRender.TimeLine_StrokeWidth

	svg.Lines = append(svg.Lines, timeLine)

	//
	// tick lines
	//

	//
	// Dates
	//

	// Date offset. Where is the date below the TimeLine
	DateYOffset := 20.0

	// put a date for every year
	for year := ganttToRender.Start.Year(); year <= ganttToRender.End.Year(); year++ {
		yearTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		durationBetweenYearAndGanttStart := yearTime.Sub(ganttToRender.Start)

		durationBetweenYearStartAndGanttStartRelativeToGanttDuration :=
			float64(durationBetweenYearAndGanttStart) / float64(ganttToRender.End.Sub(ganttToRender.Start))

		yearText := new(gongsvg_models.Text).Stage()
		yearText.Name = fmt.Sprintf("%d", year)
		yearText.Content = yearText.Name
		yearText.X = XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenYearStartAndGanttStartRelativeToGanttDuration
		yearText.Y = yTimeLine + DateYOffset
		yearText.Color = "black"
		yearText.FillOpacity = 1.0
		svg.Texts = append(svg.Texts, yearText)

		// log.Printf("year %d", year)
		//
		// draw the line
		//
		lineForYear := new(gongsvg_models.Line).Stage()
		lineForYear.Name = yearText.Name
		svg.Lines = append(svg.Lines, lineForYear)
		lineForYear.X1 = yearText.X
		lineForYear.X2 = lineForYear.X1
		lineForYear.Y1 = YTopMargin
		lineForYear.Y2 = yTimeLine
		lineForYear.Stroke = "black"
		lineForYear.StrokeWidth = 0.25
		lineForYear.StrokeDashArray = "4 4"
	}

	//
	// Lanes
	//
	currentY := YTopMargin

	mapLane_TextY := make(map[*gonggantt_models.Lane]float64, 0)
	mapLane_TopY := make(map[*gonggantt_models.Lane]float64, 0)

	//
	// Sort Lanes according to the Order
	//
	sort.Slice(ganttToRender.Lanes, func(i, j int) bool {
		return ganttToRender.Lanes[i].Order < ganttToRender.Lanes[j].Order
	})

	laneIndex := 0
	for _, lane := range ganttToRender.Lanes {

		laneSVG := new(gongsvg_models.Rect).Stage()
		svg.Rects = append(svg.Rects, laneSVG)
		laneSVG.Name = lane.Name
		laneSVG.X = XLeftLanes
		laneSVG.Y = currentY
		mapLane_TopY[lane] = currentY

		laneSVG.Width = XRightMargin - XLeftLanes
		laneSVG.Height = LaneHeight

		laneSVG.Color = "grey"
		laneSVG.FillOpacity = 0.1

		if laneIndex%2 == 0 {
			laneSVG.Color = "black"
		}
		laneIndex = laneIndex + 1
		laneSVG.StrokeWidth = 1.5

		laneText := new(gongsvg_models.Text).Stage()
		laneText.Name = lane.Name
		laneText.Content = laneText.Name
		laneText.X = XLeftText
		laneText.Y = currentY + LaneHeight/2.0 + TextHeight/2.0
		mapLane_TextY[lane] = laneText.Y
		laneText.Color = "black"
		laneText.FillOpacity = 1.0
		svg.Texts = append(svg.Texts, laneText)

		//
		// Bar
		//
		for _, bar := range lane.Bars {
			barSVG := new(gongsvg_models.Rect).Stage()
			svg.Rects = append(svg.Rects, barSVG)
			barSVG.Name = bar.Name

			durationBetweenBarStartAndGanttStart := bar.Start.Sub(ganttToRender.Start)
			durationBetweenBarStartAndGanttStartRelativeToGanttDuration :=
				float64(durationBetweenBarStartAndGanttStart) / float64(ganttToRender.End.Sub(ganttToRender.Start))
			// log.Printf("Duration is %f", durationBetweenBarStartAndGanttStartRelativeToGanttDuration)

			durationBetweenBarEndAndBarStart := bar.End.Sub(bar.Start)
			durationBetweenBarEndAndBarStartRelativeToGanttDuration :=
				float64(durationBetweenBarEndAndBarStart) / float64(ganttToRender.End.Sub(ganttToRender.Start))
			// log.Printf("Relative Duration is %f", durationBetweenBarEndAndBarStartRelativeToGanttDuration)

			barSVG.X = XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenBarStartAndGanttStartRelativeToGanttDuration
			barSVG.Y = currentY + (LaneHeight-barHeigth)/2.0
			barSVG.Height = barHeigth
			barSVG.Width = (XRightMargin - XLeftLanes) * durationBetweenBarEndAndBarStartRelativeToGanttDuration

			barSVG.Color = "blue"
			barSVG.FillOpacity = 0.1
			barSVG.Stroke = "blue"
			barSVG.StrokeWidth = 0.5

			// bar text
			barText := new(gongsvg_models.Text).Stage()
			barText.Name = bar.Name
			barText.Content = barText.Name
			barText.X = barSVG.X + XLeftText
			barText.Y = currentY + LaneHeight/2.0 + TextHeight/2.0
			barText.Color = "black"
			barText.FillOpacity = 1.0
			svg.Texts = append(svg.Texts, barText)
		}

		currentY = currentY + LaneHeight

	}

	//
	// Milestones
	//
	for _, milestone := range ganttToRender.Milestones {

		durationBetweenMilestoneAndGanttStart := milestone.Date.Sub(ganttToRender.Start)
		durationBetweenMilestoneAndGanttStartRelativeToGanttDuration :=
			float64(durationBetweenMilestoneAndGanttStart) / float64(ganttToRender.End.Sub(ganttToRender.Start))

		//
		// draw the line
		//
		line := new(gongsvg_models.Line).Stage()
		line.Name = milestone.Name
		svg.Lines = append(svg.Lines, line)
		line.X1 = XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenMilestoneAndGanttStartRelativeToGanttDuration
		line.X2 = line.X1
		line.Y1 = YTopMargin
		line.Y2 = yTimeLine
		line.Stroke = "black"
		line.StrokeWidth = 0.5
		line.StrokeDashArray = "2 2"

		//
		// draw diamond
		//
		diamondWidth := 18.0
		for _, diamondAndTextAnchor := range milestone.DiamonfAndTextAnchors {

			diamond := new(gongsvg_models.Rect).Stage()
			svg.Rects = append(svg.Rects, diamond)
			diamond.Name = milestone.Name
			diamond.X = line.X1 - diamondWidth/2.0
			diamond.Y = mapLane_TextY[diamondAndTextAnchor] - diamondWidth + LaneHeight/2.0
			diamond.Width = diamondWidth
			diamond.Height = diamondWidth
			diamond.Color = "red"
			diamond.FillOpacity = 0.4
			diamond.Transform = fmt.Sprintf("rotate(%d %d %d)", 45, int64(diamond.X+diamondWidth/2.0), int64(diamond.Y+diamondWidth/2.0))

			// bar text
			milestoneText := new(gongsvg_models.Text).Stage()
			milestoneText.Name = milestone.Name
			milestoneText.Content = milestoneText.Name
			milestoneText.X = diamond.X + XLeftText + diamondWidth
			milestoneText.Y = diamond.Y + TextHeight/2.0 + 5 // manual fine tuning
			milestoneText.Color = "black"
			milestoneText.FillOpacity = 1.0
			svg.Texts = append(svg.Texts, milestoneText)
		}
	}

	//
	// Groups of Lanes
	//
	for _, group := range ganttToRender.Groups {

		if len(group.GroupLanes) == 0 {
			continue
		}

		groupSVG := new(gongsvg_models.Rect).Stage()
		svg.Rects = append(svg.Rects, groupSVG)
		groupSVG.Name = group.Name

		// compute X from list of lane
		groupTopY := math.MaxFloat64
		groupBottomY := 0.0
		for _, lane := range group.GroupLanes {
			if groupTopY > mapLane_TopY[lane] {
				groupTopY = mapLane_TopY[lane]
			}
			if groupBottomY < (mapLane_TopY[lane] + ganttToRender.LaneHeight) {
				groupBottomY = mapLane_TopY[lane] + ganttToRender.LaneHeight
			}
		}

		groupSVG.X = 0
		groupSVG.Y = groupTopY
		groupSVG.Width = ganttToRender.XRightMargin - groupSVG.X
		groupSVG.Height = groupBottomY - groupTopY

		groupSVG.Stroke = ganttToRender.Group_Stroke
		groupSVG.StrokeWidth = ganttToRender.Group_StrokeWidth
		groupSVG.StrokeDashArray = ganttToRender.Group_StrokeDashArray

		// text
		groupText := new(gongsvg_models.Text).Stage()
		groupText.Name = group.Name
		groupText.Content = groupText.Name
		groupText.X = XLeftText
		groupText.Y = groupTopY + ganttToRender.DateYOffset
		groupText.Color = "blue"
		groupText.FillOpacity = 0.5
		svg.Texts = append(svg.Texts, groupText)
	}

	gongsvg_models.Stage.Commit()

	// log.Printf("Before Commit")
}
