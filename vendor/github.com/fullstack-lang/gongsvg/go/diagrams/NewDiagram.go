package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsvg/go/models"
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

	"ref_models.Animate": &(ref_models.Animate{}),

	"ref_models.Circle": &(ref_models.Circle{}),

	"ref_models.Circle.Animations": (ref_models.Circle{}).Animations,

	"ref_models.Circle.CX": (ref_models.Circle{}).CX,

	"ref_models.Circle.CY": (ref_models.Circle{}).CY,

	"ref_models.Circle.Name": (ref_models.Circle{}).Name,

	"ref_models.Circle.Radius": (ref_models.Circle{}).Radius,

	"ref_models.Ellipse": &(ref_models.Ellipse{}),

	"ref_models.Layer": &(ref_models.Layer{}),

	"ref_models.Layer.Circles": (ref_models.Layer{}).Circles,

	"ref_models.Layer.Display": (ref_models.Layer{}).Display,

	"ref_models.Layer.Ellipses": (ref_models.Layer{}).Ellipses,

	"ref_models.Layer.Lines": (ref_models.Layer{}).Lines,

	"ref_models.Layer.Name": (ref_models.Layer{}).Name,

	"ref_models.Layer.Paths": (ref_models.Layer{}).Paths,

	"ref_models.Layer.Polygones": (ref_models.Layer{}).Polygones,

	"ref_models.Layer.Polylines": (ref_models.Layer{}).Polylines,

	"ref_models.Layer.Rects": (ref_models.Layer{}).Rects,

	"ref_models.Layer.Texts": (ref_models.Layer{}).Texts,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Path": &(ref_models.Path{}),

	"ref_models.Polygone": &(ref_models.Polygone{}),

	"ref_models.Polyline": &(ref_models.Polyline{}),

	"ref_models.Rect": &(ref_models.Rect{}),

	"ref_models.SVG": &(ref_models.SVG{}),

	"ref_models.SVG.Layers": (ref_models.SVG{}).Layers,

	"ref_models.SVG.Name": (ref_models.SVG{}).Name,

	"ref_models.Text": &(ref_models.Text{}),
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_CX := (&models.Field{Name: `CX`}).Stage(stage)
	__Field__000001_CY := (&models.Field{Name: `CY`}).Stage(stage)
	__Field__000002_Display := (&models.Field{Name: `Display`}).Stage(stage)
	__Field__000003_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000005_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000006_Radius := (&models.Field{Name: `Radius`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Animate := (&models.GongStructShape{Name: `NewDiagram-Animate`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Circle := (&models.GongStructShape{Name: `NewDiagram-Circle`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_Ellipse := (&models.GongStructShape{Name: `NewDiagram-Ellipse`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_Line := (&models.GongStructShape{Name: `NewDiagram-Line`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_Path := (&models.GongStructShape{Name: `NewDiagram-Path`}).Stage(stage)
	__GongStructShape__000005_NewDiagram_Polygone := (&models.GongStructShape{Name: `NewDiagram-Polygone`}).Stage(stage)
	__GongStructShape__000006_NewDiagram_Polyline := (&models.GongStructShape{Name: `NewDiagram-Polyline`}).Stage(stage)
	__GongStructShape__000007_NewDiagram_Rect := (&models.GongStructShape{Name: `NewDiagram-Rect`}).Stage(stage)
	__GongStructShape__000008_NewDiagram_SVG := (&models.GongStructShape{Name: `NewDiagram-SVG`}).Stage(stage)
	__GongStructShape__000009_NewDiagram_SVG := (&models.GongStructShape{Name: `NewDiagram-SVG`}).Stage(stage)
	__GongStructShape__000010_NewDiagram_Text := (&models.GongStructShape{Name: `NewDiagram-Text`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Animations := (&models.Link{Name: `Animations`}).Stage(stage)
	__Link__000001_Circles := (&models.Link{Name: `Circles`}).Stage(stage)
	__Link__000002_Ellipses := (&models.Link{Name: `Ellipses`}).Stage(stage)
	__Link__000003_Layers := (&models.Link{Name: `Layers`}).Stage(stage)
	__Link__000004_Lines := (&models.Link{Name: `Lines`}).Stage(stage)
	__Link__000005_Paths := (&models.Link{Name: `Paths`}).Stage(stage)
	__Link__000006_Polygones := (&models.Link{Name: `Polygones`}).Stage(stage)
	__Link__000007_Polylines := (&models.Link{Name: `Polylines`}).Stage(stage)
	__Link__000008_Rects := (&models.Link{Name: `Rects`}).Stage(stage)
	__Link__000009_Texts := (&models.Link{Name: `Texts`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_Animate := (&models.Position{Name: `Pos-NewDiagram-Animate`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Circle := (&models.Position{Name: `Pos-NewDiagram-Circle`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Ellipse := (&models.Position{Name: `Pos-NewDiagram-Ellipse`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_Line := (&models.Position{Name: `Pos-NewDiagram-Line`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_Path := (&models.Position{Name: `Pos-NewDiagram-Path`}).Stage(stage)
	__Position__000005_Pos_NewDiagram_Polygone := (&models.Position{Name: `Pos-NewDiagram-Polygone`}).Stage(stage)
	__Position__000006_Pos_NewDiagram_Polyline := (&models.Position{Name: `Pos-NewDiagram-Polyline`}).Stage(stage)
	__Position__000007_Pos_NewDiagram_Rect := (&models.Position{Name: `Pos-NewDiagram-Rect`}).Stage(stage)
	__Position__000008_Pos_NewDiagram_SVG := (&models.Position{Name: `Pos-NewDiagram-SVG`}).Stage(stage)
	__Position__000009_Pos_NewDiagram_SVG := (&models.Position{Name: `Pos-NewDiagram-SVG`}).Stage(stage)
	__Position__000010_Pos_NewDiagram_Text := (&models.Position{Name: `Pos-NewDiagram-Text`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_Animate := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Circle and NewDiagram-Animate`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Circle := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Circle`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Ellipse := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Ellipse`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Line := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Line`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Path := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Path`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polygone := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Polygone`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polyline := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Polyline`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_SVG := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-SVG`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Text := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Text`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_CX.Name = `CX`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.CX]
	__Field__000000_CX.Identifier = `ref_models.Circle.CX`
	__Field__000000_CX.FieldTypeAsString = ``
	__Field__000000_CX.Structname = `Circle`
	__Field__000000_CX.Fieldtypename = `float64`

	// Field values setup
	__Field__000001_CY.Name = `CY`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.CY]
	__Field__000001_CY.Identifier = `ref_models.Circle.CY`
	__Field__000001_CY.FieldTypeAsString = ``
	__Field__000001_CY.Structname = `Circle`
	__Field__000001_CY.Fieldtypename = `float64`

	// Field values setup
	__Field__000002_Display.Name = `Display`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Display]
	__Field__000002_Display.Identifier = `ref_models.Layer.Display`
	__Field__000002_Display.FieldTypeAsString = ``
	__Field__000002_Display.Structname = `SVG`
	__Field__000002_Display.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.Name]
	__Field__000003_Name.Identifier = `ref_models.SVG.Name`
	__Field__000003_Name.FieldTypeAsString = ``
	__Field__000003_Name.Structname = `SVG`
	__Field__000003_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000004_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Name]
	__Field__000004_Name.Identifier = `ref_models.Layer.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `SVG`
	__Field__000004_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000005_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Name]
	__Field__000005_Name.Identifier = `ref_models.Circle.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `Circle`
	__Field__000005_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Radius.Name = `Radius`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Radius]
	__Field__000006_Radius.Identifier = `ref_models.Circle.Radius`
	__Field__000006_Radius.FieldTypeAsString = ``
	__Field__000006_Radius.Structname = `Circle`
	__Field__000006_Radius.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Animate.Name = `NewDiagram-Animate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__GongStructShape__000000_NewDiagram_Animate.Identifier = `ref_models.Animate`
	__GongStructShape__000000_NewDiagram_Animate.ShowNbInstances = true
	__GongStructShape__000000_NewDiagram_Animate.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Animate.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Animate.Heigth = 63.000000
	__GongStructShape__000000_NewDiagram_Animate.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Circle.Name = `NewDiagram-Circle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle]
	__GongStructShape__000001_NewDiagram_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000001_NewDiagram_Circle.ShowNbInstances = true
	__GongStructShape__000001_NewDiagram_Circle.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Circle.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Circle.Heigth = 123.000000
	__GongStructShape__000001_NewDiagram_Circle.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Ellipse.Name = `NewDiagram-Ellipse`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Ellipse]
	__GongStructShape__000002_NewDiagram_Ellipse.Identifier = `ref_models.Ellipse`
	__GongStructShape__000002_NewDiagram_Ellipse.ShowNbInstances = true
	__GongStructShape__000002_NewDiagram_Ellipse.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Ellipse.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Ellipse.Heigth = 63.000000
	__GongStructShape__000002_NewDiagram_Ellipse.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_Line.Name = `NewDiagram-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000003_NewDiagram_Line.Identifier = `ref_models.Line`
	__GongStructShape__000003_NewDiagram_Line.ShowNbInstances = true
	__GongStructShape__000003_NewDiagram_Line.NbInstances = 0
	__GongStructShape__000003_NewDiagram_Line.Width = 240.000000
	__GongStructShape__000003_NewDiagram_Line.Heigth = 63.000000
	__GongStructShape__000003_NewDiagram_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_Path.Name = `NewDiagram-Path`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Path]
	__GongStructShape__000004_NewDiagram_Path.Identifier = `ref_models.Path`
	__GongStructShape__000004_NewDiagram_Path.ShowNbInstances = true
	__GongStructShape__000004_NewDiagram_Path.NbInstances = 0
	__GongStructShape__000004_NewDiagram_Path.Width = 240.000000
	__GongStructShape__000004_NewDiagram_Path.Heigth = 63.000000
	__GongStructShape__000004_NewDiagram_Path.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NewDiagram_Polygone.Name = `NewDiagram-Polygone`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polygone]
	__GongStructShape__000005_NewDiagram_Polygone.Identifier = `ref_models.Polygone`
	__GongStructShape__000005_NewDiagram_Polygone.ShowNbInstances = true
	__GongStructShape__000005_NewDiagram_Polygone.NbInstances = 0
	__GongStructShape__000005_NewDiagram_Polygone.Width = 240.000000
	__GongStructShape__000005_NewDiagram_Polygone.Heigth = 63.000000
	__GongStructShape__000005_NewDiagram_Polygone.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NewDiagram_Polyline.Name = `NewDiagram-Polyline`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polyline]
	__GongStructShape__000006_NewDiagram_Polyline.Identifier = `ref_models.Polyline`
	__GongStructShape__000006_NewDiagram_Polyline.ShowNbInstances = true
	__GongStructShape__000006_NewDiagram_Polyline.NbInstances = 0
	__GongStructShape__000006_NewDiagram_Polyline.Width = 240.000000
	__GongStructShape__000006_NewDiagram_Polyline.Heigth = 63.000000
	__GongStructShape__000006_NewDiagram_Polyline.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_NewDiagram_Rect.Name = `NewDiagram-Rect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__GongStructShape__000007_NewDiagram_Rect.Identifier = `ref_models.Rect`
	__GongStructShape__000007_NewDiagram_Rect.ShowNbInstances = true
	__GongStructShape__000007_NewDiagram_Rect.NbInstances = 147
	__GongStructShape__000007_NewDiagram_Rect.Width = 240.000000
	__GongStructShape__000007_NewDiagram_Rect.Heigth = 63.000000
	__GongStructShape__000007_NewDiagram_Rect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_NewDiagram_SVG.Name = `NewDiagram-SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG]
	__GongStructShape__000008_NewDiagram_SVG.Identifier = `ref_models.SVG`
	__GongStructShape__000008_NewDiagram_SVG.ShowNbInstances = false
	__GongStructShape__000008_NewDiagram_SVG.NbInstances = 0
	__GongStructShape__000008_NewDiagram_SVG.Width = 240.000000
	__GongStructShape__000008_NewDiagram_SVG.Heigth = 78.000000
	__GongStructShape__000008_NewDiagram_SVG.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000009_NewDiagram_SVG.Name = `NewDiagram-SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__GongStructShape__000009_NewDiagram_SVG.Identifier = `ref_models.Layer`
	__GongStructShape__000009_NewDiagram_SVG.ShowNbInstances = true
	__GongStructShape__000009_NewDiagram_SVG.NbInstances = 2
	__GongStructShape__000009_NewDiagram_SVG.Width = 240.000000
	__GongStructShape__000009_NewDiagram_SVG.Heigth = 93.000000
	__GongStructShape__000009_NewDiagram_SVG.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000010_NewDiagram_Text.Name = `NewDiagram-Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__GongStructShape__000010_NewDiagram_Text.Identifier = `ref_models.Text`
	__GongStructShape__000010_NewDiagram_Text.ShowNbInstances = true
	__GongStructShape__000010_NewDiagram_Text.NbInstances = 147
	__GongStructShape__000010_NewDiagram_Text.Width = 240.000000
	__GongStructShape__000010_NewDiagram_Text.Heigth = 63.000000
	__GongStructShape__000010_NewDiagram_Text.IsSelected = false

	// Link values setup
	__Link__000000_Animations.Name = `Animations`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Animations]
	__Link__000000_Animations.Identifier = `ref_models.Circle.Animations`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__Link__000000_Animations.Fieldtypename = `ref_models.Animate`
	__Link__000000_Animations.TargetMultiplicity = models.MANY
	__Link__000000_Animations.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000001_Circles.Name = `Circles`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Circles]
	__Link__000001_Circles.Identifier = `ref_models.Layer.Circles`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle]
	__Link__000001_Circles.Fieldtypename = `ref_models.Circle`
	__Link__000001_Circles.TargetMultiplicity = models.MANY
	__Link__000001_Circles.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000002_Ellipses.Name = `Ellipses`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Ellipses]
	__Link__000002_Ellipses.Identifier = `ref_models.Layer.Ellipses`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Ellipse]
	__Link__000002_Ellipses.Fieldtypename = `ref_models.Ellipse`
	__Link__000002_Ellipses.TargetMultiplicity = models.MANY
	__Link__000002_Ellipses.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000003_Layers.Name = `Layers`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.Layers]
	__Link__000003_Layers.Identifier = `ref_models.SVG.Layers`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__Link__000003_Layers.Fieldtypename = `ref_models.Layer`
	__Link__000003_Layers.TargetMultiplicity = models.MANY
	__Link__000003_Layers.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000004_Lines.Name = `Lines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Lines]
	__Link__000004_Lines.Identifier = `ref_models.Layer.Lines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__Link__000004_Lines.Fieldtypename = `ref_models.Line`
	__Link__000004_Lines.TargetMultiplicity = models.MANY
	__Link__000004_Lines.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000005_Paths.Name = `Paths`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Paths]
	__Link__000005_Paths.Identifier = `ref_models.Layer.Paths`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Path]
	__Link__000005_Paths.Fieldtypename = `ref_models.Path`
	__Link__000005_Paths.TargetMultiplicity = models.MANY
	__Link__000005_Paths.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000006_Polygones.Name = `Polygones`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Polygones]
	__Link__000006_Polygones.Identifier = `ref_models.Layer.Polygones`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polygone]
	__Link__000006_Polygones.Fieldtypename = `ref_models.Polygone`
	__Link__000006_Polygones.TargetMultiplicity = models.MANY
	__Link__000006_Polygones.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000007_Polylines.Name = `Polylines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Polylines]
	__Link__000007_Polylines.Identifier = `ref_models.Layer.Polylines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polyline]
	__Link__000007_Polylines.Fieldtypename = `ref_models.Polyline`
	__Link__000007_Polylines.TargetMultiplicity = models.MANY
	__Link__000007_Polylines.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000008_Rects.Name = `Rects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Rects]
	__Link__000008_Rects.Identifier = `ref_models.Layer.Rects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000008_Rects.Fieldtypename = `ref_models.Rect`
	__Link__000008_Rects.TargetMultiplicity = models.MANY
	__Link__000008_Rects.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000009_Texts.Name = `Texts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Texts]
	__Link__000009_Texts.Identifier = `ref_models.Layer.Texts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__Link__000009_Texts.Fieldtypename = `ref_models.Text`
	__Link__000009_Texts.TargetMultiplicity = models.MANY
	__Link__000009_Texts.SourceMultiplicity = models.ZERO_ONE

	// Position values setup
	__Position__000000_Pos_NewDiagram_Animate.X = 1170.000000
	__Position__000000_Pos_NewDiagram_Animate.Y = 360.000000
	__Position__000000_Pos_NewDiagram_Animate.Name = `Pos-NewDiagram-Animate`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Circle.X = 780.000000
	__Position__000001_Pos_NewDiagram_Circle.Y = 680.000000
	__Position__000001_Pos_NewDiagram_Circle.Name = `Pos-NewDiagram-Circle`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Ellipse.X = 780.000000
	__Position__000002_Pos_NewDiagram_Ellipse.Y = 500.000000
	__Position__000002_Pos_NewDiagram_Ellipse.Name = `Pos-NewDiagram-Ellipse`

	// Position values setup
	__Position__000003_Pos_NewDiagram_Line.X = 780.000000
	__Position__000003_Pos_NewDiagram_Line.Y = 590.000000
	__Position__000003_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Position values setup
	__Position__000004_Pos_NewDiagram_Path.X = 780.000000
	__Position__000004_Pos_NewDiagram_Path.Y = 240.000000
	__Position__000004_Pos_NewDiagram_Path.Name = `Pos-NewDiagram-Path`

	// Position values setup
	__Position__000005_Pos_NewDiagram_Polygone.X = 780.000000
	__Position__000005_Pos_NewDiagram_Polygone.Y = 410.000000
	__Position__000005_Pos_NewDiagram_Polygone.Name = `Pos-NewDiagram-Polygone`

	// Position values setup
	__Position__000006_Pos_NewDiagram_Polyline.X = 780.000000
	__Position__000006_Pos_NewDiagram_Polyline.Y = 330.000000
	__Position__000006_Pos_NewDiagram_Polyline.Name = `Pos-NewDiagram-Polyline`

	// Position values setup
	__Position__000007_Pos_NewDiagram_Rect.X = 780.000000
	__Position__000007_Pos_NewDiagram_Rect.Y = 150.000000
	__Position__000007_Pos_NewDiagram_Rect.Name = `Pos-NewDiagram-Rect`

	// Position values setup
	__Position__000008_Pos_NewDiagram_SVG.X = 100.000000
	__Position__000008_Pos_NewDiagram_SVG.Y = 360.000000
	__Position__000008_Pos_NewDiagram_SVG.Name = `Pos-NewDiagram-SVG`

	// Position values setup
	__Position__000009_Pos_NewDiagram_SVG.X = 22.000000
	__Position__000009_Pos_NewDiagram_SVG.Y = 33.000000
	__Position__000009_Pos_NewDiagram_SVG.Name = `Pos-NewDiagram-SVG`

	// Position values setup
	__Position__000010_Pos_NewDiagram_Text.X = 780.000000
	__Position__000010_Pos_NewDiagram_Text.Y = 60.000000
	__Position__000010_Pos_NewDiagram_Text.Name = `Pos-NewDiagram-Text`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_Animate.X = 1285.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_Animate.Y = 701.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_Animate.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Circle and NewDiagram-Animate`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Circle.X = 430.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Circle.Y = 706.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Circle.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Circle`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Ellipse.X = 460.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Ellipse.Y = 529.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Ellipse.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Ellipse`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Line.X = 450.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Line.Y = 624.000000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Line.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Line`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Path.X = 500.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Path.Y = 269.000000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Path.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Path`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polygone.X = 480.000000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polygone.Y = 441.500000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polygone.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Polygone`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polyline.X = 490.000000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polyline.Y = 361.500000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polyline.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Polyline`

	// Vertice values setup
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.X = 530.000000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.Y = 181.500000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_SVG.X = 76.000000
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_SVG.Y = 168.000000
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_SVG.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-SVG`

	// Vertice values setup
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Text.X = 550.000000
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Text.Y = 86.500000
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Text.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Text`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Animate)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Ellipse)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_Line)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_Path)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000005_NewDiagram_Polygone)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000007_NewDiagram_Rect)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000009_NewDiagram_SVG)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000010_NewDiagram_Text)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_Polyline)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Circle)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000008_NewDiagram_SVG)
	__GongStructShape__000000_NewDiagram_Animate.Position = __Position__000000_Pos_NewDiagram_Animate
	__GongStructShape__000001_NewDiagram_Circle.Position = __Position__000001_Pos_NewDiagram_Circle
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000005_Name)
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000000_CX)
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000001_CY)
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000006_Radius)
	__GongStructShape__000001_NewDiagram_Circle.Links = append(__GongStructShape__000001_NewDiagram_Circle.Links, __Link__000000_Animations)
	__GongStructShape__000002_NewDiagram_Ellipse.Position = __Position__000002_Pos_NewDiagram_Ellipse
	__GongStructShape__000003_NewDiagram_Line.Position = __Position__000003_Pos_NewDiagram_Line
	__GongStructShape__000004_NewDiagram_Path.Position = __Position__000004_Pos_NewDiagram_Path
	__GongStructShape__000005_NewDiagram_Polygone.Position = __Position__000005_Pos_NewDiagram_Polygone
	__GongStructShape__000006_NewDiagram_Polyline.Position = __Position__000006_Pos_NewDiagram_Polyline
	__GongStructShape__000007_NewDiagram_Rect.Position = __Position__000007_Pos_NewDiagram_Rect
	__GongStructShape__000008_NewDiagram_SVG.Position = __Position__000009_Pos_NewDiagram_SVG
	__GongStructShape__000008_NewDiagram_SVG.Fields = append(__GongStructShape__000008_NewDiagram_SVG.Fields, __Field__000003_Name)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000003_Layers)
	__GongStructShape__000009_NewDiagram_SVG.Position = __Position__000008_Pos_NewDiagram_SVG
	__GongStructShape__000009_NewDiagram_SVG.Fields = append(__GongStructShape__000009_NewDiagram_SVG.Fields, __Field__000002_Display)
	__GongStructShape__000009_NewDiagram_SVG.Fields = append(__GongStructShape__000009_NewDiagram_SVG.Fields, __Field__000004_Name)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000001_Circles)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000002_Ellipses)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000004_Lines)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000005_Paths)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000006_Polygones)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000007_Polylines)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000008_Rects)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000009_Texts)
	__GongStructShape__000010_NewDiagram_Text.Position = __Position__000010_Pos_NewDiagram_Text
	__Link__000000_Animations.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Circle_and_NewDiagram_Animate
	__Link__000001_Circles.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Circle
	__Link__000002_Ellipses.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Ellipse
	__Link__000003_Layers.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_SVG
	__Link__000004_Lines.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Line
	__Link__000005_Paths.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Path
	__Link__000006_Polygones.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polygone
	__Link__000007_Polylines.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Polyline
	__Link__000008_Rects.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect
	__Link__000009_Texts.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Text
}


