// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for AnchorType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (anchortype AnchorType) ToString() (res string) {

	// migration of former implementation of enum
	switch anchortype {
	// insertion code per enum code
	case ANCHOR_TOP:
		res = "ANCHOR_TOP"
	case ANCHOR_BOTTOM:
		res = "ANCHOR_BOTTOM"
	case ANCHOR_LEFT:
		res = "ANCHOR_LEFT"
	case ANCHOR_RIGHT:
		res = "ANCHOR_RIGHT"
	case ANCHOR_CENTER:
		res = "ANCHOR_CENTER"
	}
	return
}

func (anchortype *AnchorType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ANCHOR_TOP":
		*anchortype = ANCHOR_TOP
	case "ANCHOR_BOTTOM":
		*anchortype = ANCHOR_BOTTOM
	case "ANCHOR_LEFT":
		*anchortype = ANCHOR_LEFT
	case "ANCHOR_RIGHT":
		*anchortype = ANCHOR_RIGHT
	case "ANCHOR_CENTER":
		*anchortype = ANCHOR_CENTER
	default:
		return errUnkownEnum
	}
	return
}

func (anchortype *AnchorType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ANCHOR_TOP":
		*anchortype = ANCHOR_TOP
	case "ANCHOR_BOTTOM":
		*anchortype = ANCHOR_BOTTOM
	case "ANCHOR_LEFT":
		*anchortype = ANCHOR_LEFT
	case "ANCHOR_RIGHT":
		*anchortype = ANCHOR_RIGHT
	case "ANCHOR_CENTER":
		*anchortype = ANCHOR_CENTER
	default:
		return errUnkownEnum
	}
	return
}

func (anchortype *AnchorType) ToCodeString() (res string) {

	switch *anchortype {
	// insertion code per enum code
	case ANCHOR_TOP:
		res = "ANCHOR_TOP"
	case ANCHOR_BOTTOM:
		res = "ANCHOR_BOTTOM"
	case ANCHOR_LEFT:
		res = "ANCHOR_LEFT"
	case ANCHOR_RIGHT:
		res = "ANCHOR_RIGHT"
	case ANCHOR_CENTER:
		res = "ANCHOR_CENTER"
	}
	return
}

func (anchortype AnchorType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ANCHOR_TOP")
	res = append(res, "ANCHOR_BOTTOM")
	res = append(res, "ANCHOR_LEFT")
	res = append(res, "ANCHOR_RIGHT")
	res = append(res, "ANCHOR_CENTER")

	return
}

func (anchortype AnchorType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ANCHOR_TOP")
	res = append(res, "ANCHOR_BOTTOM")
	res = append(res, "ANCHOR_LEFT")
	res = append(res, "ANCHOR_RIGHT")
	res = append(res, "ANCHOR_CENTER")

	return
}

// Utility function for ColorType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (colortype ColorType) ToString() (res string) {

	// migration of former implementation of enum
	switch colortype {
	// insertion code per enum code
	case Aliceblue:
		res = "aliceblue"
	case Antiquewhite:
		res = "antiquewhite"
	case Aqua:
		res = "aqua"
	case Aquamarine:
		res = "aquamarine "
	case Azure:
		res = "azure"
	case Beige:
		res = "beige"
	case Bisque:
		res = "bisque "
	case Black:
		res = "black"
	case Blanchedalmond:
		res = "blanchedalmond "
	case Blue:
		res = "blue"
	case Blueviolet:
		res = "blueviolet "
	case Brown:
		res = "brown"
	case Burlywood:
		res = "burlywood"
	case Cadetblue:
		res = "cadetblue"
	case Chartreuse:
		res = "chartreuse "
	case Chocolate:
		res = "chocolate"
	case Coral:
		res = "coral"
	case Cornflowerblue:
		res = "cornflowerblue "
	case Cornsilk:
		res = "cornsilk"
	case Crimson:
		res = "crimson"
	case Cyan:
		res = "cyan"
	case Darkblue:
		res = "darkblue"
	case Darkcyan:
		res = "darkcyan"
	case Darkgoldenrod:
		res = "darkgoldenrod"
	case Darkgray:
		res = "darkgray"
	case Darkgreen:
		res = "darkgreen"
	case Darkgrey:
		res = "darkgrey"
	case Darkkhaki:
		res = "darkkhaki"
	case Darkmagenta:
		res = "darkmagenta"
	case Darkolivegreen:
		res = "darkolivegreen "
	case Darkorange:
		res = "darkorange "
	case Darkorchid:
		res = "darkorchid "
	case Darkred:
		res = "darkred"
	case Darksalmon:
		res = "darksalmon "
	case Darkseagreen:
		res = "darkseagreen"
	case Darkslateblue:
		res = "darkslateblue"
	case Darkslategray:
		res = "darkslategray"
	case Darkslategrey:
		res = "darkslategrey"
	case Darkturquoise:
		res = "darkturquoise"
	case Darkviolet:
		res = "darkviolet "
	case Deeppink:
		res = "deeppink"
	case Deepskyblue:
		res = "deepskyblue"
	case Dimgray:
		res = "dimgray"
	case Dimgrey:
		res = "dimgrey"
	case Dodgerblue:
		res = "dodgerblue "
	case Firebrick:
		res = "firebrick"
	case Floralwhite:
		res = "floralwhite"
	case Forestgreen:
		res = "forestgreen"
	case Fuchsia:
		res = "fuchsia"
	case Gainsboro:
		res = "gainsboro"
	case Ghostwhite:
		res = "ghostwhite "
	case Gold:
		res = "gold"
	case Goldenrod:
		res = "goldenrod"
	case Gray:
		res = "gray"
	case Green:
		res = "green"
	case Greenyellow:
		res = "greenyellow"
	case Grey:
		res = "grey"
	case Honeydew:
		res = "honeydew"
	case Hotpink:
		res = "hotpink"
	case Indianred:
		res = "indianred"
	case Indigo:
		res = "indigo "
	case Ivory:
		res = "ivory"
	case Khaki:
		res = "khaki"
	case Lavender:
		res = "lavender"
	case Lavenderblush:
		res = "lavenderblush"
	case Lawngreen:
		res = "lawngreen"
	case Lemonchiffon:
		res = "lemonchiffon"
	case Lightblue:
		res = "lightblue"
	case Lightcoral:
		res = "lightcoral "
	case Lightcyan:
		res = "lightcyan"
	case Lightgoldenrodyellow:
		res = "lightgoldenrodyellow"
	case Lightgray:
		res = "lightgray"
	case Lightgreen:
		res = "lightgreen "
	case Lightgrey:
		res = "lightgrey"
	case Lightpink:
		res = "lightpink"
	case Lightsalmon:
		res = "lightsalmon"
	case Lightseagreen:
		res = "lightseagreen"
	case Lightskyblue:
		res = "lightskyblue"
	case Lightslategray:
		res = "lightslategray "
	case Lightslategrey:
		res = "lightslategrey "
	case Lightsteelblue:
		res = "lightsteelblue "
	case Lightyellow:
		res = "lightyellow"
	case Lime:
		res = "lime"
	case Limegreen:
		res = "limegreen"
	case Linen:
		res = "linen"
	case Magenta:
		res = "magenta"
	case Maroon:
		res = "maroon "
	case Mediumaquamarine:
		res = "mediumaquamarine"
	case Mediumblue:
		res = "mediumblue "
	case Mediumorchid:
		res = "mediumorchid"
	case Mediumpurple:
		res = "mediumpurple"
	case Mediumseagreen:
		res = "mediumseagreen "
	case Mediumslateblue:
		res = "mediumslateblue"
	case Mediumspringgreen:
		res = "mediumspringgreen"
	case Mediumturquoise:
		res = "mediumturquoise"
	case Mediumvioletred:
		res = "mediumvioletred"
	case Midnightblue:
		res = "midnightblue"
	case Mintcream:
		res = "mintcream"
	case Mistyrose:
		res = "mistyrose"
	case Moccasin:
		res = "moccasin"
	case Navajowhite:
		res = "navajowhite"
	case Navy:
		res = "navy"
	case Oldlace:
		res = "oldlace"
	case Olive:
		res = "olive"
	case Olivedrab:
		res = "olivedrab"
	case Orange:
		res = "orange "
	case Orangered:
		res = "orangered"
	case Orchid:
		res = "orchid "
	case Palegoldenrod:
		res = "palegoldenrod"
	case Palegreen:
		res = "palegreen"
	case Paleturquoise:
		res = "paleturquoise"
	case Palevioletred:
		res = "palevioletred"
	case Papayawhip:
		res = "papayawhip "
	case Peachpuff:
		res = "peachpuff"
	case Peru:
		res = "peru"
	case Pink:
		res = "pink"
	case Plum:
		res = "plum"
	case Powderblue:
		res = "powderblue "
	case Purple:
		res = "purple "
	case Red:
		res = "red"
	case Rosybrown:
		res = "rosybrown"
	case Royalblue:
		res = "royalblue"
	case Saddlebrown:
		res = "saddlebrown"
	case Salmon:
		res = "salmon "
	case Sandybrown:
		res = "sandybrown "
	case Seagreen:
		res = "seagreen"
	case Seashell:
		res = "seashell"
	case Sienna:
		res = "sienna "
	case Silver:
		res = "silver "
	case Skyblue:
		res = "skyblue"
	case Slateblue:
		res = "slateblue"
	case Slategray:
		res = "slategray"
	case Slategrey:
		res = "slategrey"
	case Snow:
		res = "snow"
	case Springgreen:
		res = "springgreen"
	case Steelblue:
		res = "steelblue"
	case Tan:
		res = "tan"
	case Teal:
		res = "teal"
	case Thistle:
		res = "thistle"
	case Tomato:
		res = "tomato "
	case Turquoise:
		res = "turquoise"
	case Violet:
		res = "violet "
	case Wheat:
		res = "wheat"
	case White:
		res = "white"
	case Whitesmoke:
		res = "whitesmoke "
	case Yellow:
		res = "yellow "
	case Yellowgreen:
		res = "yellowgreen"
	}
	return
}

