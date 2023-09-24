// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongsvg/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	playground *Playground
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	playground *Playground,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.playground = playground
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
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
	if nodeImplGongstruct.gongStruct.GetName() == "Animate" {
		fillUpTable[models.Animate](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Circle" {
		fillUpTable[models.Circle](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Ellipse" {
		fillUpTable[models.Ellipse](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Layer" {
		fillUpTable[models.Layer](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Line" {
		fillUpTable[models.Line](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Link" {
		fillUpTable[models.Link](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "LinkAnchoredText" {
		fillUpTable[models.LinkAnchoredText](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Path" {
		fillUpTable[models.Path](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Point" {
		fillUpTable[models.Point](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polygone" {
		fillUpTable[models.Polygone](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Polyline" {
		fillUpTable[models.Polyline](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Rect" {
		fillUpTable[models.Rect](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredRect" {
		fillUpTable[models.RectAnchoredRect](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectAnchoredText" {
		fillUpTable[models.RectAnchoredText](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RectLinkLink" {
		fillUpTable[models.RectLinkLink](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SVG" {
		fillUpTable[models.SVG](nodeImplGongstruct.playground)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Text" {
		fillUpTable[models.Text](nodeImplGongstruct.playground)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.playground.tableStage.Commit()
}
