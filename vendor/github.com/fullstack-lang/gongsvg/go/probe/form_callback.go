// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
	"github.com/fullstack-lang/gongsvg/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__AnimateFormCallback(
	animate *models.Animate,
	playground *Playground,
) (animateFormCallback *AnimateFormCallback) {
	animateFormCallback = new(AnimateFormCallback)
	animateFormCallback.playground = playground
	animateFormCallback.animate = animate

	animateFormCallback.CreationMode = (animate == nil)

	return
}

type AnimateFormCallback struct {
	animate *models.Animate

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (animateFormCallback *AnimateFormCallback) OnSave() {

	log.Println("AnimateFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	animateFormCallback.playground.formStage.Checkout()

	if animateFormCallback.animate == nil {
		animateFormCallback.animate = new(models.Animate).Stage(animateFormCallback.playground.stageOfInterest)
	}
	animate_ := animateFormCallback.animate
	_ = animate_

	// get the formGroup
	formGroup := animateFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(animate_.Name), formDiv)
		case "AttributeName":
			FormDivBasicFieldToField(&(animate_.AttributeName), formDiv)
		case "Values":
			FormDivBasicFieldToField(&(animate_.Values), formDiv)
		case "Dur":
			FormDivBasicFieldToField(&(animate_.Dur), formDiv)
		case "RepeatCount":
			FormDivBasicFieldToField(&(animate_.RepeatCount), formDiv)
		case "Circle:Animations":
			// we need to retrieve the field owner before the change
			var pastCircleOwner *models.Circle
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Circle"
			rf.Fieldname = "Animations"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastCircleOwner = reverseFieldOwner.(*models.Circle)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCircleOwner != nil {
					idx := slices.Index(pastCircleOwner.Animations, animate_)
					pastCircleOwner.Animations = slices.Delete(pastCircleOwner.Animations, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _circle := range *models.GetGongstructInstancesSet[models.Circle](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _circle.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCircleOwner := _circle // we have a match
						if pastCircleOwner != nil {
							if newCircleOwner != pastCircleOwner {
								idx := slices.Index(pastCircleOwner.Animations, animate_)
								pastCircleOwner.Animations = slices.Delete(pastCircleOwner.Animations, idx, idx+1)
								newCircleOwner.Animations = append(newCircleOwner.Animations, animate_)
							}
						} else {
							newCircleOwner.Animations = append(newCircleOwner.Animations, animate_)
						}
					}
				}
			}
		case "Ellipse:Animates":
			// we need to retrieve the field owner before the change
			var pastEllipseOwner *models.Ellipse
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Ellipse"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastEllipseOwner = reverseFieldOwner.(*models.Ellipse)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastEllipseOwner != nil {
					idx := slices.Index(pastEllipseOwner.Animates, animate_)
					pastEllipseOwner.Animates = slices.Delete(pastEllipseOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _ellipse := range *models.GetGongstructInstancesSet[models.Ellipse](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _ellipse.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newEllipseOwner := _ellipse // we have a match
						if pastEllipseOwner != nil {
							if newEllipseOwner != pastEllipseOwner {
								idx := slices.Index(pastEllipseOwner.Animates, animate_)
								pastEllipseOwner.Animates = slices.Delete(pastEllipseOwner.Animates, idx, idx+1)
								newEllipseOwner.Animates = append(newEllipseOwner.Animates, animate_)
							}
						} else {
							newEllipseOwner.Animates = append(newEllipseOwner.Animates, animate_)
						}
					}
				}
			}
		case "Line:Animates":
			// we need to retrieve the field owner before the change
			var pastLineOwner *models.Line
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Line"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastLineOwner = reverseFieldOwner.(*models.Line)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLineOwner != nil {
					idx := slices.Index(pastLineOwner.Animates, animate_)
					pastLineOwner.Animates = slices.Delete(pastLineOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _line := range *models.GetGongstructInstancesSet[models.Line](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _line.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLineOwner := _line // we have a match
						if pastLineOwner != nil {
							if newLineOwner != pastLineOwner {
								idx := slices.Index(pastLineOwner.Animates, animate_)
								pastLineOwner.Animates = slices.Delete(pastLineOwner.Animates, idx, idx+1)
								newLineOwner.Animates = append(newLineOwner.Animates, animate_)
							}
						} else {
							newLineOwner.Animates = append(newLineOwner.Animates, animate_)
						}
					}
				}
			}
		case "LinkAnchoredText:Animates":
			// we need to retrieve the field owner before the change
			var pastLinkAnchoredTextOwner *models.LinkAnchoredText
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "LinkAnchoredText"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastLinkAnchoredTextOwner = reverseFieldOwner.(*models.LinkAnchoredText)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLinkAnchoredTextOwner != nil {
					idx := slices.Index(pastLinkAnchoredTextOwner.Animates, animate_)
					pastLinkAnchoredTextOwner.Animates = slices.Delete(pastLinkAnchoredTextOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _linkanchoredtext := range *models.GetGongstructInstancesSet[models.LinkAnchoredText](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _linkanchoredtext.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLinkAnchoredTextOwner := _linkanchoredtext // we have a match
						if pastLinkAnchoredTextOwner != nil {
							if newLinkAnchoredTextOwner != pastLinkAnchoredTextOwner {
								idx := slices.Index(pastLinkAnchoredTextOwner.Animates, animate_)
								pastLinkAnchoredTextOwner.Animates = slices.Delete(pastLinkAnchoredTextOwner.Animates, idx, idx+1)
								newLinkAnchoredTextOwner.Animates = append(newLinkAnchoredTextOwner.Animates, animate_)
							}
						} else {
							newLinkAnchoredTextOwner.Animates = append(newLinkAnchoredTextOwner.Animates, animate_)
						}
					}
				}
			}
		case "Path:Animates":
			// we need to retrieve the field owner before the change
			var pastPathOwner *models.Path
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Path"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastPathOwner = reverseFieldOwner.(*models.Path)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastPathOwner != nil {
					idx := slices.Index(pastPathOwner.Animates, animate_)
					pastPathOwner.Animates = slices.Delete(pastPathOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _path := range *models.GetGongstructInstancesSet[models.Path](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _path.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newPathOwner := _path // we have a match
						if pastPathOwner != nil {
							if newPathOwner != pastPathOwner {
								idx := slices.Index(pastPathOwner.Animates, animate_)
								pastPathOwner.Animates = slices.Delete(pastPathOwner.Animates, idx, idx+1)
								newPathOwner.Animates = append(newPathOwner.Animates, animate_)
							}
						} else {
							newPathOwner.Animates = append(newPathOwner.Animates, animate_)
						}
					}
				}
			}
		case "Polygone:Animates":
			// we need to retrieve the field owner before the change
			var pastPolygoneOwner *models.Polygone
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Polygone"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastPolygoneOwner = reverseFieldOwner.(*models.Polygone)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastPolygoneOwner != nil {
					idx := slices.Index(pastPolygoneOwner.Animates, animate_)
					pastPolygoneOwner.Animates = slices.Delete(pastPolygoneOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _polygone := range *models.GetGongstructInstancesSet[models.Polygone](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _polygone.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newPolygoneOwner := _polygone // we have a match
						if pastPolygoneOwner != nil {
							if newPolygoneOwner != pastPolygoneOwner {
								idx := slices.Index(pastPolygoneOwner.Animates, animate_)
								pastPolygoneOwner.Animates = slices.Delete(pastPolygoneOwner.Animates, idx, idx+1)
								newPolygoneOwner.Animates = append(newPolygoneOwner.Animates, animate_)
							}
						} else {
							newPolygoneOwner.Animates = append(newPolygoneOwner.Animates, animate_)
						}
					}
				}
			}
		case "Polyline:Animates":
			// we need to retrieve the field owner before the change
			var pastPolylineOwner *models.Polyline
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Polyline"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastPolylineOwner = reverseFieldOwner.(*models.Polyline)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastPolylineOwner != nil {
					idx := slices.Index(pastPolylineOwner.Animates, animate_)
					pastPolylineOwner.Animates = slices.Delete(pastPolylineOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _polyline := range *models.GetGongstructInstancesSet[models.Polyline](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _polyline.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newPolylineOwner := _polyline // we have a match
						if pastPolylineOwner != nil {
							if newPolylineOwner != pastPolylineOwner {
								idx := slices.Index(pastPolylineOwner.Animates, animate_)
								pastPolylineOwner.Animates = slices.Delete(pastPolylineOwner.Animates, idx, idx+1)
								newPolylineOwner.Animates = append(newPolylineOwner.Animates, animate_)
							}
						} else {
							newPolylineOwner.Animates = append(newPolylineOwner.Animates, animate_)
						}
					}
				}
			}
		case "Rect:Animations":
			// we need to retrieve the field owner before the change
			var pastRectOwner *models.Rect
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "Animations"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastRectOwner = reverseFieldOwner.(*models.Rect)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRectOwner != nil {
					idx := slices.Index(pastRectOwner.Animations, animate_)
					pastRectOwner.Animations = slices.Delete(pastRectOwner.Animations, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _rect := range *models.GetGongstructInstancesSet[models.Rect](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _rect.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRectOwner := _rect // we have a match
						if pastRectOwner != nil {
							if newRectOwner != pastRectOwner {
								idx := slices.Index(pastRectOwner.Animations, animate_)
								pastRectOwner.Animations = slices.Delete(pastRectOwner.Animations, idx, idx+1)
								newRectOwner.Animations = append(newRectOwner.Animations, animate_)
							}
						} else {
							newRectOwner.Animations = append(newRectOwner.Animations, animate_)
						}
					}
				}
			}
		case "RectAnchoredText:Animates":
			// we need to retrieve the field owner before the change
			var pastRectAnchoredTextOwner *models.RectAnchoredText
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "RectAnchoredText"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastRectAnchoredTextOwner = reverseFieldOwner.(*models.RectAnchoredText)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRectAnchoredTextOwner != nil {
					idx := slices.Index(pastRectAnchoredTextOwner.Animates, animate_)
					pastRectAnchoredTextOwner.Animates = slices.Delete(pastRectAnchoredTextOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _rectanchoredtext := range *models.GetGongstructInstancesSet[models.RectAnchoredText](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _rectanchoredtext.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRectAnchoredTextOwner := _rectanchoredtext // we have a match
						if pastRectAnchoredTextOwner != nil {
							if newRectAnchoredTextOwner != pastRectAnchoredTextOwner {
								idx := slices.Index(pastRectAnchoredTextOwner.Animates, animate_)
								pastRectAnchoredTextOwner.Animates = slices.Delete(pastRectAnchoredTextOwner.Animates, idx, idx+1)
								newRectAnchoredTextOwner.Animates = append(newRectAnchoredTextOwner.Animates, animate_)
							}
						} else {
							newRectAnchoredTextOwner.Animates = append(newRectAnchoredTextOwner.Animates, animate_)
						}
					}
				}
			}
		case "Text:Animates":
			// we need to retrieve the field owner before the change
			var pastTextOwner *models.Text
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Text"
			rf.Fieldname = "Animates"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				animateFormCallback.playground.stageOfInterest,
				animateFormCallback.playground.backRepoOfInterest,
				animate_,
				&rf)

			if reverseFieldOwner != nil {
				pastTextOwner = reverseFieldOwner.(*models.Text)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTextOwner != nil {
					idx := slices.Index(pastTextOwner.Animates, animate_)
					pastTextOwner.Animates = slices.Delete(pastTextOwner.Animates, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _text := range *models.GetGongstructInstancesSet[models.Text](animateFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _text.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTextOwner := _text // we have a match
						if pastTextOwner != nil {
							if newTextOwner != pastTextOwner {
								idx := slices.Index(pastTextOwner.Animates, animate_)
								pastTextOwner.Animates = slices.Delete(pastTextOwner.Animates, idx, idx+1)
								newTextOwner.Animates = append(newTextOwner.Animates, animate_)
							}
						} else {
							newTextOwner.Animates = append(newTextOwner.Animates, animate_)
						}
					}
				}
			}
		}
	}

	animateFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Animate](
		animateFormCallback.playground,
	)
	animateFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if animateFormCallback.CreationMode {
		animateFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__AnimateFormCallback(
				nil,
				animateFormCallback.playground,
			),
		}).Stage(animateFormCallback.playground.formStage)
		animate := new(models.Animate)
		FillUpForm(animate, newFormGroup, animateFormCallback.playground)
		animateFormCallback.playground.formStage.Commit()
	}

	fillUpTree(animateFormCallback.playground)
}
func __gong__New__CircleFormCallback(
	circle *models.Circle,
	playground *Playground,
) (circleFormCallback *CircleFormCallback) {
	circleFormCallback = new(CircleFormCallback)
	circleFormCallback.playground = playground
	circleFormCallback.circle = circle

	circleFormCallback.CreationMode = (circle == nil)

	return
}

