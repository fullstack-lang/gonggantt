package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gonggantt/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram2 models.StageStruct
var ___dummy__Time_NewDiagram2 time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram2 ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram2 map[string]any = map[string]any{
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

	"ref_models.GONG__ENUM_CAST_INT": ref_models.GONG__ENUM_CAST_INT,

	"ref_models.GONG__ENUM_CAST_STRING": ref_models.GONG__ENUM_CAST_STRING,

	"ref_models.GONG__ExpressionType": ref_models.GONG__ExpressionType(""),

	"ref_models.GONG__FIELD_OR_CONST_VALUE": ref_models.GONG__FIELD_OR_CONST_VALUE,

	"ref_models.GONG__FIELD_VALUE": ref_models.GONG__FIELD_VALUE,

	"ref_models.GONG__IDENTIFIER_CONST": ref_models.GONG__IDENTIFIER_CONST,

	"ref_models.GONG__STRUCT_INSTANCE": ref_models.GONG__STRUCT_INSTANCE,

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
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagram2Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_End := (&models.Field{Name: `End`}).Stage()
	__Field__000001_Start := (&models.Field{Name: `Start`}).Stage()

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Bar := (&models.GongStructShape{Name: `NewDiagram-Bar`}).Stage()
	__GongStructShape__000001_NewDiagram_Gantt := (&models.GongStructShape{Name: `NewDiagram-Gantt`}).Stage()
	__GongStructShape__000002_NewDiagram_Lane := (&models.GongStructShape{Name: `NewDiagram-Lane`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Bars := (&models.Link{Name: `Bars`}).Stage()
	__Link__000001_Lanes := (&models.Link{Name: `Lanes`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Bar := (&models.Position{Name: `Pos-NewDiagram-Bar`}).Stage()
	__Position__000001_Pos_NewDiagram_Gantt := (&models.Position{Name: `Pos-NewDiagram-Gantt`}).Stage()
	__Position__000002_Pos_NewDiagram_Lane := (&models.Position{Name: `Pos-NewDiagram-Lane`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Lane := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Gantt and NewDiagram-Lane`}).Stage()
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Lane_and_NewDiagram_Bar := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Lane and NewDiagram-Bar`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.End]
	__Field__000000_End.Identifier = `ref_models.Bar.End`
	__Field__000000_End.FieldTypeAsString = ``
	__Field__000000_End.Structname = `Bar`
	__Field__000000_End.Fieldtypename = `Time`

	// Field values setup
	__Field__000001_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.Start]
	__Field__000001_Start.Identifier = `ref_models.Bar.Start`
	__Field__000001_Start.FieldTypeAsString = ``
	__Field__000001_Start.Structname = `Bar`
	__Field__000001_Start.Fieldtypename = `Time`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Bar.Name = `NewDiagram-Bar`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar]
	__GongStructShape__000000_NewDiagram_Bar.Identifier = `ref_models.Bar`
	__GongStructShape__000000_NewDiagram_Bar.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Bar.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Bar.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Bar.Heigth = 93.000000
	__GongStructShape__000000_NewDiagram_Bar.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Gantt.Name = `NewDiagram-Gantt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt]
	__GongStructShape__000001_NewDiagram_Gantt.Identifier = `ref_models.Gantt`
	__GongStructShape__000001_NewDiagram_Gantt.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Gantt.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Gantt.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Gantt.Heigth = 63.000000
	__GongStructShape__000001_NewDiagram_Gantt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Lane.Name = `NewDiagram-Lane`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane]
	__GongStructShape__000002_NewDiagram_Lane.Identifier = `ref_models.Lane`
	__GongStructShape__000002_NewDiagram_Lane.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_Lane.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Lane.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Lane.Heigth = 63.000000
	__GongStructShape__000002_NewDiagram_Lane.IsSelected = false

	// Link values setup
	__Link__000000_Bars.Name = `Bars`
	__Link__000000_Bars.Structname = `Lane`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane.Bars]
	__Link__000000_Bars.Identifier = `ref_models.Lane.Bars`
	__Link__000000_Bars.Fieldtypename = `Bar`
	__Link__000000_Bars.TargetMultiplicity = models.MANY
	__Link__000000_Bars.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000001_Lanes.Name = `Lanes`
	__Link__000001_Lanes.Structname = `Gantt`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Lanes]
	__Link__000001_Lanes.Identifier = `ref_models.Gantt.Lanes`
	__Link__000001_Lanes.Fieldtypename = `Lane`
	__Link__000001_Lanes.TargetMultiplicity = models.MANY
	__Link__000001_Lanes.SourceMultiplicity = models.ZERO_ONE

	// Position values setup
	__Position__000000_Pos_NewDiagram_Bar.X = 650.000000
	__Position__000000_Pos_NewDiagram_Bar.Y = 60.000000
	__Position__000000_Pos_NewDiagram_Bar.Name = `Pos-NewDiagram-Bar`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Gantt.X = 40.000000
	__Position__000001_Pos_NewDiagram_Gantt.Y = 60.000000
	__Position__000001_Pos_NewDiagram_Gantt.Name = `Pos-NewDiagram-Gantt`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Lane.X = 350.000000
	__Position__000002_Pos_NewDiagram_Lane.Y = 60.000000
	__Position__000002_Pos_NewDiagram_Lane.Name = `Pos-NewDiagram-Lane`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Lane.X = 284.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Lane.Y = 187.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Lane.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Gantt and NewDiagram-Lane`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Lane_and_NewDiagram_Bar.X = 614.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Lane_and_NewDiagram_Bar.Y = 203.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Lane_and_NewDiagram_Bar.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Lane and NewDiagram-Bar`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Gantt)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Lane)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Bar)
	__GongStructShape__000000_NewDiagram_Bar.Position = __Position__000000_Pos_NewDiagram_Bar
	__GongStructShape__000000_NewDiagram_Bar.Fields = append(__GongStructShape__000000_NewDiagram_Bar.Fields, __Field__000001_Start)
	__GongStructShape__000000_NewDiagram_Bar.Fields = append(__GongStructShape__000000_NewDiagram_Bar.Fields, __Field__000000_End)
	__GongStructShape__000001_NewDiagram_Gantt.Position = __Position__000001_Pos_NewDiagram_Gantt
	__GongStructShape__000001_NewDiagram_Gantt.Links = append(__GongStructShape__000001_NewDiagram_Gantt.Links, __Link__000001_Lanes)
	__GongStructShape__000002_NewDiagram_Lane.Position = __Position__000002_Pos_NewDiagram_Lane
	__GongStructShape__000002_NewDiagram_Lane.Links = append(__GongStructShape__000002_NewDiagram_Lane.Links, __Link__000000_Bars)
	__Link__000000_Bars.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Lane_and_NewDiagram_Bar
	__Link__000001_Lanes.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Lane
}