func (colortype *ColorType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "aliceblue":
		*colortype = Aliceblue
	case "antiquewhite":
		*colortype = Antiquewhite
	case "aqua":
		*colortype = Aqua
	case "aquamarine ":
		*colortype = Aquamarine
	case "azure":
		*colortype = Azure
	case "beige":
		*colortype = Beige
	case "bisque ":
		*colortype = Bisque
	case "black":
		*colortype = Black
	case "blanchedalmond ":
		*colortype = Blanchedalmond
	case "blue":
		*colortype = Blue
	case "blueviolet ":
		*colortype = Blueviolet
	case "brown":
		*colortype = Brown
	case "burlywood":
		*colortype = Burlywood
	case "cadetblue":
		*colortype = Cadetblue
	case "chartreuse ":
		*colortype = Chartreuse
	case "chocolate":
		*colortype = Chocolate
	case "coral":
		*colortype = Coral
	case "cornflowerblue ":
		*colortype = Cornflowerblue
	case "cornsilk":
		*colortype = Cornsilk
	case "crimson":
		*colortype = Crimson
	case "cyan":
		*colortype = Cyan
	case "darkblue":
		*colortype = Darkblue
	case "darkcyan":
		*colortype = Darkcyan
	case "darkgoldenrod":
		*colortype = Darkgoldenrod
	case "darkgray":
		*colortype = Darkgray
	case "darkgreen":
		*colortype = Darkgreen
	case "darkgrey":
		*colortype = Darkgrey
	case "darkkhaki":
		*colortype = Darkkhaki
	case "darkmagenta":
		*colortype = Darkmagenta
	case "darkolivegreen ":
		*colortype = Darkolivegreen
	case "darkorange ":
		*colortype = Darkorange
	case "darkorchid ":
		*colortype = Darkorchid
	case "darkred":
		*colortype = Darkred
	case "darksalmon ":
		*colortype = Darksalmon
	case "darkseagreen":
		*colortype = Darkseagreen
	case "darkslateblue":
		*colortype = Darkslateblue
	case "darkslategray":
		*colortype = Darkslategray
	case "darkslategrey":
		*colortype = Darkslategrey
	case "darkturquoise":
		*colortype = Darkturquoise
	case "darkviolet ":
		*colortype = Darkviolet
	case "deeppink":
		*colortype = Deeppink
	case "deepskyblue":
		*colortype = Deepskyblue
	case "dimgray":
		*colortype = Dimgray
	case "dimgrey":
		*colortype = Dimgrey
	case "dodgerblue ":
		*colortype = Dodgerblue
	case "firebrick":
		*colortype = Firebrick
	case "floralwhite":
		*colortype = Floralwhite
	case "forestgreen":
		*colortype = Forestgreen
	case "fuchsia":
		*colortype = Fuchsia
	case "gainsboro":
		*colortype = Gainsboro
	case "ghostwhite ":
		*colortype = Ghostwhite
	case "gold":
		*colortype = Gold
	case "goldenrod":
		*colortype = Goldenrod
	case "gray":
		*colortype = Gray
	case "green":
		*colortype = Green
	case "greenyellow":
		*colortype = Greenyellow
	case "grey":
		*colortype = Grey
	case "honeydew":
		*colortype = Honeydew
	case "hotpink":
		*colortype = Hotpink
	case "indianred":
		*colortype = Indianred
	case "indigo ":
		*colortype = Indigo
	case "ivory":
		*colortype = Ivory
	case "khaki":
		*colortype = Khaki
	case "lavender":
		*colortype = Lavender
	case "lavenderblush":
		*colortype = Lavenderblush
	case "lawngreen":
		*colortype = Lawngreen
	case "lemonchiffon":
		*colortype = Lemonchiffon
	case "lightblue":
		*colortype = Lightblue
	case "lightcoral ":
		*colortype = Lightcoral
	case "lightcyan":
		*colortype = Lightcyan
	case "lightgoldenrodyellow":
		*colortype = Lightgoldenrodyellow
	case "lightgray":
		*colortype = Lightgray
	case "lightgreen ":
		*colortype = Lightgreen
	case "lightgrey":
		*colortype = Lightgrey
	case "lightpink":
		*colortype = Lightpink
	case "lightsalmon":
		*colortype = Lightsalmon
	case "lightseagreen":
		*colortype = Lightseagreen
	case "lightskyblue":
		*colortype = Lightskyblue
	case "lightslategray ":
		*colortype = Lightslategray
	case "lightslategrey ":
		*colortype = Lightslategrey
	case "lightsteelblue ":
		*colortype = Lightsteelblue
	case "lightyellow":
		*colortype = Lightyellow
	case "lime":
		*colortype = Lime
	case "limegreen":
		*colortype = Limegreen
	case "linen":
		*colortype = Linen
	case "magenta":
		*colortype = Magenta
	case "maroon ":
		*colortype = Maroon
	case "mediumaquamarine":
		*colortype = Mediumaquamarine
	case "mediumblue ":
		*colortype = Mediumblue
	case "mediumorchid":
		*colortype = Mediumorchid
	case "mediumpurple":
		*colortype = Mediumpurple
	case "mediumseagreen ":
		*colortype = Mediumseagreen
	case "mediumslateblue":
		*colortype = Mediumslateblue
	case "mediumspringgreen":
		*colortype = Mediumspringgreen
	case "mediumturquoise":
		*colortype = Mediumturquoise
	case "mediumvioletred":
		*colortype = Mediumvioletred
	case "midnightblue":
		*colortype = Midnightblue
	case "mintcream":
		*colortype = Mintcream
	case "mistyrose":
		*colortype = Mistyrose
	case "moccasin":
		*colortype = Moccasin
	case "navajowhite":
		*colortype = Navajowhite
	case "navy":
		*colortype = Navy
	case "oldlace":
		*colortype = Oldlace
	case "olive":
		*colortype = Olive
	case "olivedrab":
		*colortype = Olivedrab
	case "orange ":
		*colortype = Orange
	case "orangered":
		*colortype = Orangered
	case "orchid ":
		*colortype = Orchid
	case "palegoldenrod":
		*colortype = Palegoldenrod
	case "palegreen":
		*colortype = Palegreen
	case "paleturquoise":
		*colortype = Paleturquoise
	case "palevioletred":
		*colortype = Palevioletred
	case "papayawhip ":
		*colortype = Papayawhip
	case "peachpuff":
		*colortype = Peachpuff
	case "peru":
		*colortype = Peru
	case "pink":
		*colortype = Pink
	case "plum":
		*colortype = Plum
	case "powderblue ":
		*colortype = Powderblue
	case "purple ":
		*colortype = Purple
	case "red":
		*colortype = Red
	case "rosybrown":
		*colortype = Rosybrown
	case "royalblue":
		*colortype = Royalblue
	case "saddlebrown":
		*colortype = Saddlebrown
	case "salmon ":
		*colortype = Salmon
	case "sandybrown ":
		*colortype = Sandybrown
	case "seagreen":
		*colortype = Seagreen
	case "seashell":
		*colortype = Seashell
	case "sienna ":
		*colortype = Sienna
	case "silver ":
		*colortype = Silver
	case "skyblue":
		*colortype = Skyblue
	case "slateblue":
		*colortype = Slateblue
	case "slategray":
		*colortype = Slategray
	case "slategrey":
		*colortype = Slategrey
	case "snow":
		*colortype = Snow
	case "springgreen":
		*colortype = Springgreen
	case "steelblue":
		*colortype = Steelblue
	case "tan":
		*colortype = Tan
	case "teal":
		*colortype = Teal
	case "thistle":
		*colortype = Thistle
	case "tomato ":
		*colortype = Tomato
	case "turquoise":
		*colortype = Turquoise
	case "violet ":
		*colortype = Violet
	case "wheat":
		*colortype = Wheat
	case "white":
		*colortype = White
	case "whitesmoke ":
		*colortype = Whitesmoke
	case "yellow ":
		*colortype = Yellow
	case "yellowgreen":
		*colortype = Yellowgreen
	default:
		return errUnkownEnum
	}
	return
}

