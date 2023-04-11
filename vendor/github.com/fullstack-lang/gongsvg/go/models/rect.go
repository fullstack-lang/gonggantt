package models

import "log"

type Rect struct {
	Name string

	X, Y, Width, Height, RX float64
	Presentation

	Animations []*Animate

	IsSelectable bool // alllow selected
	IsSelected   bool

	CanHaveHorizontalHandles bool // allows HasHorizontalHandles
	HasHorizontalHandles     bool // is true when selected

	CanMoveHorizontaly bool
	CanMoveVerticaly   bool

	Impl RectImplInterface
}

// OnAfterUpdate, notice that rect == stagedRect
func (rect *Rect) OnAfterUpdate(stage *StageStruct, _, frontRect *Rect) {

	log.Println("Rect, OnAfterUpdate", rect.Name)

	// behavior logic
	if frontRect.IsSelected != rect.IsSelected {
		rect.IsSelected = frontRect.IsSelected
		if frontRect.IsSelected && frontRect.CanHaveHorizontalHandles {
			rect.HasHorizontalHandles = true
		} else {
			rect.HasHorizontalHandles = false
		}
		rect.Commit(stage)
	}

	if rect.X != frontRect.X ||
		rect.Y != frontRect.Y ||
		rect.Width != frontRect.Width ||
		rect.Height != frontRect.Height {

		rect.X = frontRect.X
		rect.Y = frontRect.Y
		rect.Width = frontRect.Width
		rect.Height = frontRect.Height

		rect.Commit(stage)

		if rect.Impl != nil {
			rect.Impl.RectUpdated(rect)
		}
	}
}
