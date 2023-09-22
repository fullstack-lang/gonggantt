package probe

import (
	"fmt"
	"sort"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gonggantt/go/models"
)

func fillUpTree(
	playground *Playground,
) {
	playground.treeStage.Reset()

	// create tree
	sidebar := (&gongtree_models.Tree{Name: "gong"}).Stage(playground.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](playground.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", playground.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&gongtree_models.Node{Name: name}).Stage(playground.treeStage)


		nodeGongstruct.IsExpanded = false
		switch gongStruct.Name {
		// insertion point
		case "Arrow":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Arrow](playground.stageOfInterest)
			for arrow := range set {
				nodeInstance := (&gongtree_models.Node{Name: arrow.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(arrow, "Arrow", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Bar":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Bar](playground.stageOfInterest)
			for bar := range set {
				nodeInstance := (&gongtree_models.Node{Name: bar.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(bar, "Bar", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Gantt":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Gantt](playground.stageOfInterest)
			for gantt := range set {
				nodeInstance := (&gongtree_models.Node{Name: gantt.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(gantt, "Gantt", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Group":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Group](playground.stageOfInterest)
			for group := range set {
				nodeInstance := (&gongtree_models.Node{Name: group.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(group, "Group", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Lane":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Lane](playground.stageOfInterest)
			for lane := range set {
				nodeInstance := (&gongtree_models.Node{Name: lane.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(lane, "Lane", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "LaneUse":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.LaneUse](playground.stageOfInterest)
			for laneuse := range set {
				nodeInstance := (&gongtree_models.Node{Name: laneuse.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(laneuse, "LaneUse", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Milestone":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Milestone](playground.stageOfInterest)
			for milestone := range set {
				nodeInstance := (&gongtree_models.Node{Name: milestone.GetName()}).Stage(playground.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(milestone, "Milestone", playground)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}	
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, playground)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(playground.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			playground,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}
	playground.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	playground     *Playground
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	playground *Playground) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.playground = playground
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.playground,
	)
}
