package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gonggantt/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Arrow": &(ref_models.Arrow{}),

	"ref_models.Arrow.From": (ref_models.Arrow{}).From,

	"ref_models.Arrow.Name": (ref_models.Arrow{}).Name,

	"ref_models.Arrow.OptionnalColor": (ref_models.Arrow{}).OptionnalColor,

	"ref_models.Arrow.OptionnalStroke": (ref_models.Arrow{}).OptionnalStroke,

	"ref_models.Arrow.To": (ref_models.Arrow{}).To,

	"ref_models.Bar": &(ref_models.Bar{}),

	"ref_models.Bar.End": (ref_models.Bar{}).End,

	"ref_models.Bar.FillOpacity": (ref_models.Bar{}).FillOpacity,

	"ref_models.Bar.Name": (ref_models.Bar{}).Name,

	"ref_models.Bar.OptionnalColor": (ref_models.Bar{}).OptionnalColor,

	"ref_models.Bar.OptionnalStroke": (ref_models.Bar{}).OptionnalStroke,

	"ref_models.Bar.Start": (ref_models.Bar{}).Start,

	"ref_models.Bar.StrokeDashArray": (ref_models.Bar{}).StrokeDashArray,

	"ref_models.Bar.StrokeWidth": (ref_models.Bar{}).StrokeWidth,

	"ref_models.Gantt": &(ref_models.Gantt{}),

	"ref_models.Gantt.AlignOnStartEndOnYearStart": (ref_models.Gantt{}).AlignOnStartEndOnYearStart,

	"ref_models.Gantt.ArrowLengthToTheRightOfStartBar": (ref_models.Gantt{}).ArrowLengthToTheRightOfStartBar,

	"ref_models.Gantt.ArrowTipLenght": (ref_models.Gantt{}).ArrowTipLenght,

	"ref_models.Gantt.Arrows": (ref_models.Gantt{}).Arrows,

	"ref_models.Gantt.ComputedEnd": (ref_models.Gantt{}).ComputedEnd,

	"ref_models.Gantt.ComputedStart": (ref_models.Gantt{}).ComputedStart,

	"ref_models.Gantt.DateYOffset": (ref_models.Gantt{}).DateYOffset,

	"ref_models.Gantt.Group_Stroke": (ref_models.Gantt{}).Group_Stroke,

	"ref_models.Gantt.Group_StrokeDashArray": (ref_models.Gantt{}).Group_StrokeDashArray,

	"ref_models.Gantt.Group_StrokeWidth": (ref_models.Gantt{}).Group_StrokeWidth,

	"ref_models.Gantt.Groups": (ref_models.Gantt{}).Groups,

	"ref_models.Gantt.LaneHeight": (ref_models.Gantt{}).LaneHeight,

	"ref_models.Gantt.Lanes": (ref_models.Gantt{}).Lanes,

	"ref_models.Gantt.ManualEnd": (ref_models.Gantt{}).ManualEnd,

	"ref_models.Gantt.ManualStart": (ref_models.Gantt{}).ManualStart,

	"ref_models.Gantt.Milestones": (ref_models.Gantt{}).Milestones,

	"ref_models.Gantt.Name": (ref_models.Gantt{}).Name,

	"ref_models.Gantt.RatioBarToLaneHeight": (ref_models.Gantt{}).RatioBarToLaneHeight,

	"ref_models.Gantt.TextHeight": (ref_models.Gantt{}).TextHeight,

	"ref_models.Gantt.TimeLine_Color": (ref_models.Gantt{}).TimeLine_Color,

	"ref_models.Gantt.TimeLine_FillOpacity": (ref_models.Gantt{}).TimeLine_FillOpacity,

	"ref_models.Gantt.TimeLine_Stroke": (ref_models.Gantt{}).TimeLine_Stroke,

	"ref_models.Gantt.TimeLine_StrokeWidth": (ref_models.Gantt{}).TimeLine_StrokeWidth,

	"ref_models.Gantt.UseManualStartAndEndDates": (ref_models.Gantt{}).UseManualStartAndEndDates,

	"ref_models.Gantt.XLeftLanes": (ref_models.Gantt{}).XLeftLanes,

	"ref_models.Gantt.XLeftText": (ref_models.Gantt{}).XLeftText,

	"ref_models.Gantt.XRightMargin": (ref_models.Gantt{}).XRightMargin,

	"ref_models.Gantt.YTopMargin": (ref_models.Gantt{}).YTopMargin,

	"ref_models.Group": &(ref_models.Group{}),

	"ref_models.Group.GroupLanes": (ref_models.Group{}).GroupLanes,

	"ref_models.Group.Name": (ref_models.Group{}).Name,

	"ref_models.Lane": &(ref_models.Lane{}),

	"ref_models.Lane.Bars": (ref_models.Lane{}).Bars,

	"ref_models.Lane.Name": (ref_models.Lane{}).Name,

	"ref_models.Lane.Order": (ref_models.Lane{}).Order,

	"ref_models.LaneUse": &(ref_models.LaneUse{}),

	"ref_models.LaneUse.Lane": (ref_models.LaneUse{}).Lane,

	"ref_models.LaneUse.Name": (ref_models.LaneUse{}).Name,

	"ref_models.Milestone": &(ref_models.Milestone{}),

	"ref_models.Milestone.Date": (ref_models.Milestone{}).Date,

	"ref_models.Milestone.DisplayVerticalBar": (ref_models.Milestone{}).DisplayVerticalBar,

	"ref_models.Milestone.LanesToDisplayMilestoneUse": (ref_models.Milestone{}).LanesToDisplayMilestoneUse,

	"ref_models.Milestone.Name": (ref_models.Milestone{}).Name,

	"ref_models.NoteOnTheModel": ref_models.NoteOnTheModel,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Button

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Setup of pointers
}


