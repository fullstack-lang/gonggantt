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
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage()

	// Declarations of staged instances of Classshape
	__Classshape__000000_Classshape0000 := (&models.Classshape{Name: `Classshape0000`}).Stage()
	__Classshape__000001_Classshape0001 := (&models.Classshape{Name: `Classshape0001`}).Stage()
	__Classshape__000002_Classshape0002 := (&models.Classshape{Name: `Classshape0002`}).Stage()
	__Classshape__000003_Classshape0003 := (&models.Classshape{Name: `Classshape0003`}).Stage()
	__Classshape__000004_Classshape0004 := (&models.Classshape{Name: `Classshape0004`}).Stage()
	__Classshape__000005_Classshape0005 := (&models.Classshape{Name: `Classshape0005`}).Stage()
	__Classshape__000006_Classshape0006 := (&models.Classshape{Name: `Classshape0006`}).Stage()

	// Declarations of staged instances of DiagramPackage
	__DiagramPackage__000000_gonggantt_diagrams := (&models.DiagramPackage{Name: `gonggantt_diagrams`}).Stage()

	// Declarations of staged instances of Field
	__Field__000000_AlignOnStartEndOnYearStart := (&models.Field{Name: `AlignOnStartEndOnYearStart`}).Stage()
	__Field__000001_ArrowLengthToTheRightOfStartBar := (&models.Field{Name: `ArrowLengthToTheRightOfStartBar`}).Stage()
	__Field__000002_ArrowTipLenght := (&models.Field{Name: `ArrowTipLenght`}).Stage()
	__Field__000003_ComputedStart := (&models.Field{Name: `ComputedStart`}).Stage()
	__Field__000004_Date := (&models.Field{Name: `Date`}).Stage()
	__Field__000005_DateYOffset := (&models.Field{Name: `DateYOffset`}).Stage()
	__Field__000006_DisplayVerticalBar := (&models.Field{Name: `DisplayVerticalBar`}).Stage()
	__Field__000007_End := (&models.Field{Name: `End`}).Stage()
	__Field__000008_Group_Stroke := (&models.Field{Name: `Group_Stroke`}).Stage()
	__Field__000009_Group_StrokeDashArray := (&models.Field{Name: `Group_StrokeDashArray`}).Stage()
	__Field__000010_Group_StrokeWidth := (&models.Field{Name: `Group_StrokeWidth`}).Stage()
	__Field__000011_LaneHeight := (&models.Field{Name: `LaneHeight`}).Stage()
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000014_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000015_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000016_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000017_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000018_OptionnalColor := (&models.Field{Name: `OptionnalColor`}).Stage()
	__Field__000019_OptionnalColor := (&models.Field{Name: `OptionnalColor`}).Stage()
	__Field__000020_OptionnalStroke := (&models.Field{Name: `OptionnalStroke`}).Stage()
	__Field__000021_OptionnalStroke := (&models.Field{Name: `OptionnalStroke`}).Stage()
	__Field__000022_Order := (&models.Field{Name: `Order`}).Stage()
	__Field__000023_RatioBarToLaneHeight := (&models.Field{Name: `RatioBarToLaneHeight`}).Stage()
	__Field__000024_Start := (&models.Field{Name: `Start`}).Stage()
	__Field__000025_TextHeight := (&models.Field{Name: `TextHeight`}).Stage()
	__Field__000026_TimeLine_Color := (&models.Field{Name: `TimeLine_Color`}).Stage()
	__Field__000027_TimeLine_FillOpacity := (&models.Field{Name: `TimeLine_FillOpacity`}).Stage()
	__Field__000028_TimeLine_Stroke := (&models.Field{Name: `TimeLine_Stroke`}).Stage()
	__Field__000029_TimeLine_StrokeWidth := (&models.Field{Name: `TimeLine_StrokeWidth`}).Stage()
	__Field__000030_XLeftLanes := (&models.Field{Name: `XLeftLanes`}).Stage()
	__Field__000031_XRightMargin := (&models.Field{Name: `XRightMargin`}).Stage()
	__Field__000032_YTopMargin := (&models.Field{Name: `YTopMargin`}).Stage()

	// Declarations of staged instances of Link
	__Link__000000_Arrows := (&models.Link{Name: `Arrows`}).Stage()
	__Link__000001_Bars := (&models.Link{Name: `Bars`}).Stage()
	__Link__000002_From := (&models.Link{Name: `From`}).Stage()
	__Link__000003_GroupLanes := (&models.Link{Name: `GroupLanes`}).Stage()
	__Link__000004_Groups := (&models.Link{Name: `Groups`}).Stage()
	__Link__000005_Lane := (&models.Link{Name: `Lane`}).Stage()
	__Link__000006_Lanes := (&models.Link{Name: `Lanes`}).Stage()
	__Link__000007_LanesToDisplayMilestoneUse := (&models.Link{Name: `LanesToDisplayMilestoneUse`}).Stage()
	__Link__000008_Milestones := (&models.Link{Name: `Milestones`}).Stage()
	__Link__000009_To := (&models.Link{Name: `To`}).Stage()

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteLink

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of Position
	__Position__000000_Position_0000 := (&models.Position{Name: `Position-0000`}).Stage()
	__Position__000001_Position_0001 := (&models.Position{Name: `Position-0001`}).Stage()
	__Position__000002_Position_0002 := (&models.Position{Name: `Position-0002`}).Stage()
	__Position__000003_Position_0003 := (&models.Position{Name: `Position-0003`}).Stage()
	__Position__000004_Position_0004 := (&models.Position{Name: `Position-0004`}).Stage()
	__Position__000005_Position_0005 := (&models.Position{Name: `Position-0005`}).Stage()
	__Position__000006_Position_0006 := (&models.Position{Name: `Position-0006`}).Stage()

	// Declarations of staged instances of Reference
	__Reference__000000_Arrow := (&models.Reference{Name: `Arrow`}).Stage()
	__Reference__000001_Bar := (&models.Reference{Name: `Bar`}).Stage()
	__Reference__000002_Gantt := (&models.Reference{Name: `Gantt`}).Stage()
	__Reference__000003_Group := (&models.Reference{Name: `Group`}).Stage()
	__Reference__000004_Lane := (&models.Reference{Name: `Lane`}).Stage()
	__Reference__000005_LaneUse := (&models.Reference{Name: `LaneUse`}).Stage()
	__Reference__000006_Milestone := (&models.Reference{Name: `Milestone`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Vertice_0000 := (&models.Vertice{Name: `Vertice-0000`}).Stage()
	__Vertice__000001_Vertice_0001 := (&models.Vertice{Name: `Vertice-0001`}).Stage()
	__Vertice__000002_Vertice_0002 := (&models.Vertice{Name: `Vertice-0002`}).Stage()
	__Vertice__000003_Vertice_0003 := (&models.Vertice{Name: `Vertice-0003`}).Stage()
	__Vertice__000004_Vertice_0004 := (&models.Vertice{Name: `Vertice-0004`}).Stage()
	__Vertice__000005_Vertice_0005 := (&models.Vertice{Name: `Vertice-0005`}).Stage()
	__Vertice__000006_Vertice_0006 := (&models.Vertice{Name: `Vertice-0006`}).Stage()
	__Vertice__000007_Vertice_0007 := (&models.Vertice{Name: `Vertice-0007`}).Stage()
	__Vertice__000008_Vertice_0008 := (&models.Vertice{Name: `Vertice-0008`}).Stage()
	__Vertice__000009_Vertice_0009 := (&models.Vertice{Name: `Vertice-0009`}).Stage()

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	// Classshape values setup
	__Classshape__000000_Classshape0000.Name = `Classshape0000`
	__Classshape__000000_Classshape0000.ReferenceName = `Arrow`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow]
	__Classshape__000000_Classshape0000.Identifier = `ref_models.Arrow`
	__Classshape__000000_Classshape0000.ShowNbInstances = true
	__Classshape__000000_Classshape0000.NbInstances = 0
	__Classshape__000000_Classshape0000.Width = 240.000000
	__Classshape__000000_Classshape0000.Heigth = 93.000000
	__Classshape__000000_Classshape0000.IsSelected = false

	// Classshape values setup
	__Classshape__000001_Classshape0001.Name = `Classshape0001`
	__Classshape__000001_Classshape0001.ReferenceName = `Bar`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar]
	__Classshape__000001_Classshape0001.Identifier = `ref_models.Bar`
	__Classshape__000001_Classshape0001.ShowNbInstances = true
	__Classshape__000001_Classshape0001.NbInstances = 0
	__Classshape__000001_Classshape0001.Width = 240.000000
	__Classshape__000001_Classshape0001.Heigth = 123.000000
	__Classshape__000001_Classshape0001.IsSelected = false

	// Classshape values setup
	__Classshape__000002_Classshape0002.Name = `Classshape0002`
	__Classshape__000002_Classshape0002.ReferenceName = `Gantt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt]
	__Classshape__000002_Classshape0002.Identifier = `ref_models.Gantt`
	__Classshape__000002_Classshape0002.ShowNbInstances = true
	__Classshape__000002_Classshape0002.NbInstances = 0
	__Classshape__000002_Classshape0002.Width = 240.000000
	__Classshape__000002_Classshape0002.Heigth = 333.000000
	__Classshape__000002_Classshape0002.IsSelected = false

	// Classshape values setup
	__Classshape__000003_Classshape0003.Name = `Classshape0003`
	__Classshape__000003_Classshape0003.ReferenceName = `Group`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Group]
	__Classshape__000003_Classshape0003.Identifier = `ref_models.Group`
	__Classshape__000003_Classshape0003.ShowNbInstances = true
	__Classshape__000003_Classshape0003.NbInstances = 0
	__Classshape__000003_Classshape0003.Width = 240.000000
	__Classshape__000003_Classshape0003.Heigth = 63.000000
	__Classshape__000003_Classshape0003.IsSelected = false

	// Classshape values setup
	__Classshape__000004_Classshape0004.Name = `Classshape0004`
	__Classshape__000004_Classshape0004.ReferenceName = `Lane`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane]
	__Classshape__000004_Classshape0004.Identifier = `ref_models.Lane`
	__Classshape__000004_Classshape0004.ShowNbInstances = true
	__Classshape__000004_Classshape0004.NbInstances = 0
	__Classshape__000004_Classshape0004.Width = 240.000000
	__Classshape__000004_Classshape0004.Heigth = 78.000000
	__Classshape__000004_Classshape0004.IsSelected = false

	// Classshape values setup
	__Classshape__000005_Classshape0005.Name = `Classshape0005`
	__Classshape__000005_Classshape0005.ReferenceName = `LaneUse`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LaneUse]
	__Classshape__000005_Classshape0005.Identifier = `ref_models.LaneUse`
	__Classshape__000005_Classshape0005.ShowNbInstances = true
	__Classshape__000005_Classshape0005.NbInstances = 0
	__Classshape__000005_Classshape0005.Width = 240.000000
	__Classshape__000005_Classshape0005.Heigth = 48.000000
	__Classshape__000005_Classshape0005.IsSelected = false

	// Classshape values setup
	__Classshape__000006_Classshape0006.Name = `Classshape0006`
	__Classshape__000006_Classshape0006.ReferenceName = `Milestone`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Milestone]
	__Classshape__000006_Classshape0006.Identifier = `ref_models.Milestone`
	__Classshape__000006_Classshape0006.ShowNbInstances = true
	__Classshape__000006_Classshape0006.NbInstances = 0
	__Classshape__000006_Classshape0006.Width = 240.000000
	__Classshape__000006_Classshape0006.Heigth = 93.000000
	__Classshape__000006_Classshape0006.IsSelected = false

	// DiagramPackage values setup
	__DiagramPackage__000000_gonggantt_diagrams.Name = `gonggantt_diagrams`
	__DiagramPackage__000000_gonggantt_diagrams.Path = `github.com/fullstack-lang/gonggantt/go/models`
	__DiagramPackage__000000_gonggantt_diagrams.GongModelPath = `github.com/fullstack-lang/gonggantt/go/models`
	__DiagramPackage__000000_gonggantt_diagrams.IsEditable = true
	__DiagramPackage__000000_gonggantt_diagrams.IsReloaded = false
	__DiagramPackage__000000_gonggantt_diagrams.AbsolutePathToDiagramPackage = `/Users/thomaspeugeot/go/src/github.com/fullstack-lang/gonggantt/go/diagrams`

	// Field values setup
	__Field__000000_AlignOnStartEndOnYearStart.Name = `AlignOnStartEndOnYearStart`
	__Field__000000_AlignOnStartEndOnYearStart.Fieldname = `AlignOnStartEndOnYearStart`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.AlignOnStartEndOnYearStart]
	__Field__000000_AlignOnStartEndOnYearStart.Identifier = `ref_models.Gantt.AlignOnStartEndOnYearStart`
	__Field__000000_AlignOnStartEndOnYearStart.FieldTypeAsString = ``
	__Field__000000_AlignOnStartEndOnYearStart.Structname = `Gantt`
	__Field__000000_AlignOnStartEndOnYearStart.Fieldtypename = `bool`

	// Field values setup
	__Field__000001_ArrowLengthToTheRightOfStartBar.Name = `ArrowLengthToTheRightOfStartBar`
	__Field__000001_ArrowLengthToTheRightOfStartBar.Fieldname = `ArrowLengthToTheRightOfStartBar`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.ArrowLengthToTheRightOfStartBar]
	__Field__000001_ArrowLengthToTheRightOfStartBar.Identifier = `ref_models.Gantt.ArrowLengthToTheRightOfStartBar`
	__Field__000001_ArrowLengthToTheRightOfStartBar.FieldTypeAsString = ``
	__Field__000001_ArrowLengthToTheRightOfStartBar.Structname = `Gantt`
	__Field__000001_ArrowLengthToTheRightOfStartBar.Fieldtypename = `float64`

	// Field values setup
	__Field__000002_ArrowTipLenght.Name = `ArrowTipLenght`
	__Field__000002_ArrowTipLenght.Fieldname = `ArrowTipLenght`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.ArrowTipLenght]
	__Field__000002_ArrowTipLenght.Identifier = `ref_models.Gantt.ArrowTipLenght`
	__Field__000002_ArrowTipLenght.FieldTypeAsString = ``
	__Field__000002_ArrowTipLenght.Structname = `Gantt`
	__Field__000002_ArrowTipLenght.Fieldtypename = `float64`

	// Field values setup
	__Field__000003_ComputedStart.Name = `ComputedStart`
	__Field__000003_ComputedStart.Fieldname = `ComputedStart`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.ComputedStart]
	__Field__000003_ComputedStart.Identifier = `ref_models.Gantt.ComputedStart`
	__Field__000003_ComputedStart.FieldTypeAsString = ``
	__Field__000003_ComputedStart.Structname = `Gantt`
	__Field__000003_ComputedStart.Fieldtypename = `Time`

	// Field values setup
	__Field__000004_Date.Name = `Date`
	__Field__000004_Date.Fieldname = `Date`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Milestone.Date]
	__Field__000004_Date.Identifier = `ref_models.Milestone.Date`
	__Field__000004_Date.FieldTypeAsString = ``
	__Field__000004_Date.Structname = `Milestone`
	__Field__000004_Date.Fieldtypename = `Time`

	// Field values setup
	__Field__000005_DateYOffset.Name = `DateYOffset`
	__Field__000005_DateYOffset.Fieldname = `DateYOffset`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.DateYOffset]
	__Field__000005_DateYOffset.Identifier = `ref_models.Gantt.DateYOffset`
	__Field__000005_DateYOffset.FieldTypeAsString = ``
	__Field__000005_DateYOffset.Structname = `Gantt`
	__Field__000005_DateYOffset.Fieldtypename = `float64`

	// Field values setup
	__Field__000006_DisplayVerticalBar.Name = `DisplayVerticalBar`
	__Field__000006_DisplayVerticalBar.Fieldname = `DisplayVerticalBar`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Milestone.DisplayVerticalBar]
	__Field__000006_DisplayVerticalBar.Identifier = `ref_models.Milestone.DisplayVerticalBar`
	__Field__000006_DisplayVerticalBar.FieldTypeAsString = ``
	__Field__000006_DisplayVerticalBar.Structname = `Milestone`
	__Field__000006_DisplayVerticalBar.Fieldtypename = `bool`

	// Field values setup
	__Field__000007_End.Name = `End`
	__Field__000007_End.Fieldname = `End`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.End]
	__Field__000007_End.Identifier = `ref_models.Bar.End`
	__Field__000007_End.FieldTypeAsString = ``
	__Field__000007_End.Structname = `Bar`
	__Field__000007_End.Fieldtypename = `Time`

	// Field values setup
	__Field__000008_Group_Stroke.Name = `Group_Stroke`
	__Field__000008_Group_Stroke.Fieldname = `Group_Stroke`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Group_Stroke]
	__Field__000008_Group_Stroke.Identifier = `ref_models.Gantt.Group_Stroke`
	__Field__000008_Group_Stroke.FieldTypeAsString = ``
	__Field__000008_Group_Stroke.Structname = `Gantt`
	__Field__000008_Group_Stroke.Fieldtypename = `string`

	// Field values setup
	__Field__000009_Group_StrokeDashArray.Name = `Group_StrokeDashArray`
	__Field__000009_Group_StrokeDashArray.Fieldname = `Group_StrokeDashArray`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Group_StrokeDashArray]
	__Field__000009_Group_StrokeDashArray.Identifier = `ref_models.Gantt.Group_StrokeDashArray`
	__Field__000009_Group_StrokeDashArray.FieldTypeAsString = ``
	__Field__000009_Group_StrokeDashArray.Structname = `Gantt`
	__Field__000009_Group_StrokeDashArray.Fieldtypename = `string`

	// Field values setup
	__Field__000010_Group_StrokeWidth.Name = `Group_StrokeWidth`
	__Field__000010_Group_StrokeWidth.Fieldname = `Group_StrokeWidth`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Group_StrokeWidth]
	__Field__000010_Group_StrokeWidth.Identifier = `ref_models.Gantt.Group_StrokeWidth`
	__Field__000010_Group_StrokeWidth.FieldTypeAsString = ``
	__Field__000010_Group_StrokeWidth.Structname = `Gantt`
	__Field__000010_Group_StrokeWidth.Fieldtypename = `float64`

	// Field values setup
	__Field__000011_LaneHeight.Name = `LaneHeight`
	__Field__000011_LaneHeight.Fieldname = `LaneHeight`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.LaneHeight]
	__Field__000011_LaneHeight.Identifier = `ref_models.Gantt.LaneHeight`
	__Field__000011_LaneHeight.FieldTypeAsString = ``
	__Field__000011_LaneHeight.Structname = `Gantt`
	__Field__000011_LaneHeight.Fieldtypename = `float64`

	// Field values setup
	__Field__000012_Name.Name = `Name`
	__Field__000012_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Milestone.Name]
	__Field__000012_Name.Identifier = `ref_models.Milestone.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `Milestone`
	__Field__000012_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000013_Name.Name = `Name`
	__Field__000013_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Name]
	__Field__000013_Name.Identifier = `ref_models.Gantt.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Gantt`
	__Field__000013_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000014_Name.Name = `Name`
	__Field__000014_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.Name]
	__Field__000014_Name.Identifier = `ref_models.Bar.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `Bar`
	__Field__000014_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000015_Name.Name = `Name`
	__Field__000015_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.Name]
	__Field__000015_Name.Identifier = `ref_models.Arrow.Name`
	__Field__000015_Name.FieldTypeAsString = ``
	__Field__000015_Name.Structname = `Arrow`
	__Field__000015_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000016_Name.Name = `Name`
	__Field__000016_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Group.Name]
	__Field__000016_Name.Identifier = `ref_models.Group.Name`
	__Field__000016_Name.FieldTypeAsString = ``
	__Field__000016_Name.Structname = `Group`
	__Field__000016_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000017_Name.Name = `Name`
	__Field__000017_Name.Fieldname = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane.Name]
	__Field__000017_Name.Identifier = `ref_models.Lane.Name`
	__Field__000017_Name.FieldTypeAsString = ``
	__Field__000017_Name.Structname = `Lane`
	__Field__000017_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000018_OptionnalColor.Name = `OptionnalColor`
	__Field__000018_OptionnalColor.Fieldname = `OptionnalColor`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.OptionnalColor]
	__Field__000018_OptionnalColor.Identifier = `ref_models.Arrow.OptionnalColor`
	__Field__000018_OptionnalColor.FieldTypeAsString = ``
	__Field__000018_OptionnalColor.Structname = `Arrow`
	__Field__000018_OptionnalColor.Fieldtypename = `string`

	// Field values setup
	__Field__000019_OptionnalColor.Name = `OptionnalColor`
	__Field__000019_OptionnalColor.Fieldname = `OptionnalColor`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.OptionnalColor]
	__Field__000019_OptionnalColor.Identifier = `ref_models.Bar.OptionnalColor`
	__Field__000019_OptionnalColor.FieldTypeAsString = ``
	__Field__000019_OptionnalColor.Structname = `Bar`
	__Field__000019_OptionnalColor.Fieldtypename = `string`

	// Field values setup
	__Field__000020_OptionnalStroke.Name = `OptionnalStroke`
	__Field__000020_OptionnalStroke.Fieldname = `OptionnalStroke`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.OptionnalStroke]
	__Field__000020_OptionnalStroke.Identifier = `ref_models.Arrow.OptionnalStroke`
	__Field__000020_OptionnalStroke.FieldTypeAsString = ``
	__Field__000020_OptionnalStroke.Structname = `Arrow`
	__Field__000020_OptionnalStroke.Fieldtypename = `string`

	// Field values setup
	__Field__000021_OptionnalStroke.Name = `OptionnalStroke`
	__Field__000021_OptionnalStroke.Fieldname = `OptionnalStroke`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.OptionnalStroke]
	__Field__000021_OptionnalStroke.Identifier = `ref_models.Bar.OptionnalStroke`
	__Field__000021_OptionnalStroke.FieldTypeAsString = ``
	__Field__000021_OptionnalStroke.Structname = `Bar`
	__Field__000021_OptionnalStroke.Fieldtypename = `string`

	// Field values setup
	__Field__000022_Order.Name = `Order`
	__Field__000022_Order.Fieldname = `Order`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane.Order]
	__Field__000022_Order.Identifier = `ref_models.Lane.Order`
	__Field__000022_Order.FieldTypeAsString = ``
	__Field__000022_Order.Structname = `Lane`
	__Field__000022_Order.Fieldtypename = `int`

	// Field values setup
	__Field__000023_RatioBarToLaneHeight.Name = `RatioBarToLaneHeight`
	__Field__000023_RatioBarToLaneHeight.Fieldname = `RatioBarToLaneHeight`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.RatioBarToLaneHeight]
	__Field__000023_RatioBarToLaneHeight.Identifier = `ref_models.Gantt.RatioBarToLaneHeight`
	__Field__000023_RatioBarToLaneHeight.FieldTypeAsString = ``
	__Field__000023_RatioBarToLaneHeight.Structname = `Gantt`
	__Field__000023_RatioBarToLaneHeight.Fieldtypename = `float64`

	// Field values setup
	__Field__000024_Start.Name = `Start`
	__Field__000024_Start.Fieldname = `Start`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Bar.Start]
	__Field__000024_Start.Identifier = `ref_models.Bar.Start`
	__Field__000024_Start.FieldTypeAsString = ``
	__Field__000024_Start.Structname = `Bar`
	__Field__000024_Start.Fieldtypename = `Time`

	// Field values setup
	__Field__000025_TextHeight.Name = `TextHeight`
	__Field__000025_TextHeight.Fieldname = `TextHeight`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.TextHeight]
	__Field__000025_TextHeight.Identifier = `ref_models.Gantt.TextHeight`
	__Field__000025_TextHeight.FieldTypeAsString = ``
	__Field__000025_TextHeight.Structname = `Gantt`
	__Field__000025_TextHeight.Fieldtypename = `float64`

	// Field values setup
	__Field__000026_TimeLine_Color.Name = `TimeLine_Color`
	__Field__000026_TimeLine_Color.Fieldname = `TimeLine_Color`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.TimeLine_Color]
	__Field__000026_TimeLine_Color.Identifier = `ref_models.Gantt.TimeLine_Color`
	__Field__000026_TimeLine_Color.FieldTypeAsString = ``
	__Field__000026_TimeLine_Color.Structname = `Gantt`
	__Field__000026_TimeLine_Color.Fieldtypename = `string`

	// Field values setup
	__Field__000027_TimeLine_FillOpacity.Name = `TimeLine_FillOpacity`
	__Field__000027_TimeLine_FillOpacity.Fieldname = `TimeLine_FillOpacity`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.TimeLine_FillOpacity]
	__Field__000027_TimeLine_FillOpacity.Identifier = `ref_models.Gantt.TimeLine_FillOpacity`
	__Field__000027_TimeLine_FillOpacity.FieldTypeAsString = ``
	__Field__000027_TimeLine_FillOpacity.Structname = `Gantt`
	__Field__000027_TimeLine_FillOpacity.Fieldtypename = `float64`

	// Field values setup
	__Field__000028_TimeLine_Stroke.Name = `TimeLine_Stroke`
	__Field__000028_TimeLine_Stroke.Fieldname = `TimeLine_Stroke`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.TimeLine_Stroke]
	__Field__000028_TimeLine_Stroke.Identifier = `ref_models.Gantt.TimeLine_Stroke`
	__Field__000028_TimeLine_Stroke.FieldTypeAsString = ``
	__Field__000028_TimeLine_Stroke.Structname = `Gantt`
	__Field__000028_TimeLine_Stroke.Fieldtypename = `string`

	// Field values setup
	__Field__000029_TimeLine_StrokeWidth.Name = `TimeLine_StrokeWidth`
	__Field__000029_TimeLine_StrokeWidth.Fieldname = `TimeLine_StrokeWidth`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.TimeLine_StrokeWidth]
	__Field__000029_TimeLine_StrokeWidth.Identifier = `ref_models.Gantt.TimeLine_StrokeWidth`
	__Field__000029_TimeLine_StrokeWidth.FieldTypeAsString = ``
	__Field__000029_TimeLine_StrokeWidth.Structname = `Gantt`
	__Field__000029_TimeLine_StrokeWidth.Fieldtypename = `float64`

	// Field values setup
	__Field__000030_XLeftLanes.Name = `XLeftLanes`
	__Field__000030_XLeftLanes.Fieldname = `XLeftLanes`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.XLeftLanes]
	__Field__000030_XLeftLanes.Identifier = `ref_models.Gantt.XLeftLanes`
	__Field__000030_XLeftLanes.FieldTypeAsString = ``
	__Field__000030_XLeftLanes.Structname = `Gantt`
	__Field__000030_XLeftLanes.Fieldtypename = `float64`

	// Field values setup
	__Field__000031_XRightMargin.Name = `XRightMargin`
	__Field__000031_XRightMargin.Fieldname = `XRightMargin`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.XRightMargin]
	__Field__000031_XRightMargin.Identifier = `ref_models.Gantt.XRightMargin`
	__Field__000031_XRightMargin.FieldTypeAsString = ``
	__Field__000031_XRightMargin.Structname = `Gantt`
	__Field__000031_XRightMargin.Fieldtypename = `float64`

	// Field values setup
	__Field__000032_YTopMargin.Name = `YTopMargin`
	__Field__000032_YTopMargin.Fieldname = `YTopMargin`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.YTopMargin]
	__Field__000032_YTopMargin.Identifier = `ref_models.Gantt.YTopMargin`
	__Field__000032_YTopMargin.FieldTypeAsString = ``
	__Field__000032_YTopMargin.Structname = `Gantt`
	__Field__000032_YTopMargin.Fieldtypename = `float64`

	// Link values setup
	__Link__000000_Arrows.Name = `Arrows`
	__Link__000000_Arrows.Fieldname = `Arrows`
	__Link__000000_Arrows.Structname = `Gantt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Arrows]
	__Link__000000_Arrows.Identifier = `ref_models.Gantt.Arrows`
	__Link__000000_Arrows.Fieldtypename = `Arrow`
	__Link__000000_Arrows.TargetMultiplicity = models.MANY
	__Link__000000_Arrows.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000001_Bars.Name = `Bars`
	__Link__000001_Bars.Fieldname = `Bars`
	__Link__000001_Bars.Structname = `Lane`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Lane.Bars]
	__Link__000001_Bars.Identifier = `ref_models.Lane.Bars`
	__Link__000001_Bars.Fieldtypename = `Bar`
	__Link__000001_Bars.TargetMultiplicity = models.MANY
	__Link__000001_Bars.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000002_From.Name = `From`
	__Link__000002_From.Fieldname = `From`
	__Link__000002_From.Structname = `Arrow`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.From]
	__Link__000002_From.Identifier = `ref_models.Arrow.From`
	__Link__000002_From.Fieldtypename = `Bar`
	__Link__000002_From.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_From.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000003_GroupLanes.Name = `GroupLanes`
	__Link__000003_GroupLanes.Fieldname = `GroupLanes`
	__Link__000003_GroupLanes.Structname = `Group`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Group.GroupLanes]
	__Link__000003_GroupLanes.Identifier = `ref_models.Group.GroupLanes`
	__Link__000003_GroupLanes.Fieldtypename = `Lane`
	__Link__000003_GroupLanes.TargetMultiplicity = models.MANY
	__Link__000003_GroupLanes.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000004_Groups.Name = `Groups`
	__Link__000004_Groups.Fieldname = `Groups`
	__Link__000004_Groups.Structname = `Gantt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Groups]
	__Link__000004_Groups.Identifier = `ref_models.Gantt.Groups`
	__Link__000004_Groups.Fieldtypename = `Group`
	__Link__000004_Groups.TargetMultiplicity = models.MANY
	__Link__000004_Groups.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000005_Lane.Name = `Lane`
	__Link__000005_Lane.Fieldname = `Lane`
	__Link__000005_Lane.Structname = `LaneUse`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.LaneUse.Lane]
	__Link__000005_Lane.Identifier = `ref_models.LaneUse.Lane`
	__Link__000005_Lane.Fieldtypename = `Lane`
	__Link__000005_Lane.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_Lane.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000006_Lanes.Name = `Lanes`
	__Link__000006_Lanes.Fieldname = `Lanes`
	__Link__000006_Lanes.Structname = `Gantt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Lanes]
	__Link__000006_Lanes.Identifier = `ref_models.Gantt.Lanes`
	__Link__000006_Lanes.Fieldtypename = `Lane`
	__Link__000006_Lanes.TargetMultiplicity = models.MANY
	__Link__000006_Lanes.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000007_LanesToDisplayMilestoneUse.Name = `LanesToDisplayMilestoneUse`
	__Link__000007_LanesToDisplayMilestoneUse.Fieldname = `LanesToDisplayMilestoneUse`
	__Link__000007_LanesToDisplayMilestoneUse.Structname = `Milestone`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Milestone.LanesToDisplayMilestoneUse]
	__Link__000007_LanesToDisplayMilestoneUse.Identifier = `ref_models.Milestone.LanesToDisplayMilestoneUse`
	__Link__000007_LanesToDisplayMilestoneUse.Fieldtypename = `LaneUse`
	__Link__000007_LanesToDisplayMilestoneUse.TargetMultiplicity = models.MANY
	__Link__000007_LanesToDisplayMilestoneUse.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000008_Milestones.Name = `Milestones`
	__Link__000008_Milestones.Fieldname = `Milestones`
	__Link__000008_Milestones.Structname = `Gantt`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Gantt.Milestones]
	__Link__000008_Milestones.Identifier = `ref_models.Gantt.Milestones`
	__Link__000008_Milestones.Fieldtypename = `Milestone`
	__Link__000008_Milestones.TargetMultiplicity = models.MANY
	__Link__000008_Milestones.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000009_To.Name = `To`
	__Link__000009_To.Fieldname = `To`
	__Link__000009_To.Structname = `Arrow`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Arrow.To]
	__Link__000009_To.Identifier = `ref_models.Arrow.To`
	__Link__000009_To.Fieldtypename = `Bar`
	__Link__000009_To.TargetMultiplicity = models.ZERO_ONE
	__Link__000009_To.SourceMultiplicity = models.MANY

	// Position values setup
	__Position__000000_Position_0000.X = 570.000000
	__Position__000000_Position_0000.Y = 280.000000
	__Position__000000_Position_0000.Name = `Position-0000`

	// Position values setup
	__Position__000001_Position_0001.X = 950.000000
	__Position__000001_Position_0001.Y = 140.000000
	__Position__000001_Position_0001.Name = `Position-0001`

	// Position values setup
	__Position__000002_Position_0002.X = 60.000000
	__Position__000002_Position_0002.Y = 40.000000
	__Position__000002_Position_0002.Name = `Position-0002`

	// Position values setup
	__Position__000003_Position_0003.X = 570.000000
	__Position__000003_Position_0003.Y = 50.000000
	__Position__000003_Position_0003.Name = `Position-0003`

	// Position values setup
	__Position__000004_Position_0004.X = 570.000000
	__Position__000004_Position_0004.Y = 140.000000
	__Position__000004_Position_0004.Name = `Position-0004`

	// Position values setup
	__Position__000005_Position_0005.X = 930.000000
	__Position__000005_Position_0005.Y = 520.000000
	__Position__000005_Position_0005.Name = `Position-0005`

	// Position values setup
	__Position__000006_Position_0006.X = 570.000000
	__Position__000006_Position_0006.Y = 400.000000
	__Position__000006_Position_0006.Name = `Position-0006`

	// Reference values setup
	__Reference__000000_Arrow.Name = `Arrow`
	__Reference__000000_Arrow.NbInstances = 0
	__Reference__000000_Arrow.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000001_Bar.Name = `Bar`
	__Reference__000001_Bar.NbInstances = 0
	__Reference__000001_Bar.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000002_Gantt.Name = `Gantt`
	__Reference__000002_Gantt.NbInstances = 0
	__Reference__000002_Gantt.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000003_Group.Name = `Group`
	__Reference__000003_Group.NbInstances = 0
	__Reference__000003_Group.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000004_Lane.Name = `Lane`
	__Reference__000004_Lane.NbInstances = 0
	__Reference__000004_Lane.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000005_LaneUse.Name = `LaneUse`
	__Reference__000005_LaneUse.NbInstances = 0
	__Reference__000005_LaneUse.Type = models.REFERENCE_GONG_STRUCT

	// Reference values setup
	__Reference__000006_Milestone.Name = `Milestone`
	__Reference__000006_Milestone.NbInstances = 0
	__Reference__000006_Milestone.Type = models.REFERENCE_GONG_STRUCT

	// Vertice values setup
	__Vertice__000000_Vertice_0000.X = 1010.000000
	__Vertice__000000_Vertice_0000.Y = 386.500000
	__Vertice__000000_Vertice_0000.Name = `Vertice-0000`

	// Vertice values setup
	__Vertice__000001_Vertice_0001.X = 1190.000000
	__Vertice__000001_Vertice_0001.Y = 356.500000
	__Vertice__000001_Vertice_0001.Name = `Vertice-0001`

	// Vertice values setup
	__Vertice__000002_Vertice_0002.X = 445.000000
	__Vertice__000002_Vertice_0002.Y = 261.500000
	__Vertice__000002_Vertice_0002.Name = `Vertice-0002`

	// Vertice values setup
	__Vertice__000003_Vertice_0003.X = 428.500000
	__Vertice__000003_Vertice_0003.Y = 107.000000
	__Vertice__000003_Vertice_0003.Name = `Vertice-0003`

	// Vertice values setup
	__Vertice__000004_Vertice_0004.X = 411.500000
	__Vertice__000004_Vertice_0004.Y = 175.000000
	__Vertice__000004_Vertice_0004.Name = `Vertice-0004`

	// Vertice values setup
	__Vertice__000005_Vertice_0005.X = 435.000000
	__Vertice__000005_Vertice_0005.Y = 361.500000
	__Vertice__000005_Vertice_0005.Name = `Vertice-0005`

	// Vertice values setup
	__Vertice__000006_Vertice_0006.X = 1080.000000
	__Vertice__000006_Vertice_0006.Y = 76.500000
	__Vertice__000006_Vertice_0006.Name = `Vertice-0006`

	// Vertice values setup
	__Vertice__000007_Vertice_0007.X = 870.000000
	__Vertice__000007_Vertice_0007.Y = 171.500000
	__Vertice__000007_Vertice_0007.Name = `Vertice-0007`

	// Vertice values setup
	__Vertice__000008_Vertice_0008.X = 885.000000
	__Vertice__000008_Vertice_0008.Y = 306.500000
	__Vertice__000008_Vertice_0008.Name = `Vertice-0008`

	// Vertice values setup
	__Vertice__000009_Vertice_0009.X = 510.000000
	__Vertice__000009_Vertice_0009.Y = 559.000000
	__Vertice__000009_Vertice_0009.Name = `Vertice-0009`

	// Setup of pointers
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000000_Classshape0000)
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000001_Classshape0001)
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000002_Classshape0002)
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000003_Classshape0003)
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000004_Classshape0004)
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000005_Classshape0005)
	__Classdiagram__000000_Default.Classshapes = append(__Classdiagram__000000_Default.Classshapes, __Classshape__000006_Classshape0006)
	__Classshape__000000_Classshape0000.Position = __Position__000000_Position_0000
	__Classshape__000000_Classshape0000.Reference = __Reference__000000_Arrow
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000015_Name)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000018_OptionnalColor)
	__Classshape__000000_Classshape0000.Fields = append(__Classshape__000000_Classshape0000.Fields, __Field__000020_OptionnalStroke)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000002_From)
	__Classshape__000000_Classshape0000.Links = append(__Classshape__000000_Classshape0000.Links, __Link__000009_To)
	__Classshape__000001_Classshape0001.Position = __Position__000001_Position_0001
	__Classshape__000001_Classshape0001.Reference = __Reference__000001_Bar
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000007_End)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000014_Name)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000019_OptionnalColor)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000021_OptionnalStroke)
	__Classshape__000001_Classshape0001.Fields = append(__Classshape__000001_Classshape0001.Fields, __Field__000024_Start)
	__Classshape__000002_Classshape0002.Position = __Position__000002_Position_0002
	__Classshape__000002_Classshape0002.Reference = __Reference__000002_Gantt
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000000_AlignOnStartEndOnYearStart)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000001_ArrowLengthToTheRightOfStartBar)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000002_ArrowTipLenght)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000003_ComputedStart)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000005_DateYOffset)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000008_Group_Stroke)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000009_Group_StrokeDashArray)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000010_Group_StrokeWidth)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000011_LaneHeight)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000013_Name)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000023_RatioBarToLaneHeight)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000025_TextHeight)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000026_TimeLine_Color)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000027_TimeLine_FillOpacity)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000028_TimeLine_Stroke)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000029_TimeLine_StrokeWidth)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000030_XLeftLanes)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000031_XRightMargin)
	__Classshape__000002_Classshape0002.Fields = append(__Classshape__000002_Classshape0002.Fields, __Field__000032_YTopMargin)
	__Classshape__000002_Classshape0002.Links = append(__Classshape__000002_Classshape0002.Links, __Link__000000_Arrows)
	__Classshape__000002_Classshape0002.Links = append(__Classshape__000002_Classshape0002.Links, __Link__000004_Groups)
	__Classshape__000002_Classshape0002.Links = append(__Classshape__000002_Classshape0002.Links, __Link__000006_Lanes)
	__Classshape__000002_Classshape0002.Links = append(__Classshape__000002_Classshape0002.Links, __Link__000008_Milestones)
	__Classshape__000003_Classshape0003.Position = __Position__000003_Position_0003
	__Classshape__000003_Classshape0003.Reference = __Reference__000003_Group
	__Classshape__000003_Classshape0003.Fields = append(__Classshape__000003_Classshape0003.Fields, __Field__000016_Name)
	__Classshape__000003_Classshape0003.Links = append(__Classshape__000003_Classshape0003.Links, __Link__000003_GroupLanes)
	__Classshape__000004_Classshape0004.Position = __Position__000004_Position_0004
	__Classshape__000004_Classshape0004.Reference = __Reference__000004_Lane
	__Classshape__000004_Classshape0004.Fields = append(__Classshape__000004_Classshape0004.Fields, __Field__000017_Name)
	__Classshape__000004_Classshape0004.Fields = append(__Classshape__000004_Classshape0004.Fields, __Field__000022_Order)
	__Classshape__000004_Classshape0004.Links = append(__Classshape__000004_Classshape0004.Links, __Link__000001_Bars)
	__Classshape__000005_Classshape0005.Position = __Position__000005_Position_0005
	__Classshape__000005_Classshape0005.Reference = __Reference__000005_LaneUse
	__Classshape__000005_Classshape0005.Links = append(__Classshape__000005_Classshape0005.Links, __Link__000005_Lane)
	__Classshape__000006_Classshape0006.Position = __Position__000006_Position_0006
	__Classshape__000006_Classshape0006.Reference = __Reference__000006_Milestone
	__Classshape__000006_Classshape0006.Fields = append(__Classshape__000006_Classshape0006.Fields, __Field__000004_Date)
	__Classshape__000006_Classshape0006.Fields = append(__Classshape__000006_Classshape0006.Fields, __Field__000006_DisplayVerticalBar)
	__Classshape__000006_Classshape0006.Fields = append(__Classshape__000006_Classshape0006.Fields, __Field__000012_Name)
	__Classshape__000006_Classshape0006.Links = append(__Classshape__000006_Classshape0006.Links, __Link__000007_LanesToDisplayMilestoneUse)
	__DiagramPackage__000000_gonggantt_diagrams.Classdiagrams = append(__DiagramPackage__000000_gonggantt_diagrams.Classdiagrams, __Classdiagram__000000_Default)
	__Link__000000_Arrows.Middlevertice = __Vertice__000002_Vertice_0002
	__Link__000001_Bars.Middlevertice = __Vertice__000007_Vertice_0007
	__Link__000002_From.Middlevertice = __Vertice__000000_Vertice_0000
	__Link__000003_GroupLanes.Middlevertice = __Vertice__000006_Vertice_0006
	__Link__000004_Groups.Middlevertice = __Vertice__000003_Vertice_0003
	__Link__000005_Lane.Middlevertice = __Vertice__000008_Vertice_0008
	__Link__000006_Lanes.Middlevertice = __Vertice__000004_Vertice_0004
	__Link__000007_LanesToDisplayMilestoneUse.Middlevertice = __Vertice__000009_Vertice_0009
	__Link__000008_Milestones.Middlevertice = __Vertice__000005_Vertice_0005
	__Link__000009_To.Middlevertice = __Vertice__000001_Vertice_0001
}