type CircleFormCallback struct {
	circle *models.Circle

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (circleFormCallback *CircleFormCallback) OnSave() {

	log.Println("CircleFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	circleFormCallback.playground.formStage.Checkout()

	if circleFormCallback.circle == nil {
		circleFormCallback.circle = new(models.Circle).Stage(circleFormCallback.playground.stageOfInterest)
	}
	circle_ := circleFormCallback.circle
	_ = circle_

	// get the formGroup
	formGroup := circleFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(circle_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(circle_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(circle_.CY), formDiv)
		case "Radius":
			FormDivBasicFieldToField(&(circle_.Radius), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(circle_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(circle_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(circle_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(circle_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(circle_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(circle_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(circle_.Transform), formDiv)
		case "Layer:Circles":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Circles"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				circleFormCallback.playground.stageOfInterest,
				circleFormCallback.playground.backRepoOfInterest,
				circle_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Circles, circle_)
					pastLayerOwner.Circles = slices.Delete(pastLayerOwner.Circles, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](circleFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Circles, circle_)
								pastLayerOwner.Circles = slices.Delete(pastLayerOwner.Circles, idx, idx+1)
								newLayerOwner.Circles = append(newLayerOwner.Circles, circle_)
							}
						} else {
							newLayerOwner.Circles = append(newLayerOwner.Circles, circle_)
						}
					}
				}
			}
		}
	}

	circleFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Circle](
		circleFormCallback.playground,
	)
	circleFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if circleFormCallback.CreationMode {
		circleFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CircleFormCallback(
				nil,
				circleFormCallback.playground,
			),
		}).Stage(circleFormCallback.playground.formStage)
		circle := new(models.Circle)
		FillUpForm(circle, newFormGroup, circleFormCallback.playground)
		circleFormCallback.playground.formStage.Commit()
	}

	fillUpTree(circleFormCallback.playground)
}
func __gong__New__EllipseFormCallback(
	ellipse *models.Ellipse,
	playground *Playground,
) (ellipseFormCallback *EllipseFormCallback) {
	ellipseFormCallback = new(EllipseFormCallback)
	ellipseFormCallback.playground = playground
	ellipseFormCallback.ellipse = ellipse

	ellipseFormCallback.CreationMode = (ellipse == nil)

	return
}

