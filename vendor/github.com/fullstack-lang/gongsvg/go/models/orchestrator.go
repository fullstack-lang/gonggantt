package models

// Orchestrator
type Orchestrator struct {
}

func (orchestrator *Orchestrator) OnAfterUpdate(
	gongsvgStage *StageStruct,
	stagedRect, backRepoRect *Rect) {

	stagedRect.OnAfterUpdate(gongsvgStage, stagedRect, backRepoRect)

}