func (colortype *ColorType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Aliceblue":
		*colortype = Aliceblue
	case "Antiquewhite":
		*colortype = Antiquewhite
	case "Aqua":
		*colortype = Aqua
	case "Aquamarine":
		*colortype = Aquamarine
	case "Azure":
		*colortype = Azure
	case "Beige":
		*colortype = Beige
	case "Bisque":
		*colortype = Bisque
	case "Black":
		*colortype = Black
	case "Blanchedalmond":
		*colortype = Blanchedalmond
	case "Blue":
		*colortype = Blue
	case "Blueviolet":
		*colortype = Blueviolet
	case "Brown":
		*colortype = Brown
	case "Burlywood":
		*colortype = Burlywood
	case "Cadetblue":
		*colortype = Cadetblue
	case "Chartreuse":
		*colortype = Chartreuse
	case "Chocolate":
		*colortype = Chocolate
	case "Coral":
		*colortype = Coral
	case "Cornflowerblue":
		*colortype = Cornflowerblue
	case "Cornsilk":
		*colortype = Cornsilk
	case "Crimson":
		*colortype = Crimson
	case "Cyan":
		*colortype = Cyan
	case "Darkblue":
		*colortype = Darkblue
	case "Darkcyan":
		*colortype = Darkcyan
	case "Darkgoldenrod":
		*colortype = Darkgoldenrod
	case "Darkgray":
		*colortype = Darkgray
	case "Darkgreen":
		*colortype = Darkgreen
	case "Darkgrey":
		*colortype = Darkgrey
	case "Darkkhaki":
		*colortype = Darkkhaki
	case "Darkmagenta":
		*colortype = Darkmagenta
	case "Darkolivegreen":
		*colortype = Darkolivegreen
	case "Darkorange":
		*colortype = Darkorange
	case "Darkorchid":
		*colortype = Darkorchid
	case "Darkred":
		*colortype = Darkred
	case "Darksalmon":
		*colortype = Darksalmon
	case "Darkseagreen":
		*colortype = Darkseagreen
	case "Darkslateblue":
		*colortype = Darkslateblue
	case "Darkslategray":
		*colortype = Darkslategray
	case "Darkslategrey":
		*colortype = Darkslategrey
	case "Darkturquoise":
		*colortype = Darkturquoise
	case "Darkviolet":
		*colortype = Darkviolet
	case "Deeppink":
		*colortype = Deeppink
	case "Deepskyblue":
		*colortype = Deepskyblue
	case "Dimgray":
		*colortype = Dimgray
	case "Dimgrey":
		*colortype = Dimgrey
	case "Dodgerblue":
		*colortype = Dodgerblue
	case "Firebrick":
		*colortype = Firebrick
	case "Floralwhite":
		*colortype = Floralwhite
	case "Forestgreen":
		*colortype = Forestgreen
	case "Fuchsia":
		*colortype = Fuchsia
	case "Gainsboro":
		*colortype = Gainsboro
	case "Ghostwhite":
		*colortype = Ghostwhite
	case "Gold":
		*colortype = Gold
	case "Goldenrod":
		*colortype = Goldenrod
	case "Gray":
		*colortype = Gray
	case "Green":
		*colortype = Green
	case "Greenyellow":
		*colortype = Greenyellow
	case "Grey":
		*colortype = Grey
	case "Honeydew":
		*colortype = Honeydew
	case "Hotpink":
		*colortype = Hotpink
	case "Indianred":
		*colortype = Indianred
	case "Indigo":
		*colortype = Indigo
	case "Ivory":
		*colortype = Ivory
	case "Khaki":
		*colortype = Khaki
	case "Lavender":
		*colortype = Lavender
	case "Lavenderblush":
		*colortype = Lavenderblush
	case "Lawngreen":
		*colortype = Lawngreen
	case "Lemonchiffon":
		*colortype = Lemonchiffon
	case "Lightblue":
		*colortype = Lightblue
	case "Lightcoral":
		*colortype = Lightcoral
	case "Lightcyan":
		*colortype = Lightcyan
	case "Lightgoldenrodyellow":
		*colortype = Lightgoldenrodyellow
	case "Lightgray":
		*colortype = Lightgray
	case "Lightgreen":
		*colortype = Lightgreen
	case "Lightgrey":
		*colortype = Lightgrey
	case "Lightpink":
		*colortype = Lightpink
	case "Lightsalmon":
		*colortype = Lightsalmon
	case "Lightseagreen":
		*colortype = Lightseagreen
	case "Lightskyblue":
		*colortype = Lightskyblue
	case "Lightslategray":
		*colortype = Lightslategray
	case "Lightslategrey":
		*colortype = Lightslategrey
	case "Lightsteelblue":
		*colortype = Lightsteelblue
	case "Lightyellow":
		*colortype = Lightyellow
	case "Lime":
		*colortype = Lime
	case "Limegreen":
		*colortype = Limegreen
	case "Linen":
		*colortype = Linen
	case "Magenta":
		*colortype = Magenta
	case "Maroon":
		*colortype = Maroon
	case "Mediumaquamarine":
		*colortype = Mediumaquamarine
	case "Mediumblue":
		*colortype = Mediumblue
	case "Mediumorchid":
		*colortype = Mediumorchid
	case "Mediumpurple":
		*colortype = Mediumpurple
	case "Mediumseagreen":
		*colortype = Mediumseagreen
	case "Mediumslateblue":
		*colortype = Mediumslateblue
	case "Mediumspringgreen":
		*colortype = Mediumspringgreen
	case "Mediumturquoise":
		*colortype = Mediumturquoise
	case "Mediumvioletred":
		*colortype = Mediumvioletred
	case "Midnightblue":
		*colortype = Midnightblue
	case "Mintcream":
		*colortype = Mintcream
	case "Mistyrose":
		*colortype = Mistyrose
	case "Moccasin":
		*colortype = Moccasin
	case "Navajowhite":
		*colortype = Navajowhite
	case "Navy":
		*colortype = Navy
	case "Oldlace":
		*colortype = Oldlace
	case "Olive":
		*colortype = Olive
	case "Olivedrab":
		*colortype = Olivedrab
	case "Orange":
		*colortype = Orange
	case "Orangered":
		*colortype = Orangered
	case "Orchid":
		*colortype = Orchid
	case "Palegoldenrod":
		*colortype = Palegoldenrod
	case "Palegreen":
		*colortype = Palegreen
	case "Paleturquoise":
		*colortype = Paleturquoise
	case "Palevioletred":
		*colortype = Palevioletred
	case "Papayawhip":
		*colortype = Papayawhip
	case "Peachpuff":
		*colortype = Peachpuff
	case "Peru":
		*colortype = Peru
	case "Pink":
		*colortype = Pink
	case "Plum":
		*colortype = Plum
	case "Powderblue":
		*colortype = Powderblue
	case "Purple":
		*colortype = Purple
	case "Red":
		*colortype = Red
	case "Rosybrown":
		*colortype = Rosybrown
	case "Royalblue":
		*colortype = Royalblue
	case "Saddlebrown":
		*colortype = Saddlebrown
	case "Salmon":
		*colortype = Salmon
	case "Sandybrown":
		*colortype = Sandybrown
	case "Seagreen":
		*colortype = Seagreen
	case "Seashell":
		*colortype = Seashell
	case "Sienna":
		*colortype = Sienna
	case "Silver":
		*colortype = Silver
	case "Skyblue":
		*colortype = Skyblue
	case "Slateblue":
		*colortype = Slateblue
	case "Slategray":
		*colortype = Slategray
	case "Slategrey":
		*colortype = Slategrey
	case "Snow":
		*colortype = Snow
	case "Springgreen":
		*colortype = Springgreen
	case "Steelblue":
		*colortype = Steelblue
	case "Tan":
		*colortype = Tan
	case "Teal":
		*colortype = Teal
	case "Thistle":
		*colortype = Thistle
	case "Tomato":
		*colortype = Tomato
	case "Turquoise":
		*colortype = Turquoise
	case "Violet":
		*colortype = Violet
	case "Wheat":
		*colortype = Wheat
	case "White":
		*colortype = White
	case "Whitesmoke":
		*colortype = Whitesmoke
	case "Yellow":
		*colortype = Yellow
	case "Yellowgreen":
		*colortype = Yellowgreen
	default:
		return errUnkownEnum
	}
	return
}