type EllipseFormCallback struct {
	ellipse *models.Ellipse

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (ellipseFormCallback *EllipseFormCallback) OnSave() {

	log.Println("EllipseFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	ellipseFormCallback.playground.formStage.Checkout()

	if ellipseFormCallback.ellipse == nil {
		ellipseFormCallback.ellipse = new(models.Ellipse).Stage(ellipseFormCallback.playground.stageOfInterest)
	}
	ellipse_ := ellipseFormCallback.ellipse
	_ = ellipse_

	// get the formGroup
	formGroup := ellipseFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(ellipse_.Name), formDiv)
		case "CX":
			FormDivBasicFieldToField(&(ellipse_.CX), formDiv)
		case "CY":
			FormDivBasicFieldToField(&(ellipse_.CY), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(ellipse_.RX), formDiv)
		case "RY":
			FormDivBasicFieldToField(&(ellipse_.RY), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(ellipse_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(ellipse_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(ellipse_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(ellipse_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(ellipse_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(ellipse_.Transform), formDiv)
		case "Layer:Ellipses":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Ellipses"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				ellipseFormCallback.playground.stageOfInterest,
				ellipseFormCallback.playground.backRepoOfInterest,
				ellipse_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Ellipses, ellipse_)
					pastLayerOwner.Ellipses = slices.Delete(pastLayerOwner.Ellipses, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](ellipseFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Ellipses, ellipse_)
								pastLayerOwner.Ellipses = slices.Delete(pastLayerOwner.Ellipses, idx, idx+1)
								newLayerOwner.Ellipses = append(newLayerOwner.Ellipses, ellipse_)
							}
						} else {
							newLayerOwner.Ellipses = append(newLayerOwner.Ellipses, ellipse_)
						}
					}
				}
			}
		}
	}

	ellipseFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Ellipse](
		ellipseFormCallback.playground,
	)
	ellipseFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if ellipseFormCallback.CreationMode {
		ellipseFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__EllipseFormCallback(
				nil,
				ellipseFormCallback.playground,
			),
		}).Stage(ellipseFormCallback.playground.formStage)
		ellipse := new(models.Ellipse)
		FillUpForm(ellipse, newFormGroup, ellipseFormCallback.playground)
		ellipseFormCallback.playground.formStage.Commit()
	}

	fillUpTree(ellipseFormCallback.playground)
}
func __gong__New__LayerFormCallback(
	layer *models.Layer,
	playground *Playground,
) (layerFormCallback *LayerFormCallback) {
	layerFormCallback = new(LayerFormCallback)
	layerFormCallback.playground = playground
	layerFormCallback.layer = layer

	layerFormCallback.CreationMode = (layer == nil)

	return
}

type LayerFormCallback struct {
	layer *models.Layer

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (layerFormCallback *LayerFormCallback) OnSave() {

	log.Println("LayerFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	layerFormCallback.playground.formStage.Checkout()

	if layerFormCallback.layer == nil {
		layerFormCallback.layer = new(models.Layer).Stage(layerFormCallback.playground.stageOfInterest)
	}
	layer_ := layerFormCallback.layer
	_ = layer_

	// get the formGroup
	formGroup := layerFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Display":
			FormDivBasicFieldToField(&(layer_.Display), formDiv)
		case "Name":
			FormDivBasicFieldToField(&(layer_.Name), formDiv)
		case "SVG:Layers":
			// we need to retrieve the field owner before the change
			var pastSVGOwner *models.SVG
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "SVG"
			rf.Fieldname = "Layers"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				layerFormCallback.playground.stageOfInterest,
				layerFormCallback.playground.backRepoOfInterest,
				layer_,
				&rf)

			if reverseFieldOwner != nil {
				pastSVGOwner = reverseFieldOwner.(*models.SVG)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastSVGOwner != nil {
					idx := slices.Index(pastSVGOwner.Layers, layer_)
					pastSVGOwner.Layers = slices.Delete(pastSVGOwner.Layers, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _svg := range *models.GetGongstructInstancesSet[models.SVG](layerFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _svg.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newSVGOwner := _svg // we have a match
						if pastSVGOwner != nil {
							if newSVGOwner != pastSVGOwner {
								idx := slices.Index(pastSVGOwner.Layers, layer_)
								pastSVGOwner.Layers = slices.Delete(pastSVGOwner.Layers, idx, idx+1)
								newSVGOwner.Layers = append(newSVGOwner.Layers, layer_)
							}
						} else {
							newSVGOwner.Layers = append(newSVGOwner.Layers, layer_)
						}
					}
				}
			}
		}
	}

	layerFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Layer](
		layerFormCallback.playground,
	)
	layerFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if layerFormCallback.CreationMode {
		layerFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LayerFormCallback(
				nil,
				layerFormCallback.playground,
			),
		}).Stage(layerFormCallback.playground.formStage)
		layer := new(models.Layer)
		FillUpForm(layer, newFormGroup, layerFormCallback.playground)
		layerFormCallback.playground.formStage.Commit()
	}

	fillUpTree(layerFormCallback.playground)
}
func __gong__New__LineFormCallback(
	line *models.Line,
	playground *Playground,
) (lineFormCallback *LineFormCallback) {
	lineFormCallback = new(LineFormCallback)
	lineFormCallback.playground = playground
	lineFormCallback.line = line

	lineFormCallback.CreationMode = (line == nil)

	return
}

type LineFormCallback struct {
	line *models.Line

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (lineFormCallback *LineFormCallback) OnSave() {

	log.Println("LineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	lineFormCallback.playground.formStage.Checkout()

	if lineFormCallback.line == nil {
		lineFormCallback.line = new(models.Line).Stage(lineFormCallback.playground.stageOfInterest)
	}
	line_ := lineFormCallback.line
	_ = line_

	// get the formGroup
	formGroup := lineFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(line_.Name), formDiv)
		case "X1":
			FormDivBasicFieldToField(&(line_.X1), formDiv)
		case "Y1":
			FormDivBasicFieldToField(&(line_.Y1), formDiv)
		case "X2":
			FormDivBasicFieldToField(&(line_.X2), formDiv)
		case "Y2":
			FormDivBasicFieldToField(&(line_.Y2), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(line_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(line_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(line_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(line_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(line_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(line_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(line_.Transform), formDiv)
		case "MouseClickX":
			FormDivBasicFieldToField(&(line_.MouseClickX), formDiv)
		case "MouseClickY":
			FormDivBasicFieldToField(&(line_.MouseClickY), formDiv)
		case "Layer:Lines":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Lines"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				lineFormCallback.playground.stageOfInterest,
				lineFormCallback.playground.backRepoOfInterest,
				line_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Lines, line_)
					pastLayerOwner.Lines = slices.Delete(pastLayerOwner.Lines, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](lineFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Lines, line_)
								pastLayerOwner.Lines = slices.Delete(pastLayerOwner.Lines, idx, idx+1)
								newLayerOwner.Lines = append(newLayerOwner.Lines, line_)
							}
						} else {
							newLayerOwner.Lines = append(newLayerOwner.Lines, line_)
						}
					}
				}
			}
		}
	}

	lineFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Line](
		lineFormCallback.playground,
	)
	lineFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if lineFormCallback.CreationMode {
		lineFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LineFormCallback(
				nil,
				lineFormCallback.playground,
			),
		}).Stage(lineFormCallback.playground.formStage)
		line := new(models.Line)
		FillUpForm(line, newFormGroup, lineFormCallback.playground)
		lineFormCallback.playground.formStage.Commit()
	}

	fillUpTree(lineFormCallback.playground)
}
func __gong__New__LinkFormCallback(
	link *models.Link,
	playground *Playground,
) (linkFormCallback *LinkFormCallback) {
	linkFormCallback = new(LinkFormCallback)
	linkFormCallback.playground = playground
	linkFormCallback.link = link

	linkFormCallback.CreationMode = (link == nil)

	return
}

