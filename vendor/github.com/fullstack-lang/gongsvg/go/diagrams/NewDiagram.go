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

	"ref_models.ANCHOR_BOTTOM": ref_models.ANCHOR_BOTTOM,

	"ref_models.ANCHOR_CENTER": ref_models.ANCHOR_CENTER,

	"ref_models.ANCHOR_LEFT": ref_models.ANCHOR_LEFT,

	"ref_models.ANCHOR_RIGHT": ref_models.ANCHOR_RIGHT,

	"ref_models.ANCHOR_TOP": ref_models.ANCHOR_TOP,

	"ref_models.Aliceblue": ref_models.Aliceblue,

	"ref_models.AnchorType": ref_models.AnchorType(""),

	"ref_models.AnchoredText": &(ref_models.AnchoredText{}),

	"ref_models.AnchoredText.Animates": (ref_models.AnchoredText{}).Animates,

	"ref_models.AnchoredText.Color": (ref_models.AnchoredText{}).Color,

	"ref_models.AnchoredText.Content": (ref_models.AnchoredText{}).Content,

	"ref_models.AnchoredText.FillOpacity": (ref_models.AnchoredText{}).FillOpacity,

	"ref_models.AnchoredText.Name": (ref_models.AnchoredText{}).Name,

	"ref_models.AnchoredText.Stroke": (ref_models.AnchoredText{}).Stroke,

	"ref_models.AnchoredText.StrokeDashArray": (ref_models.AnchoredText{}).StrokeDashArray,

	"ref_models.AnchoredText.StrokeDashArrayWhenSelected": (ref_models.AnchoredText{}).StrokeDashArrayWhenSelected,

	"ref_models.AnchoredText.StrokeWidth": (ref_models.AnchoredText{}).StrokeWidth,

	"ref_models.AnchoredText.Transform": (ref_models.AnchoredText{}).Transform,

	"ref_models.AnchoredText.X_Offset": (ref_models.AnchoredText{}).X_Offset,

	"ref_models.AnchoredText.Y_Offset": (ref_models.AnchoredText{}).Y_Offset,

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

	"ref_models.Circle.StrokeDashArrayWhenSelected": (ref_models.Circle{}).StrokeDashArrayWhenSelected,

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

	"ref_models.Ellipse.StrokeDashArrayWhenSelected": (ref_models.Ellipse{}).StrokeDashArrayWhenSelected,

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

	"ref_models.LINK_TYPE_FLOATING_ORTHOGONAL": ref_models.LINK_TYPE_FLOATING_ORTHOGONAL,

	"ref_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS": ref_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS,

	"ref_models.Lavender": ref_models.Lavender,

	"ref_models.Lavenderblush": ref_models.Lavenderblush,

	"ref_models.Lawngreen": ref_models.Lawngreen,

	"ref_models.Layer": &(ref_models.Layer{}),

	"ref_models.Layer.Circles": (ref_models.Layer{}).Circles,

	"ref_models.Layer.Display": (ref_models.Layer{}).Display,

	"ref_models.Layer.Ellipses": (ref_models.Layer{}).Ellipses,

	"ref_models.Layer.Lines": (ref_models.Layer{}).Lines,

	"ref_models.Layer.Links": (ref_models.Layer{}).Links,

	"ref_models.Layer.Name": (ref_models.Layer{}).Name,

	"ref_models.Layer.Paths": (ref_models.Layer{}).Paths,

	"ref_models.Layer.Polygones": (ref_models.Layer{}).Polygones,

	"ref_models.Layer.Polylines": (ref_models.Layer{}).Polylines,

	"ref_models.Layer.RectLinkLinks": (ref_models.Layer{}).RectLinkLinks,

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

	"ref_models.Line.StrokeDashArrayWhenSelected": (ref_models.Line{}).StrokeDashArrayWhenSelected,

	"ref_models.Line.StrokeWidth": (ref_models.Line{}).StrokeWidth,

	"ref_models.Line.Transform": (ref_models.Line{}).Transform,

	"ref_models.Line.X1": (ref_models.Line{}).X1,

	"ref_models.Line.X2": (ref_models.Line{}).X2,

	"ref_models.Line.Y1": (ref_models.Line{}).Y1,

	"ref_models.Line.Y2": (ref_models.Line{}).Y2,

	"ref_models.Linen": ref_models.Linen,

	"ref_models.Link": &(ref_models.Link{}),

	"ref_models.Link.Color": (ref_models.Link{}).Color,

	"ref_models.Link.ControlPoints": (ref_models.Link{}).ControlPoints,

	"ref_models.Link.CornerOffsetRatio": (ref_models.Link{}).CornerOffsetRatio,

	"ref_models.Link.CornerRadius": (ref_models.Link{}).CornerRadius,

	"ref_models.Link.End": (ref_models.Link{}).End,

	"ref_models.Link.EndAnchorType": (ref_models.Link{}).EndAnchorType,

	"ref_models.Link.EndArrowSize": (ref_models.Link{}).EndArrowSize,

	"ref_models.Link.EndOrientation": (ref_models.Link{}).EndOrientation,

	"ref_models.Link.EndRatio": (ref_models.Link{}).EndRatio,

	"ref_models.Link.FillOpacity": (ref_models.Link{}).FillOpacity,

	"ref_models.Link.HasEndArrow": (ref_models.Link{}).HasEndArrow,

	"ref_models.Link.Name": (ref_models.Link{}).Name,

	"ref_models.Link.Start": (ref_models.Link{}).Start,

	"ref_models.Link.StartAnchorType": (ref_models.Link{}).StartAnchorType,

	"ref_models.Link.StartOrientation": (ref_models.Link{}).StartOrientation,

	"ref_models.Link.StartRatio": (ref_models.Link{}).StartRatio,

	"ref_models.Link.Stroke": (ref_models.Link{}).Stroke,

	"ref_models.Link.StrokeDashArray": (ref_models.Link{}).StrokeDashArray,

	"ref_models.Link.StrokeDashArrayWhenSelected": (ref_models.Link{}).StrokeDashArrayWhenSelected,

	"ref_models.Link.StrokeWidth": (ref_models.Link{}).StrokeWidth,

	"ref_models.Link.TextAtArrowEnd": (ref_models.Link{}).TextAtArrowEnd,

	"ref_models.Link.TextAtArrowStart": (ref_models.Link{}).TextAtArrowStart,

	"ref_models.Link.Transform": (ref_models.Link{}).Transform,

	"ref_models.Link.Type": (ref_models.Link{}).Type,

	"ref_models.LinkType": ref_models.LinkType(""),

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

	"ref_models.ORIENTATION_HORIZONTAL": ref_models.ORIENTATION_HORIZONTAL,

	"ref_models.ORIENTATION_VERTICAL": ref_models.ORIENTATION_VERTICAL,

	"ref_models.Oldlace": ref_models.Oldlace,

	"ref_models.Olive": ref_models.Olive,

	"ref_models.Olivedrab": ref_models.Olivedrab,

	"ref_models.Orange": ref_models.Orange,

	"ref_models.Orangered": ref_models.Orangered,

	"ref_models.Orchid": ref_models.Orchid,

	"ref_models.OrientationType": ref_models.OrientationType(""),

	"ref_models.POSITION_ON_ARROW_END": ref_models.POSITION_ON_ARROW_END,

	"ref_models.POSITION_ON_ARROW_START": ref_models.POSITION_ON_ARROW_START,

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

	"ref_models.Path.StrokeDashArrayWhenSelected": (ref_models.Path{}).StrokeDashArrayWhenSelected,

	"ref_models.Path.StrokeWidth": (ref_models.Path{}).StrokeWidth,

	"ref_models.Path.Transform": (ref_models.Path{}).Transform,

	"ref_models.Peachpuff": ref_models.Peachpuff,

	"ref_models.Peru": ref_models.Peru,

	"ref_models.Pink": ref_models.Pink,

	"ref_models.Plum": ref_models.Plum,

	"ref_models.Point": &(ref_models.Point{}),

	"ref_models.Point.Name": (ref_models.Point{}).Name,

	"ref_models.Point.X": (ref_models.Point{}).X,

	"ref_models.Point.Y": (ref_models.Point{}).Y,

	"ref_models.Polygone": &(ref_models.Polygone{}),

	"ref_models.Polygone.Animates": (ref_models.Polygone{}).Animates,

	"ref_models.Polygone.Color": (ref_models.Polygone{}).Color,

	"ref_models.Polygone.FillOpacity": (ref_models.Polygone{}).FillOpacity,

	"ref_models.Polygone.Name": (ref_models.Polygone{}).Name,

	"ref_models.Polygone.Points": (ref_models.Polygone{}).Points,

	"ref_models.Polygone.Stroke": (ref_models.Polygone{}).Stroke,

	"ref_models.Polygone.StrokeDashArray": (ref_models.Polygone{}).StrokeDashArray,

	"ref_models.Polygone.StrokeDashArrayWhenSelected": (ref_models.Polygone{}).StrokeDashArrayWhenSelected,

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

	"ref_models.Polyline.StrokeDashArrayWhenSelected": (ref_models.Polyline{}).StrokeDashArrayWhenSelected,

	"ref_models.Polyline.StrokeWidth": (ref_models.Polyline{}).StrokeWidth,

	"ref_models.Polyline.Transform": (ref_models.Polyline{}).Transform,

	"ref_models.PositionOnArrowType": ref_models.PositionOnArrowType(""),

	"ref_models.Powderblue": ref_models.Powderblue,

	"ref_models.Purple": ref_models.Purple,

	"ref_models.RECT_ANCHOR_BOTTOM": ref_models.RECT_ANCHOR_BOTTOM,

	"ref_models.RECT_ANCHOR_BOTTOM_LEFT": ref_models.RECT_ANCHOR_BOTTOM_LEFT,

	"ref_models.RECT_ANCHOR_BOTTOM_RIGHT": ref_models.RECT_ANCHOR_BOTTOM_RIGHT,

	"ref_models.RECT_ANCHOR_CENTER": ref_models.RECT_ANCHOR_CENTER,

	"ref_models.RECT_ANCHOR_LEFT": ref_models.RECT_ANCHOR_LEFT,

	"ref_models.RECT_ANCHOR_RIGHT": ref_models.RECT_ANCHOR_RIGHT,

	"ref_models.RECT_ANCHOR_TOP": ref_models.RECT_ANCHOR_TOP,

	"ref_models.RECT_ANCHOR_TOP_LEFT": ref_models.RECT_ANCHOR_TOP_LEFT,

	"ref_models.RECT_ANCHOR_TOP_RIGHT": ref_models.RECT_ANCHOR_TOP_RIGHT,

	"ref_models.Rect": &(ref_models.Rect{}),

	"ref_models.Rect.Animations": (ref_models.Rect{}).Animations,

	"ref_models.Rect.CanHaveBottomHandle": (ref_models.Rect{}).CanHaveBottomHandle,

	"ref_models.Rect.CanHaveLeftHandle": (ref_models.Rect{}).CanHaveLeftHandle,

	"ref_models.Rect.CanHaveRightHandle": (ref_models.Rect{}).CanHaveRightHandle,

	"ref_models.Rect.CanHaveTopHandle": (ref_models.Rect{}).CanHaveTopHandle,

	"ref_models.Rect.CanMoveHorizontaly": (ref_models.Rect{}).CanMoveHorizontaly,

	"ref_models.Rect.CanMoveVerticaly": (ref_models.Rect{}).CanMoveVerticaly,

	"ref_models.Rect.Color": (ref_models.Rect{}).Color,

	"ref_models.Rect.FillOpacity": (ref_models.Rect{}).FillOpacity,

	"ref_models.Rect.HasBottomHandle": (ref_models.Rect{}).HasBottomHandle,

	"ref_models.Rect.HasLeftHandle": (ref_models.Rect{}).HasLeftHandle,

	"ref_models.Rect.HasRightHandle": (ref_models.Rect{}).HasRightHandle,

	"ref_models.Rect.HasTopHandle": (ref_models.Rect{}).HasTopHandle,

	"ref_models.Rect.Height": (ref_models.Rect{}).Height,

	"ref_models.Rect.IsSelectable": (ref_models.Rect{}).IsSelectable,

	"ref_models.Rect.IsSelected": (ref_models.Rect{}).IsSelected,

	"ref_models.Rect.Name": (ref_models.Rect{}).Name,

	"ref_models.Rect.RX": (ref_models.Rect{}).RX,

	"ref_models.Rect.RectAnchoredRects": (ref_models.Rect{}).RectAnchoredRects,

	"ref_models.Rect.RectAnchoredTexts": (ref_models.Rect{}).RectAnchoredTexts,

	"ref_models.Rect.Stroke": (ref_models.Rect{}).Stroke,

	"ref_models.Rect.StrokeDashArray": (ref_models.Rect{}).StrokeDashArray,

	"ref_models.Rect.StrokeDashArrayWhenSelected": (ref_models.Rect{}).StrokeDashArrayWhenSelected,

	"ref_models.Rect.StrokeWidth": (ref_models.Rect{}).StrokeWidth,

	"ref_models.Rect.Transform": (ref_models.Rect{}).Transform,

	"ref_models.Rect.Width": (ref_models.Rect{}).Width,

	"ref_models.Rect.X": (ref_models.Rect{}).X,

	"ref_models.Rect.Y": (ref_models.Rect{}).Y,

	"ref_models.RectAnchorType": ref_models.RectAnchorType(""),

	"ref_models.RectAnchoredRect": &(ref_models.RectAnchoredRect{}),

	"ref_models.RectAnchoredRect.Color": (ref_models.RectAnchoredRect{}).Color,

	"ref_models.RectAnchoredRect.FillOpacity": (ref_models.RectAnchoredRect{}).FillOpacity,

	"ref_models.RectAnchoredRect.Height": (ref_models.RectAnchoredRect{}).Height,

	"ref_models.RectAnchoredRect.HeightFollowRect": (ref_models.RectAnchoredRect{}).HeightFollowRect,

	"ref_models.RectAnchoredRect.Name": (ref_models.RectAnchoredRect{}).Name,

	"ref_models.RectAnchoredRect.RX": (ref_models.RectAnchoredRect{}).RX,

	"ref_models.RectAnchoredRect.RectAnchorType": (ref_models.RectAnchoredRect{}).RectAnchorType,

	"ref_models.RectAnchoredRect.Stroke": (ref_models.RectAnchoredRect{}).Stroke,

	"ref_models.RectAnchoredRect.StrokeDashArray": (ref_models.RectAnchoredRect{}).StrokeDashArray,

	"ref_models.RectAnchoredRect.StrokeDashArrayWhenSelected": (ref_models.RectAnchoredRect{}).StrokeDashArrayWhenSelected,

	"ref_models.RectAnchoredRect.StrokeWidth": (ref_models.RectAnchoredRect{}).StrokeWidth,

	"ref_models.RectAnchoredRect.Transform": (ref_models.RectAnchoredRect{}).Transform,

	"ref_models.RectAnchoredRect.Width": (ref_models.RectAnchoredRect{}).Width,

	"ref_models.RectAnchoredRect.WidthFollowRect": (ref_models.RectAnchoredRect{}).WidthFollowRect,

	"ref_models.RectAnchoredRect.X": (ref_models.RectAnchoredRect{}).X,

	"ref_models.RectAnchoredRect.X_Offset": (ref_models.RectAnchoredRect{}).X_Offset,

	"ref_models.RectAnchoredRect.Y": (ref_models.RectAnchoredRect{}).Y,

	"ref_models.RectAnchoredRect.Y_Offset": (ref_models.RectAnchoredRect{}).Y_Offset,

	"ref_models.RectAnchoredText": &(ref_models.RectAnchoredText{}),

	"ref_models.RectAnchoredText.Animates": (ref_models.RectAnchoredText{}).Animates,

	"ref_models.RectAnchoredText.Color": (ref_models.RectAnchoredText{}).Color,

	"ref_models.RectAnchoredText.Content": (ref_models.RectAnchoredText{}).Content,

	"ref_models.RectAnchoredText.FillOpacity": (ref_models.RectAnchoredText{}).FillOpacity,

	"ref_models.RectAnchoredText.FontSize": (ref_models.RectAnchoredText{}).FontSize,

	"ref_models.RectAnchoredText.FontWeight": (ref_models.RectAnchoredText{}).FontWeight,

	"ref_models.RectAnchoredText.Name": (ref_models.RectAnchoredText{}).Name,

	"ref_models.RectAnchoredText.RectAnchorType": (ref_models.RectAnchoredText{}).RectAnchorType,

	"ref_models.RectAnchoredText.Stroke": (ref_models.RectAnchoredText{}).Stroke,

	"ref_models.RectAnchoredText.StrokeDashArray": (ref_models.RectAnchoredText{}).StrokeDashArray,

	"ref_models.RectAnchoredText.StrokeDashArrayWhenSelected": (ref_models.RectAnchoredText{}).StrokeDashArrayWhenSelected,

	"ref_models.RectAnchoredText.StrokeWidth": (ref_models.RectAnchoredText{}).StrokeWidth,

	"ref_models.RectAnchoredText.TextAnchorType": (ref_models.RectAnchoredText{}).TextAnchorType,

	"ref_models.RectAnchoredText.Transform": (ref_models.RectAnchoredText{}).Transform,

	"ref_models.RectAnchoredText.X_Offset": (ref_models.RectAnchoredText{}).X_Offset,

	"ref_models.RectAnchoredText.Y_Offset": (ref_models.RectAnchoredText{}).Y_Offset,

	"ref_models.RectLinkLink": &(ref_models.RectLinkLink{}),

	"ref_models.RectLinkLink.Color": (ref_models.RectLinkLink{}).Color,

	"ref_models.RectLinkLink.End": (ref_models.RectLinkLink{}).End,

	"ref_models.RectLinkLink.FillOpacity": (ref_models.RectLinkLink{}).FillOpacity,

	"ref_models.RectLinkLink.Name": (ref_models.RectLinkLink{}).Name,

	"ref_models.RectLinkLink.Start": (ref_models.RectLinkLink{}).Start,

	"ref_models.RectLinkLink.Stroke": (ref_models.RectLinkLink{}).Stroke,

	"ref_models.RectLinkLink.StrokeDashArray": (ref_models.RectLinkLink{}).StrokeDashArray,

	"ref_models.RectLinkLink.StrokeDashArrayWhenSelected": (ref_models.RectLinkLink{}).StrokeDashArrayWhenSelected,

	"ref_models.RectLinkLink.StrokeWidth": (ref_models.RectLinkLink{}).StrokeWidth,

	"ref_models.RectLinkLink.TargetAnchorPosition": (ref_models.RectLinkLink{}).TargetAnchorPosition,

	"ref_models.RectLinkLink.Transform": (ref_models.RectLinkLink{}).Transform,

	"ref_models.Red": ref_models.Red,

	"ref_models.Rosybrown": ref_models.Rosybrown,

	"ref_models.Royalblue": ref_models.Royalblue,

	"ref_models.SIDE_BOTTOM": ref_models.SIDE_BOTTOM,

	"ref_models.SIDE_LEFT": ref_models.SIDE_LEFT,

	"ref_models.SIDE_RIGHT": ref_models.SIDE_RIGHT,

	"ref_models.SIDE_TOP": ref_models.SIDE_TOP,

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

	"ref_models.SideType": ref_models.SideType(""),

	"ref_models.Sienna": ref_models.Sienna,

	"ref_models.Silver": ref_models.Silver,

	"ref_models.Skyblue": ref_models.Skyblue,

	"ref_models.Slateblue": ref_models.Slateblue,

	"ref_models.Slategray": ref_models.Slategray,

	"ref_models.Slategrey": ref_models.Slategrey,

	"ref_models.Snow": ref_models.Snow,

	"ref_models.Springgreen": ref_models.Springgreen,

	"ref_models.Steelblue": ref_models.Steelblue,

	"ref_models.TEXT_ANCHOR_CENTER": ref_models.TEXT_ANCHOR_CENTER,

	"ref_models.TEXT_ANCHOR_LEFT": ref_models.TEXT_ANCHOR_LEFT,

	"ref_models.TEXT_ANCHOR_RIGHT": ref_models.TEXT_ANCHOR_RIGHT,

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

	"ref_models.Text.StrokeDashArrayWhenSelected": (ref_models.Text{}).StrokeDashArrayWhenSelected,

	"ref_models.Text.StrokeWidth": (ref_models.Text{}).StrokeWidth,

	"ref_models.Text.Transform": (ref_models.Text{}).Transform,

	"ref_models.Text.X": (ref_models.Text{}).X,

	"ref_models.Text.Y": (ref_models.Text{}).Y,

	"ref_models.TextAnchorType": ref_models.TextAnchorType(""),

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
	__Field__000000_AttributeName := (&models.Field{Name: `AttributeName`}).Stage(stage)
	__Field__000001_CanHaveBottomHandle := (&models.Field{Name: `CanHaveBottomHandle`}).Stage(stage)
	__Field__000002_CanHaveLeftHandle := (&models.Field{Name: `CanHaveLeftHandle`}).Stage(stage)
	__Field__000003_CanHaveRightHandle := (&models.Field{Name: `CanHaveRightHandle`}).Stage(stage)
	__Field__000004_CanHaveTopHandle := (&models.Field{Name: `CanHaveTopHandle`}).Stage(stage)
	__Field__000005_CanMoveHorizontaly := (&models.Field{Name: `CanMoveHorizontaly`}).Stage(stage)
	__Field__000006_CanMoveVerticaly := (&models.Field{Name: `CanMoveVerticaly`}).Stage(stage)
	__Field__000007_Color := (&models.Field{Name: `Color`}).Stage(stage)
	__Field__000008_Content := (&models.Field{Name: `Content`}).Stage(stage)
	__Field__000009_CornerOffsetRatio := (&models.Field{Name: `CornerOffsetRatio`}).Stage(stage)
	__Field__000010_CornerRadius := (&models.Field{Name: `CornerRadius`}).Stage(stage)
	__Field__000011_Display := (&models.Field{Name: `Display`}).Stage(stage)
	__Field__000012_Dur := (&models.Field{Name: `Dur`}).Stage(stage)
	__Field__000013_EndAnchorType := (&models.Field{Name: `EndAnchorType`}).Stage(stage)
	__Field__000014_EndArrowSize := (&models.Field{Name: `EndArrowSize`}).Stage(stage)
	__Field__000015_EndOrientation := (&models.Field{Name: `EndOrientation`}).Stage(stage)
	__Field__000016_EndRatio := (&models.Field{Name: `EndRatio`}).Stage(stage)
	__Field__000017_HasBottomHandle := (&models.Field{Name: `HasBottomHandle`}).Stage(stage)
	__Field__000018_HasLeftHandle := (&models.Field{Name: `HasLeftHandle`}).Stage(stage)
	__Field__000019_HasRightHandle := (&models.Field{Name: `HasRightHandle`}).Stage(stage)
	__Field__000020_HasTopHandle := (&models.Field{Name: `HasTopHandle`}).Stage(stage)
	__Field__000021_Height := (&models.Field{Name: `Height`}).Stage(stage)
	__Field__000022_IsSelectable := (&models.Field{Name: `IsSelectable`}).Stage(stage)
	__Field__000023_IsSelected := (&models.Field{Name: `IsSelected`}).Stage(stage)
	__Field__000024_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000025_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000026_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000027_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000028_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000029_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000030_Points := (&models.Field{Name: `Points`}).Stage(stage)
	__Field__000031_RX := (&models.Field{Name: `RX`}).Stage(stage)
	__Field__000032_RepeatCount := (&models.Field{Name: `RepeatCount`}).Stage(stage)
	__Field__000033_Stroke := (&models.Field{Name: `Stroke`}).Stage(stage)
	__Field__000034_Values := (&models.Field{Name: `Values`}).Stage(stage)
	__Field__000035_Width := (&models.Field{Name: `Width`}).Stage(stage)
	__Field__000036_X := (&models.Field{Name: `X`}).Stage(stage)
	__Field__000037_X_Offset := (&models.Field{Name: `X_Offset`}).Stage(stage)
	__Field__000038_Y := (&models.Field{Name: `Y`}).Stage(stage)
	__Field__000039_Y_Offset := (&models.Field{Name: `Y_Offset`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_AnchoredText := (&models.GongStructShape{Name: `NewDiagram-AnchoredText`}).Stage(stage)
	__GongStructShape__000001_NewDiagram_Animate := (&models.GongStructShape{Name: `NewDiagram-Animate`}).Stage(stage)
	__GongStructShape__000002_NewDiagram_Circle := (&models.GongStructShape{Name: `NewDiagram-Circle`}).Stage(stage)
	__GongStructShape__000003_NewDiagram_Ellipse := (&models.GongStructShape{Name: `NewDiagram-Ellipse`}).Stage(stage)
	__GongStructShape__000004_NewDiagram_Layer := (&models.GongStructShape{Name: `NewDiagram-Layer`}).Stage(stage)
	__GongStructShape__000005_NewDiagram_Line := (&models.GongStructShape{Name: `NewDiagram-Line`}).Stage(stage)
	__GongStructShape__000006_NewDiagram_Link := (&models.GongStructShape{Name: `NewDiagram-Link`}).Stage(stage)
	__GongStructShape__000007_NewDiagram_Path := (&models.GongStructShape{Name: `NewDiagram-Path`}).Stage(stage)
	__GongStructShape__000008_NewDiagram_Point := (&models.GongStructShape{Name: `NewDiagram-Point`}).Stage(stage)
	__GongStructShape__000009_NewDiagram_Polygone := (&models.GongStructShape{Name: `NewDiagram-Polygone`}).Stage(stage)
	__GongStructShape__000010_NewDiagram_Polyline := (&models.GongStructShape{Name: `NewDiagram-Polyline`}).Stage(stage)
	__GongStructShape__000011_NewDiagram_Rect := (&models.GongStructShape{Name: `NewDiagram-Rect`}).Stage(stage)
	__GongStructShape__000012_NewDiagram_RectAnchoredRect := (&models.GongStructShape{Name: `NewDiagram-RectAnchoredRect`}).Stage(stage)
	__GongStructShape__000013_NewDiagram_RectAnchoredText := (&models.GongStructShape{Name: `NewDiagram-RectAnchoredText`}).Stage(stage)
	__GongStructShape__000014_NewDiagram_RectLinkLink := (&models.GongStructShape{Name: `NewDiagram-RectLinkLink`}).Stage(stage)
	__GongStructShape__000015_NewDiagram_SVG := (&models.GongStructShape{Name: `NewDiagram-SVG`}).Stage(stage)
	__GongStructShape__000016_NewDiagram_Text := (&models.GongStructShape{Name: `NewDiagram-Text`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Circles := (&models.Link{Name: `Circles`}).Stage(stage)
	__Link__000001_Ellipses := (&models.Link{Name: `Ellipses`}).Stage(stage)
	__Link__000002_End := (&models.Link{Name: `End`}).Stage(stage)
	__Link__000003_EndRect := (&models.Link{Name: `EndRect`}).Stage(stage)
	__Link__000004_Layers := (&models.Link{Name: `Layers`}).Stage(stage)
	__Link__000005_Lines := (&models.Link{Name: `Lines`}).Stage(stage)
	__Link__000006_Links := (&models.Link{Name: `Links`}).Stage(stage)
	__Link__000007_Paths := (&models.Link{Name: `Paths`}).Stage(stage)
	__Link__000008_Polygones := (&models.Link{Name: `Polygones`}).Stage(stage)
	__Link__000009_Polylines := (&models.Link{Name: `Polylines`}).Stage(stage)
	__Link__000010_RectAnchoredRects := (&models.Link{Name: `RectAnchoredRects`}).Stage(stage)
	__Link__000011_RectAnchoredTexts := (&models.Link{Name: `RectAnchoredTexts`}).Stage(stage)
	__Link__000012_RectLinkLinks := (&models.Link{Name: `RectLinkLinks`}).Stage(stage)
	__Link__000013_Rects := (&models.Link{Name: `Rects`}).Stage(stage)
	__Link__000014_Start := (&models.Link{Name: `Start`}).Stage(stage)
	__Link__000015_StartRect := (&models.Link{Name: `StartRect`}).Stage(stage)
	__Link__000016_Texts := (&models.Link{Name: `Texts`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_AnchoredText := (&models.Position{Name: `Pos-NewDiagram-AnchoredText`}).Stage(stage)
	__Position__000001_Pos_NewDiagram_Animate := (&models.Position{Name: `Pos-NewDiagram-Animate`}).Stage(stage)
	__Position__000002_Pos_NewDiagram_Circle := (&models.Position{Name: `Pos-NewDiagram-Circle`}).Stage(stage)
	__Position__000003_Pos_NewDiagram_Ellipse := (&models.Position{Name: `Pos-NewDiagram-Ellipse`}).Stage(stage)
	__Position__000004_Pos_NewDiagram_Layer := (&models.Position{Name: `Pos-NewDiagram-Layer`}).Stage(stage)
	__Position__000005_Pos_NewDiagram_Line := (&models.Position{Name: `Pos-NewDiagram-Line`}).Stage(stage)
	__Position__000006_Pos_NewDiagram_Link := (&models.Position{Name: `Pos-NewDiagram-Link`}).Stage(stage)
	__Position__000007_Pos_NewDiagram_Path := (&models.Position{Name: `Pos-NewDiagram-Path`}).Stage(stage)
	__Position__000008_Pos_NewDiagram_Point := (&models.Position{Name: `Pos-NewDiagram-Point`}).Stage(stage)
	__Position__000009_Pos_NewDiagram_Polygone := (&models.Position{Name: `Pos-NewDiagram-Polygone`}).Stage(stage)
	__Position__000010_Pos_NewDiagram_Polyline := (&models.Position{Name: `Pos-NewDiagram-Polyline`}).Stage(stage)
	__Position__000011_Pos_NewDiagram_Rect := (&models.Position{Name: `Pos-NewDiagram-Rect`}).Stage(stage)
	__Position__000012_Pos_NewDiagram_RectAnchoredRect := (&models.Position{Name: `Pos-NewDiagram-RectAnchoredRect`}).Stage(stage)
	__Position__000013_Pos_NewDiagram_RectAnchoredText := (&models.Position{Name: `Pos-NewDiagram-RectAnchoredText`}).Stage(stage)
	__Position__000014_Pos_NewDiagram_RectLinkLink := (&models.Position{Name: `Pos-NewDiagram-RectLinkLink`}).Stage(stage)
	__Position__000015_Pos_NewDiagram_SVG := (&models.Position{Name: `Pos-NewDiagram-SVG`}).Stage(stage)
	__Position__000016_Pos_NewDiagram_Text := (&models.Position{Name: `Pos-NewDiagram-Text`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Circle := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Circle`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Ellipse := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Ellipse`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Line := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Line`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Link`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Path := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Path`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polygone := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Polygone`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polyline := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Polyline`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-RectLinkLink`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Text := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Text`}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000012_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredRect`}).Stage(stage)
	__Vertice__000013_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredText`}).Stage(stage)
	__Vertice__000014_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Layer`}).Stage(stage)
	__Vertice__000015_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Rect`}).Stage(stage)
	__Vertice__000016_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect := (&models.Vertice{Name: `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Rect`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_AttributeName.Name = `AttributeName`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.AttributeName]
	__Field__000000_AttributeName.Identifier = `ref_models.Animate.AttributeName`
	__Field__000000_AttributeName.FieldTypeAsString = ``
	__Field__000000_AttributeName.Structname = `Animate`
	__Field__000000_AttributeName.Fieldtypename = `string`

	// Field values setup
	__Field__000001_CanHaveBottomHandle.Name = `CanHaveBottomHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveBottomHandle]
	__Field__000001_CanHaveBottomHandle.Identifier = `ref_models.Rect.CanHaveBottomHandle`
	__Field__000001_CanHaveBottomHandle.FieldTypeAsString = ``
	__Field__000001_CanHaveBottomHandle.Structname = `Rect`
	__Field__000001_CanHaveBottomHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000002_CanHaveLeftHandle.Name = `CanHaveLeftHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveLeftHandle]
	__Field__000002_CanHaveLeftHandle.Identifier = `ref_models.Rect.CanHaveLeftHandle`
	__Field__000002_CanHaveLeftHandle.FieldTypeAsString = ``
	__Field__000002_CanHaveLeftHandle.Structname = `Rect`
	__Field__000002_CanHaveLeftHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_CanHaveRightHandle.Name = `CanHaveRightHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveRightHandle]
	__Field__000003_CanHaveRightHandle.Identifier = `ref_models.Rect.CanHaveRightHandle`
	__Field__000003_CanHaveRightHandle.FieldTypeAsString = ``
	__Field__000003_CanHaveRightHandle.Structname = `Rect`
	__Field__000003_CanHaveRightHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000004_CanHaveTopHandle.Name = `CanHaveTopHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanHaveTopHandle]
	__Field__000004_CanHaveTopHandle.Identifier = `ref_models.Rect.CanHaveTopHandle`
	__Field__000004_CanHaveTopHandle.FieldTypeAsString = ``
	__Field__000004_CanHaveTopHandle.Structname = `Rect`
	__Field__000004_CanHaveTopHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000005_CanMoveHorizontaly.Name = `CanMoveHorizontaly`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanMoveHorizontaly]
	__Field__000005_CanMoveHorizontaly.Identifier = `ref_models.Rect.CanMoveHorizontaly`
	__Field__000005_CanMoveHorizontaly.FieldTypeAsString = ``
	__Field__000005_CanMoveHorizontaly.Structname = `Rect`
	__Field__000005_CanMoveHorizontaly.Fieldtypename = `bool`

	// Field values setup
	__Field__000006_CanMoveVerticaly.Name = `CanMoveVerticaly`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.CanMoveVerticaly]
	__Field__000006_CanMoveVerticaly.Identifier = `ref_models.Rect.CanMoveVerticaly`
	__Field__000006_CanMoveVerticaly.FieldTypeAsString = ``
	__Field__000006_CanMoveVerticaly.Structname = `Rect`
	__Field__000006_CanMoveVerticaly.Fieldtypename = `bool`

	// Field values setup
	__Field__000007_Color.Name = `Color`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AnchoredText.Color]
	__Field__000007_Color.Identifier = `ref_models.AnchoredText.Color`
	__Field__000007_Color.FieldTypeAsString = ``
	__Field__000007_Color.Structname = `AnchoredText`
	__Field__000007_Color.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Content.Name = `Content`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AnchoredText.Content]
	__Field__000008_Content.Identifier = `ref_models.AnchoredText.Content`
	__Field__000008_Content.FieldTypeAsString = ``
	__Field__000008_Content.Structname = `AnchoredText`
	__Field__000008_Content.Fieldtypename = `string`

	// Field values setup
	__Field__000009_CornerOffsetRatio.Name = `CornerOffsetRatio`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.CornerOffsetRatio]
	__Field__000009_CornerOffsetRatio.Identifier = `ref_models.Link.CornerOffsetRatio`
	__Field__000009_CornerOffsetRatio.FieldTypeAsString = ``
	__Field__000009_CornerOffsetRatio.Structname = `Link`
	__Field__000009_CornerOffsetRatio.Fieldtypename = `float64`

	// Field values setup
	__Field__000010_CornerRadius.Name = `CornerRadius`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.CornerRadius]
	__Field__000010_CornerRadius.Identifier = `ref_models.Link.CornerRadius`
	__Field__000010_CornerRadius.FieldTypeAsString = ``
	__Field__000010_CornerRadius.Structname = `Link`
	__Field__000010_CornerRadius.Fieldtypename = `float64`

	// Field values setup
	__Field__000011_Display.Name = `Display`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Display]
	__Field__000011_Display.Identifier = `ref_models.Layer.Display`
	__Field__000011_Display.FieldTypeAsString = ``
	__Field__000011_Display.Structname = `Layer`
	__Field__000011_Display.Fieldtypename = `bool`

	// Field values setup
	__Field__000012_Dur.Name = `Dur`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.Dur]
	__Field__000012_Dur.Identifier = `ref_models.Animate.Dur`
	__Field__000012_Dur.FieldTypeAsString = ``
	__Field__000012_Dur.Structname = `Animate`
	__Field__000012_Dur.Fieldtypename = `string`

	// Field values setup
	__Field__000013_EndAnchorType.Name = `EndAnchorType`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndAnchorType]
	__Field__000013_EndAnchorType.Identifier = `ref_models.Link.EndAnchorType`
	__Field__000013_EndAnchorType.FieldTypeAsString = ``
	__Field__000013_EndAnchorType.Structname = `Link`
	__Field__000013_EndAnchorType.Fieldtypename = `AnchorType`

	// Field values setup
	__Field__000014_EndArrowSize.Name = `EndArrowSize`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndArrowSize]
	__Field__000014_EndArrowSize.Identifier = `ref_models.Link.EndArrowSize`
	__Field__000014_EndArrowSize.FieldTypeAsString = ``
	__Field__000014_EndArrowSize.Structname = `Link`
	__Field__000014_EndArrowSize.Fieldtypename = `float64`

	// Field values setup
	__Field__000015_EndOrientation.Name = `EndOrientation`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndOrientation]
	__Field__000015_EndOrientation.Identifier = `ref_models.Link.EndOrientation`
	__Field__000015_EndOrientation.FieldTypeAsString = ``
	__Field__000015_EndOrientation.Structname = `Link`
	__Field__000015_EndOrientation.Fieldtypename = `OrientationType`

	// Field values setup
	__Field__000016_EndRatio.Name = `EndRatio`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.EndRatio]
	__Field__000016_EndRatio.Identifier = `ref_models.Link.EndRatio`
	__Field__000016_EndRatio.FieldTypeAsString = ``
	__Field__000016_EndRatio.Structname = `Link`
	__Field__000016_EndRatio.Fieldtypename = `float64`

	// Field values setup
	__Field__000017_HasBottomHandle.Name = `HasBottomHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasBottomHandle]
	__Field__000017_HasBottomHandle.Identifier = `ref_models.Rect.HasBottomHandle`
	__Field__000017_HasBottomHandle.FieldTypeAsString = ``
	__Field__000017_HasBottomHandle.Structname = `Rect`
	__Field__000017_HasBottomHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000018_HasLeftHandle.Name = `HasLeftHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasLeftHandle]
	__Field__000018_HasLeftHandle.Identifier = `ref_models.Rect.HasLeftHandle`
	__Field__000018_HasLeftHandle.FieldTypeAsString = ``
	__Field__000018_HasLeftHandle.Structname = `Rect`
	__Field__000018_HasLeftHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000019_HasRightHandle.Name = `HasRightHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasRightHandle]
	__Field__000019_HasRightHandle.Identifier = `ref_models.Rect.HasRightHandle`
	__Field__000019_HasRightHandle.FieldTypeAsString = ``
	__Field__000019_HasRightHandle.Structname = `Rect`
	__Field__000019_HasRightHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000020_HasTopHandle.Name = `HasTopHandle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.HasTopHandle]
	__Field__000020_HasTopHandle.Identifier = `ref_models.Rect.HasTopHandle`
	__Field__000020_HasTopHandle.FieldTypeAsString = ``
	__Field__000020_HasTopHandle.Structname = `Rect`
	__Field__000020_HasTopHandle.Fieldtypename = `bool`

	// Field values setup
	__Field__000021_Height.Name = `Height`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Height]
	__Field__000021_Height.Identifier = `ref_models.Rect.Height`
	__Field__000021_Height.FieldTypeAsString = ``
	__Field__000021_Height.Structname = `Rect`
	__Field__000021_Height.Fieldtypename = `float64`

	// Field values setup
	__Field__000022_IsSelectable.Name = `IsSelectable`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.IsSelectable]
	__Field__000022_IsSelectable.Identifier = `ref_models.Rect.IsSelectable`
	__Field__000022_IsSelectable.FieldTypeAsString = ``
	__Field__000022_IsSelectable.Structname = `Rect`
	__Field__000022_IsSelectable.Fieldtypename = `bool`

	// Field values setup
	__Field__000023_IsSelected.Name = `IsSelected`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.IsSelected]
	__Field__000023_IsSelected.Identifier = `ref_models.Rect.IsSelected`
	__Field__000023_IsSelected.FieldTypeAsString = ``
	__Field__000023_IsSelected.Structname = `Rect`
	__Field__000023_IsSelected.Fieldtypename = `bool`

	// Field values setup
	__Field__000024_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.Name]
	__Field__000024_Name.Identifier = `ref_models.Animate.Name`
	__Field__000024_Name.FieldTypeAsString = ``
	__Field__000024_Name.Structname = `Animate`
	__Field__000024_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000025_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polyline.Name]
	__Field__000025_Name.Identifier = `ref_models.Polyline.Name`
	__Field__000025_Name.FieldTypeAsString = ``
	__Field__000025_Name.Structname = `Polyline`
	__Field__000025_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000026_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AnchoredText.Name]
	__Field__000026_Name.Identifier = `ref_models.AnchoredText.Name`
	__Field__000026_Name.FieldTypeAsString = ``
	__Field__000026_Name.Structname = `AnchoredText`
	__Field__000026_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000027_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Name]
	__Field__000027_Name.Identifier = `ref_models.Layer.Name`
	__Field__000027_Name.FieldTypeAsString = ``
	__Field__000027_Name.Structname = `Layer`
	__Field__000027_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000028_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.Name]
	__Field__000028_Name.Identifier = `ref_models.SVG.Name`
	__Field__000028_Name.FieldTypeAsString = ``
	__Field__000028_Name.Structname = `SVG`
	__Field__000028_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000029_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Name]
	__Field__000029_Name.Identifier = `ref_models.Rect.Name`
	__Field__000029_Name.FieldTypeAsString = ``
	__Field__000029_Name.Structname = `Rect`
	__Field__000029_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000030_Points.Name = `Points`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polyline.Points]
	__Field__000030_Points.Identifier = `ref_models.Polyline.Points`
	__Field__000030_Points.FieldTypeAsString = ``
	__Field__000030_Points.Structname = `Polyline`
	__Field__000030_Points.Fieldtypename = `string`

	// Field values setup
	__Field__000031_RX.Name = `RX`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.RX]
	__Field__000031_RX.Identifier = `ref_models.Rect.RX`
	__Field__000031_RX.FieldTypeAsString = ``
	__Field__000031_RX.Structname = `Rect`
	__Field__000031_RX.Fieldtypename = `float64`

	// Field values setup
	__Field__000032_RepeatCount.Name = `RepeatCount`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.RepeatCount]
	__Field__000032_RepeatCount.Identifier = `ref_models.Animate.RepeatCount`
	__Field__000032_RepeatCount.FieldTypeAsString = ``
	__Field__000032_RepeatCount.Structname = `Animate`
	__Field__000032_RepeatCount.Fieldtypename = `string`

	// Field values setup
	__Field__000033_Stroke.Name = `Stroke`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AnchoredText.Stroke]
	__Field__000033_Stroke.Identifier = `ref_models.AnchoredText.Stroke`
	__Field__000033_Stroke.FieldTypeAsString = ``
	__Field__000033_Stroke.Structname = `AnchoredText`
	__Field__000033_Stroke.Fieldtypename = `string`

	// Field values setup
	__Field__000034_Values.Name = `Values`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate.Values]
	__Field__000034_Values.Identifier = `ref_models.Animate.Values`
	__Field__000034_Values.FieldTypeAsString = ``
	__Field__000034_Values.Structname = `Animate`
	__Field__000034_Values.Fieldtypename = `string`

	// Field values setup
	__Field__000035_Width.Name = `Width`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Width]
	__Field__000035_Width.Identifier = `ref_models.Rect.Width`
	__Field__000035_Width.FieldTypeAsString = ``
	__Field__000035_Width.Structname = `Rect`
	__Field__000035_Width.Fieldtypename = `float64`

	// Field values setup
	__Field__000036_X.Name = `X`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.X]
	__Field__000036_X.Identifier = `ref_models.Rect.X`
	__Field__000036_X.FieldTypeAsString = ``
	__Field__000036_X.Structname = `Rect`
	__Field__000036_X.Fieldtypename = `float64`

	// Field values setup
	__Field__000037_X_Offset.Name = `X_Offset`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AnchoredText.X_Offset]
	__Field__000037_X_Offset.Identifier = `ref_models.AnchoredText.X_Offset`
	__Field__000037_X_Offset.FieldTypeAsString = ``
	__Field__000037_X_Offset.Structname = `AnchoredText`
	__Field__000037_X_Offset.Fieldtypename = `float64`

	// Field values setup
	__Field__000038_Y.Name = `Y`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.Y]
	__Field__000038_Y.Identifier = `ref_models.Rect.Y`
	__Field__000038_Y.FieldTypeAsString = ``
	__Field__000038_Y.Structname = `Rect`
	__Field__000038_Y.Fieldtypename = `float64`

	// Field values setup
	__Field__000039_Y_Offset.Name = `Y_Offset`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AnchoredText.Y_Offset]
	__Field__000039_Y_Offset.Identifier = `ref_models.AnchoredText.Y_Offset`
	__Field__000039_Y_Offset.FieldTypeAsString = ``
	__Field__000039_Y_Offset.Structname = `AnchoredText`
	__Field__000039_Y_Offset.Fieldtypename = `float64`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_AnchoredText.Name = `NewDiagram-AnchoredText`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.AnchoredText]
	__GongStructShape__000000_NewDiagram_AnchoredText.Identifier = `ref_models.AnchoredText`
	__GongStructShape__000000_NewDiagram_AnchoredText.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_AnchoredText.NbInstances = 0
	__GongStructShape__000000_NewDiagram_AnchoredText.Width = 240.000000
	__GongStructShape__000000_NewDiagram_AnchoredText.Heigth = 153.000000
	__GongStructShape__000000_NewDiagram_AnchoredText.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Animate.Name = `NewDiagram-Animate`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Animate]
	__GongStructShape__000001_NewDiagram_Animate.Identifier = `ref_models.Animate`
	__GongStructShape__000001_NewDiagram_Animate.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Animate.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Animate.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Animate.Heigth = 138.000000
	__GongStructShape__000001_NewDiagram_Animate.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_Circle.Name = `NewDiagram-Circle`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle]
	__GongStructShape__000002_NewDiagram_Circle.Identifier = `ref_models.Circle`
	__GongStructShape__000002_NewDiagram_Circle.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_Circle.NbInstances = 0
	__GongStructShape__000002_NewDiagram_Circle.Width = 240.000000
	__GongStructShape__000002_NewDiagram_Circle.Heigth = 63.000000
	__GongStructShape__000002_NewDiagram_Circle.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_Ellipse.Name = `NewDiagram-Ellipse`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Ellipse]
	__GongStructShape__000003_NewDiagram_Ellipse.Identifier = `ref_models.Ellipse`
	__GongStructShape__000003_NewDiagram_Ellipse.ShowNbInstances = false
	__GongStructShape__000003_NewDiagram_Ellipse.NbInstances = 0
	__GongStructShape__000003_NewDiagram_Ellipse.Width = 240.000000
	__GongStructShape__000003_NewDiagram_Ellipse.Heigth = 63.000000
	__GongStructShape__000003_NewDiagram_Ellipse.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000004_NewDiagram_Layer.Name = `NewDiagram-Layer`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__GongStructShape__000004_NewDiagram_Layer.Identifier = `ref_models.Layer`
	__GongStructShape__000004_NewDiagram_Layer.ShowNbInstances = false
	__GongStructShape__000004_NewDiagram_Layer.NbInstances = 0
	__GongStructShape__000004_NewDiagram_Layer.Width = 240.000000
	__GongStructShape__000004_NewDiagram_Layer.Heigth = 93.000000
	__GongStructShape__000004_NewDiagram_Layer.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000005_NewDiagram_Line.Name = `NewDiagram-Line`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__GongStructShape__000005_NewDiagram_Line.Identifier = `ref_models.Line`
	__GongStructShape__000005_NewDiagram_Line.ShowNbInstances = false
	__GongStructShape__000005_NewDiagram_Line.NbInstances = 0
	__GongStructShape__000005_NewDiagram_Line.Width = 240.000000
	__GongStructShape__000005_NewDiagram_Line.Heigth = 63.000000
	__GongStructShape__000005_NewDiagram_Line.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000006_NewDiagram_Link.Name = `NewDiagram-Link`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__GongStructShape__000006_NewDiagram_Link.Identifier = `ref_models.Link`
	__GongStructShape__000006_NewDiagram_Link.ShowNbInstances = false
	__GongStructShape__000006_NewDiagram_Link.NbInstances = 0
	__GongStructShape__000006_NewDiagram_Link.Width = 240.000000
	__GongStructShape__000006_NewDiagram_Link.Heigth = 153.000000
	__GongStructShape__000006_NewDiagram_Link.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000007_NewDiagram_Path.Name = `NewDiagram-Path`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Path]
	__GongStructShape__000007_NewDiagram_Path.Identifier = `ref_models.Path`
	__GongStructShape__000007_NewDiagram_Path.ShowNbInstances = false
	__GongStructShape__000007_NewDiagram_Path.NbInstances = 0
	__GongStructShape__000007_NewDiagram_Path.Width = 240.000000
	__GongStructShape__000007_NewDiagram_Path.Heigth = 63.000000
	__GongStructShape__000007_NewDiagram_Path.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000008_NewDiagram_Point.Name = `NewDiagram-Point`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Point]
	__GongStructShape__000008_NewDiagram_Point.Identifier = `ref_models.Point`
	__GongStructShape__000008_NewDiagram_Point.ShowNbInstances = false
	__GongStructShape__000008_NewDiagram_Point.NbInstances = 0
	__GongStructShape__000008_NewDiagram_Point.Width = 240.000000
	__GongStructShape__000008_NewDiagram_Point.Heigth = 63.000000
	__GongStructShape__000008_NewDiagram_Point.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000009_NewDiagram_Polygone.Name = `NewDiagram-Polygone`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polygone]
	__GongStructShape__000009_NewDiagram_Polygone.Identifier = `ref_models.Polygone`
	__GongStructShape__000009_NewDiagram_Polygone.ShowNbInstances = false
	__GongStructShape__000009_NewDiagram_Polygone.NbInstances = 0
	__GongStructShape__000009_NewDiagram_Polygone.Width = 240.000000
	__GongStructShape__000009_NewDiagram_Polygone.Heigth = 63.000000
	__GongStructShape__000009_NewDiagram_Polygone.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000010_NewDiagram_Polyline.Name = `NewDiagram-Polyline`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polyline]
	__GongStructShape__000010_NewDiagram_Polyline.Identifier = `ref_models.Polyline`
	__GongStructShape__000010_NewDiagram_Polyline.ShowNbInstances = false
	__GongStructShape__000010_NewDiagram_Polyline.NbInstances = 0
	__GongStructShape__000010_NewDiagram_Polyline.Width = 240.000000
	__GongStructShape__000010_NewDiagram_Polyline.Heigth = 93.000000
	__GongStructShape__000010_NewDiagram_Polyline.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000011_NewDiagram_Rect.Name = `NewDiagram-Rect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__GongStructShape__000011_NewDiagram_Rect.Identifier = `ref_models.Rect`
	__GongStructShape__000011_NewDiagram_Rect.ShowNbInstances = false
	__GongStructShape__000011_NewDiagram_Rect.NbInstances = 0
	__GongStructShape__000011_NewDiagram_Rect.Width = 240.000000
	__GongStructShape__000011_NewDiagram_Rect.Heigth = 333.000000
	__GongStructShape__000011_NewDiagram_Rect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.Name = `NewDiagram-RectAnchoredRect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredRect]
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.Identifier = `ref_models.RectAnchoredRect`
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.ShowNbInstances = false
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.NbInstances = 0
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.Width = 240.000000
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.Heigth = 63.000000
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000013_NewDiagram_RectAnchoredText.Name = `NewDiagram-RectAnchoredText`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredText]
	__GongStructShape__000013_NewDiagram_RectAnchoredText.Identifier = `ref_models.RectAnchoredText`
	__GongStructShape__000013_NewDiagram_RectAnchoredText.ShowNbInstances = false
	__GongStructShape__000013_NewDiagram_RectAnchoredText.NbInstances = 0
	__GongStructShape__000013_NewDiagram_RectAnchoredText.Width = 240.000000
	__GongStructShape__000013_NewDiagram_RectAnchoredText.Heigth = 63.000000
	__GongStructShape__000013_NewDiagram_RectAnchoredText.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000014_NewDiagram_RectLinkLink.Name = `NewDiagram-RectLinkLink`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink]
	__GongStructShape__000014_NewDiagram_RectLinkLink.Identifier = `ref_models.RectLinkLink`
	__GongStructShape__000014_NewDiagram_RectLinkLink.ShowNbInstances = false
	__GongStructShape__000014_NewDiagram_RectLinkLink.NbInstances = 0
	__GongStructShape__000014_NewDiagram_RectLinkLink.Width = 240.000000
	__GongStructShape__000014_NewDiagram_RectLinkLink.Heigth = 63.000000
	__GongStructShape__000014_NewDiagram_RectLinkLink.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000015_NewDiagram_SVG.Name = `NewDiagram-SVG`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG]
	__GongStructShape__000015_NewDiagram_SVG.Identifier = `ref_models.SVG`
	__GongStructShape__000015_NewDiagram_SVG.ShowNbInstances = false
	__GongStructShape__000015_NewDiagram_SVG.NbInstances = 0
	__GongStructShape__000015_NewDiagram_SVG.Width = 240.000000
	__GongStructShape__000015_NewDiagram_SVG.Heigth = 78.000000
	__GongStructShape__000015_NewDiagram_SVG.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000016_NewDiagram_Text.Name = `NewDiagram-Text`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__GongStructShape__000016_NewDiagram_Text.Identifier = `ref_models.Text`
	__GongStructShape__000016_NewDiagram_Text.ShowNbInstances = false
	__GongStructShape__000016_NewDiagram_Text.NbInstances = 0
	__GongStructShape__000016_NewDiagram_Text.Width = 240.000000
	__GongStructShape__000016_NewDiagram_Text.Heigth = 63.000000
	__GongStructShape__000016_NewDiagram_Text.IsSelected = false

	// Link values setup
	__Link__000000_Circles.Name = `Circles`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Circles]
	__Link__000000_Circles.Identifier = `ref_models.Layer.Circles`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Circle]
	__Link__000000_Circles.Fieldtypename = `ref_models.Circle`
	__Link__000000_Circles.TargetMultiplicity = models.MANY
	__Link__000000_Circles.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000001_Ellipses.Name = `Ellipses`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Ellipses]
	__Link__000001_Ellipses.Identifier = `ref_models.Layer.Ellipses`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Ellipse]
	__Link__000001_Ellipses.Fieldtypename = `ref_models.Ellipse`
	__Link__000001_Ellipses.TargetMultiplicity = models.MANY
	__Link__000001_Ellipses.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000002_End.Name = `End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.End]
	__Link__000002_End.Identifier = `ref_models.Link.End`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000002_End.Fieldtypename = `ref_models.Rect`
	__Link__000002_End.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_End.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000003_EndRect.Name = `EndRect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.EndRect]
	__Link__000003_EndRect.Identifier = `ref_models.SVG.EndRect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000003_EndRect.Fieldtypename = `ref_models.Rect`
	__Link__000003_EndRect.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_EndRect.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000004_Layers.Name = `Layers`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.Layers]
	__Link__000004_Layers.Identifier = `ref_models.SVG.Layers`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer]
	__Link__000004_Layers.Fieldtypename = `ref_models.Layer`
	__Link__000004_Layers.TargetMultiplicity = models.MANY
	__Link__000004_Layers.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000005_Lines.Name = `Lines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Lines]
	__Link__000005_Lines.Identifier = `ref_models.Layer.Lines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Line]
	__Link__000005_Lines.Fieldtypename = `ref_models.Line`
	__Link__000005_Lines.TargetMultiplicity = models.MANY
	__Link__000005_Lines.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000006_Links.Name = `Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Links]
	__Link__000006_Links.Identifier = `ref_models.Layer.Links`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link]
	__Link__000006_Links.Fieldtypename = `ref_models.Link`
	__Link__000006_Links.TargetMultiplicity = models.MANY
	__Link__000006_Links.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000007_Paths.Name = `Paths`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Paths]
	__Link__000007_Paths.Identifier = `ref_models.Layer.Paths`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Path]
	__Link__000007_Paths.Fieldtypename = `ref_models.Path`
	__Link__000007_Paths.TargetMultiplicity = models.MANY
	__Link__000007_Paths.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000008_Polygones.Name = `Polygones`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Polygones]
	__Link__000008_Polygones.Identifier = `ref_models.Layer.Polygones`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polygone]
	__Link__000008_Polygones.Fieldtypename = `ref_models.Polygone`
	__Link__000008_Polygones.TargetMultiplicity = models.MANY
	__Link__000008_Polygones.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000009_Polylines.Name = `Polylines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Polylines]
	__Link__000009_Polylines.Identifier = `ref_models.Layer.Polylines`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Polyline]
	__Link__000009_Polylines.Fieldtypename = `ref_models.Polyline`
	__Link__000009_Polylines.TargetMultiplicity = models.MANY
	__Link__000009_Polylines.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000010_RectAnchoredRects.Name = `RectAnchoredRects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.RectAnchoredRects]
	__Link__000010_RectAnchoredRects.Identifier = `ref_models.Rect.RectAnchoredRects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredRect]
	__Link__000010_RectAnchoredRects.Fieldtypename = `ref_models.RectAnchoredRect`
	__Link__000010_RectAnchoredRects.TargetMultiplicity = models.MANY
	__Link__000010_RectAnchoredRects.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000011_RectAnchoredTexts.Name = `RectAnchoredTexts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect.RectAnchoredTexts]
	__Link__000011_RectAnchoredTexts.Identifier = `ref_models.Rect.RectAnchoredTexts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectAnchoredText]
	__Link__000011_RectAnchoredTexts.Fieldtypename = `ref_models.RectAnchoredText`
	__Link__000011_RectAnchoredTexts.TargetMultiplicity = models.MANY
	__Link__000011_RectAnchoredTexts.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000012_RectLinkLinks.Name = `RectLinkLinks`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.RectLinkLinks]
	__Link__000012_RectLinkLinks.Identifier = `ref_models.Layer.RectLinkLinks`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.RectLinkLink]
	__Link__000012_RectLinkLinks.Fieldtypename = `ref_models.RectLinkLink`
	__Link__000012_RectLinkLinks.TargetMultiplicity = models.MANY
	__Link__000012_RectLinkLinks.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000013_Rects.Name = `Rects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Rects]
	__Link__000013_Rects.Identifier = `ref_models.Layer.Rects`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000013_Rects.Fieldtypename = `ref_models.Rect`
	__Link__000013_Rects.TargetMultiplicity = models.MANY
	__Link__000013_Rects.SourceMultiplicity = models.ZERO_ONE

	// Link values setup
	__Link__000014_Start.Name = `Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Link.Start]
	__Link__000014_Start.Identifier = `ref_models.Link.Start`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000014_Start.Fieldtypename = `ref_models.Rect`
	__Link__000014_Start.TargetMultiplicity = models.ZERO_ONE
	__Link__000014_Start.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000015_StartRect.Name = `StartRect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.SVG.StartRect]
	__Link__000015_StartRect.Identifier = `ref_models.SVG.StartRect`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Rect]
	__Link__000015_StartRect.Fieldtypename = `ref_models.Rect`
	__Link__000015_StartRect.TargetMultiplicity = models.ZERO_ONE
	__Link__000015_StartRect.SourceMultiplicity = models.MANY

	// Link values setup
	__Link__000016_Texts.Name = `Texts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Layer.Texts]
	__Link__000016_Texts.Identifier = `ref_models.Layer.Texts`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Text]
	__Link__000016_Texts.Fieldtypename = `ref_models.Text`
	__Link__000016_Texts.TargetMultiplicity = models.MANY
	__Link__000016_Texts.SourceMultiplicity = models.ZERO_ONE

	// Position values setup
	__Position__000000_Pos_NewDiagram_AnchoredText.X = 30.000000
	__Position__000000_Pos_NewDiagram_AnchoredText.Y = 750.000000
	__Position__000000_Pos_NewDiagram_AnchoredText.Name = `Pos-NewDiagram-AnchoredText`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Animate.X = 1100.000000
	__Position__000001_Pos_NewDiagram_Animate.Y = 310.000000
	__Position__000001_Pos_NewDiagram_Animate.Name = `Pos-NewDiagram-Animate`

	// Position values setup
	__Position__000002_Pos_NewDiagram_Circle.X = 560.000000
	__Position__000002_Pos_NewDiagram_Circle.Y = 420.000000
	__Position__000002_Pos_NewDiagram_Circle.Name = `Pos-NewDiagram-Circle`

	// Position values setup
	__Position__000003_Pos_NewDiagram_Ellipse.X = 550.000000
	__Position__000003_Pos_NewDiagram_Ellipse.Y = 320.000000
	__Position__000003_Pos_NewDiagram_Ellipse.Name = `Pos-NewDiagram-Ellipse`

	// Position values setup
	__Position__000004_Pos_NewDiagram_Layer.X = 80.000000
	__Position__000004_Pos_NewDiagram_Layer.Y = 100.000000
	__Position__000004_Pos_NewDiagram_Layer.Name = `Pos-NewDiagram-Layer`

	// Position values setup
	__Position__000005_Pos_NewDiagram_Line.X = 550.000000
	__Position__000005_Pos_NewDiagram_Line.Y = 170.000000
	__Position__000005_Pos_NewDiagram_Line.Name = `Pos-NewDiagram-Line`

	// Position values setup
	__Position__000006_Pos_NewDiagram_Link.X = 560.000000
	__Position__000006_Pos_NewDiagram_Link.Y = 490.000000
	__Position__000006_Pos_NewDiagram_Link.Name = `Pos-NewDiagram-Link`

	// Position values setup
	__Position__000007_Pos_NewDiagram_Path.X = 540.000000
	__Position__000007_Pos_NewDiagram_Path.Y = 250.000000
	__Position__000007_Pos_NewDiagram_Path.Name = `Pos-NewDiagram-Path`

	// Position values setup
	__Position__000008_Pos_NewDiagram_Point.X = 970.000000
	__Position__000008_Pos_NewDiagram_Point.Y = 190.000000
	__Position__000008_Pos_NewDiagram_Point.Name = `Pos-NewDiagram-Point`

	// Position values setup
	__Position__000009_Pos_NewDiagram_Polygone.X = 540.000000
	__Position__000009_Pos_NewDiagram_Polygone.Y = 20.000000
	__Position__000009_Pos_NewDiagram_Polygone.Name = `Pos-NewDiagram-Polygone`

	// Position values setup
	__Position__000010_Pos_NewDiagram_Polyline.X = 870.000000
	__Position__000010_Pos_NewDiagram_Polyline.Y = 70.000000
	__Position__000010_Pos_NewDiagram_Polyline.Name = `Pos-NewDiagram-Polyline`

	// Position values setup
	__Position__000011_Pos_NewDiagram_Rect.X = 560.000000
	__Position__000011_Pos_NewDiagram_Rect.Y = 760.000000
	__Position__000011_Pos_NewDiagram_Rect.Name = `Pos-NewDiagram-Rect`

	// Position values setup
	__Position__000012_Pos_NewDiagram_RectAnchoredRect.X = 1110.000000
	__Position__000012_Pos_NewDiagram_RectAnchoredRect.Y = 820.000000
	__Position__000012_Pos_NewDiagram_RectAnchoredRect.Name = `Pos-NewDiagram-RectAnchoredRect`

	// Position values setup
	__Position__000013_Pos_NewDiagram_RectAnchoredText.X = 1090.000000
	__Position__000013_Pos_NewDiagram_RectAnchoredText.Y = 660.000000
	__Position__000013_Pos_NewDiagram_RectAnchoredText.Name = `Pos-NewDiagram-RectAnchoredText`

	// Position values setup
	__Position__000014_Pos_NewDiagram_RectLinkLink.X = 520.000000
	__Position__000014_Pos_NewDiagram_RectLinkLink.Y = 670.000000
	__Position__000014_Pos_NewDiagram_RectLinkLink.Name = `Pos-NewDiagram-RectLinkLink`

	// Position values setup
	__Position__000015_Pos_NewDiagram_SVG.X = 40.000000
	__Position__000015_Pos_NewDiagram_SVG.Y = 520.000000
	__Position__000015_Pos_NewDiagram_SVG.Name = `Pos-NewDiagram-SVG`

	// Position values setup
	__Position__000016_Pos_NewDiagram_Text.X = 40.000000
	__Position__000016_Pos_NewDiagram_Text.Y = 440.000000
	__Position__000016_Pos_NewDiagram_Text.Name = `Pos-NewDiagram-Text`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Circle.X = 361.000000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Circle.Y = 445.500000
	__Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Circle.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Circle`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Ellipse.X = 388.000000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Ellipse.Y = 321.500000
	__Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Ellipse.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Ellipse`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Line.X = 421.000000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Line.Y = 186.500000
	__Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Line.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Line`

	// Vertice values setup
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link.X = 384.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link.Y = 548.500000
	__Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Link`

	// Vertice values setup
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Path.X = 421.500000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Path.Y = 260.500000
	__Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Path.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Path`

	// Vertice values setup
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polygone.X = 446.000000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polygone.Y = 102.500000
	__Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polygone.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Polygone`

	// Vertice values setup
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polyline.X = 679.500000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polyline.Y = 84.000000
	__Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polyline.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Polyline`

	// Vertice values setup
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect.X = 348.000000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect.Y = 684.000000
	__Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink.X = 405.500000
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink.Y = 633.500000
	__Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-RectLinkLink`

	// Vertice values setup
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Text.X = 149.500000
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Text.Y = 317.000000
	__Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Text.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Layer and NewDiagram-Text`

	// Vertice values setup
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.X = 955.000000
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Y = 516.500000
	__Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.X = 965.000000
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Y = 576.500000
	__Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Link and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000012_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect.X = 1025.000000
	__Vertice__000012_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect.Y = 744.000000
	__Vertice__000012_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredRect`

	// Vertice values setup
	__Vertice__000013_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText.X = 965.000000
	__Vertice__000013_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText.Y = 619.000000
	__Vertice__000013_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-Rect and NewDiagram-RectAnchoredText`

	// Vertice values setup
	__Vertice__000014_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer.X = 53.000000
	__Vertice__000014_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer.Y = 284.000000
	__Vertice__000014_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Layer`

	// Vertice values setup
	__Vertice__000015_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.X = 260.000000
	__Vertice__000015_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.Y = 794.000000
	__Vertice__000015_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Rect`

	// Vertice values setup
	__Vertice__000016_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.X = 270.000000
	__Vertice__000016_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.Y = 714.000000
	__Vertice__000016_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect.Name = `Verticle in class diagram NewDiagram in middle between NewDiagram-SVG and NewDiagram-Rect`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_AnchoredText)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Animate)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_Ellipse)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_Circle)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000004_NewDiagram_Layer)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000005_NewDiagram_Line)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000006_NewDiagram_Link)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000007_NewDiagram_Path)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000008_NewDiagram_Point)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000010_NewDiagram_Polyline)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000011_NewDiagram_Rect)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000012_NewDiagram_RectAnchoredRect)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000013_NewDiagram_RectAnchoredText)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000014_NewDiagram_RectLinkLink)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000015_NewDiagram_SVG)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000016_NewDiagram_Text)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000009_NewDiagram_Polygone)
	__GongStructShape__000000_NewDiagram_AnchoredText.Position = __Position__000000_Pos_NewDiagram_AnchoredText
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000026_Name)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000008_Content)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000037_X_Offset)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000039_Y_Offset)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000007_Color)
	__GongStructShape__000000_NewDiagram_AnchoredText.Fields = append(__GongStructShape__000000_NewDiagram_AnchoredText.Fields, __Field__000033_Stroke)
	__GongStructShape__000001_NewDiagram_Animate.Position = __Position__000001_Pos_NewDiagram_Animate
	__GongStructShape__000001_NewDiagram_Animate.Fields = append(__GongStructShape__000001_NewDiagram_Animate.Fields, __Field__000024_Name)
	__GongStructShape__000001_NewDiagram_Animate.Fields = append(__GongStructShape__000001_NewDiagram_Animate.Fields, __Field__000000_AttributeName)
	__GongStructShape__000001_NewDiagram_Animate.Fields = append(__GongStructShape__000001_NewDiagram_Animate.Fields, __Field__000034_Values)
	__GongStructShape__000001_NewDiagram_Animate.Fields = append(__GongStructShape__000001_NewDiagram_Animate.Fields, __Field__000012_Dur)
	__GongStructShape__000001_NewDiagram_Animate.Fields = append(__GongStructShape__000001_NewDiagram_Animate.Fields, __Field__000032_RepeatCount)
	__GongStructShape__000002_NewDiagram_Circle.Position = __Position__000002_Pos_NewDiagram_Circle
	__GongStructShape__000003_NewDiagram_Ellipse.Position = __Position__000003_Pos_NewDiagram_Ellipse
	__GongStructShape__000004_NewDiagram_Layer.Position = __Position__000004_Pos_NewDiagram_Layer
	__GongStructShape__000004_NewDiagram_Layer.Fields = append(__GongStructShape__000004_NewDiagram_Layer.Fields, __Field__000011_Display)
	__GongStructShape__000004_NewDiagram_Layer.Fields = append(__GongStructShape__000004_NewDiagram_Layer.Fields, __Field__000027_Name)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000000_Circles)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000001_Ellipses)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000007_Paths)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000009_Polylines)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000008_Polygones)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000012_RectLinkLinks)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000016_Texts)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000005_Lines)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000006_Links)
	__GongStructShape__000004_NewDiagram_Layer.Links = append(__GongStructShape__000004_NewDiagram_Layer.Links, __Link__000013_Rects)
	__GongStructShape__000005_NewDiagram_Line.Position = __Position__000005_Pos_NewDiagram_Line
	__GongStructShape__000006_NewDiagram_Link.Position = __Position__000006_Pos_NewDiagram_Link
	__GongStructShape__000006_NewDiagram_Link.Fields = append(__GongStructShape__000006_NewDiagram_Link.Fields, __Field__000013_EndAnchorType)
	__GongStructShape__000006_NewDiagram_Link.Fields = append(__GongStructShape__000006_NewDiagram_Link.Fields, __Field__000015_EndOrientation)
	__GongStructShape__000006_NewDiagram_Link.Fields = append(__GongStructShape__000006_NewDiagram_Link.Fields, __Field__000016_EndRatio)
	__GongStructShape__000006_NewDiagram_Link.Fields = append(__GongStructShape__000006_NewDiagram_Link.Fields, __Field__000009_CornerOffsetRatio)
	__GongStructShape__000006_NewDiagram_Link.Fields = append(__GongStructShape__000006_NewDiagram_Link.Fields, __Field__000010_CornerRadius)
	__GongStructShape__000006_NewDiagram_Link.Fields = append(__GongStructShape__000006_NewDiagram_Link.Fields, __Field__000014_EndArrowSize)
	__GongStructShape__000006_NewDiagram_Link.Links = append(__GongStructShape__000006_NewDiagram_Link.Links, __Link__000002_End)
	__GongStructShape__000006_NewDiagram_Link.Links = append(__GongStructShape__000006_NewDiagram_Link.Links, __Link__000014_Start)
	__GongStructShape__000007_NewDiagram_Path.Position = __Position__000007_Pos_NewDiagram_Path
	__GongStructShape__000008_NewDiagram_Point.Position = __Position__000008_Pos_NewDiagram_Point
	__GongStructShape__000009_NewDiagram_Polygone.Position = __Position__000009_Pos_NewDiagram_Polygone
	__GongStructShape__000010_NewDiagram_Polyline.Position = __Position__000010_Pos_NewDiagram_Polyline
	__GongStructShape__000010_NewDiagram_Polyline.Fields = append(__GongStructShape__000010_NewDiagram_Polyline.Fields, __Field__000025_Name)
	__GongStructShape__000010_NewDiagram_Polyline.Fields = append(__GongStructShape__000010_NewDiagram_Polyline.Fields, __Field__000030_Points)
	__GongStructShape__000011_NewDiagram_Rect.Position = __Position__000011_Pos_NewDiagram_Rect
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000029_Name)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000036_X)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000038_Y)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000035_Width)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000021_Height)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000031_RX)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000022_IsSelectable)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000023_IsSelected)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000002_CanHaveLeftHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000018_HasLeftHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000003_CanHaveRightHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000019_HasRightHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000004_CanHaveTopHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000020_HasTopHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000001_CanHaveBottomHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000017_HasBottomHandle)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000005_CanMoveHorizontaly)
	__GongStructShape__000011_NewDiagram_Rect.Fields = append(__GongStructShape__000011_NewDiagram_Rect.Fields, __Field__000006_CanMoveVerticaly)
	__GongStructShape__000011_NewDiagram_Rect.Links = append(__GongStructShape__000011_NewDiagram_Rect.Links, __Link__000010_RectAnchoredRects)
	__GongStructShape__000011_NewDiagram_Rect.Links = append(__GongStructShape__000011_NewDiagram_Rect.Links, __Link__000011_RectAnchoredTexts)
	__GongStructShape__000012_NewDiagram_RectAnchoredRect.Position = __Position__000012_Pos_NewDiagram_RectAnchoredRect
	__GongStructShape__000013_NewDiagram_RectAnchoredText.Position = __Position__000013_Pos_NewDiagram_RectAnchoredText
	__GongStructShape__000014_NewDiagram_RectLinkLink.Position = __Position__000014_Pos_NewDiagram_RectLinkLink
	__GongStructShape__000015_NewDiagram_SVG.Position = __Position__000015_Pos_NewDiagram_SVG
	__GongStructShape__000015_NewDiagram_SVG.Fields = append(__GongStructShape__000015_NewDiagram_SVG.Fields, __Field__000028_Name)
	__GongStructShape__000015_NewDiagram_SVG.Links = append(__GongStructShape__000015_NewDiagram_SVG.Links, __Link__000004_Layers)
	__GongStructShape__000015_NewDiagram_SVG.Links = append(__GongStructShape__000015_NewDiagram_SVG.Links, __Link__000015_StartRect)
	__GongStructShape__000015_NewDiagram_SVG.Links = append(__GongStructShape__000015_NewDiagram_SVG.Links, __Link__000003_EndRect)
	__GongStructShape__000016_NewDiagram_Text.Position = __Position__000016_Pos_NewDiagram_Text
	__Link__000000_Circles.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Circle
	__Link__000001_Ellipses.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Ellipse
	__Link__000002_End.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect
	__Link__000003_EndRect.Middlevertice = __Vertice__000015_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect
	__Link__000004_Layers.Middlevertice = __Vertice__000014_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Layer
	__Link__000005_Lines.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Line
	__Link__000006_Links.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Link
	__Link__000007_Paths.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Path
	__Link__000008_Polygones.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polygone
	__Link__000009_Polylines.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Polyline
	__Link__000010_RectAnchoredRects.Middlevertice = __Vertice__000012_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredRect
	__Link__000011_RectAnchoredTexts.Middlevertice = __Vertice__000013_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Rect_and_NewDiagram_RectAnchoredText
	__Link__000012_RectLinkLinks.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_RectLinkLink
	__Link__000013_Rects.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Rect
	__Link__000014_Start.Middlevertice = __Vertice__000011_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Link_and_NewDiagram_Rect
	__Link__000015_StartRect.Middlevertice = __Vertice__000016_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_SVG_and_NewDiagram_Rect
	__Link__000016_Texts.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_NewDiagram_in_middle_between_NewDiagram_Layer_and_NewDiagram_Text
}