func (colortype *ColorType) ToCodeString() (res string) {

	switch *colortype {
	// insertion code per enum code
	case Aliceblue:
		res = "Aliceblue"
	case Antiquewhite:
		res = "Antiquewhite"
	case Aqua:
		res = "Aqua"
	case Aquamarine:
		res = "Aquamarine"
	case Azure:
		res = "Azure"
	case Beige:
		res = "Beige"
	case Bisque:
		res = "Bisque"
	case Black:
		res = "Black"
	case Blanchedalmond:
		res = "Blanchedalmond"
	case Blue:
		res = "Blue"
	case Blueviolet:
		res = "Blueviolet"
	case Brown:
		res = "Brown"
	case Burlywood:
		res = "Burlywood"
	case Cadetblue:
		res = "Cadetblue"
	case Chartreuse:
		res = "Chartreuse"
	case Chocolate:
		res = "Chocolate"
	case Coral:
		res = "Coral"
	case Cornflowerblue:
		res = "Cornflowerblue"
	case Cornsilk:
		res = "Cornsilk"
	case Crimson:
		res = "Crimson"
	case Cyan:
		res = "Cyan"
	case Darkblue:
		res = "Darkblue"
	case Darkcyan:
		res = "Darkcyan"
	case Darkgoldenrod:
		res = "Darkgoldenrod"
	case Darkgray:
		res = "Darkgray"
	case Darkgreen:
		res = "Darkgreen"
	case Darkgrey:
		res = "Darkgrey"
	case Darkkhaki:
		res = "Darkkhaki"
	case Darkmagenta:
		res = "Darkmagenta"
	case Darkolivegreen:
		res = "Darkolivegreen"
	case Darkorange:
		res = "Darkorange"
	case Darkorchid:
		res = "Darkorchid"
	case Darkred:
		res = "Darkred"
	case Darksalmon:
		res = "Darksalmon"
	case Darkseagreen:
		res = "Darkseagreen"
	case Darkslateblue:
		res = "Darkslateblue"
	case Darkslategray:
		res = "Darkslategray"
	case Darkslategrey:
		res = "Darkslategrey"
	case Darkturquoise:
		res = "Darkturquoise"
	case Darkviolet:
		res = "Darkviolet"
	case Deeppink:
		res = "Deeppink"
	case Deepskyblue:
		res = "Deepskyblue"
	case Dimgray:
		res = "Dimgray"
	case Dimgrey:
		res = "Dimgrey"
	case Dodgerblue:
		res = "Dodgerblue"
	case Firebrick:
		res = "Firebrick"
	case Floralwhite:
		res = "Floralwhite"
	case Forestgreen:
		res = "Forestgreen"
	case Fuchsia:
		res = "Fuchsia"
	case Gainsboro:
		res = "Gainsboro"
	case Ghostwhite:
		res = "Ghostwhite"
	case Gold:
		res = "Gold"
	case Goldenrod:
		res = "Goldenrod"
	case Gray:
		res = "Gray"
	case Green:
		res = "Green"
	case Greenyellow:
		res = "Greenyellow"
	case Grey:
		res = "Grey"
	case Honeydew:
		res = "Honeydew"
	case Hotpink:
		res = "Hotpink"
	case Indianred:
		res = "Indianred"
	case Indigo:
		res = "Indigo"
	case Ivory:
		res = "Ivory"
	case Khaki:
		res = "Khaki"
	case Lavender:
		res = "Lavender"
	case Lavenderblush:
		res = "Lavenderblush"
	case Lawngreen:
		res = "Lawngreen"
	case Lemonchiffon:
		res = "Lemonchiffon"
	case Lightblue:
		res = "Lightblue"
	case Lightcoral:
		res = "Lightcoral"
	case Lightcyan:
		res = "Lightcyan"
	case Lightgoldenrodyellow:
		res = "Lightgoldenrodyellow"
	case Lightgray:
		res = "Lightgray"
	case Lightgreen:
		res = "Lightgreen"
	case Lightgrey:
		res = "Lightgrey"
	case Lightpink:
		res = "Lightpink"
	case Lightsalmon:
		res = "Lightsalmon"
	case Lightseagreen:
		res = "Lightseagreen"
	case Lightskyblue:
		res = "Lightskyblue"
	case Lightslategray:
		res = "Lightslategray"
	case Lightslategrey:
		res = "Lightslategrey"
	case Lightsteelblue:
		res = "Lightsteelblue"
	case Lightyellow:
		res = "Lightyellow"
	case Lime:
		res = "Lime"
	case Limegreen:
		res = "Limegreen"
	case Linen:
		res = "Linen"
	case Magenta:
		res = "Magenta"
	case Maroon:
		res = "Maroon"
	case Mediumaquamarine:
		res = "Mediumaquamarine"
	case Mediumblue:
		res = "Mediumblue"
	case Mediumorchid:
		res = "Mediumorchid"
	case Mediumpurple:
		res = "Mediumpurple"
	case Mediumseagreen:
		res = "Mediumseagreen"
	case Mediumslateblue:
		res = "Mediumslateblue"
	case Mediumspringgreen:
		res = "Mediumspringgreen"
	case Mediumturquoise:
		res = "Mediumturquoise"
	case Mediumvioletred:
		res = "Mediumvioletred"
	case Midnightblue:
		res = "Midnightblue"
	case Mintcream:
		res = "Mintcream"
	case Mistyrose:
		res = "Mistyrose"
	case Moccasin:
		res = "Moccasin"
	case Navajowhite:
		res = "Navajowhite"
	case Navy:
		res = "Navy"
	case Oldlace:
		res = "Oldlace"
	case Olive:
		res = "Olive"
	case Olivedrab:
		res = "Olivedrab"
	case Orange:
		res = "Orange"
	case Orangered:
		res = "Orangered"
	case Orchid:
		res = "Orchid"
	case Palegoldenrod:
		res = "Palegoldenrod"
	case Palegreen:
		res = "Palegreen"
	case Paleturquoise:
		res = "Paleturquoise"
	case Palevioletred:
		res = "Palevioletred"
	case Papayawhip:
		res = "Papayawhip"
	case Peachpuff:
		res = "Peachpuff"
	case Peru:
		res = "Peru"
	case Pink:
		res = "Pink"
	case Plum:
		res = "Plum"
	case Powderblue:
		res = "Powderblue"
	case Purple:
		res = "Purple"
	case Red:
		res = "Red"
	case Rosybrown:
		res = "Rosybrown"
	case Royalblue:
		res = "Royalblue"
	case Saddlebrown:
		res = "Saddlebrown"
	case Salmon:
		res = "Salmon"
	case Sandybrown:
		res = "Sandybrown"
	case Seagreen:
		res = "Seagreen"
	case Seashell:
		res = "Seashell"
	case Sienna:
		res = "Sienna"
	case Silver:
		res = "Silver"
	case Skyblue:
		res = "Skyblue"
	case Slateblue:
		res = "Slateblue"
	case Slategray:
		res = "Slategray"
	case Slategrey:
		res = "Slategrey"
	case Snow:
		res = "Snow"
	case Springgreen:
		res = "Springgreen"
	case Steelblue:
		res = "Steelblue"
	case Tan:
		res = "Tan"
	case Teal:
		res = "Teal"
	case Thistle:
		res = "Thistle"
	case Tomato:
		res = "Tomato"
	case Turquoise:
		res = "Turquoise"
	case Violet:
		res = "Violet"
	case Wheat:
		res = "Wheat"
	case White:
		res = "White"
	case Whitesmoke:
		res = "Whitesmoke"
	case Yellow:
		res = "Yellow"
	case Yellowgreen:
		res = "Yellowgreen"
	}
	return
}