type LinkFormCallback struct {
	link *models.Link

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (linkFormCallback *LinkFormCallback) OnSave() {

	log.Println("LinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkFormCallback.playground.formStage.Checkout()

	if linkFormCallback.link == nil {
		linkFormCallback.link = new(models.Link).Stage(linkFormCallback.playground.stageOfInterest)
	}
	link_ := linkFormCallback.link
	_ = link_

	// get the formGroup
	formGroup := linkFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(link_.Name), formDiv)
		case "Type":
			FormDivEnumStringFieldToField(&(link_.Type), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(link_.Start), linkFormCallback.playground.stageOfInterest, formDiv)
		case "StartAnchorType":
			FormDivEnumStringFieldToField(&(link_.StartAnchorType), formDiv)
		case "End":
			FormDivSelectFieldToField(&(link_.End), linkFormCallback.playground.stageOfInterest, formDiv)
		case "EndAnchorType":
			FormDivEnumStringFieldToField(&(link_.EndAnchorType), formDiv)
		case "StartOrientation":
			FormDivEnumStringFieldToField(&(link_.StartOrientation), formDiv)
		case "StartRatio":
			FormDivBasicFieldToField(&(link_.StartRatio), formDiv)
		case "EndOrientation":
			FormDivEnumStringFieldToField(&(link_.EndOrientation), formDiv)
		case "EndRatio":
			FormDivBasicFieldToField(&(link_.EndRatio), formDiv)
		case "CornerOffsetRatio":
			FormDivBasicFieldToField(&(link_.CornerOffsetRatio), formDiv)
		case "CornerRadius":
			FormDivBasicFieldToField(&(link_.CornerRadius), formDiv)
		case "HasEndArrow":
			FormDivBasicFieldToField(&(link_.HasEndArrow), formDiv)
		case "EndArrowSize":
			FormDivBasicFieldToField(&(link_.EndArrowSize), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(link_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(link_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(link_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(link_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(link_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(link_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(link_.Transform), formDiv)
		case "Layer:Links":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Links"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				linkFormCallback.playground.stageOfInterest,
				linkFormCallback.playground.backRepoOfInterest,
				link_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Links, link_)
					pastLayerOwner.Links = slices.Delete(pastLayerOwner.Links, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](linkFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Links, link_)
								pastLayerOwner.Links = slices.Delete(pastLayerOwner.Links, idx, idx+1)
								newLayerOwner.Links = append(newLayerOwner.Links, link_)
							}
						} else {
							newLayerOwner.Links = append(newLayerOwner.Links, link_)
						}
					}
				}
			}
		}
	}

	linkFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Link](
		linkFormCallback.playground,
	)
	linkFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if linkFormCallback.CreationMode {
		linkFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LinkFormCallback(
				nil,
				linkFormCallback.playground,
			),
		}).Stage(linkFormCallback.playground.formStage)
		link := new(models.Link)
		FillUpForm(link, newFormGroup, linkFormCallback.playground)
		linkFormCallback.playground.formStage.Commit()
	}

	fillUpTree(linkFormCallback.playground)
}
func __gong__New__LinkAnchoredTextFormCallback(
	linkanchoredtext *models.LinkAnchoredText,
	playground *Playground,
) (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) {
	linkanchoredtextFormCallback = new(LinkAnchoredTextFormCallback)
	linkanchoredtextFormCallback.playground = playground
	linkanchoredtextFormCallback.linkanchoredtext = linkanchoredtext

	linkanchoredtextFormCallback.CreationMode = (linkanchoredtext == nil)

	return
}

type LinkAnchoredTextFormCallback struct {
	linkanchoredtext *models.LinkAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (linkanchoredtextFormCallback *LinkAnchoredTextFormCallback) OnSave() {

	log.Println("LinkAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	linkanchoredtextFormCallback.playground.formStage.Checkout()

	if linkanchoredtextFormCallback.linkanchoredtext == nil {
		linkanchoredtextFormCallback.linkanchoredtext = new(models.LinkAnchoredText).Stage(linkanchoredtextFormCallback.playground.stageOfInterest)
	}
	linkanchoredtext_ := linkanchoredtextFormCallback.linkanchoredtext
	_ = linkanchoredtext_

	// get the formGroup
	formGroup := linkanchoredtextFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(linkanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(linkanchoredtext_.Content), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(linkanchoredtext_.Y_Offset), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(linkanchoredtext_.FontWeight), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(linkanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(linkanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(linkanchoredtext_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(linkanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(linkanchoredtext_.Transform), formDiv)
		case "Link:TextAtArrowEnd":
			// we need to retrieve the field owner before the change
			var pastLinkOwner *models.Link
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "TextAtArrowEnd"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				linkanchoredtextFormCallback.playground.stageOfInterest,
				linkanchoredtextFormCallback.playground.backRepoOfInterest,
				linkanchoredtext_,
				&rf)

			if reverseFieldOwner != nil {
				pastLinkOwner = reverseFieldOwner.(*models.Link)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLinkOwner != nil {
					idx := slices.Index(pastLinkOwner.TextAtArrowEnd, linkanchoredtext_)
					pastLinkOwner.TextAtArrowEnd = slices.Delete(pastLinkOwner.TextAtArrowEnd, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _link := range *models.GetGongstructInstancesSet[models.Link](linkanchoredtextFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _link.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLinkOwner := _link // we have a match
						if pastLinkOwner != nil {
							if newLinkOwner != pastLinkOwner {
								idx := slices.Index(pastLinkOwner.TextAtArrowEnd, linkanchoredtext_)
								pastLinkOwner.TextAtArrowEnd = slices.Delete(pastLinkOwner.TextAtArrowEnd, idx, idx+1)
								newLinkOwner.TextAtArrowEnd = append(newLinkOwner.TextAtArrowEnd, linkanchoredtext_)
							}
						} else {
							newLinkOwner.TextAtArrowEnd = append(newLinkOwner.TextAtArrowEnd, linkanchoredtext_)
						}
					}
				}
			}
		case "Link:TextAtArrowStart":
			// we need to retrieve the field owner before the change
			var pastLinkOwner *models.Link
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "TextAtArrowStart"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				linkanchoredtextFormCallback.playground.stageOfInterest,
				linkanchoredtextFormCallback.playground.backRepoOfInterest,
				linkanchoredtext_,
				&rf)

			if reverseFieldOwner != nil {
				pastLinkOwner = reverseFieldOwner.(*models.Link)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLinkOwner != nil {
					idx := slices.Index(pastLinkOwner.TextAtArrowStart, linkanchoredtext_)
					pastLinkOwner.TextAtArrowStart = slices.Delete(pastLinkOwner.TextAtArrowStart, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _link := range *models.GetGongstructInstancesSet[models.Link](linkanchoredtextFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _link.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLinkOwner := _link // we have a match
						if pastLinkOwner != nil {
							if newLinkOwner != pastLinkOwner {
								idx := slices.Index(pastLinkOwner.TextAtArrowStart, linkanchoredtext_)
								pastLinkOwner.TextAtArrowStart = slices.Delete(pastLinkOwner.TextAtArrowStart, idx, idx+1)
								newLinkOwner.TextAtArrowStart = append(newLinkOwner.TextAtArrowStart, linkanchoredtext_)
							}
						} else {
							newLinkOwner.TextAtArrowStart = append(newLinkOwner.TextAtArrowStart, linkanchoredtext_)
						}
					}
				}
			}
		}
	}

	linkanchoredtextFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.LinkAnchoredText](
		linkanchoredtextFormCallback.playground,
	)
	linkanchoredtextFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if linkanchoredtextFormCallback.CreationMode {
		linkanchoredtextFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__LinkAnchoredTextFormCallback(
				nil,
				linkanchoredtextFormCallback.playground,
			),
		}).Stage(linkanchoredtextFormCallback.playground.formStage)
		linkanchoredtext := new(models.LinkAnchoredText)
		FillUpForm(linkanchoredtext, newFormGroup, linkanchoredtextFormCallback.playground)
		linkanchoredtextFormCallback.playground.formStage.Commit()
	}

	fillUpTree(linkanchoredtextFormCallback.playground)
}
func __gong__New__PathFormCallback(
	path *models.Path,
	playground *Playground,
) (pathFormCallback *PathFormCallback) {
	pathFormCallback = new(PathFormCallback)
	pathFormCallback.playground = playground
	pathFormCallback.path = path

	pathFormCallback.CreationMode = (path == nil)

	return
}

type PathFormCallback struct {
	path *models.Path

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (pathFormCallback *PathFormCallback) OnSave() {

	log.Println("PathFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pathFormCallback.playground.formStage.Checkout()

	if pathFormCallback.path == nil {
		pathFormCallback.path = new(models.Path).Stage(pathFormCallback.playground.stageOfInterest)
	}
	path_ := pathFormCallback.path
	_ = path_

	// get the formGroup
	formGroup := pathFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(path_.Name), formDiv)
		case "Definition":
			FormDivBasicFieldToField(&(path_.Definition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(path_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(path_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(path_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(path_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(path_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(path_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(path_.Transform), formDiv)
		case "Layer:Paths":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Paths"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				pathFormCallback.playground.stageOfInterest,
				pathFormCallback.playground.backRepoOfInterest,
				path_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Paths, path_)
					pastLayerOwner.Paths = slices.Delete(pastLayerOwner.Paths, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](pathFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Paths, path_)
								pastLayerOwner.Paths = slices.Delete(pastLayerOwner.Paths, idx, idx+1)
								newLayerOwner.Paths = append(newLayerOwner.Paths, path_)
							}
						} else {
							newLayerOwner.Paths = append(newLayerOwner.Paths, path_)
						}
					}
				}
			}
		}
	}

	pathFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Path](
		pathFormCallback.playground,
	)
	pathFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if pathFormCallback.CreationMode {
		pathFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__PathFormCallback(
				nil,
				pathFormCallback.playground,
			),
		}).Stage(pathFormCallback.playground.formStage)
		path := new(models.Path)
		FillUpForm(path, newFormGroup, pathFormCallback.playground)
		pathFormCallback.playground.formStage.Commit()
	}

	fillUpTree(pathFormCallback.playground)
}
func __gong__New__PointFormCallback(
	point *models.Point,
	playground *Playground,
) (pointFormCallback *PointFormCallback) {
	pointFormCallback = new(PointFormCallback)
	pointFormCallback.playground = playground
	pointFormCallback.point = point

	pointFormCallback.CreationMode = (point == nil)

	return
}

