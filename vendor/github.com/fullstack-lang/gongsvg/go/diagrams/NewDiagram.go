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

	"ref_models.Aliceblue": ref_models.Aliceblue,

	"ref_models.Animate": &(ref_models.Animate{}),

	"ref_models.Animate.AttributeName": (ref_models.Animate{}).AttributeName,

	"ref_models.Animate.Dur": (ref_models.Animate{}).Dur,

	"ref_models.Animate.Name": (ref_models.Animate{}).Name,

	"ref_models.Animate.RepeatCount": (ref_models.Animate{}).RepeatCount,

	"ref_models.Animate.Values": (ref_models.Animate{}).Values,

	"ref_models.Antiquewhite": ref_models.Antiquewhite,

	"ref_models.Aqua": ref_models.Aqua,

	"ref_models.Aquamarine": ref_models.Aquamarine,

	"ref_models.Azure": ref_models.Azure,

	"ref_models.Beige": ref_models.Beige,

	"ref_models.Bisque": ref_models.Bisque,

	"ref_models.Black": ref_models.Black,

	"ref_models.Blanchedalmond": ref_models.Blanchedalmond,

	"ref_models.Blue": ref_models.Blue,

	"ref_models.Blueviolet": ref_models.Blueviolet,

	"ref_models.Brown": ref_models.Brown,

	"ref_models.Burlywood": ref_models.Burlywood,

	"ref_models.Cadetblue": ref_models.Cadetblue,

	"ref_models.Chartreuse": ref_models.Chartreuse,

	"ref_models.Chocolate": ref_models.Chocolate,

	"ref_models.Circle": &(ref_models.Circle{}),

	"ref_models.Circle.Animations": (ref_models.Circle{}).Animations,

	"ref_models.Circle.CX": (ref_models.Circle{}).CX,

	"ref_models.Circle.CY": (ref_models.Circle{}).CY,

	"ref_models.Circle.Color": (ref_models.Circle{}).Color,

	"ref_models.Circle.FillOpacity": (ref_models.Circle{}).FillOpacity,

	"ref_models.Circle.Name": (ref_models.Circle{}).Name,

	"ref_models.Circle.Radius": (ref_models.Circle{}).Radius,

	"ref_models.Circle.Stroke": (ref_models.Circle{}).Stroke,

	"ref_models.Circle.StrokeDashArray": (ref_models.Circle{}).StrokeDashArray,

	"ref_models.Circle.StrokeWidth": (ref_models.Circle{}).StrokeWidth,

	"ref_models.Circle.Transform": (ref_models.Circle{}).Transform,

	"ref_models.ColorType": ref_models.ColorType(""),

	"ref_models.Coral": ref_models.Coral,

	"ref_models.Cornflowerblue": ref_models.Cornflowerblue,

	"ref_models.Cornsilk": ref_models.Cornsilk,

	"ref_models.Crimson": ref_models.Crimson,

	"ref_models.Cyan": ref_models.Cyan,

	"ref_models.DRAWING_LINE": ref_models.DRAWING_LINE,

	"ref_models.Darkblue": ref_models.Darkblue,

	"ref_models.Darkcyan": ref_models.Darkcyan,

	"ref_models.Darkgoldenrod": ref_models.Darkgoldenrod,

	"ref_models.Darkgray": ref_models.Darkgray,

	"ref_models.Darkgreen": ref_models.Darkgreen,

	"ref_models.Darkgrey": ref_models.Darkgrey,

	"ref_models.Darkkhaki": ref_models.Darkkhaki,

	"ref_models.Darkmagenta": ref_models.Darkmagenta,

	"ref_models.Darkolivegreen": ref_models.Darkolivegreen,

	"ref_models.Darkorange": ref_models.Darkorange,

	"ref_models.Darkorchid": ref_models.Darkorchid,

	"ref_models.Darkred": ref_models.Darkred,

	"ref_models.Darksalmon": ref_models.Darksalmon,

	"ref_models.Darkseagreen": ref_models.Darkseagreen,

	"ref_models.Darkslateblue": ref_models.Darkslateblue,

	"ref_models.Darkslategray": ref_models.Darkslategray,

	"ref_models.Darkslategrey": ref_models.Darkslategrey,

	"ref_models.Darkturquoise": ref_models.Darkturquoise,

	"ref_models.Darkviolet": ref_models.Darkviolet,

	"ref_models.Deeppink": ref_models.Deeppink,

	"ref_models.Deepskyblue": ref_models.Deepskyblue,

	"ref_models.Dimgray": ref_models.Dimgray,

	"ref_models.Dimgrey": ref_models.Dimgrey,

	"ref_models.Dodgerblue": ref_models.Dodgerblue,

	"ref_models.DrawingState": ref_models.DrawingState(""),

	"ref_models.Ellipse": &(ref_models.Ellipse{}),

	"ref_models.Ellipse.Animates": (ref_models.Ellipse{}).Animates,

	"ref_models.Ellipse.CX": (ref_models.Ellipse{}).CX,

	"ref_models.Ellipse.CY": (ref_models.Ellipse{}).CY,

	"ref_models.Ellipse.Color": (ref_models.Ellipse{}).Color,

	"ref_models.Ellipse.FillOpacity": (ref_models.Ellipse{}).FillOpacity,

	"ref_models.Ellipse.Name": (ref_models.Ellipse{}).Name,

	"ref_models.Ellipse.RX": (ref_models.Ellipse{}).RX,

	"ref_models.Ellipse.RY": (ref_models.Ellipse{}).RY,

	"ref_models.Ellipse.Stroke": (ref_models.Ellipse{}).Stroke,

	"ref_models.Ellipse.StrokeDashArray": (ref_models.Ellipse{}).StrokeDashArray,

	"ref_models.Ellipse.StrokeWidth": (ref_models.Ellipse{}).StrokeWidth,

	"ref_models.Ellipse.Transform": (ref_models.Ellipse{}).Transform,

	"ref_models.Firebrick": ref_models.Firebrick,

	"ref_models.Floralwhite": ref_models.Floralwhite,

	"ref_models.Forestgreen": ref_models.Forestgreen,

	"ref_models.Fuchsia": ref_models.Fuchsia,

	"ref_models.Gainsboro": ref_models.Gainsboro,

	"ref_models.Ghostwhite": ref_models.Ghostwhite,

	"ref_models.Gold": ref_models.Gold,

	"ref_models.Goldenrod": ref_models.Goldenrod,

	"ref_models.Gray": ref_models.Gray,

	"ref_models.Green": ref_models.Green,

	"ref_models.Greenyellow": ref_models.Greenyellow,

	"ref_models.Grey": ref_models.Grey,

	"ref_models.Honeydew": ref_models.Honeydew,

	"ref_models.Hotpink": ref_models.Hotpink,

	"ref_models.Indianred": ref_models.Indianred,

	"ref_models.Indigo": ref_models.Indigo,

	"ref_models.Ivory": ref_models.Ivory,

	"ref_models.Khaki": ref_models.Khaki,

	"ref_models.Lavender": ref_models.Lavender,

	"ref_models.Lavenderblush": ref_models.Lavenderblush,

	"ref_models.Lawngreen": ref_models.Lawngreen,

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

	"ref_models.Lemonchiffon": ref_models.Lemonchiffon,

	"ref_models.Lightblue": ref_models.Lightblue,

	"ref_models.Lightcoral": ref_models.Lightcoral,

	"ref_models.Lightcyan": ref_models.Lightcyan,

	"ref_models.Lightgoldenrodyellow": ref_models.Lightgoldenrodyellow,

	"ref_models.Lightgray": ref_models.Lightgray,

	"ref_models.Lightgreen": ref_models.Lightgreen,

	"ref_models.Lightgrey": ref_models.Lightgrey,

	"ref_models.Lightpink": ref_models.Lightpink,

	"ref_models.Lightsalmon": ref_models.Lightsalmon,

	"ref_models.Lightseagreen": ref_models.Lightseagreen,

	"ref_models.Lightskyblue": ref_models.Lightskyblue,

	"ref_models.Lightslategray": ref_models.Lightslategray,

	"ref_models.Lightslategrey": ref_models.Lightslategrey,

	"ref_models.Lightsteelblue": ref_models.Lightsteelblue,

	"ref_models.Lightyellow": ref_models.Lightyellow,

	"ref_models.Lime": ref_models.Lime,

	"ref_models.Limegreen": ref_models.Limegreen,

	"ref_models.Line": &(ref_models.Line{}),

	"ref_models.Line.Animates": (ref_models.Line{}).Animates,

	"ref_models.Line.Color": (ref_models.Line{}).Color,

	"ref_models.Line.FillOpacity": (ref_models.Line{}).FillOpacity,

	"ref_models.Line.MouseClickX": (ref_models.Line{}).MouseClickX,

	"ref_models.Line.MouseClickY": (ref_models.Line{}).MouseClickY,

	"ref_models.Line.Name": (ref_models.Line{}).Name,

	"ref_models.Line.Stroke": (ref_models.Line{}).Stroke,

	"ref_models.Line.StrokeDashArray": (ref_models.Line{}).StrokeDashArray,

	"ref_models.Line.StrokeWidth": (ref_models.Line{}).StrokeWidth,

	"ref_models.Line.Transform": (ref_models.Line{}).Transform,

	"ref_models.Line.X1": (ref_models.Line{}).X1,

	"ref_models.Line.X2": (ref_models.Line{}).X2,

	"ref_models.Line.Y1": (ref_models.Line{}).Y1,

	"ref_models.Line.Y2": (ref_models.Line{}).Y2,

	"ref_models.Linen": ref_models.Linen,

	"ref_models.Magenta": ref_models.Magenta,

	"ref_models.Maroon": ref_models.Maroon,

	"ref_models.Mediumaquamarine": ref_models.Mediumaquamarine,

	"ref_models.Mediumblue": ref_models.Mediumblue,

	"ref_models.Mediumorchid": ref_models.Mediumorchid,

	"ref_models.Mediumpurple": ref_models.Mediumpurple,

	"ref_models.Mediumseagreen": ref_models.Mediumseagreen,

	"ref_models.Mediumslateblue": ref_models.Mediumslateblue,

	"ref_models.Mediumspringgreen": ref_models.Mediumspringgreen,

	"ref_models.Mediumturquoise": ref_models.Mediumturquoise,

	"ref_models.Mediumvioletred": ref_models.Mediumvioletred,

	"ref_models.Midnightblue": ref_models.Midnightblue,

	"ref_models.Mintcream": ref_models.Mintcream,

	"ref_models.Mistyrose": ref_models.Mistyrose,

	"ref_models.Moccasin": ref_models.Moccasin,

	"ref_models.NOT_DRAWING_LINE": ref_models.NOT_DRAWING_LINE,

	"ref_models.Navajowhite": ref_models.Navajowhite,

	"ref_models.Navy": ref_models.Navy,

	"ref_models.Oldlace": ref_models.Oldlace,

	"ref_models.Olive": ref_models.Olive,

	"ref_models.Olivedrab": ref_models.Olivedrab,

	"ref_models.Orange": ref_models.Orange,

	"ref_models.Orangered": ref_models.Orangered,

	"ref_models.Orchid": ref_models.Orchid,

	"ref_models.Palegoldenrod": ref_models.Palegoldenrod,

	"ref_models.Palegreen": ref_models.Palegreen,

	"ref_models.Paleturquoise": ref_models.Paleturquoise,

	"ref_models.Palevioletred": ref_models.Palevioletred,

	"ref_models.Papayawhip": ref_models.Papayawhip,

	"ref_models.Path": &(ref_models.Path{}),

	"ref_models.Path.Animates": (ref_models.Path{}).Animates,

	"ref_models.Path.Color": (ref_models.Path{}).Color,

	"ref_models.Path.Definition": (ref_models.Path{}).Definition,

	"ref_models.Path.FillOpacity": (ref_models.Path{}).FillOpacity,

	"ref_models.Path.Name": (ref_models.Path{}).Name,

	"ref_models.Path.Stroke": (ref_models.Path{}).Stroke,

	"ref_models.Path.StrokeDashArray": (ref_models.Path{}).StrokeDashArray,

	"ref_models.Path.StrokeWidth": (ref_models.Path{}).StrokeWidth,

	"ref_models.Path.Transform": (ref_models.Path{}).Transform,

	"ref_models.Peachpuff": ref_models.Peachpuff,

	"ref_models.Peru": ref_models.Peru,

	"ref_models.Pink": ref_models.Pink,

	"ref_models.Plum": ref_models.Plum,

	"ref_models.Polygone": &(ref_models.Polygone{}),

	"ref_models.Polygone.Animates": (ref_models.Polygone{}).Animates,

	"ref_models.Polygone.Color": (ref_models.Polygone{}).Color,

	"ref_models.Polygone.FillOpacity": (ref_models.Polygone{}).FillOpacity,

	"ref_models.Polygone.Name": (ref_models.Polygone{}).Name,

	"ref_models.Polygone.Points": (ref_models.Polygone{}).Points,

	"ref_models.Polygone.Stroke": (ref_models.Polygone{}).Stroke,

	"ref_models.Polygone.StrokeDashArray": (ref_models.Polygone{}).StrokeDashArray,

	"ref_models.Polygone.StrokeWidth": (ref_models.Polygone{}).StrokeWidth,

	"ref_models.Polygone.Transform": (ref_models.Polygone{}).Transform,

	"ref_models.Polyline": &(ref_models.Polyline{}),

	"ref_models.Polyline.Animates": (ref_models.Polyline{}).Animates,

	"ref_models.Polyline.Color": (ref_models.Polyline{}).Color,

	"ref_models.Polyline.FillOpacity": (ref_models.Polyline{}).FillOpacity,

	"ref_models.Polyline.Name": (ref_models.Polyline{}).Name,

	"ref_models.Polyline.Points": (ref_models.Polyline{}).Points,

	"ref_models.Polyline.Stroke": (ref_models.Polyline{}).Stroke,

	"ref_models.Polyline.StrokeDashArray": (ref_models.Polyline{}).StrokeDashArray,

	"ref_models.Polyline.StrokeWidth": (ref_models.Polyline{}).StrokeWidth,

	"ref_models.Polyline.Transform": (ref_models.Polyline{}).Transform,

	"ref_models.Powderblue": ref_models.Powderblue,

	"ref_models.Purple": ref_models.Purple,

	"ref_models.Rect": &(ref_models.Rect{}),

	"ref_models.Rect.Animations": (ref_models.Rect{}).Animations,

	"ref_models.Rect.CanHaveHorizontalHandles": (ref_models.Rect{}).CanHaveLeftHandle,

	"ref_models.Rect.CanMoveHorizontaly": (ref_models.Rect{}).CanMoveHorizontaly,

	"ref_models.Rect.CanMoveVerticaly": (ref_models.Rect{}).CanMoveVerticaly,

	"ref_models.Rect.Color": (ref_models.Rect{}).Color,

	"ref_models.Rect.FillOpacity": (ref_models.Rect{}).FillOpacity,

	"ref_models.Rect.HasHorizontalHandles": (ref_models.Rect{}).HasLeftHandle,

	"ref_models.Rect.Height": (ref_models.Rect{}).Height,

	"ref_models.Rect.IsSelectable": (ref_models.Rect{}).IsSelectable,

	"ref_models.Rect.IsSelected": (ref_models.Rect{}).IsSelected,

	"ref_models.Rect.Name": (ref_models.Rect{}).Name,

	"ref_models.Rect.RX": (ref_models.Rect{}).RX,

	"ref_models.Rect.Stroke": (ref_models.Rect{}).Stroke,

	"ref_models.Rect.StrokeDashArray": (ref_models.Rect{}).StrokeDashArray,

	"ref_models.Rect.StrokeWidth": (ref_models.Rect{}).StrokeWidth,

	"ref_models.Rect.Transform": (ref_models.Rect{}).Transform,

	"ref_models.Rect.Width": (ref_models.Rect{}).Width,

	"ref_models.Rect.X": (ref_models.Rect{}).X,

	"ref_models.Rect.Y": (ref_models.Rect{}).Y,

	"ref_models.Red": ref_models.Red,

	"ref_models.Rosybrown": ref_models.Rosybrown,

	"ref_models.Royalblue": ref_models.Royalblue,

	"ref_models.SVG": &(ref_models.SVG{}),

	"ref_models.SVG.DrawingState": (ref_models.SVG{}).DrawingState,

	"ref_models.SVG.EndRect": (ref_models.SVG{}).EndRect,

	"ref_models.SVG.Layers": (ref_models.SVG{}).Layers,

	"ref_models.SVG.Name": (ref_models.SVG{}).Name,

	"ref_models.SVG.StartRect": (ref_models.SVG{}).StartRect,

	"ref_models.Saddlebrown": ref_models.Saddlebrown,

	"ref_models.Salmon": ref_models.Salmon,

	"ref_models.Sandybrown": ref_models.Sandybrown,

	"ref_models.Seagreen": ref_models.Seagreen,

	"ref_models.Seashell": ref_models.Seashell,

	"ref_models.Sienna": ref_models.Sienna,

	"ref_models.Silver": ref_models.Silver,

	"ref_models.Skyblue": ref_models.Skyblue,

	"ref_models.Slateblue": ref_models.Slateblue,

	"ref_models.Slategray": ref_models.Slategray,

	"ref_models.Slategrey": ref_models.Slategrey,

	"ref_models.Snow": ref_models.Snow,

	"ref_models.Springgreen": ref_models.Springgreen,

	"ref_models.Steelblue": ref_models.Steelblue,

	"ref_models.Tan": ref_models.Tan,

	"ref_models.Teal": ref_models.Teal,

	"ref_models.Text": &(ref_models.Text{}),

	"ref_models.Text.Animates": (ref_models.Text{}).Animates,

	"ref_models.Text.Color": (ref_models.Text{}).Color,

	"ref_models.Text.Content": (ref_models.Text{}).Content,

	"ref_models.Text.FillOpacity": (ref_models.Text{}).FillOpacity,

	"ref_models.Text.Name": (ref_models.Text{}).Name,

	"ref_models.Text.Stroke": (ref_models.Text{}).Stroke,

	"ref_models.Text.StrokeDashArray": (ref_models.Text{}).StrokeDashArray,

	"ref_models.Text.StrokeWidth": (ref_models.Text{}).StrokeWidth,

	"ref_models.Text.Transform": (ref_models.Text{}).Transform,

	"ref_models.Text.X": (ref_models.Text{}).X,

	"ref_models.Text.Y": (ref_models.Text{}).Y,

	"ref_models.Thistle": ref_models.Thistle,

	"ref_models.Tomato": ref_models.Tomato,

	"ref_models.Turquoise": ref_models.Turquoise,

	"ref_models.Violet": ref_models.Violet,

	"ref_models.Wheat": ref_models.Wheat,

	"ref_models.White": ref_models.White,

	"ref_models.Whitesmoke": ref_models.Whitesmoke,

	"ref_models.Yellow": ref_models.Yellow,

	"ref_models.Yellowgreen": ref_models.Yellowgreen,
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
	__Field__000003_MouseClickX := (&models.Field{Name: `MouseClickX`}).Stage(stage)
	__Field__000004_MouseClickY := (&models.Field{Name: `MouseClickY`}).Stage(stage)
	__Field__000005_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000006_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000007_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000008_Radius := (&models.Field{Name: `Radius`}).Stage(stage)
	__Field__000009_StrokeWidth := (&models.Field{Name: `StrokeWidth`}).Stage(stage)
	__Field__000010_X1 := (&models.Field{Name: `X1`}).Stage(stage)
	__Field__000011_X2 := (&models.Field{Name: `X2`}).Stage(stage)
	__Field__000012_Y1 := (&models.Field{Name: `Y1`}).Stage(stage)
	__Field__000013_Y2 := (&models.Field{Name: `Y2`}).Stage(stage)

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
	__Field__000003_MouseClickX.Name = `MouseClickX`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.MouseClickX]
	__Field__000003_MouseClickX.Identifier = `ref_models.Line.MouseClickX`
	__Field__000003_MouseClickX.FieldTypeAsString = ``
	__Field__000003_MouseClickX.Structname = `Line`
	__Field__000003_MouseClickX.Fieldtypename = `float64`

	// Field values setup
	__Field__000004_MouseClickY.Name = `MouseClickY`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.MouseClickY]
	__Field__000004_MouseClickY.Identifier = `ref_models.Line.MouseClickY`
	__Field__000004_MouseClickY.FieldTypeAsString = ``
	__Field__000004_MouseClickY.Structname = `Line`
	__Field__000004_MouseClickY.Fieldtypename = `float64`

	// Field values setup
	__Field__000005_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Name]
	__Field__000005_Name.Identifier = `ref_models.Circle.Name`
	__Field__000005_Name.FieldTypeAsString = ``
	__Field__000005_Name.Structname = `Circle`
	__Field__000005_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000006_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.Name]
	__Field__000006_Name.Identifier = `ref_models.SVG.Name`
	__Field__000006_Name.FieldTypeAsString = ``
	__Field__000006_Name.Structname = `SVG`
	__Field__000006_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Name]
	__Field__000007_Name.Identifier = `ref_models.Layer.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = `SVG`
	__Field__000007_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Radius.Name = `Radius`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle.Radius]
	__Field__000008_Radius.Identifier = `ref_models.Circle.Radius`
	__Field__000008_Radius.FieldTypeAsString = ``
	__Field__000008_Radius.Structname = `Circle`
	__Field__000008_Radius.Fieldtypename = `float64`

	// Field values setup
	__Field__000009_StrokeWidth.Name = `StrokeWidth`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.StrokeWidth]
	__Field__000009_StrokeWidth.Identifier = `ref_models.Line.StrokeWidth`
	__Field__000009_StrokeWidth.FieldTypeAsString = ``
	__Field__000009_StrokeWidth.Structname = `Line`
	__Field__000009_StrokeWidth.Fieldtypename = `float64`

	// Field values setup
	__Field__000010_X1.Name = `X1`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.X1]
	__Field__000010_X1.Identifier = `ref_models.Line.X1`
	__Field__000010_X1.FieldTypeAsString = ``
	__Field__000010_X1.Structname = `Line`
	__Field__000010_X1.Fieldtypename = `float64`

	// Field values setup
	__Field__000011_X2.Name = `X2`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.X2]
	__Field__000011_X2.Identifier = `ref_models.Line.X2`
	__Field__000011_X2.FieldTypeAsString = ``
	__Field__000011_X2.Structname = `Line`
	__Field__000011_X2.Fieldtypename = `float64`

	// Field values setup
	__Field__000012_Y1.Name = `Y1`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Y1]
	__Field__000012_Y1.Identifier = `ref_models.Line.Y1`
	__Field__000012_Y1.FieldTypeAsString = ``
	__Field__000012_Y1.Structname = `Line`
	__Field__000012_Y1.Fieldtypename = `float64`

	// Field values setup
	__Field__000013_Y2.Name = `Y2`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line.Y2]
	__Field__000013_Y2.Identifier = `ref_models.Line.Y2`
	__Field__000013_Y2.FieldTypeAsString = ``
	__Field__000013_Y2.Structname = `Line`
	__Field__000013_Y2.Fieldtypename = `float64`

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
	__GongStructShape__000001_NewDiagram_Circle.NbInstances = 1
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
	__GongStructShape__000003_NewDiagram_Line.NbInstances = 1
	__GongStructShape__000003_NewDiagram_Line.Width = 240.000000
	__GongStructShape__000003_NewDiagram_Line.Heigth = 168.000000
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
	__GongStructShape__000007_NewDiagram_Rect.NbInstances = 2
	__GongStructShape__000007_NewDiagram_Rect.Width = 240.000000
	__GongStructShape__000007_NewDiagram_Rect.Heigth = 63.000000
	__GongStructShape__000007_NewDiagram_Rect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_NewDiagram_SVG.Name = `NewDiagram-SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__GongStructShape__000008_NewDiagram_SVG.Identifier = `ref_models.Layer`
	__GongStructShape__000008_NewDiagram_SVG.ShowNbInstances = true
	__GongStructShape__000008_NewDiagram_SVG.NbInstances = 4
	__GongStructShape__000008_NewDiagram_SVG.Width = 240.000000
	__GongStructShape__000008_NewDiagram_SVG.Heigth = 93.000000
	__GongStructShape__000008_NewDiagram_SVG.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000009_NewDiagram_SVG.Name = `NewDiagram-SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG]
	__GongStructShape__000009_NewDiagram_SVG.Identifier = `ref_models.SVG`
	__GongStructShape__000009_NewDiagram_SVG.ShowNbInstances = true
	__GongStructShape__000009_NewDiagram_SVG.NbInstances = 1
	__GongStructShape__000009_NewDiagram_SVG.Width = 240.000000
	__GongStructShape__000009_NewDiagram_SVG.Heigth = 78.000000
	__GongStructShape__000009_NewDiagram_SVG.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000010_NewDiagram_Text.Name = `NewDiagram-Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__GongStructShape__000010_NewDiagram_Text.Identifier = `ref_models.Text`
	__GongStructShape__000010_NewDiagram_Text.ShowNbInstances = true
	__GongStructShape__000010_NewDiagram_Text.NbInstances = 1
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
	__Position__000001_Pos_NewDiagram_Circle.Y = 770.000000
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
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000008_NewDiagram_SVG)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000010_NewDiagram_Text)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_Polyline)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Circle)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000009_NewDiagram_SVG)
	__GongStructShape__000000_NewDiagram_Animate.Position = __Position__000000_Pos_NewDiagram_Animate
	__GongStructShape__000001_NewDiagram_Circle.Position = __Position__000001_Pos_NewDiagram_Circle
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000005_Name)
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000000_CX)
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000001_CY)
	__GongStructShape__000001_NewDiagram_Circle.Fields = append(__GongStructShape__000001_NewDiagram_Circle.Fields, __Field__000008_Radius)
	__GongStructShape__000001_NewDiagram_Circle.Links = append(__GongStructShape__000001_NewDiagram_Circle.Links, __Link__000000_Animations)
	__GongStructShape__000002_NewDiagram_Ellipse.Position = __Position__000002_Pos_NewDiagram_Ellipse
	__GongStructShape__000003_NewDiagram_Line.Position = __Position__000003_Pos_NewDiagram_Line
	__GongStructShape__000003_NewDiagram_Line.Fields = append(__GongStructShape__000003_NewDiagram_Line.Fields, __Field__000010_X1)
	__GongStructShape__000003_NewDiagram_Line.Fields = append(__GongStructShape__000003_NewDiagram_Line.Fields, __Field__000012_Y1)
	__GongStructShape__000003_NewDiagram_Line.Fields = append(__GongStructShape__000003_NewDiagram_Line.Fields, __Field__000011_X2)
	__GongStructShape__000003_NewDiagram_Line.Fields = append(__GongStructShape__000003_NewDiagram_Line.Fields, __Field__000013_Y2)
	__GongStructShape__000003_NewDiagram_Line.Fields = append(__GongStructShape__000003_NewDiagram_Line.Fields, __Field__000009_StrokeWidth)
	__GongStructShape__000003_NewDiagram_Line.Fields = append(__GongStructShape__000003_NewDiagram_Line.Fields, __Field__000003_MouseClickX)
	__GongStructShape__000003_NewDiagram_Line.Fields = append(__GongStructShape__000003_NewDiagram_Line.Fields, __Field__000004_MouseClickY)
	__GongStructShape__000004_NewDiagram_Path.Position = __Position__000004_Pos_NewDiagram_Path
	__GongStructShape__000005_NewDiagram_Polygone.Position = __Position__000005_Pos_NewDiagram_Polygone
	__GongStructShape__000006_NewDiagram_Polyline.Position = __Position__000006_Pos_NewDiagram_Polyline
	__GongStructShape__000007_NewDiagram_Rect.Position = __Position__000007_Pos_NewDiagram_Rect
	__GongStructShape__000008_NewDiagram_SVG.Position = __Position__000008_Pos_NewDiagram_SVG
	__GongStructShape__000008_NewDiagram_SVG.Fields = append(__GongStructShape__000008_NewDiagram_SVG.Fields, __Field__000002_Display)
	__GongStructShape__000008_NewDiagram_SVG.Fields = append(__GongStructShape__000008_NewDiagram_SVG.Fields, __Field__000007_Name)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000001_Circles)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000002_Ellipses)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000004_Lines)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000005_Paths)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000006_Polygones)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000007_Polylines)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000008_Rects)
	__GongStructShape__000008_NewDiagram_SVG.Links = append(__GongStructShape__000008_NewDiagram_SVG.Links, __Link__000009_Texts)
	__GongStructShape__000009_NewDiagram_SVG.Position = __Position__000009_Pos_NewDiagram_SVG
	__GongStructShape__000009_NewDiagram_SVG.Fields = append(__GongStructShape__000009_NewDiagram_SVG.Fields, __Field__000006_Name)
	__GongStructShape__000009_NewDiagram_SVG.Links = append(__GongStructShape__000009_NewDiagram_SVG.Links, __Link__000003_Layers)
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
