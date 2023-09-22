// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.playground = playground
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Arrow:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Bar:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Gantt:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Group:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Lane:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.LaneUse:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Milestone:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.playground.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.playground)
	fillUpTree(cellDeleteIconImpl.playground)
	cellDeleteIconImpl.playground.tableStage.Commit()
}