type PointFormCallback struct {
	point *models.Point

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (pointFormCallback *PointFormCallback) OnSave() {

	log.Println("PointFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	pointFormCallback.playground.formStage.Checkout()

	if pointFormCallback.point == nil {
		pointFormCallback.point = new(models.Point).Stage(pointFormCallback.playground.stageOfInterest)
	}
	point_ := pointFormCallback.point
	_ = point_

	// get the formGroup
	formGroup := pointFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(point_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(point_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(point_.Y), formDiv)
		case "Link:ControlPoints":
			// we need to retrieve the field owner before the change
			var pastLinkOwner *models.Link
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Link"
			rf.Fieldname = "ControlPoints"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				pointFormCallback.playground.stageOfInterest,
				pointFormCallback.playground.backRepoOfInterest,
				point_,
				&rf)

			if reverseFieldOwner != nil {
				pastLinkOwner = reverseFieldOwner.(*models.Link)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLinkOwner != nil {
					idx := slices.Index(pastLinkOwner.ControlPoints, point_)
					pastLinkOwner.ControlPoints = slices.Delete(pastLinkOwner.ControlPoints, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _link := range *models.GetGongstructInstancesSet[models.Link](pointFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _link.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLinkOwner := _link // we have a match
						if pastLinkOwner != nil {
							if newLinkOwner != pastLinkOwner {
								idx := slices.Index(pastLinkOwner.ControlPoints, point_)
								pastLinkOwner.ControlPoints = slices.Delete(pastLinkOwner.ControlPoints, idx, idx+1)
								newLinkOwner.ControlPoints = append(newLinkOwner.ControlPoints, point_)
							}
						} else {
							newLinkOwner.ControlPoints = append(newLinkOwner.ControlPoints, point_)
						}
					}
				}
			}
		}
	}

	pointFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Point](
		pointFormCallback.playground,
	)
	pointFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if pointFormCallback.CreationMode {
		pointFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__PointFormCallback(
				nil,
				pointFormCallback.playground,
			),
		}).Stage(pointFormCallback.playground.formStage)
		point := new(models.Point)
		FillUpForm(point, newFormGroup, pointFormCallback.playground)
		pointFormCallback.playground.formStage.Commit()
	}

	fillUpTree(pointFormCallback.playground)
}
func __gong__New__PolygoneFormCallback(
	polygone *models.Polygone,
	playground *Playground,
) (polygoneFormCallback *PolygoneFormCallback) {
	polygoneFormCallback = new(PolygoneFormCallback)
	polygoneFormCallback.playground = playground
	polygoneFormCallback.polygone = polygone

	polygoneFormCallback.CreationMode = (polygone == nil)

	return
}

type PolygoneFormCallback struct {
	polygone *models.Polygone

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (polygoneFormCallback *PolygoneFormCallback) OnSave() {

	log.Println("PolygoneFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polygoneFormCallback.playground.formStage.Checkout()

	if polygoneFormCallback.polygone == nil {
		polygoneFormCallback.polygone = new(models.Polygone).Stage(polygoneFormCallback.playground.stageOfInterest)
	}
	polygone_ := polygoneFormCallback.polygone
	_ = polygone_

	// get the formGroup
	formGroup := polygoneFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polygone_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polygone_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polygone_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polygone_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polygone_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polygone_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polygone_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polygone_.Transform), formDiv)
		case "Layer:Polygones":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Polygones"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				polygoneFormCallback.playground.stageOfInterest,
				polygoneFormCallback.playground.backRepoOfInterest,
				polygone_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Polygones, polygone_)
					pastLayerOwner.Polygones = slices.Delete(pastLayerOwner.Polygones, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](polygoneFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Polygones, polygone_)
								pastLayerOwner.Polygones = slices.Delete(pastLayerOwner.Polygones, idx, idx+1)
								newLayerOwner.Polygones = append(newLayerOwner.Polygones, polygone_)
							}
						} else {
							newLayerOwner.Polygones = append(newLayerOwner.Polygones, polygone_)
						}
					}
				}
			}
		}
	}

	polygoneFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Polygone](
		polygoneFormCallback.playground,
	)
	polygoneFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if polygoneFormCallback.CreationMode {
		polygoneFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__PolygoneFormCallback(
				nil,
				polygoneFormCallback.playground,
			),
		}).Stage(polygoneFormCallback.playground.formStage)
		polygone := new(models.Polygone)
		FillUpForm(polygone, newFormGroup, polygoneFormCallback.playground)
		polygoneFormCallback.playground.formStage.Commit()
	}

	fillUpTree(polygoneFormCallback.playground)
}
func __gong__New__PolylineFormCallback(
	polyline *models.Polyline,
	playground *Playground,
) (polylineFormCallback *PolylineFormCallback) {
	polylineFormCallback = new(PolylineFormCallback)
	polylineFormCallback.playground = playground
	polylineFormCallback.polyline = polyline

	polylineFormCallback.CreationMode = (polyline == nil)

	return
}

