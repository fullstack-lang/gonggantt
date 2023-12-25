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

	"ref_models.Bar.ComputedDuration": (ref_models.Bar{}).ComputedDuration,

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

	"ref_models.Gantt.ComputedDuration": (ref_models.Gantt{}).ComputedDuration,

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

	"ref_models.GanttStackName": ref_models.GanttStackName,

	"ref_models.GanttStacksNames": ref_models.GanttStacksNames(""),

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

	"ref_models.SvgStackName": ref_models.SvgStackName,
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

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_End := (&models.Field{Name: `End`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000002_Start := (&models.Field{Name: `Start`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Arrow := (&models.GongStructShape{Name: `Default-Arrow`}).Stage(stage)
	__GongStructShape__000001_Default_Bar := (&models.GongStructShape{Name: `Default-Bar`}).Stage(stage)
	__GongStructShape__000002_Default_Gantt := (&models.GongStructShape{Name: `Default-Gantt`}).Stage(stage)
	__GongStructShape__000003_Default_Lane := (&models.GongStructShape{Name: `Default-Lane`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Bars := (&models.Link{Name: `Bars`}).Stage(stage)
	__Link__000001_From := (&models.Link{Name: `From`}).Stage(stage)
	__Link__000002_Lanes := (&models.Link{Name: `Lanes`}).Stage(stage)

	// Declarations of staged instances of NoteShape
	__NoteShape__000000_NoteOnTheModel := (&models.NoteShape{Name: `NoteOnTheModel`}).Stage(stage)

	// Declarations of staged instances of NoteShapeLink
	__NoteShapeLink__000000_Bar := (&models.NoteShapeLink{Name: `Bar`}).Stage(stage)
	__NoteShapeLink__000001_Gantt := (&models.NoteShapeLink{Name: `Gantt`}).Stage(stage)
	__NoteShapeLink__000002_Lane := (&models.NoteShapeLink{Name: `Lane`}).Stage(stage)

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Arrow := (&models.Position{Name: `Pos-Default-Arrow`}).Stage(stage)
	__Position__000001_Pos_Default_Bar := (&models.Position{Name: `Pos-Default-Bar`}).Stage(stage)
	__Position__000002_Pos_Default_Gantt := (&models.Position{Name: `Pos-Default-Gantt`}).Stage(stage)
	__Position__000003_Pos_Default_Lane := (&models.Position{Name: `Pos-Default-Lane`}).Stage(stage)

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Arrow_and_Default_Bar := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Arrow and Default-Bar`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Gantt_and_Default_Lane := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Gantt and Default-Lane`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Lane_and_Default_Bar := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Lane and Default-Bar`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true

	// Field values setup
	__Field__000000_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.End]
	__Field__000000_End.Identifier = `ref_models.Bar.End`
	__Field__000000_End.FieldTypeAsString = ``
	__Field__000000_End.Structname = `Bar`
	__Field__000000_End.Fieldtypename = `Time`

	// Field values setup
	__Field__000001_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.Name]
	__Field__000001_Name.Identifier = `ref_models.Arrow.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `Arrow`
	__Field__000001_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000002_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.Start]
	__Field__000002_Start.Identifier = `ref_models.Bar.Start`
	__Field__000002_Start.FieldTypeAsString = ``
	__Field__000002_Start.Structname = `Bar`
	__Field__000002_Start.Fieldtypename = `Time`

	// GongStructShape values setup
	__GongStructShape__000000_Default_Arrow.Name = `Default-Arrow`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow]
	__GongStructShape__000000_Default_Arrow.Identifier = `ref_models.Arrow`
	__GongStructShape__000000_Default_Arrow.ShowNbInstances = false
	__GongStructShape__000000_Default_Arrow.NbInstances = 0
	__GongStructShape__000000_Default_Arrow.Width = 240.000000
	__GongStructShape__000000_Default_Arrow.Height = 78.000000
	__GongStructShape__000000_Default_Arrow.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_Bar.Name = `Default-Bar`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar]
	__GongStructShape__000001_Default_Bar.Identifier = `ref_models.Bar`
	__GongStructShape__000001_Default_Bar.ShowNbInstances = true
	__GongStructShape__000001_Default_Bar.NbInstances = 4
	__GongStructShape__000001_Default_Bar.Width = 240.000000
	__GongStructShape__000001_Default_Bar.Height = 93.000000
	__GongStructShape__000001_Default_Bar.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_Gantt.Name = `Default-Gantt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt]
	__GongStructShape__000002_Default_Gantt.Identifier = `ref_models.Gantt`
	__GongStructShape__000002_Default_Gantt.ShowNbInstances = true
	__GongStructShape__000002_Default_Gantt.NbInstances = 1
	__GongStructShape__000002_Default_Gantt.Width = 240.000000
	__GongStructShape__000002_Default_Gantt.Height = 63.000000
	__GongStructShape__000002_Default_Gantt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_Default_Lane.Name = `Default-Lane`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane]
	__GongStructShape__000003_Default_Lane.Identifier = `ref_models.Lane`
	__GongStructShape__000003_Default_Lane.ShowNbInstances = true
	__GongStructShape__000003_Default_Lane.NbInstances = 4
	__GongStructShape__000003_Default_Lane.Width = 240.000000
	__GongStructShape__000003_Default_Lane.Height = 63.000000
	__GongStructShape__000003_Default_Lane.IsSelected = false

	// Link values setup
	__Link__000000_Bars.Name = `Bars`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane.Bars]
	__Link__000000_Bars.Identifier = `ref_models.Lane.Bars`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar]
	__Link__000000_Bars.Fieldtypename = `ref_models.Bar`
	__Link__000000_Bars.FieldOffsetX = -50.000000
	__Link__000000_Bars.FieldOffsetY = -16.000000
	__Link__000000_Bars.TargetMultiplicity = models.MANY
	__Link__000000_Bars.TargetMultiplicityOffsetX = -31.000000
	__Link__000000_Bars.TargetMultiplicityOffsetY = 24.000000
	__Link__000000_Bars.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Bars.SourceMultiplicityOffsetX = 20.000000
	__Link__000000_Bars.SourceMultiplicityOffsetY = 26.000000
	__Link__000000_Bars.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Bars.StartRatio = 0.500000
	__Link__000000_Bars.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Bars.EndRatio = 0.408602
	__Link__000000_Bars.CornerOffsetRatio = 1.334444

	// Link values setup
	__Link__000001_From.Name = `From`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.From]
	__Link__000001_From.Identifier = `ref_models.Arrow.From`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar]
	__Link__000001_From.Fieldtypename = `ref_models.Bar`
	__Link__000001_From.FieldOffsetX = -50.000000
	__Link__000001_From.FieldOffsetY = -16.000000
	__Link__000001_From.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_From.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_From.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_From.SourceMultiplicity = models.MANY
	__Link__000001_From.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_From.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_From.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_From.StartRatio = 0.500000
	__Link__000001_From.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_From.EndRatio = 0.500000
	__Link__000001_From.CornerOffsetRatio = 1.301667

	// Link values setup
	__Link__000002_Lanes.Name = `Lanes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Lanes]
	__Link__000002_Lanes.Identifier = `ref_models.Gantt.Lanes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane]
	__Link__000002_Lanes.Fieldtypename = `ref_models.Lane`
	__Link__000002_Lanes.FieldOffsetX = -87.000000
	__Link__000002_Lanes.FieldOffsetY = -16.000000
	__Link__000002_Lanes.TargetMultiplicity = models.MANY
	__Link__000002_Lanes.TargetMultiplicityOffsetX = -46.000000
	__Link__000002_Lanes.TargetMultiplicityOffsetY = 26.000000
	__Link__000002_Lanes.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_Lanes.SourceMultiplicityOffsetX = 21.000000
	__Link__000002_Lanes.SourceMultiplicityOffsetY = 18.000000
	__Link__000002_Lanes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Lanes.StartRatio = 0.500000
	__Link__000002_Lanes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_Lanes.EndRatio = 0.500000
	__Link__000002_Lanes.CornerOffsetRatio = 1.388611

	// NoteShape values setup
	__NoteShape__000000_NoteOnTheModel.Name = `NoteOnTheModel`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.NoteOnTheModel]
	__NoteShape__000000_NoteOnTheModel.Identifier = `ref_models.NoteOnTheModel`
	__NoteShape__000000_NoteOnTheModel.Body = `A Gantt Diagram
is an instance of [models.Gantt] that has
a slice of lanes ([models.Lane]).

Each lane vertical position is defined by it order field [models.Lane.Order].
Each has a slice of tasks

A task is displayed as a bar ([models.Bar]) with a start and end date
[models.Bar.StartDate] and [models.Bar.EndDate]).
`
	__NoteShape__000000_NoteOnTheModel.BodyHTML = `<p>A Gantt Diagram
is an instance of <a href="/models#Gantt">models.Gantt</a> that has
a slice of lanes (<a href="/models#Lane">models.Lane</a>).
<p>Each lane vertical position is defined by it order field <a href="/models#Lane.Order">models.Lane.Order</a>.
Each has a slice of tasks
<p>A task is displayed as a bar (<a href="/models#Bar">models.Bar</a>) with a start and end date
<a href="/models#Bar.StartDate">models.Bar.StartDate</a> and <a href="/models#Bar.EndDate">models.Bar.EndDate</a>).
`
	__NoteShape__000000_NoteOnTheModel.X = 101.000000
	__NoteShape__000000_NoteOnTheModel.Y = 349.000000
	__NoteShape__000000_NoteOnTheModel.Width = 653.999939
	__NoteShape__000000_NoteOnTheModel.Height = 181.000000
	__NoteShape__000000_NoteOnTheModel.Matched = false

	// NoteShapeLink values setup
	__NoteShapeLink__000000_Bar.Name = `Bar`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar]
	__NoteShapeLink__000000_Bar.Identifier = `ref_models.Bar`
	__NoteShapeLink__000000_Bar.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000001_Gantt.Name = `Gantt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt]
	__NoteShapeLink__000001_Gantt.Identifier = `ref_models.Gantt`
	__NoteShapeLink__000001_Gantt.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// NoteShapeLink values setup
	__NoteShapeLink__000002_Lane.Name = `Lane`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane]
	__NoteShapeLink__000002_Lane.Identifier = `ref_models.Lane`
	__NoteShapeLink__000002_Lane.Type = models.NOTE_SHAPE_LINK_TO_GONG_STRUCT_OR_ENUM_SHAPE

	// Position values setup
	__Position__000000_Pos_Default_Arrow.X = 933.000000
	__Position__000000_Pos_Default_Arrow.Y = 312.000000
	__Position__000000_Pos_Default_Arrow.Name = `Pos-Default-Arrow`

	// Position values setup
	__Position__000001_Pos_Default_Bar.X = 941.000000
	__Position__000001_Pos_Default_Bar.Y = 93.000000
	__Position__000001_Pos_Default_Bar.Name = `Pos-Default-Bar`

	// Position values setup
	__Position__000002_Pos_Default_Gantt.X = 34.000000
	__Position__000002_Pos_Default_Gantt.Y = 100.000000
	__Position__000002_Pos_Default_Gantt.Name = `Pos-Default-Gantt`

	// Position values setup
	__Position__000003_Pos_Default_Lane.X = 506.000000
	__Position__000003_Pos_Default_Lane.Y = 96.000000
	__Position__000003_Pos_Default_Lane.Name = `Pos-Default-Lane`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Arrow_and_Default_Bar.X = 851.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Arrow_and_Default_Bar.Y = 199.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Arrow_and_Default_Bar.Name = `Verticle in class diagram Default in middle between Default-Arrow and Default-Bar`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Gantt_and_Default_Lane.X = 651.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Gantt_and_Default_Lane.Y = 130.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Gantt_and_Default_Lane.Name = `Verticle in class diagram Default in middle between Default-Gantt and Default-Lane`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Lane_and_Default_Bar.X = 1083.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Lane_and_Default_Bar.Y = 126.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Lane_and_Default_Bar.Name = `Verticle in class diagram Default in middle between Default-Lane and Default-Bar`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Bar)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Gantt)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_Lane)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Arrow)
	__Classdiagram__000000_Default.NoteShapes = append(__Classdiagram__000000_Default.NoteShapes, __NoteShape__000000_NoteOnTheModel)
	__GongStructShape__000000_Default_Arrow.Position = __Position__000000_Pos_Default_Arrow
	__GongStructShape__000000_Default_Arrow.Fields = append(__GongStructShape__000000_Default_Arrow.Fields, __Field__000001_Name)
	__GongStructShape__000000_Default_Arrow.Links = append(__GongStructShape__000000_Default_Arrow.Links, __Link__000001_From)
	__GongStructShape__000001_Default_Bar.Position = __Position__000001_Pos_Default_Bar
	__GongStructShape__000001_Default_Bar.Fields = append(__GongStructShape__000001_Default_Bar.Fields, __Field__000002_Start)
	__GongStructShape__000001_Default_Bar.Fields = append(__GongStructShape__000001_Default_Bar.Fields, __Field__000000_End)
	__GongStructShape__000002_Default_Gantt.Position = __Position__000002_Pos_Default_Gantt
	__GongStructShape__000002_Default_Gantt.Links = append(__GongStructShape__000002_Default_Gantt.Links, __Link__000002_Lanes)
	__GongStructShape__000003_Default_Lane.Position = __Position__000003_Pos_Default_Lane
	__GongStructShape__000003_Default_Lane.Links = append(__GongStructShape__000003_Default_Lane.Links, __Link__000000_Bars)
	__Link__000000_Bars.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Lane_and_Default_Bar
	__Link__000001_From.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Arrow_and_Default_Bar
	__Link__000002_Lanes.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Gantt_and_Default_Lane
	__NoteShape__000000_NoteOnTheModel.NoteShapeLinks = append(__NoteShape__000000_NoteOnTheModel.NoteShapeLinks, __NoteShapeLink__000000_Bar)
	__NoteShape__000000_NoteOnTheModel.NoteShapeLinks = append(__NoteShape__000000_NoteOnTheModel.NoteShapeLinks, __NoteShapeLink__000001_Gantt)
	__NoteShape__000000_NoteOnTheModel.NoteShapeLinks = append(__NoteShape__000000_NoteOnTheModel.NoteShapeLinks, __NoteShapeLink__000002_Lane)
}


