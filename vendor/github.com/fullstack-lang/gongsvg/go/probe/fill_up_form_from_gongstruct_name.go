// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Animate":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Animate Form",
			OnSave: __gong__New__AnimateFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		animate := new(models.Animate)
		FillUpForm(animate, formGroup, playground)
	case "Circle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Circle Form",
			OnSave: __gong__New__CircleFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		circle := new(models.Circle)
		FillUpForm(circle, formGroup, playground)
	case "Ellipse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Ellipse Form",
			OnSave: __gong__New__EllipseFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		ellipse := new(models.Ellipse)
		FillUpForm(ellipse, formGroup, playground)
	case "Layer":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Layer Form",
			OnSave: __gong__New__LayerFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		layer := new(models.Layer)
		FillUpForm(layer, formGroup, playground)
	case "Line":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Line Form",
			OnSave: __gong__New__LineFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		line := new(models.Line)
		FillUpForm(line, formGroup, playground)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Link Form",
			OnSave: __gong__New__LinkFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		link := new(models.Link)
		FillUpForm(link, formGroup, playground)
	case "LinkAnchoredText":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " LinkAnchoredText Form",
			OnSave: __gong__New__LinkAnchoredTextFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		linkanchoredtext := new(models.LinkAnchoredText)
		FillUpForm(linkanchoredtext, formGroup, playground)
	case "Path":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Path Form",
			OnSave: __gong__New__PathFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		path := new(models.Path)
		FillUpForm(path, formGroup, playground)
	case "Point":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Point Form",
			OnSave: __gong__New__PointFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		point := new(models.Point)
		FillUpForm(point, formGroup, playground)
	case "Polygone":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Polygone Form",
			OnSave: __gong__New__PolygoneFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		polygone := new(models.Polygone)
		FillUpForm(polygone, formGroup, playground)
	case "Polyline":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Polyline Form",
			OnSave: __gong__New__PolylineFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		polyline := new(models.Polyline)
		FillUpForm(polyline, formGroup, playground)
	case "Rect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Rect Form",
			OnSave: __gong__New__RectFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		rect := new(models.Rect)
		FillUpForm(rect, formGroup, playground)
	case "RectAnchoredRect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " RectAnchoredRect Form",
			OnSave: __gong__New__RectAnchoredRectFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		rectanchoredrect := new(models.RectAnchoredRect)
		FillUpForm(rectanchoredrect, formGroup, playground)
	case "RectAnchoredText":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " RectAnchoredText Form",
			OnSave: __gong__New__RectAnchoredTextFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		rectanchoredtext := new(models.RectAnchoredText)
		FillUpForm(rectanchoredtext, formGroup, playground)
	case "RectLinkLink":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " RectLinkLink Form",
			OnSave: __gong__New__RectLinkLinkFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		rectlinklink := new(models.RectLinkLink)
		FillUpForm(rectlinklink, formGroup, playground)
	case "SVG":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " SVG Form",
			OnSave: __gong__New__SVGFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		svg := new(models.SVG)
		FillUpForm(svg, formGroup, playground)
	case "Text":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Text Form",
			OnSave: __gong__New__TextFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		text := new(models.Text)
		FillUpForm(text, formGroup, playground)
	}
	formStage.Commit()
}