type PolylineFormCallback struct {
	polyline *models.Polyline

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (polylineFormCallback *PolylineFormCallback) OnSave() {

	log.Println("PolylineFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	polylineFormCallback.playground.formStage.Checkout()

	if polylineFormCallback.polyline == nil {
		polylineFormCallback.polyline = new(models.Polyline).Stage(polylineFormCallback.playground.stageOfInterest)
	}
	polyline_ := polylineFormCallback.polyline
	_ = polyline_

	// get the formGroup
	formGroup := polylineFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(polyline_.Name), formDiv)
		case "Points":
			FormDivBasicFieldToField(&(polyline_.Points), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(polyline_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(polyline_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(polyline_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(polyline_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(polyline_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(polyline_.Transform), formDiv)
		case "Layer:Polylines":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Polylines"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				polylineFormCallback.playground.stageOfInterest,
				polylineFormCallback.playground.backRepoOfInterest,
				polyline_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Polylines, polyline_)
					pastLayerOwner.Polylines = slices.Delete(pastLayerOwner.Polylines, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](polylineFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Polylines, polyline_)
								pastLayerOwner.Polylines = slices.Delete(pastLayerOwner.Polylines, idx, idx+1)
								newLayerOwner.Polylines = append(newLayerOwner.Polylines, polyline_)
							}
						} else {
							newLayerOwner.Polylines = append(newLayerOwner.Polylines, polyline_)
						}
					}
				}
			}
		}
	}

	polylineFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Polyline](
		polylineFormCallback.playground,
	)
	polylineFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if polylineFormCallback.CreationMode {
		polylineFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__PolylineFormCallback(
				nil,
				polylineFormCallback.playground,
			),
		}).Stage(polylineFormCallback.playground.formStage)
		polyline := new(models.Polyline)
		FillUpForm(polyline, newFormGroup, polylineFormCallback.playground)
		polylineFormCallback.playground.formStage.Commit()
	}

	fillUpTree(polylineFormCallback.playground)
}
func __gong__New__RectFormCallback(
	rect *models.Rect,
	playground *Playground,
) (rectFormCallback *RectFormCallback) {
	rectFormCallback = new(RectFormCallback)
	rectFormCallback.playground = playground
	rectFormCallback.rect = rect

	rectFormCallback.CreationMode = (rect == nil)

	return
}

type RectFormCallback struct {
	rect *models.Rect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectFormCallback *RectFormCallback) OnSave() {

	log.Println("RectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectFormCallback.playground.formStage.Checkout()

	if rectFormCallback.rect == nil {
		rectFormCallback.rect = new(models.Rect).Stage(rectFormCallback.playground.stageOfInterest)
	}
	rect_ := rectFormCallback.rect
	_ = rect_

	// get the formGroup
	formGroup := rectFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rect_.RX), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rect_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rect_.Transform), formDiv)
		case "IsSelectable":
			FormDivBasicFieldToField(&(rect_.IsSelectable), formDiv)
		case "IsSelected":
			FormDivBasicFieldToField(&(rect_.IsSelected), formDiv)
		case "CanHaveLeftHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveLeftHandle), formDiv)
		case "HasLeftHandle":
			FormDivBasicFieldToField(&(rect_.HasLeftHandle), formDiv)
		case "CanHaveRightHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveRightHandle), formDiv)
		case "HasRightHandle":
			FormDivBasicFieldToField(&(rect_.HasRightHandle), formDiv)
		case "CanHaveTopHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveTopHandle), formDiv)
		case "HasTopHandle":
			FormDivBasicFieldToField(&(rect_.HasTopHandle), formDiv)
		case "CanHaveBottomHandle":
			FormDivBasicFieldToField(&(rect_.CanHaveBottomHandle), formDiv)
		case "HasBottomHandle":
			FormDivBasicFieldToField(&(rect_.HasBottomHandle), formDiv)
		case "CanMoveHorizontaly":
			FormDivBasicFieldToField(&(rect_.CanMoveHorizontaly), formDiv)
		case "CanMoveVerticaly":
			FormDivBasicFieldToField(&(rect_.CanMoveVerticaly), formDiv)
		case "Layer:Rects":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Rects"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				rectFormCallback.playground.stageOfInterest,
				rectFormCallback.playground.backRepoOfInterest,
				rect_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Rects, rect_)
					pastLayerOwner.Rects = slices.Delete(pastLayerOwner.Rects, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](rectFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Rects, rect_)
								pastLayerOwner.Rects = slices.Delete(pastLayerOwner.Rects, idx, idx+1)
								newLayerOwner.Rects = append(newLayerOwner.Rects, rect_)
							}
						} else {
							newLayerOwner.Rects = append(newLayerOwner.Rects, rect_)
						}
					}
				}
			}
		}
	}

	rectFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Rect](
		rectFormCallback.playground,
	)
	rectFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectFormCallback.CreationMode {
		rectFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RectFormCallback(
				nil,
				rectFormCallback.playground,
			),
		}).Stage(rectFormCallback.playground.formStage)
		rect := new(models.Rect)
		FillUpForm(rect, newFormGroup, rectFormCallback.playground)
		rectFormCallback.playground.formStage.Commit()
	}

	fillUpTree(rectFormCallback.playground)
}
func __gong__New__RectAnchoredRectFormCallback(
	rectanchoredrect *models.RectAnchoredRect,
	playground *Playground,
) (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) {
	rectanchoredrectFormCallback = new(RectAnchoredRectFormCallback)
	rectanchoredrectFormCallback.playground = playground
	rectanchoredrectFormCallback.rectanchoredrect = rectanchoredrect

	rectanchoredrectFormCallback.CreationMode = (rectanchoredrect == nil)

	return
}