func (colortype ColorType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Aliceblue")
	res = append(res, "Antiquewhite")
	res = append(res, "Aqua")
	res = append(res, "Aquamarine")
	res = append(res, "Azure")
	res = append(res, "Beige")
	res = append(res, "Bisque")
	res = append(res, "Black")
	res = append(res, "Blanchedalmond")
	res = append(res, "Blue")
	res = append(res, "Blueviolet")
	res = append(res, "Brown")
	res = append(res, "Burlywood")
	res = append(res, "Cadetblue")
	res = append(res, "Chartreuse")
	res = append(res, "Chocolate")
	res = append(res, "Coral")
	res = append(res, "Cornflowerblue")
	res = append(res, "Cornsilk")
	res = append(res, "Crimson")
	res = append(res, "Cyan")
	res = append(res, "Darkblue")
	res = append(res, "Darkcyan")
	res = append(res, "Darkgoldenrod")
	res = append(res, "Darkgray")
	res = append(res, "Darkgreen")
	res = append(res, "Darkgrey")
	res = append(res, "Darkkhaki")
	res = append(res, "Darkmagenta")
	res = append(res, "Darkolivegreen")
	res = append(res, "Darkorange")
	res = append(res, "Darkorchid")
	res = append(res, "Darkred")
	res = append(res, "Darksalmon")
	res = append(res, "Darkseagreen")
	res = append(res, "Darkslateblue")
	res = append(res, "Darkslategray")
	res = append(res, "Darkslategrey")
	res = append(res, "Darkturquoise")
	res = append(res, "Darkviolet")
	res = append(res, "Deeppink")
	res = append(res, "Deepskyblue")
	res = append(res, "Dimgray")
	res = append(res, "Dimgrey")
	res = append(res, "Dodgerblue")
	res = append(res, "Firebrick")
	res = append(res, "Floralwhite")
	res = append(res, "Forestgreen")
	res = append(res, "Fuchsia")
	res = append(res, "Gainsboro")
	res = append(res, "Ghostwhite")
	res = append(res, "Gold")
	res = append(res, "Goldenrod")
	res = append(res, "Gray")
	res = append(res, "Green")
	res = append(res, "Greenyellow")
	res = append(res, "Grey")
	res = append(res, "Honeydew")
	res = append(res, "Hotpink")
	res = append(res, "Indianred")
	res = append(res, "Indigo")
	res = append(res, "Ivory")
	res = append(res, "Khaki")
	res = append(res, "Lavender")
	res = append(res, "Lavenderblush")
	res = append(res, "Lawngreen")
	res = append(res, "Lemonchiffon")
	res = append(res, "Lightblue")
	res = append(res, "Lightcoral")
	res = append(res, "Lightcyan")
	res = append(res, "Lightgoldenrodyellow")
	res = append(res, "Lightgray")
	res = append(res, "Lightgreen")
	res = append(res, "Lightgrey")
	res = append(res, "Lightpink")
	res = append(res, "Lightsalmon")
	res = append(res, "Lightseagreen")
	res = append(res, "Lightskyblue")
	res = append(res, "Lightslategray")
	res = append(res, "Lightslategrey")
	res = append(res, "Lightsteelblue")
	res = append(res, "Lightyellow")
	res = append(res, "Lime")
	res = append(res, "Limegreen")
	res = append(res, "Linen")
	res = append(res, "Magenta")
	res = append(res, "Maroon")
	res = append(res, "Mediumaquamarine")
	res = append(res, "Mediumblue")
	res = append(res, "Mediumorchid")
	res = append(res, "Mediumpurple")
	res = append(res, "Mediumseagreen")
	res = append(res, "Mediumslateblue")
	res = append(res, "Mediumspringgreen")
	res = append(res, "Mediumturquoise")
	res = append(res, "Mediumvioletred")
	res = append(res, "Midnightblue")
	res = append(res, "Mintcream")
	res = append(res, "Mistyrose")
	res = append(res, "Moccasin")
	res = append(res, "Navajowhite")
	res = append(res, "Navy")
	res = append(res, "Oldlace")
	res = append(res, "Olive")
	res = append(res, "Olivedrab")
	res = append(res, "Orange")
	res = append(res, "Orangered")
	res = append(res, "Orchid")
	res = append(res, "Palegoldenrod")
	res = append(res, "Palegreen")
	res = append(res, "Paleturquoise")
	res = append(res, "Palevioletred")
	res = append(res, "Papayawhip")
	res = append(res, "Peachpuff")
	res = append(res, "Peru")
	res = append(res, "Pink")
	res = append(res, "Plum")
	res = append(res, "Powderblue")
	res = append(res, "Purple")
	res = append(res, "Red")
	res = append(res, "Rosybrown")
	res = append(res, "Royalblue")
	res = append(res, "Saddlebrown")
	res = append(res, "Salmon")
	res = append(res, "Sandybrown")
	res = append(res, "Seagreen")
	res = append(res, "Seashell")
	res = append(res, "Sienna")
	res = append(res, "Silver")
	res = append(res, "Skyblue")
	res = append(res, "Slateblue")
	res = append(res, "Slategray")
	res = append(res, "Slategrey")
	res = append(res, "Snow")
	res = append(res, "Springgreen")
	res = append(res, "Steelblue")
	res = append(res, "Tan")
	res = append(res, "Teal")
	res = append(res, "Thistle")
	res = append(res, "Tomato")
	res = append(res, "Turquoise")
	res = append(res, "Violet")
	res = append(res, "Wheat")
	res = append(res, "White")
	res = append(res, "Whitesmoke")
	res = append(res, "Yellow")
	res = append(res, "Yellowgreen")

	return
}

