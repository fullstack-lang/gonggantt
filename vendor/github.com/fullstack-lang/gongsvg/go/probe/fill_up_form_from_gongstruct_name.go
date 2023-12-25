// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "Animate":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Animate Form",
			OnSave: __gong__New__AnimateFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		animate := new(models.Animate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(animate, formGroup, probe)
	case "Circle":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Circle Form",
			OnSave: __gong__New__CircleFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		circle := new(models.Circle)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(circle, formGroup, probe)
	case "Ellipse":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Ellipse Form",
			OnSave: __gong__New__EllipseFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		ellipse := new(models.Ellipse)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(ellipse, formGroup, probe)
	case "Layer":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Layer Form",
			OnSave: __gong__New__LayerFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		layer := new(models.Layer)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(layer, formGroup, probe)
	case "Line":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Line Form",
			OnSave: __gong__New__LineFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		line := new(models.Line)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(line, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Link Form",
			OnSave: __gong__New__LinkFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	case "LinkAnchoredText":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "LinkAnchoredText Form",
			OnSave: __gong__New__LinkAnchoredTextFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		linkanchoredtext := new(models.LinkAnchoredText)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(linkanchoredtext, formGroup, probe)
	case "Path":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Path Form",
			OnSave: __gong__New__PathFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		path := new(models.Path)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(path, formGroup, probe)
	case "Point":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Point Form",
			OnSave: __gong__New__PointFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		point := new(models.Point)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(point, formGroup, probe)
	case "Polygone":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Polygone Form",
			OnSave: __gong__New__PolygoneFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		polygone := new(models.Polygone)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(polygone, formGroup, probe)
	case "Polyline":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Polyline Form",
			OnSave: __gong__New__PolylineFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		polyline := new(models.Polyline)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(polyline, formGroup, probe)
	case "Rect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Rect Form",
			OnSave: __gong__New__RectFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		rect := new(models.Rect)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rect, formGroup, probe)
	case "RectAnchoredPath":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RectAnchoredPath Form",
			OnSave: __gong__New__RectAnchoredPathFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		rectanchoredpath := new(models.RectAnchoredPath)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectanchoredpath, formGroup, probe)
	case "RectAnchoredRect":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RectAnchoredRect Form",
			OnSave: __gong__New__RectAnchoredRectFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		rectanchoredrect := new(models.RectAnchoredRect)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectanchoredrect, formGroup, probe)
	case "RectAnchoredText":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RectAnchoredText Form",
			OnSave: __gong__New__RectAnchoredTextFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		rectanchoredtext := new(models.RectAnchoredText)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectanchoredtext, formGroup, probe)
	case "RectLinkLink":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "RectLinkLink Form",
			OnSave: __gong__New__RectLinkLinkFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		rectlinklink := new(models.RectLinkLink)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rectlinklink, formGroup, probe)
	case "SVG":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "SVG Form",
			OnSave: __gong__New__SVGFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		svg := new(models.SVG)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(svg, formGroup, probe)
	case "Text":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + "Text Form",
			OnSave: __gong__New__TextFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		text := new(models.Text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(text, formGroup, probe)
	}
	formStage.Commit()
}