type RectAnchoredRectFormCallback struct {
	rectanchoredrect *models.RectAnchoredRect

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectanchoredrectFormCallback *RectAnchoredRectFormCallback) OnSave() {

	log.Println("RectAnchoredRectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredrectFormCallback.playground.formStage.Checkout()

	if rectanchoredrectFormCallback.rectanchoredrect == nil {
		rectanchoredrectFormCallback.rectanchoredrect = new(models.RectAnchoredRect).Stage(rectanchoredrectFormCallback.playground.stageOfInterest)
	}
	rectanchoredrect_ := rectanchoredrectFormCallback.rectanchoredrect
	_ = rectanchoredrect_

	// get the formGroup
	formGroup := rectanchoredrectFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredrect_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(rectanchoredrect_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y), formDiv)
		case "Width":
			FormDivBasicFieldToField(&(rectanchoredrect_.Width), formDiv)
		case "Height":
			FormDivBasicFieldToField(&(rectanchoredrect_.Height), formDiv)
		case "RX":
			FormDivBasicFieldToField(&(rectanchoredrect_.RX), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredrect_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredrect_.RectAnchorType), formDiv)
		case "WidthFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.WidthFollowRect), formDiv)
		case "HeightFollowRect":
			FormDivBasicFieldToField(&(rectanchoredrect_.HeightFollowRect), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredrect_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredrect_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredrect_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredrect_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredrect_.Transform), formDiv)
		case "Rect:RectAnchoredRects":
			// we need to retrieve the field owner before the change
			var pastRectOwner *models.Rect
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "RectAnchoredRects"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				rectanchoredrectFormCallback.playground.stageOfInterest,
				rectanchoredrectFormCallback.playground.backRepoOfInterest,
				rectanchoredrect_,
				&rf)

			if reverseFieldOwner != nil {
				pastRectOwner = reverseFieldOwner.(*models.Rect)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRectOwner != nil {
					idx := slices.Index(pastRectOwner.RectAnchoredRects, rectanchoredrect_)
					pastRectOwner.RectAnchoredRects = slices.Delete(pastRectOwner.RectAnchoredRects, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _rect := range *models.GetGongstructInstancesSet[models.Rect](rectanchoredrectFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _rect.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRectOwner := _rect // we have a match
						if pastRectOwner != nil {
							if newRectOwner != pastRectOwner {
								idx := slices.Index(pastRectOwner.RectAnchoredRects, rectanchoredrect_)
								pastRectOwner.RectAnchoredRects = slices.Delete(pastRectOwner.RectAnchoredRects, idx, idx+1)
								newRectOwner.RectAnchoredRects = append(newRectOwner.RectAnchoredRects, rectanchoredrect_)
							}
						} else {
							newRectOwner.RectAnchoredRects = append(newRectOwner.RectAnchoredRects, rectanchoredrect_)
						}
					}
				}
			}
		}
	}

	rectanchoredrectFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.RectAnchoredRect](
		rectanchoredrectFormCallback.playground,
	)
	rectanchoredrectFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectanchoredrectFormCallback.CreationMode {
		rectanchoredrectFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RectAnchoredRectFormCallback(
				nil,
				rectanchoredrectFormCallback.playground,
			),
		}).Stage(rectanchoredrectFormCallback.playground.formStage)
		rectanchoredrect := new(models.RectAnchoredRect)
		FillUpForm(rectanchoredrect, newFormGroup, rectanchoredrectFormCallback.playground)
		rectanchoredrectFormCallback.playground.formStage.Commit()
	}

	fillUpTree(rectanchoredrectFormCallback.playground)
}
func __gong__New__RectAnchoredTextFormCallback(
	rectanchoredtext *models.RectAnchoredText,
	playground *Playground,
) (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) {
	rectanchoredtextFormCallback = new(RectAnchoredTextFormCallback)
	rectanchoredtextFormCallback.playground = playground
	rectanchoredtextFormCallback.rectanchoredtext = rectanchoredtext

	rectanchoredtextFormCallback.CreationMode = (rectanchoredtext == nil)

	return
}

type RectAnchoredTextFormCallback struct {
	rectanchoredtext *models.RectAnchoredText

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectanchoredtextFormCallback *RectAnchoredTextFormCallback) OnSave() {

	log.Println("RectAnchoredTextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectanchoredtextFormCallback.playground.formStage.Checkout()

	if rectanchoredtextFormCallback.rectanchoredtext == nil {
		rectanchoredtextFormCallback.rectanchoredtext = new(models.RectAnchoredText).Stage(rectanchoredtextFormCallback.playground.stageOfInterest)
	}
	rectanchoredtext_ := rectanchoredtextFormCallback.rectanchoredtext
	_ = rectanchoredtext_

	// get the formGroup
	formGroup := rectanchoredtextFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectanchoredtext_.Name), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(rectanchoredtext_.Content), formDiv)
		case "FontWeight":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontWeight), formDiv)
		case "FontSize":
			FormDivBasicFieldToField(&(rectanchoredtext_.FontSize), formDiv)
		case "X_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.X_Offset), formDiv)
		case "Y_Offset":
			FormDivBasicFieldToField(&(rectanchoredtext_.Y_Offset), formDiv)
		case "RectAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.RectAnchorType), formDiv)
		case "TextAnchorType":
			FormDivEnumStringFieldToField(&(rectanchoredtext_.TextAnchorType), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectanchoredtext_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectanchoredtext_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectanchoredtext_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectanchoredtext_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectanchoredtext_.Transform), formDiv)
		case "Rect:RectAnchoredTexts":
			// we need to retrieve the field owner before the change
			var pastRectOwner *models.Rect
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Rect"
			rf.Fieldname = "RectAnchoredTexts"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				rectanchoredtextFormCallback.playground.stageOfInterest,
				rectanchoredtextFormCallback.playground.backRepoOfInterest,
				rectanchoredtext_,
				&rf)

			if reverseFieldOwner != nil {
				pastRectOwner = reverseFieldOwner.(*models.Rect)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastRectOwner != nil {
					idx := slices.Index(pastRectOwner.RectAnchoredTexts, rectanchoredtext_)
					pastRectOwner.RectAnchoredTexts = slices.Delete(pastRectOwner.RectAnchoredTexts, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _rect := range *models.GetGongstructInstancesSet[models.Rect](rectanchoredtextFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _rect.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newRectOwner := _rect // we have a match
						if pastRectOwner != nil {
							if newRectOwner != pastRectOwner {
								idx := slices.Index(pastRectOwner.RectAnchoredTexts, rectanchoredtext_)
								pastRectOwner.RectAnchoredTexts = slices.Delete(pastRectOwner.RectAnchoredTexts, idx, idx+1)
								newRectOwner.RectAnchoredTexts = append(newRectOwner.RectAnchoredTexts, rectanchoredtext_)
							}
						} else {
							newRectOwner.RectAnchoredTexts = append(newRectOwner.RectAnchoredTexts, rectanchoredtext_)
						}
					}
				}
			}
		}
	}

	rectanchoredtextFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.RectAnchoredText](
		rectanchoredtextFormCallback.playground,
	)
	rectanchoredtextFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectanchoredtextFormCallback.CreationMode {
		rectanchoredtextFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RectAnchoredTextFormCallback(
				nil,
				rectanchoredtextFormCallback.playground,
			),
		}).Stage(rectanchoredtextFormCallback.playground.formStage)
		rectanchoredtext := new(models.RectAnchoredText)
		FillUpForm(rectanchoredtext, newFormGroup, rectanchoredtextFormCallback.playground)
		rectanchoredtextFormCallback.playground.formStage.Commit()
	}

	fillUpTree(rectanchoredtextFormCallback.playground)
}
func __gong__New__RectLinkLinkFormCallback(
	rectlinklink *models.RectLinkLink,
	playground *Playground,
) (rectlinklinkFormCallback *RectLinkLinkFormCallback) {
	rectlinklinkFormCallback = new(RectLinkLinkFormCallback)
	rectlinklinkFormCallback.playground = playground
	rectlinklinkFormCallback.rectlinklink = rectlinklink

	rectlinklinkFormCallback.CreationMode = (rectlinklink == nil)

	return
}