func (colortype ColorType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "aliceblue")
	res = append(res, "antiquewhite")
	res = append(res, "aqua")
	res = append(res, "aquamarine ")
	res = append(res, "azure")
	res = append(res, "beige")
	res = append(res, "bisque ")
	res = append(res, "black")
	res = append(res, "blanchedalmond ")
	res = append(res, "blue")
	res = append(res, "blueviolet ")
	res = append(res, "brown")
	res = append(res, "burlywood")
	res = append(res, "cadetblue")
	res = append(res, "chartreuse ")
	res = append(res, "chocolate")
	res = append(res, "coral")
	res = append(res, "cornflowerblue ")
	res = append(res, "cornsilk")
	res = append(res, "crimson")
	res = append(res, "cyan")
	res = append(res, "darkblue")
	res = append(res, "darkcyan")
	res = append(res, "darkgoldenrod")
	res = append(res, "darkgray")
	res = append(res, "darkgreen")
	res = append(res, "darkgrey")
	res = append(res, "darkkhaki")
	res = append(res, "darkmagenta")
	res = append(res, "darkolivegreen ")
	res = append(res, "darkorange ")
	res = append(res, "darkorchid ")
	res = append(res, "darkred")
	res = append(res, "darksalmon ")
	res = append(res, "darkseagreen")
	res = append(res, "darkslateblue")
	res = append(res, "darkslategray")
	res = append(res, "darkslategrey")
	res = append(res, "darkturquoise")
	res = append(res, "darkviolet ")
	res = append(res, "deeppink")
	res = append(res, "deepskyblue")
	res = append(res, "dimgray")
	res = append(res, "dimgrey")
	res = append(res, "dodgerblue ")
	res = append(res, "firebrick")
	res = append(res, "floralwhite")
	res = append(res, "forestgreen")
	res = append(res, "fuchsia")
	res = append(res, "gainsboro")
	res = append(res, "ghostwhite ")
	res = append(res, "gold")
	res = append(res, "goldenrod")
	res = append(res, "gray")
	res = append(res, "green")
	res = append(res, "greenyellow")
	res = append(res, "grey")
	res = append(res, "honeydew")
	res = append(res, "hotpink")
	res = append(res, "indianred")
	res = append(res, "indigo ")
	res = append(res, "ivory")
	res = append(res, "khaki")
	res = append(res, "lavender")
	res = append(res, "lavenderblush")
	res = append(res, "lawngreen")
	res = append(res, "lemonchiffon")
	res = append(res, "lightblue")
	res = append(res, "lightcoral ")
	res = append(res, "lightcyan")
	res = append(res, "lightgoldenrodyellow")
	res = append(res, "lightgray")
	res = append(res, "lightgreen ")
	res = append(res, "lightgrey")
	res = append(res, "lightpink")
	res = append(res, "lightsalmon")
	res = append(res, "lightseagreen")
	res = append(res, "lightskyblue")
	res = append(res, "lightslategray ")
	res = append(res, "lightslategrey ")
	res = append(res, "lightsteelblue ")
	res = append(res, "lightyellow")
	res = append(res, "lime")
	res = append(res, "limegreen")
	res = append(res, "linen")
	res = append(res, "magenta")
	res = append(res, "maroon ")
	res = append(res, "mediumaquamarine")
	res = append(res, "mediumblue ")
	res = append(res, "mediumorchid")
	res = append(res, "mediumpurple")
	res = append(res, "mediumseagreen ")
	res = append(res, "mediumslateblue")
	res = append(res, "mediumspringgreen")
	res = append(res, "mediumturquoise")
	res = append(res, "mediumvioletred")
	res = append(res, "midnightblue")
	res = append(res, "mintcream")
	res = append(res, "mistyrose")
	res = append(res, "moccasin")
	res = append(res, "navajowhite")
	res = append(res, "navy")
	res = append(res, "oldlace")
	res = append(res, "olive")
	res = append(res, "olivedrab")
	res = append(res, "orange ")
	res = append(res, "orangered")
	res = append(res, "orchid ")
	res = append(res, "palegoldenrod")
	res = append(res, "palegreen")
	res = append(res, "paleturquoise")
	res = append(res, "palevioletred")
	res = append(res, "papayawhip ")
	res = append(res, "peachpuff")
	res = append(res, "peru")
	res = append(res, "pink")
	res = append(res, "plum")
	res = append(res, "powderblue ")
	res = append(res, "purple ")
	res = append(res, "red")
	res = append(res, "rosybrown")
	res = append(res, "royalblue")
	res = append(res, "saddlebrown")
	res = append(res, "salmon ")
	res = append(res, "sandybrown ")
	res = append(res, "seagreen")
	res = append(res, "seashell")
	res = append(res, "sienna ")
	res = append(res, "silver ")
	res = append(res, "skyblue")
	res = append(res, "slateblue")
	res = append(res, "slategray")
	res = append(res, "slategrey")
	res = append(res, "snow")
	res = append(res, "springgreen")
	res = append(res, "steelblue")
	res = append(res, "tan")
	res = append(res, "teal")
	res = append(res, "thistle")
	res = append(res, "tomato ")
	res = append(res, "turquoise")
	res = append(res, "violet ")
	res = append(res, "wheat")
	res = append(res, "white")
	res = append(res, "whitesmoke ")
	res = append(res, "yellow ")
	res = append(res, "yellowgreen")

	return
}

// Utility function for DrawingState
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (drawingstate DrawingState) ToString() (res string) {

	// migration of former implementation of enum
	switch drawingstate {
	// insertion code per enum code
	case NOT_DRAWING_LINK:
		res = "NOT_DRAWING_LINK"
	case DRAWING_LINK:
		res = "DRAWING_LINK"
	}
	return
}

func (drawingstate *DrawingState) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NOT_DRAWING_LINK":
		*drawingstate = NOT_DRAWING_LINK
	case "DRAWING_LINK":
		*drawingstate = DRAWING_LINK
	default:
		return errUnkownEnum
	}
	return
}

func (drawingstate *DrawingState) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "NOT_DRAWING_LINK":
		*drawingstate = NOT_DRAWING_LINK
	case "DRAWING_LINK":
		*drawingstate = DRAWING_LINK
	default:
		return errUnkownEnum
	}
	return
}

func (drawingstate *DrawingState) ToCodeString() (res string) {

	switch *drawingstate {
	// insertion code per enum code
	case NOT_DRAWING_LINK:
		res = "NOT_DRAWING_LINK"
	case DRAWING_LINK:
		res = "DRAWING_LINK"
	}
	return
}

func (drawingstate DrawingState) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NOT_DRAWING_LINK")
	res = append(res, "DRAWING_LINK")

	return
}

func (drawingstate DrawingState) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "NOT_DRAWING_LINK")
	res = append(res, "DRAWING_LINK")

	return
}

// Utility function for LinkAnchorType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (linkanchortype LinkAnchorType) ToString() (res string) {

	// migration of former implementation of enum
	switch linkanchortype {
	// insertion code per enum code
	case LINK_LEFT_OR_TOP:
		res = "LINK_LEFT_OR_TOP"
	case LINK_RIGHT_OR_BOTTOM:
		res = "LINK_RIGHT_OR_BOTTOM"
	}
	return
}

func (linkanchortype *LinkAnchorType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LINK_LEFT_OR_TOP":
		*linkanchortype = LINK_LEFT_OR_TOP
	case "LINK_RIGHT_OR_BOTTOM":
		*linkanchortype = LINK_RIGHT_OR_BOTTOM
	default:
		return errUnkownEnum
	}
	return
}

func (linkanchortype *LinkAnchorType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LINK_LEFT_OR_TOP":
		*linkanchortype = LINK_LEFT_OR_TOP
	case "LINK_RIGHT_OR_BOTTOM":
		*linkanchortype = LINK_RIGHT_OR_BOTTOM
	default:
		return errUnkownEnum
	}
	return
}

func (linkanchortype *LinkAnchorType) ToCodeString() (res string) {

	switch *linkanchortype {
	// insertion code per enum code
	case LINK_LEFT_OR_TOP:
		res = "LINK_LEFT_OR_TOP"
	case LINK_RIGHT_OR_BOTTOM:
		res = "LINK_RIGHT_OR_BOTTOM"
	}
	return
}

func (linkanchortype LinkAnchorType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LINK_LEFT_OR_TOP")
	res = append(res, "LINK_RIGHT_OR_BOTTOM")

	return
}

func (linkanchortype LinkAnchorType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LINK_LEFT_OR_TOP")
	res = append(res, "LINK_RIGHT_OR_BOTTOM")

	return
}

// Utility function for LinkType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (linktype LinkType) ToString() (res string) {

	// migration of former implementation of enum
	switch linktype {
	// insertion code per enum code
	case LINK_TYPE_LINE_WITH_CONTROL_POINTS:
		res = "LINK_TYPE_LINE_WITH_CONTROL_POINTS"
	case LINK_TYPE_FLOATING_ORTHOGONAL:
		res = "LINK_TYPE_FLOATING_ORTHOGONAL"
	}
	return
}

