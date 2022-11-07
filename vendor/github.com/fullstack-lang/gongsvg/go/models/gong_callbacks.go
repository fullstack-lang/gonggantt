package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Animate:
		if stage.OnAfterAnimateCreateCallback != nil {
			stage.OnAfterAnimateCreateCallback.OnAfterCreate(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleCreateCallback != nil {
			stage.OnAfterCircleCreateCallback.OnAfterCreate(stage, target)
		}
	case *Ellipse:
		if stage.OnAfterEllipseCreateCallback != nil {
			stage.OnAfterEllipseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Line:
		if stage.OnAfterLineCreateCallback != nil {
			stage.OnAfterLineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Path:
		if stage.OnAfterPathCreateCallback != nil {
			stage.OnAfterPathCreateCallback.OnAfterCreate(stage, target)
		}
	case *Polygone:
		if stage.OnAfterPolygoneCreateCallback != nil {
			stage.OnAfterPolygoneCreateCallback.OnAfterCreate(stage, target)
		}
	case *Polyline:
		if stage.OnAfterPolylineCreateCallback != nil {
			stage.OnAfterPolylineCreateCallback.OnAfterCreate(stage, target)
		}
	case *Rect:
		if stage.OnAfterRectCreateCallback != nil {
			stage.OnAfterRectCreateCallback.OnAfterCreate(stage, target)
		}
	case *SVG:
		if stage.OnAfterSVGCreateCallback != nil {
			stage.OnAfterSVGCreateCallback.OnAfterCreate(stage, target)
		}
	case *Text:
		if stage.OnAfterTextCreateCallback != nil {
			stage.OnAfterTextCreateCallback.OnAfterCreate(stage, target)
		}
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Animate:
		newTarget := any(new).(*Animate)
		if stage.OnAfterAnimateUpdateCallback != nil {
			stage.OnAfterAnimateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Circle:
		newTarget := any(new).(*Circle)
		if stage.OnAfterCircleUpdateCallback != nil {
			stage.OnAfterCircleUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Ellipse:
		newTarget := any(new).(*Ellipse)
		if stage.OnAfterEllipseUpdateCallback != nil {
			stage.OnAfterEllipseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Line:
		newTarget := any(new).(*Line)
		if stage.OnAfterLineUpdateCallback != nil {
			stage.OnAfterLineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Path:
		newTarget := any(new).(*Path)
		if stage.OnAfterPathUpdateCallback != nil {
			stage.OnAfterPathUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Polygone:
		newTarget := any(new).(*Polygone)
		if stage.OnAfterPolygoneUpdateCallback != nil {
			stage.OnAfterPolygoneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Polyline:
		newTarget := any(new).(*Polyline)
		if stage.OnAfterPolylineUpdateCallback != nil {
			stage.OnAfterPolylineUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Rect:
		newTarget := any(new).(*Rect)
		if stage.OnAfterRectUpdateCallback != nil {
			stage.OnAfterRectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SVG:
		newTarget := any(new).(*SVG)
		if stage.OnAfterSVGUpdateCallback != nil {
			stage.OnAfterSVGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Text:
		newTarget := any(new).(*Text)
		if stage.OnAfterTextUpdateCallback != nil {
			stage.OnAfterTextUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Animate:
		if stage.OnAfterAnimateDeleteCallback != nil {
			staged := any(staged).(*Animate)
			stage.OnAfterAnimateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Circle:
		if stage.OnAfterCircleDeleteCallback != nil {
			staged := any(staged).(*Circle)
			stage.OnAfterCircleDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Ellipse:
		if stage.OnAfterEllipseDeleteCallback != nil {
			staged := any(staged).(*Ellipse)
			stage.OnAfterEllipseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Line:
		if stage.OnAfterLineDeleteCallback != nil {
			staged := any(staged).(*Line)
			stage.OnAfterLineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Path:
		if stage.OnAfterPathDeleteCallback != nil {
			staged := any(staged).(*Path)
			stage.OnAfterPathDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Polygone:
		if stage.OnAfterPolygoneDeleteCallback != nil {
			staged := any(staged).(*Polygone)
			stage.OnAfterPolygoneDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Polyline:
		if stage.OnAfterPolylineDeleteCallback != nil {
			staged := any(staged).(*Polyline)
			stage.OnAfterPolylineDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Rect:
		if stage.OnAfterRectDeleteCallback != nil {
			staged := any(staged).(*Rect)
			stage.OnAfterRectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SVG:
		if stage.OnAfterSVGDeleteCallback != nil {
			staged := any(staged).(*SVG)
			stage.OnAfterSVGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Text:
		if stage.OnAfterTextDeleteCallback != nil {
			staged := any(staged).(*Text)
			stage.OnAfterTextDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Animate:
		if stage.OnAfterAnimateReadCallback != nil {
			stage.OnAfterAnimateReadCallback.OnAfterRead(stage, target)
		}
	case *Circle:
		if stage.OnAfterCircleReadCallback != nil {
			stage.OnAfterCircleReadCallback.OnAfterRead(stage, target)
		}
	case *Ellipse:
		if stage.OnAfterEllipseReadCallback != nil {
			stage.OnAfterEllipseReadCallback.OnAfterRead(stage, target)
		}
	case *Line:
		if stage.OnAfterLineReadCallback != nil {
			stage.OnAfterLineReadCallback.OnAfterRead(stage, target)
		}
	case *Path:
		if stage.OnAfterPathReadCallback != nil {
			stage.OnAfterPathReadCallback.OnAfterRead(stage, target)
		}
	case *Polygone:
		if stage.OnAfterPolygoneReadCallback != nil {
			stage.OnAfterPolygoneReadCallback.OnAfterRead(stage, target)
		}
	case *Polyline:
		if stage.OnAfterPolylineReadCallback != nil {
			stage.OnAfterPolylineReadCallback.OnAfterRead(stage, target)
		}
	case *Rect:
		if stage.OnAfterRectReadCallback != nil {
			stage.OnAfterRectReadCallback.OnAfterRead(stage, target)
		}
	case *SVG:
		if stage.OnAfterSVGReadCallback != nil {
			stage.OnAfterSVGReadCallback.OnAfterRead(stage, target)
		}
	case *Text:
		if stage.OnAfterTextReadCallback != nil {
			stage.OnAfterTextReadCallback.OnAfterRead(stage, target)
		}
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateUpdateCallback = any(callback).(OnAfterUpdateInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleUpdateCallback = any(callback).(OnAfterUpdateInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseUpdateCallback = any(callback).(OnAfterUpdateInterface[Ellipse])
	
	case *Line:
		stage.OnAfterLineUpdateCallback = any(callback).(OnAfterUpdateInterface[Line])
	
	case *Path:
		stage.OnAfterPathUpdateCallback = any(callback).(OnAfterUpdateInterface[Path])
	
	case *Polygone:
		stage.OnAfterPolygoneUpdateCallback = any(callback).(OnAfterUpdateInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineUpdateCallback = any(callback).(OnAfterUpdateInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectUpdateCallback = any(callback).(OnAfterUpdateInterface[Rect])
	
	case *SVG:
		stage.OnAfterSVGUpdateCallback = any(callback).(OnAfterUpdateInterface[SVG])
	
	case *Text:
		stage.OnAfterTextUpdateCallback = any(callback).(OnAfterUpdateInterface[Text])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateCreateCallback = any(callback).(OnAfterCreateInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleCreateCallback = any(callback).(OnAfterCreateInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseCreateCallback = any(callback).(OnAfterCreateInterface[Ellipse])
	
	case *Line:
		stage.OnAfterLineCreateCallback = any(callback).(OnAfterCreateInterface[Line])
	
	case *Path:
		stage.OnAfterPathCreateCallback = any(callback).(OnAfterCreateInterface[Path])
	
	case *Polygone:
		stage.OnAfterPolygoneCreateCallback = any(callback).(OnAfterCreateInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineCreateCallback = any(callback).(OnAfterCreateInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectCreateCallback = any(callback).(OnAfterCreateInterface[Rect])
	
	case *SVG:
		stage.OnAfterSVGCreateCallback = any(callback).(OnAfterCreateInterface[SVG])
	
	case *Text:
		stage.OnAfterTextCreateCallback = any(callback).(OnAfterCreateInterface[Text])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateDeleteCallback = any(callback).(OnAfterDeleteInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleDeleteCallback = any(callback).(OnAfterDeleteInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseDeleteCallback = any(callback).(OnAfterDeleteInterface[Ellipse])
	
	case *Line:
		stage.OnAfterLineDeleteCallback = any(callback).(OnAfterDeleteInterface[Line])
	
	case *Path:
		stage.OnAfterPathDeleteCallback = any(callback).(OnAfterDeleteInterface[Path])
	
	case *Polygone:
		stage.OnAfterPolygoneDeleteCallback = any(callback).(OnAfterDeleteInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineDeleteCallback = any(callback).(OnAfterDeleteInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectDeleteCallback = any(callback).(OnAfterDeleteInterface[Rect])
	
	case *SVG:
		stage.OnAfterSVGDeleteCallback = any(callback).(OnAfterDeleteInterface[SVG])
	
	case *Text:
		stage.OnAfterTextDeleteCallback = any(callback).(OnAfterDeleteInterface[Text])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Animate:
		stage.OnAfterAnimateReadCallback = any(callback).(OnAfterReadInterface[Animate])
	
	case *Circle:
		stage.OnAfterCircleReadCallback = any(callback).(OnAfterReadInterface[Circle])
	
	case *Ellipse:
		stage.OnAfterEllipseReadCallback = any(callback).(OnAfterReadInterface[Ellipse])
	
	case *Line:
		stage.OnAfterLineReadCallback = any(callback).(OnAfterReadInterface[Line])
	
	case *Path:
		stage.OnAfterPathReadCallback = any(callback).(OnAfterReadInterface[Path])
	
	case *Polygone:
		stage.OnAfterPolygoneReadCallback = any(callback).(OnAfterReadInterface[Polygone])
	
	case *Polyline:
		stage.OnAfterPolylineReadCallback = any(callback).(OnAfterReadInterface[Polyline])
	
	case *Rect:
		stage.OnAfterRectReadCallback = any(callback).(OnAfterReadInterface[Rect])
	
	case *SVG:
		stage.OnAfterSVGReadCallback = any(callback).(OnAfterReadInterface[SVG])
	
	case *Text:
		stage.OnAfterTextReadCallback = any(callback).(OnAfterReadInterface[Text])
	
	}
}