type RectLinkLinkFormCallback struct {
	rectlinklink *models.RectLinkLink

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (rectlinklinkFormCallback *RectLinkLinkFormCallback) OnSave() {

	log.Println("RectLinkLinkFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	rectlinklinkFormCallback.playground.formStage.Checkout()

	if rectlinklinkFormCallback.rectlinklink == nil {
		rectlinklinkFormCallback.rectlinklink = new(models.RectLinkLink).Stage(rectlinklinkFormCallback.playground.stageOfInterest)
	}
	rectlinklink_ := rectlinklinkFormCallback.rectlinklink
	_ = rectlinklink_

	// get the formGroup
	formGroup := rectlinklinkFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(rectlinklink_.Name), formDiv)
		case "Start":
			FormDivSelectFieldToField(&(rectlinklink_.Start), rectlinklinkFormCallback.playground.stageOfInterest, formDiv)
		case "End":
			FormDivSelectFieldToField(&(rectlinklink_.End), rectlinklinkFormCallback.playground.stageOfInterest, formDiv)
		case "TargetAnchorPosition":
			FormDivBasicFieldToField(&(rectlinklink_.TargetAnchorPosition), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(rectlinklink_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(rectlinklink_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(rectlinklink_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(rectlinklink_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(rectlinklink_.Transform), formDiv)
		case "Layer:RectLinkLinks":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "RectLinkLinks"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				rectlinklinkFormCallback.playground.stageOfInterest,
				rectlinklinkFormCallback.playground.backRepoOfInterest,
				rectlinklink_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.RectLinkLinks, rectlinklink_)
					pastLayerOwner.RectLinkLinks = slices.Delete(pastLayerOwner.RectLinkLinks, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](rectlinklinkFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.RectLinkLinks, rectlinklink_)
								pastLayerOwner.RectLinkLinks = slices.Delete(pastLayerOwner.RectLinkLinks, idx, idx+1)
								newLayerOwner.RectLinkLinks = append(newLayerOwner.RectLinkLinks, rectlinklink_)
							}
						} else {
							newLayerOwner.RectLinkLinks = append(newLayerOwner.RectLinkLinks, rectlinklink_)
						}
					}
				}
			}
		}
	}

	rectlinklinkFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.RectLinkLink](
		rectlinklinkFormCallback.playground,
	)
	rectlinklinkFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if rectlinklinkFormCallback.CreationMode {
		rectlinklinkFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__RectLinkLinkFormCallback(
				nil,
				rectlinklinkFormCallback.playground,
			),
		}).Stage(rectlinklinkFormCallback.playground.formStage)
		rectlinklink := new(models.RectLinkLink)
		FillUpForm(rectlinklink, newFormGroup, rectlinklinkFormCallback.playground)
		rectlinklinkFormCallback.playground.formStage.Commit()
	}

	fillUpTree(rectlinklinkFormCallback.playground)
}
func __gong__New__SVGFormCallback(
	svg *models.SVG,
	playground *Playground,
) (svgFormCallback *SVGFormCallback) {
	svgFormCallback = new(SVGFormCallback)
	svgFormCallback.playground = playground
	svgFormCallback.svg = svg

	svgFormCallback.CreationMode = (svg == nil)

	return
}

type SVGFormCallback struct {
	svg *models.SVG

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (svgFormCallback *SVGFormCallback) OnSave() {

	log.Println("SVGFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgFormCallback.playground.formStage.Checkout()

	if svgFormCallback.svg == nil {
		svgFormCallback.svg = new(models.SVG).Stage(svgFormCallback.playground.stageOfInterest)
	}
	svg_ := svgFormCallback.svg
	_ = svg_

	// get the formGroup
	formGroup := svgFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svg_.Name), formDiv)
		case "DrawingState":
			FormDivEnumStringFieldToField(&(svg_.DrawingState), formDiv)
		case "StartRect":
			FormDivSelectFieldToField(&(svg_.StartRect), svgFormCallback.playground.stageOfInterest, formDiv)
		case "EndRect":
			FormDivSelectFieldToField(&(svg_.EndRect), svgFormCallback.playground.stageOfInterest, formDiv)
		case "IsEditable":
			FormDivBasicFieldToField(&(svg_.IsEditable), formDiv)
		}
	}

	svgFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.SVG](
		svgFormCallback.playground,
	)
	svgFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if svgFormCallback.CreationMode {
		svgFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__SVGFormCallback(
				nil,
				svgFormCallback.playground,
			),
		}).Stage(svgFormCallback.playground.formStage)
		svg := new(models.SVG)
		FillUpForm(svg, newFormGroup, svgFormCallback.playground)
		svgFormCallback.playground.formStage.Commit()
	}

	fillUpTree(svgFormCallback.playground)
}
func __gong__New__TextFormCallback(
	text *models.Text,
	playground *Playground,
) (textFormCallback *TextFormCallback) {
	textFormCallback = new(TextFormCallback)
	textFormCallback.playground = playground
	textFormCallback.text = text

	textFormCallback.CreationMode = (text == nil)

	return
}

type TextFormCallback struct {
	text *models.Text

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (textFormCallback *TextFormCallback) OnSave() {

	log.Println("TextFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	textFormCallback.playground.formStage.Checkout()

	if textFormCallback.text == nil {
		textFormCallback.text = new(models.Text).Stage(textFormCallback.playground.stageOfInterest)
	}
	text_ := textFormCallback.text
	_ = text_

	// get the formGroup
	formGroup := textFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(text_.Name), formDiv)
		case "X":
			FormDivBasicFieldToField(&(text_.X), formDiv)
		case "Y":
			FormDivBasicFieldToField(&(text_.Y), formDiv)
		case "Content":
			FormDivBasicFieldToField(&(text_.Content), formDiv)
		case "Color":
			FormDivBasicFieldToField(&(text_.Color), formDiv)
		case "FillOpacity":
			FormDivBasicFieldToField(&(text_.FillOpacity), formDiv)
		case "Stroke":
			FormDivBasicFieldToField(&(text_.Stroke), formDiv)
		case "StrokeWidth":
			FormDivBasicFieldToField(&(text_.StrokeWidth), formDiv)
		case "StrokeDashArray":
			FormDivBasicFieldToField(&(text_.StrokeDashArray), formDiv)
		case "StrokeDashArrayWhenSelected":
			FormDivBasicFieldToField(&(text_.StrokeDashArrayWhenSelected), formDiv)
		case "Transform":
			FormDivBasicFieldToField(&(text_.Transform), formDiv)
		case "Layer:Texts":
			// we need to retrieve the field owner before the change
			var pastLayerOwner *models.Layer
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Layer"
			rf.Fieldname = "Texts"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				textFormCallback.playground.stageOfInterest,
				textFormCallback.playground.backRepoOfInterest,
				text_,
				&rf)

			if reverseFieldOwner != nil {
				pastLayerOwner = reverseFieldOwner.(*models.Layer)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastLayerOwner != nil {
					idx := slices.Index(pastLayerOwner.Texts, text_)
					pastLayerOwner.Texts = slices.Delete(pastLayerOwner.Texts, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _layer := range *models.GetGongstructInstancesSet[models.Layer](textFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _layer.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newLayerOwner := _layer // we have a match
						if pastLayerOwner != nil {
							if newLayerOwner != pastLayerOwner {
								idx := slices.Index(pastLayerOwner.Texts, text_)
								pastLayerOwner.Texts = slices.Delete(pastLayerOwner.Texts, idx, idx+1)
								newLayerOwner.Texts = append(newLayerOwner.Texts, text_)
							}
						} else {
							newLayerOwner.Texts = append(newLayerOwner.Texts, text_)
						}
					}
				}
			}
		}
	}

	textFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Text](
		textFormCallback.playground,
	)
	textFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if textFormCallback.CreationMode {
		textFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TextFormCallback(
				nil,
				textFormCallback.playground,
			),
		}).Stage(textFormCallback.playground.formStage)
		text := new(models.Text)
		FillUpForm(text, newFormGroup, textFormCallback.playground)
		textFormCallback.playground.formStage.Commit()
	}

	fillUpTree(textFormCallback.playground)
}