func (linktype *LinkType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LINK_TYPE_LINE_WITH_CONTROL_POINTS":
		*linktype = LINK_TYPE_LINE_WITH_CONTROL_POINTS
	case "LINK_TYPE_FLOATING_ORTHOGONAL":
		*linktype = LINK_TYPE_FLOATING_ORTHOGONAL
	default:
		return errUnkownEnum
	}
	return
}

func (linktype *LinkType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "LINK_TYPE_LINE_WITH_CONTROL_POINTS":
		*linktype = LINK_TYPE_LINE_WITH_CONTROL_POINTS
	case "LINK_TYPE_FLOATING_ORTHOGONAL":
		*linktype = LINK_TYPE_FLOATING_ORTHOGONAL
	default:
		return errUnkownEnum
	}
	return
}

func (linktype *LinkType) ToCodeString() (res string) {

	switch *linktype {
	// insertion code per enum code
	case LINK_TYPE_LINE_WITH_CONTROL_POINTS:
		res = "LINK_TYPE_LINE_WITH_CONTROL_POINTS"
	case LINK_TYPE_FLOATING_ORTHOGONAL:
		res = "LINK_TYPE_FLOATING_ORTHOGONAL"
	}
	return
}

func (linktype LinkType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LINK_TYPE_LINE_WITH_CONTROL_POINTS")
	res = append(res, "LINK_TYPE_FLOATING_ORTHOGONAL")

	return
}

func (linktype LinkType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "LINK_TYPE_LINE_WITH_CONTROL_POINTS")
	res = append(res, "LINK_TYPE_FLOATING_ORTHOGONAL")

	return
}

// Utility function for OrientationType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (orientationtype OrientationType) ToString() (res string) {

	// migration of former implementation of enum
	switch orientationtype {
	// insertion code per enum code
	case ORIENTATION_HORIZONTAL:
		res = "ORIENTATION_HORIZONTAL"
	case ORIENTATION_VERTICAL:
		res = "ORIENTATION_VERTICAL"
	}
	return
}

func (orientationtype *OrientationType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ORIENTATION_HORIZONTAL":
		*orientationtype = ORIENTATION_HORIZONTAL
	case "ORIENTATION_VERTICAL":
		*orientationtype = ORIENTATION_VERTICAL
	default:
		return errUnkownEnum
	}
	return
}

func (orientationtype *OrientationType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ORIENTATION_HORIZONTAL":
		*orientationtype = ORIENTATION_HORIZONTAL
	case "ORIENTATION_VERTICAL":
		*orientationtype = ORIENTATION_VERTICAL
	default:
		return errUnkownEnum
	}
	return
}

func (orientationtype *OrientationType) ToCodeString() (res string) {

	switch *orientationtype {
	// insertion code per enum code
	case ORIENTATION_HORIZONTAL:
		res = "ORIENTATION_HORIZONTAL"
	case ORIENTATION_VERTICAL:
		res = "ORIENTATION_VERTICAL"
	}
	return
}

func (orientationtype OrientationType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ORIENTATION_HORIZONTAL")
	res = append(res, "ORIENTATION_VERTICAL")

	return
}

func (orientationtype OrientationType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ORIENTATION_HORIZONTAL")
	res = append(res, "ORIENTATION_VERTICAL")

	return
}

// Utility function for PositionOnArrowType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (positiononarrowtype PositionOnArrowType) ToString() (res string) {

	// migration of former implementation of enum
	switch positiononarrowtype {
	// insertion code per enum code
	case POSITION_ON_ARROW_START:
		res = "POSITION_ON_ARROW_START"
	case POSITION_ON_ARROW_END:
		res = "POSITION_ON_ARROW_END"
	}
	return
}

func (positiononarrowtype *PositionOnArrowType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "POSITION_ON_ARROW_START":
		*positiononarrowtype = POSITION_ON_ARROW_START
	case "POSITION_ON_ARROW_END":
		*positiononarrowtype = POSITION_ON_ARROW_END
	default:
		return errUnkownEnum
	}
	return
}

func (positiononarrowtype *PositionOnArrowType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "POSITION_ON_ARROW_START":
		*positiononarrowtype = POSITION_ON_ARROW_START
	case "POSITION_ON_ARROW_END":
		*positiononarrowtype = POSITION_ON_ARROW_END
	default:
		return errUnkownEnum
	}
	return
}

func (positiononarrowtype *PositionOnArrowType) ToCodeString() (res string) {

	switch *positiononarrowtype {
	// insertion code per enum code
	case POSITION_ON_ARROW_START:
		res = "POSITION_ON_ARROW_START"
	case POSITION_ON_ARROW_END:
		res = "POSITION_ON_ARROW_END"
	}
	return
}

func (positiononarrowtype PositionOnArrowType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "POSITION_ON_ARROW_START")
	res = append(res, "POSITION_ON_ARROW_END")

	return
}

func (positiononarrowtype PositionOnArrowType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "POSITION_ON_ARROW_START")
	res = append(res, "POSITION_ON_ARROW_END")

	return
}

// Utility function for RectAnchorType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (rectanchortype RectAnchorType) ToString() (res string) {

	// migration of former implementation of enum
	switch rectanchortype {
	// insertion code per enum code
	case RECT_TOP:
		res = "RECT_TOP"
	case RECT_TOP_LEFT:
		res = "RECT_TOP_LEFT"
	case RECT_TOP_RIGHT:
		res = "RECT_TOP_RIGHT"
	case RECT_BOTTOM:
		res = "RECT_BOTTOM"
	case RECT_BOTTOM_LEFT:
		res = "RECT_BOTTOM_LEFT"
	case RECT_BOTTOM_LEFT_LEFT:
		res = "RECT_BOTTOM_LEFT_LEFT"
	case RECT_BOTTOM_BOTTOM_LEFT:
		res = "RECT_BOTTOM_BOTTOM_LEFT"
	case RECT_BOTTOM_RIGHT:
		res = "RECT_BOTTOM_RIGHT"
	case RECT_BOTTOM_INSIDE_RIGHT:
		res = "RECT_BOTTOM_INSIDE_RIGHT"
	case RECT_LEFT:
		res = "RECT_LEFT"
	case RECT_RIGHT:
		res = "RECT_RIGHT"
	case RECT_CENTER:
		res = "RECT_CENTER"
	}
	return
}

func (rectanchortype *RectAnchorType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RECT_TOP":
		*rectanchortype = RECT_TOP
	case "RECT_TOP_LEFT":
		*rectanchortype = RECT_TOP_LEFT
	case "RECT_TOP_RIGHT":
		*rectanchortype = RECT_TOP_RIGHT
	case "RECT_BOTTOM":
		*rectanchortype = RECT_BOTTOM
	case "RECT_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT
	case "RECT_BOTTOM_LEFT_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT_LEFT
	case "RECT_BOTTOM_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_BOTTOM_LEFT
	case "RECT_BOTTOM_RIGHT":
		*rectanchortype = RECT_BOTTOM_RIGHT
	case "RECT_BOTTOM_INSIDE_RIGHT":
		*rectanchortype = RECT_BOTTOM_INSIDE_RIGHT
	case "RECT_LEFT":
		*rectanchortype = RECT_LEFT
	case "RECT_RIGHT":
		*rectanchortype = RECT_RIGHT
	case "RECT_CENTER":
		*rectanchortype = RECT_CENTER
	default:
		return errUnkownEnum
	}
	return
}

func (rectanchortype *RectAnchorType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RECT_TOP":
		*rectanchortype = RECT_TOP
	case "RECT_TOP_LEFT":
		*rectanchortype = RECT_TOP_LEFT
	case "RECT_TOP_RIGHT":
		*rectanchortype = RECT_TOP_RIGHT
	case "RECT_BOTTOM":
		*rectanchortype = RECT_BOTTOM
	case "RECT_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT
	case "RECT_BOTTOM_LEFT_LEFT":
		*rectanchortype = RECT_BOTTOM_LEFT_LEFT
	case "RECT_BOTTOM_BOTTOM_LEFT":
		*rectanchortype = RECT_BOTTOM_BOTTOM_LEFT
	case "RECT_BOTTOM_RIGHT":
		*rectanchortype = RECT_BOTTOM_RIGHT
	case "RECT_BOTTOM_INSIDE_RIGHT":
		*rectanchortype = RECT_BOTTOM_INSIDE_RIGHT
	case "RECT_LEFT":
		*rectanchortype = RECT_LEFT
	case "RECT_RIGHT":
		*rectanchortype = RECT_RIGHT
	case "RECT_CENTER":
		*rectanchortype = RECT_CENTER
	default:
		return errUnkownEnum
	}
	return
}

