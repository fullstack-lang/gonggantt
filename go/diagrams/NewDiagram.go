package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gonggantt/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
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
func NewDiagramInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_AlignOnStartEndOnYearStart := (&models.Field{Name: `AlignOnStartEndOnYearStart`}).Stage()
	__Field__000001_ArrowLengthToTheRightOfStartBar := (&models.Field{Name: `ArrowLengthToTheRightOfStartBar`}).Stage()
	__Field__000002_ArrowTipLenght := (&models.Field{Name: `ArrowTipLenght`}).Stage()
	__Field__000003_ComputedEnd := (&models.Field{Name: `ComputedEnd`}).Stage()
	__Field__000004_ComputedStart := (&models.Field{Name: `ComputedStart`}).Stage()
	__Field__000005_End := (&models.Field{Name: `End`}).Stage()
	__Field__000006_FillOpacity := (&models.Field{Name: `FillOpacity`}).Stage()
	__Field__000007_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000008_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000009_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000010_OptionnalColor := (&models.Field{Name: `OptionnalColor`}).Stage()
	__Field__000011_OptionnalColor := (&models.Field{Name: `OptionnalColor`}).Stage()
	__Field__000012_OptionnalStroke := (&models.Field{Name: `OptionnalStroke`}).Stage()
	__Field__000013_OptionnalStroke := (&models.Field{Name: `OptionnalStroke`}).Stage()
	__Field__000014_Start := (&models.Field{Name: `Start`}).Stage()
	__Field__000015_StrokeDashArray := (&models.Field{Name: `StrokeDashArray`}).Stage()
	__Field__000016_StrokeWidth := (&models.Field{Name: `StrokeWidth`}).Stage()

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Arrow := (&models.GongStructShape{Name: `NewDiagram-Arrow`}).Stage()
	__GongStructShape__000001_NewDiagram_Bar := (&models.GongStructShape{Name: `NewDiagram-Bar`}).Stage()
	__GongStructShape__000002_NewDiagram_Gantt := (&models.GongStructShape{Name: `NewDiagram-Gantt`}).Stage()
	__GongStructShape__000003_NewDiagram_Group := (&models.GongStructShape{Name: `NewDiagram-Group`}).Stage()
	__GongStructShape__000004_NewDiagram_Lane := (&models.GongStructShape{Name: `NewDiagram-Lane`}).Stage()
	__GongStructShape__000005_NewDiagram_LaneUse := (&models.GongStructShape{Name: `NewDiagram-LaneUse`}).Stage()
	__GongStructShape__000006_NewDiagram_Milestone := (&models.GongStructShape{Name: `NewDiagram-Milestone`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Arrows := (&models.Link{Name: `Arrows`}).Stage()
	__Link__000001_From := (&models.Link{Name: `From`}).Stage()
	__Link__000002_GroupLanes := (&models.Link{Name: `GroupLanes`}).Stage()
	__Link__000003_To := (&models.Link{Name: `To`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Arrow := (&models.Position{Name: `Pos-NewDiagram-Arrow`}).Stage()
	__Position__000001_Pos_NewDiagram_Bar := (&models.Position{Name: `Pos-NewDiagram-Bar`}).Stage()
	__Position__000002_Pos_NewDiagram_Gantt := (&models.Position{Name: `Pos-NewDiagram-Gantt`}).Stage()
	__Position__000003_Pos_NewDiagram_Group := (&models.Position{Name: `Pos-NewDiagram-Group`}).Stage()
	__Position__000004_Pos_NewDiagram_Lane := (&models.Position{Name: `Pos-NewDiagram-Lane`}).Stage()
	__Position__000005_Pos_NewDiagram_LaneUse := (&models.Position{Name: `Pos-NewDiagram-LaneUse`}).Stage()
	__Position__000006_Pos_NewDiagram_Milestone := (&models.Position{Name: `Pos-NewDiagram-Milestone`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Arrow and NewDiagram-Bar`}).Stage()
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Arrow and NewDiagram-Bar`}).Stage()
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Arrow := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Gantt and NewDiagram-Arrow`}).Stage()
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Group_and_NewDiagram_Lane := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Group and NewDiagram-Lane`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_AlignOnStartEndOnYearStart.Name = `AlignOnStartEndOnYearStart`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.AlignOnStartEndOnYearStart]
	__Field__000000_AlignOnStartEndOnYearStart.Identifier = `ref_models.Gantt.AlignOnStartEndOnYearStart`
	__Field__000000_AlignOnStartEndOnYearStart.FieldTypeAsString = ``
	__Field__000000_AlignOnStartEndOnYearStart.Structname = `Gantt`
	__Field__000000_AlignOnStartEndOnYearStart.Fieldtypename = `bool`

	// Field values setup
	__Field__000001_ArrowLengthToTheRightOfStartBar.Name = `ArrowLengthToTheRightOfStartBar`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.ArrowLengthToTheRightOfStartBar]
	__Field__000001_ArrowLengthToTheRightOfStartBar.Identifier = `ref_models.Gantt.ArrowLengthToTheRightOfStartBar`
	__Field__000001_ArrowLengthToTheRightOfStartBar.FieldTypeAsString = ``
	__Field__000001_ArrowLengthToTheRightOfStartBar.Structname = `Gantt`
	__Field__000001_ArrowLengthToTheRightOfStartBar.Fieldtypename = `float64`

	// Field values setup
	__Field__000002_ArrowTipLenght.Name = `ArrowTipLenght`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.ArrowTipLenght]
	__Field__000002_ArrowTipLenght.Identifier = `ref_models.Gantt.ArrowTipLenght`
	__Field__000002_ArrowTipLenght.FieldTypeAsString = ``
	__Field__000002_ArrowTipLenght.Structname = `Gantt`
	__Field__000002_ArrowTipLenght.Fieldtypename = `float64`

	// Field values setup
	__Field__000003_ComputedEnd.Name = `ComputedEnd`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.ComputedEnd]
	__Field__000003_ComputedEnd.Identifier = `ref_models.Gantt.ComputedEnd`
	__Field__000003_ComputedEnd.FieldTypeAsString = ``
	__Field__000003_ComputedEnd.Structname = `Gantt`
	__Field__000003_ComputedEnd.Fieldtypename = `Time`

	// Field values setup
	__Field__000004_ComputedStart.Name = `ComputedStart`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.ComputedStart]
	__Field__000004_ComputedStart.Identifier = `ref_models.Gantt.ComputedStart`
	__Field__000004_ComputedStart.FieldTypeAsString = ``
	__Field__000004_ComputedStart.Structname = `Gantt`
	__Field__000004_ComputedStart.Fieldtypename = `Time`

	// Field values setup
	__Field__000005_End.Name = `End`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.End]
	__Field__000005_End.Identifier = `ref_models.Bar.End`
	__Field__000005_End.FieldTypeAsString = ``
	__Field__000005_End.Structname = `Bar`
	__Field__000005_End.Fieldtypename = `Time`

	// Field values setup
	__Field__000006_FillOpacity.Name = `FillOpacity`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.FillOpacity]
	__Field__000006_FillOpacity.Identifier = `ref_models.Bar.FillOpacity`
	__Field__000006_FillOpacity.FieldTypeAsString = ``
	__Field__000006_FillOpacity.Structname = `Bar`
	__Field__000006_FillOpacity.Fieldtypename = `float64`

	// Field values setup
	__Field__000007_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.Name]
	__Field__000007_Name.Identifier = `ref_models.Arrow.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = `Arrow`
	__Field__000007_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Group.Name]
	__Field__000008_Name.Identifier = `ref_models.Group.Name`
	__Field__000008_Name.FieldTypeAsString = ``
	__Field__000008_Name.Structname = `Group`
	__Field__000008_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000009_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.Name]
	__Field__000009_Name.Identifier = `ref_models.Bar.Name`
	__Field__000009_Name.FieldTypeAsString = ``
	__Field__000009_Name.Structname = `Bar`
	__Field__000009_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000010_OptionnalColor.Name = `OptionnalColor`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.OptionnalColor]
	__Field__000010_OptionnalColor.Identifier = `ref_models.Arrow.OptionnalColor`
	__Field__000010_OptionnalColor.FieldTypeAsString = ``
	__Field__000010_OptionnalColor.Structname = `Arrow`
	__Field__000010_OptionnalColor.Fieldtypename = `string`

	// Field values setup
	__Field__000011_OptionnalColor.Name = `OptionnalColor`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.OptionnalColor]
	__Field__000011_OptionnalColor.Identifier = `ref_models.Bar.OptionnalColor`
	__Field__000011_OptionnalColor.FieldTypeAsString = ``
	__Field__000011_OptionnalColor.Structname = `Bar`
	__Field__000011_OptionnalColor.Fieldtypename = `string`

	// Field values setup
	__Field__000012_OptionnalStroke.Name = `OptionnalStroke`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.OptionnalStroke]
	__Field__000012_OptionnalStroke.Identifier = `ref_models.Bar.OptionnalStroke`
	__Field__000012_OptionnalStroke.FieldTypeAsString = ``
	__Field__000012_OptionnalStroke.Structname = `Bar`
	__Field__000012_OptionnalStroke.Fieldtypename = `string`

	// Field values setup
	__Field__000013_OptionnalStroke.Name = `OptionnalStroke`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.OptionnalStroke]
	__Field__000013_OptionnalStroke.Identifier = `ref_models.Arrow.OptionnalStroke`
	__Field__000013_OptionnalStroke.FieldTypeAsString = ``
	__Field__000013_OptionnalStroke.Structname = `Arrow`
	__Field__000013_OptionnalStroke.Fieldtypename = `string`

	// Field values setup
	__Field__000014_Start.Name = `Start`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.Start]
	__Field__000014_Start.Identifier = `ref_models.Bar.Start`
	__Field__000014_Start.FieldTypeAsString = ``
	__Field__000014_Start.Structname = `Bar`
	__Field__000014_Start.Fieldtypename = `Time`

	// Field values setup
	__Field__000015_StrokeDashArray.Name = `StrokeDashArray`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.StrokeDashArray]
	__Field__000015_StrokeDashArray.Identifier = `ref_models.Bar.StrokeDashArray`
	__Field__000015_StrokeDashArray.FieldTypeAsString = ``
	__Field__000015_StrokeDashArray.Structname = `Bar`
	__Field__000015_StrokeDashArray.Fieldtypename = `string`

	// Field values setup
	__Field__000016_StrokeWidth.Name = `StrokeWidth`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.StrokeWidth]
	__Field__000016_StrokeWidth.Identifier = `ref_models.Bar.StrokeWidth`
	__Field__000016_StrokeWidth.FieldTypeAsString = ``
	__Field__000016_StrokeWidth.Structname = `Bar`
	__Field__000016_StrokeWidth.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Arrow.Name = `NewDiagram-Arrow`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow]
	__GongStructShape__000000_NewDiagram_Arrow.Identifier = `ref_models.Arrow`
	__GongStructShape__000000_NewDiagram_Arrow.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Arrow.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Arrow.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Arrow.Heigth = 108.000000
	__GongStructShape__000000_NewDiagram_Arrow.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Bar.Name = `NewDiagram-Bar`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar]
	__GongStructShape__000001_NewDiagram_Bar.Identifier = `ref_models.Bar`
	__GongStructShape__000001_NewDiagram_Bar.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Bar.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Bar.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Bar.Heigth = 183.000000
	__GongStructShape__000001_NewDiagram_Bar.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Gantt.Name = `NewDiagram-Gantt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt]
	__GongStructShape__000002_NewDiagram_Gantt.Identifier = `ref_models.Gantt`
	__GongStructShape__000002_NewDiagram_Gantt.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_Gantt.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Gantt.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Gantt.Heigth = 138.000000
	__GongStructShape__000002_NewDiagram_Gantt.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_Group.Name = `NewDiagram-Group`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Group]
	__GongStructShape__000003_NewDiagram_Group.Identifier = `ref_models.Group`
	__GongStructShape__000003_NewDiagram_Group.ShowNbInstances = false
	__GongStructShape__000003_NewDiagram_Group.NbInstances = 0
	__GongStructShape__000003_NewDiagram_Group.Width = 240.000000
	__GongStructShape__000003_NewDiagram_Group.Heigth = 78.000000
	__GongStructShape__000003_NewDiagram_Group.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_Lane.Name = `NewDiagram-Lane`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane]
	__GongStructShape__000004_NewDiagram_Lane.Identifier = `ref_models.Lane`
	__GongStructShape__000004_NewDiagram_Lane.ShowNbInstances = false
	__GongStructShape__000004_NewDiagram_Lane.NbInstances = 0
	__GongStructShape__000004_NewDiagram_Lane.Width = 240.000000
	__GongStructShape__000004_NewDiagram_Lane.Heigth = 63.000000
	__GongStructShape__000004_NewDiagram_Lane.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NewDiagram_LaneUse.Name = `NewDiagram-LaneUse`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LaneUse]
	__GongStructShape__000005_NewDiagram_LaneUse.Identifier = `ref_models.LaneUse`
	__GongStructShape__000005_NewDiagram_LaneUse.ShowNbInstances = false
	__GongStructShape__000005_NewDiagram_LaneUse.NbInstances = 0
	__GongStructShape__000005_NewDiagram_LaneUse.Width = 240.000000
	__GongStructShape__000005_NewDiagram_LaneUse.Heigth = 63.000000
	__GongStructShape__000005_NewDiagram_LaneUse.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NewDiagram_Milestone.Name = `NewDiagram-Milestone`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Milestone]
	__GongStructShape__000006_NewDiagram_Milestone.Identifier = `ref_models.Milestone`
	__GongStructShape__000006_NewDiagram_Milestone.ShowNbInstances = false
	__GongStructShape__000006_NewDiagram_Milestone.NbInstances = 0
	__GongStructShape__000006_NewDiagram_Milestone.Width = 240.000000
	__GongStructShape__000006_NewDiagram_Milestone.Heigth = 63.000000
	__GongStructShape__000006_NewDiagram_Milestone.IsSelected = false

	// Link values setup
	__Link__000000_Arrows.Name = `Arrows`
	__Link__000000_Arrows.Structname = `Gantt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Arrows]
	__Link__000000_Arrows.Identifier = `ref_models.Gantt.Arrows`
	__Link__000000_Arrows.Fieldtypename = `Arrow`
	__Link__000000_Arrows.TargetMultiplicity = models.MANY
	__Link__000000_Arrows.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000001_From.Name = `From`
	__Link__000001_From.Structname = `Arrow`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.From]
	__Link__000001_From.Identifier = `ref_models.Arrow.From`
	__Link__000001_From.Fieldtypename = `Bar`
	__Link__000001_From.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_From.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000002_GroupLanes.Name = `GroupLanes`
	__Link__000002_GroupLanes.Structname = `Group`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Group.GroupLanes]
	__Link__000002_GroupLanes.Identifier = `ref_models.Group.GroupLanes`
	__Link__000002_GroupLanes.Fieldtypename = `Lane`
	__Link__000002_GroupLanes.TargetMultiplicity = models.MANY
	__Link__000002_GroupLanes.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000003_To.Name = `To`
	__Link__000003_To.Structname = `Arrow`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.To]
	__Link__000003_To.Identifier = `ref_models.Arrow.To`
	__Link__000003_To.Fieldtypename = `Bar`
	__Link__000003_To.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_To.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Pos_NewDiagram_Arrow.X = 590.000000
	__Position__000000_Pos_NewDiagram_Arrow.Y = 560.000000
	__Position__000000_Pos_NewDiagram_Arrow.Name = `Pos-NewDiagram-Arrow`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Bar.X = 620.000000
	__Position__000001_Pos_NewDiagram_Bar.Y = 160.000000
	__Position__000001_Pos_NewDiagram_Bar.Name = `Pos-NewDiagram-Bar`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Gantt.X = 20.000000
	__Position__000002_Pos_NewDiagram_Gantt.Y = 160.000000
	__Position__000002_Pos_NewDiagram_Gantt.Name = `Pos-NewDiagram-Gantt`

	// Position values setup
	__Position__000003_Pos_NewDiagram_Group.X = 340.000000
	__Position__000003_Pos_NewDiagram_Group.Y = 50.000000
	__Position__000003_Pos_NewDiagram_Group.Name = `Pos-NewDiagram-Group`

	// Position values setup
	__Position__000004_Pos_NewDiagram_Lane.X = 440.000000
	__Position__000004_Pos_NewDiagram_Lane.Y = 460.000000
	__Position__000004_Pos_NewDiagram_Lane.Name = `Pos-NewDiagram-Lane`

	// Position values setup
	__Position__000005_Pos_NewDiagram_LaneUse.X = 190.000000
	__Position__000005_Pos_NewDiagram_LaneUse.Y = 240.000000
	__Position__000005_Pos_NewDiagram_LaneUse.Name = `Pos-NewDiagram-LaneUse`

	// Position values setup
	__Position__000006_Pos_NewDiagram_Milestone.X = 690.000000
	__Position__000006_Pos_NewDiagram_Milestone.Y = 400.000000
	__Position__000006_Pos_NewDiagram_Milestone.Name = `Pos-NewDiagram-Milestone`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar.X = 418.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar.Y = 86.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Arrow and NewDiagram-Bar`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar.X = 418.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar.Y = 101.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Arrow and NewDiagram-Bar`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Arrow.X = 525.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Arrow.Y = 339.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Arrow.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Gantt and NewDiagram-Arrow`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Group_and_NewDiagram_Lane.X = 383.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Group_and_NewDiagram_Lane.Y = 68.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Group_and_NewDiagram_Lane.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Group and NewDiagram-Lane`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Bar)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Gantt)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_Lane)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000005_NewDiagram_LaneUse)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_Milestone)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_Group)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Arrow)
	__GongStructShape__000000_NewDiagram_Arrow.Position = __Position__000000_Pos_NewDiagram_Arrow
	__GongStructShape__000000_NewDiagram_Arrow.Fields = append(__GongStructShape__000000_NewDiagram_Arrow.Fields, __Field__000007_Name)
	__GongStructShape__000000_NewDiagram_Arrow.Fields = append(__GongStructShape__000000_NewDiagram_Arrow.Fields, __Field__000010_OptionnalColor)
	__GongStructShape__000000_NewDiagram_Arrow.Fields = append(__GongStructShape__000000_NewDiagram_Arrow.Fields, __Field__000013_OptionnalStroke)
	__GongStructShape__000000_NewDiagram_Arrow.Links = append(__GongStructShape__000000_NewDiagram_Arrow.Links, __Link__000001_From)
	__GongStructShape__000000_NewDiagram_Arrow.Links = append(__GongStructShape__000000_NewDiagram_Arrow.Links, __Link__000003_To)
	__GongStructShape__000001_NewDiagram_Bar.Position = __Position__000001_Pos_NewDiagram_Bar
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000009_Name)
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000014_Start)
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000005_End)
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000011_OptionnalColor)
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000012_OptionnalStroke)
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000006_FillOpacity)
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000016_StrokeWidth)
	__GongStructShape__000001_NewDiagram_Bar.Fields = append(__GongStructShape__000001_NewDiagram_Bar.Fields, __Field__000015_StrokeDashArray)
	__GongStructShape__000002_NewDiagram_Gantt.Position = __Position__000002_Pos_NewDiagram_Gantt
	__GongStructShape__000002_NewDiagram_Gantt.Fields = append(__GongStructShape__000002_NewDiagram_Gantt.Fields, __Field__000004_ComputedStart)
	__GongStructShape__000002_NewDiagram_Gantt.Fields = append(__GongStructShape__000002_NewDiagram_Gantt.Fields, __Field__000003_ComputedEnd)
	__GongStructShape__000002_NewDiagram_Gantt.Fields = append(__GongStructShape__000002_NewDiagram_Gantt.Fields, __Field__000001_ArrowLengthToTheRightOfStartBar)
	__GongStructShape__000002_NewDiagram_Gantt.Fields = append(__GongStructShape__000002_NewDiagram_Gantt.Fields, __Field__000002_ArrowTipLenght)
	__GongStructShape__000002_NewDiagram_Gantt.Fields = append(__GongStructShape__000002_NewDiagram_Gantt.Fields, __Field__000000_AlignOnStartEndOnYearStart)
	__GongStructShape__000002_NewDiagram_Gantt.Links = append(__GongStructShape__000002_NewDiagram_Gantt.Links, __Link__000000_Arrows)
	__GongStructShape__000003_NewDiagram_Group.Position = __Position__000003_Pos_NewDiagram_Group
	__GongStructShape__000003_NewDiagram_Group.Fields = append(__GongStructShape__000003_NewDiagram_Group.Fields, __Field__000008_Name)
	__GongStructShape__000003_NewDiagram_Group.Links = append(__GongStructShape__000003_NewDiagram_Group.Links, __Link__000002_GroupLanes)
	__GongStructShape__000004_NewDiagram_Lane.Position = __Position__000004_Pos_NewDiagram_Lane
	__GongStructShape__000005_NewDiagram_LaneUse.Position = __Position__000005_Pos_NewDiagram_LaneUse
	__GongStructShape__000006_NewDiagram_Milestone.Position = __Position__000006_Pos_NewDiagram_Milestone
	__Link__000000_Arrows.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Gantt_and_NewDiagram_Arrow
	__Link__000001_From.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar
	__Link__000002_GroupLanes.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Group_and_NewDiagram_Lane
	__Link__000003_To.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Arrow_and_NewDiagram_Bar
}


