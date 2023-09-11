// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gonggantt/go/models"
	"github.com/fullstack-lang/gonggantt/go/orm"
)

type NodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "Arrow" {
		fillUpTable[models.Arrow](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Bar" {
		fillUpTable[models.Bar](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Gantt" {
		fillUpTable[models.Gantt](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Group" {
		fillUpTable[models.Group](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Lane" {
		fillUpTable[models.Lane](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LaneUse" {
		fillUpTable[models.LaneUse](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Milestone" {
		fillUpTable[models.Milestone](nodeImplGongstruct.playground)
	}

	nodeImplGongstruct.playground.tableStage.Commit()
}

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	playground *Playground,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.Arrow:
		fillUpTable[models.Arrow](playground)
	case *models.Bar:
		fillUpTable[models.Bar](playground)
	case *models.Gantt:
		fillUpTable[models.Gantt](playground)
	case *models.Group:
		fillUpTable[models.Group](playground)
	case *models.Lane:
		fillUpTable[models.Lane](playground)
	case *models.LaneUse:
		fillUpTable[models.LaneUse](playground)
	case *models.Milestone:
		fillUpTable[models.Milestone](playground)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	playground *Playground,
) {

	playground.tableStage.Reset()
	playground.tableStage.Commit()

	table := new(gongtable.Table).Stage(playground.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	playground.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](playground.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			playground.stageOfInterest,
			playground.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(playground.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(playground.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(playground.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance, playground)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				playground.stageOfInterest,
				playground.backRepoOfInterest,
				structInstance,
			),
		}).Stage(playground.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(playground.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(playground.tableStage)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(playground.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(playground.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.playground = playground
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	formStage := rowUpdate.playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(rowUpdate.Instance).(type) {
	// insertion point
	case *models.Arrow:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewArrowFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Bar:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewBarFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Gantt:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGanttFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Group:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewGroupFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Lane:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLaneFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.LaneUse:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewLaneUseFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	case *models.Milestone:
		formGroup := (&gongtable.FormGroup{
			Name: gongtable.FormGroupDefaultName.ToString(),
			OnSave: NewMilestoneFormCallback(
				instancesTyped,
				rowUpdate.playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, rowUpdate.playground)
	}
	formStage.Commit()

}