func (rectanchortype *RectAnchorType) ToCodeString() (res string) {

	switch *rectanchortype {
	// insertion code per enum code
	case RECT_TOP:
		res = "RECT_TOP"
	case RECT_TOP_LEFT:
		res = "RECT_TOP_LEFT"
	case RECT_TOP_RIGHT:
		res = "RECT_TOP_RIGHT"
	case RECT_BOTTOM:
		res = "RECT_BOTTOM"
	case RECT_BOTTOM_LEFT:
		res = "RECT_BOTTOM_LEFT"
	case RECT_BOTTOM_LEFT_LEFT:
		res = "RECT_BOTTOM_LEFT_LEFT"
	case RECT_BOTTOM_BOTTOM_LEFT:
		res = "RECT_BOTTOM_BOTTOM_LEFT"
	case RECT_BOTTOM_RIGHT:
		res = "RECT_BOTTOM_RIGHT"
	case RECT_BOTTOM_INSIDE_RIGHT:
		res = "RECT_BOTTOM_INSIDE_RIGHT"
	case RECT_LEFT:
		res = "RECT_LEFT"
	case RECT_RIGHT:
		res = "RECT_RIGHT"
	case RECT_CENTER:
		res = "RECT_CENTER"
	}
	return
}

func (rectanchortype RectAnchorType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RECT_TOP")
	res = append(res, "RECT_TOP_LEFT")
	res = append(res, "RECT_TOP_RIGHT")
	res = append(res, "RECT_BOTTOM")
	res = append(res, "RECT_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_LEFT_LEFT")
	res = append(res, "RECT_BOTTOM_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_RIGHT")
	res = append(res, "RECT_BOTTOM_INSIDE_RIGHT")
	res = append(res, "RECT_LEFT")
	res = append(res, "RECT_RIGHT")
	res = append(res, "RECT_CENTER")

	return
}

func (rectanchortype RectAnchorType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RECT_TOP")
	res = append(res, "RECT_TOP_LEFT")
	res = append(res, "RECT_TOP_RIGHT")
	res = append(res, "RECT_BOTTOM")
	res = append(res, "RECT_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_LEFT_LEFT")
	res = append(res, "RECT_BOTTOM_BOTTOM_LEFT")
	res = append(res, "RECT_BOTTOM_RIGHT")
	res = append(res, "RECT_BOTTOM_INSIDE_RIGHT")
	res = append(res, "RECT_LEFT")
	res = append(res, "RECT_RIGHT")
	res = append(res, "RECT_CENTER")

	return
}

// Utility function for SideType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (sidetype SideType) ToString() (res string) {

	// migration of former implementation of enum
	switch sidetype {
	// insertion code per enum code
	case SIDE_TOP:
		res = "SIDE_TOP"
	case SIDE_BOTTOM:
		res = "SIDE_BOTTOM"
	case SIDE_LEFT:
		res = "SIDE_LEFT"
	case SIDE_RIGHT:
		res = "SIDE_RIGHT"
	}
	return
}

func (sidetype *SideType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SIDE_TOP":
		*sidetype = SIDE_TOP
	case "SIDE_BOTTOM":
		*sidetype = SIDE_BOTTOM
	case "SIDE_LEFT":
		*sidetype = SIDE_LEFT
	case "SIDE_RIGHT":
		*sidetype = SIDE_RIGHT
	default:
		return errUnkownEnum
	}
	return
}

func (sidetype *SideType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SIDE_TOP":
		*sidetype = SIDE_TOP
	case "SIDE_BOTTOM":
		*sidetype = SIDE_BOTTOM
	case "SIDE_LEFT":
		*sidetype = SIDE_LEFT
	case "SIDE_RIGHT":
		*sidetype = SIDE_RIGHT
	default:
		return errUnkownEnum
	}
	return
}

func (sidetype *SideType) ToCodeString() (res string) {

	switch *sidetype {
	// insertion code per enum code
	case SIDE_TOP:
		res = "SIDE_TOP"
	case SIDE_BOTTOM:
		res = "SIDE_BOTTOM"
	case SIDE_LEFT:
		res = "SIDE_LEFT"
	case SIDE_RIGHT:
		res = "SIDE_RIGHT"
	}
	return
}

func (sidetype SideType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SIDE_TOP")
	res = append(res, "SIDE_BOTTOM")
	res = append(res, "SIDE_LEFT")
	res = append(res, "SIDE_RIGHT")

	return
}

func (sidetype SideType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SIDE_TOP")
	res = append(res, "SIDE_BOTTOM")
	res = append(res, "SIDE_LEFT")
	res = append(res, "SIDE_RIGHT")

	return
}

// Utility function for StackName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (stackname StackName) ToString() (res string) {

	// migration of former implementation of enum
	switch stackname {
	// insertion code per enum code
	case StackNameDefault:
		res = "gongsvg"
	}
	return
}

func (stackname *StackName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "gongsvg":
		*stackname = StackNameDefault
	default:
		return errUnkownEnum
	}
	return
}

func (stackname *StackName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "StackNameDefault":
		*stackname = StackNameDefault
	default:
		return errUnkownEnum
	}
	return
}

func (stackname *StackName) ToCodeString() (res string) {

	switch *stackname {
	// insertion code per enum code
	case StackNameDefault:
		res = "StackNameDefault"
	}
	return
}

func (stackname StackName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "StackNameDefault")

	return
}

func (stackname StackName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "gongsvg")

	return
}

// Utility function for TextAnchorType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (textanchortype TextAnchorType) ToString() (res string) {

	// migration of former implementation of enum
	switch textanchortype {
	// insertion code per enum code
	case TEXT_ANCHOR_START:
		res = "start"
	case TEXT_ANCHOR_CENTER:
		res = "middle"
	case TEXT_ANCHOR_END:
		res = "end"
	}
	return
}

func (textanchortype *TextAnchorType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "start":
		*textanchortype = TEXT_ANCHOR_START
	case "middle":
		*textanchortype = TEXT_ANCHOR_CENTER
	case "end":
		*textanchortype = TEXT_ANCHOR_END
	default:
		return errUnkownEnum
	}
	return
}

func (textanchortype *TextAnchorType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TEXT_ANCHOR_START":
		*textanchortype = TEXT_ANCHOR_START
	case "TEXT_ANCHOR_CENTER":
		*textanchortype = TEXT_ANCHOR_CENTER
	case "TEXT_ANCHOR_END":
		*textanchortype = TEXT_ANCHOR_END
	default:
		return errUnkownEnum
	}
	return
}

func (textanchortype *TextAnchorType) ToCodeString() (res string) {

	switch *textanchortype {
	// insertion code per enum code
	case TEXT_ANCHOR_START:
		res = "TEXT_ANCHOR_START"
	case TEXT_ANCHOR_CENTER:
		res = "TEXT_ANCHOR_CENTER"
	case TEXT_ANCHOR_END:
		res = "TEXT_ANCHOR_END"
	}
	return
}

func (textanchortype TextAnchorType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TEXT_ANCHOR_START")
	res = append(res, "TEXT_ANCHOR_CENTER")
	res = append(res, "TEXT_ANCHOR_END")

	return
}

func (textanchortype TextAnchorType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "start")
	res = append(res, "middle")
	res = append(res, "end")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | AnchorType | ColorType | DrawingState | LinkAnchorType | LinkType | OrientationType | PositionOnArrowType | RectAnchorType | SideType | StackName | TextAnchorType
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*AnchorType | *ColorType | *DrawingState | *LinkAnchorType | *LinkType | *OrientationType | *PositionOnArrowType | *RectAnchorType | *SideType | *StackName | *TextAnchorType
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
