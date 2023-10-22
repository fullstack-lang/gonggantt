// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gonggantt/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Arrows":
				if _gantt, ok := stage.Gantt_Arrows_reverseMap[inst]; ok {
					res = _gantt.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			case "Bars":
				if _lane, ok := stage.Lane_Bars_reverseMap[inst]; ok {
					res = _lane.Name
				}
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Gantt:
		tmp := GetInstanceDBFromInstance[models.Gantt, GanttDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Groups":
				if _gantt, ok := stage.Gantt_Groups_reverseMap[inst]; ok {
					res = _gantt.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Lane:
		tmp := GetInstanceDBFromInstance[models.Lane, LaneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Lanes":
				if _gantt, ok := stage.Gantt_Lanes_reverseMap[inst]; ok {
					res = _gantt.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			case "GroupLanes":
				if _group, ok := stage.Group_GroupLanes_reverseMap[inst]; ok {
					res = _group.Name
				}
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.LaneUse:
		tmp := GetInstanceDBFromInstance[models.LaneUse, LaneUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			case "LanesToDisplayMilestoneUse":
				if _milestone, ok := stage.Milestone_LanesToDisplayMilestoneUse_reverseMap[inst]; ok {
					res = _milestone.Name
				}
			}
		}

	case *models.Milestone:
		tmp := GetInstanceDBFromInstance[models.Milestone, MilestoneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Milestones":
				if _gantt, ok := stage.Gantt_Milestones_reverseMap[inst]; ok {
					res = _gantt.Name
				}
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Arrow:
		tmp := GetInstanceDBFromInstance[models.Arrow, ArrowDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Arrows":
				res = stage.Gantt_Arrows_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Bar:
		tmp := GetInstanceDBFromInstance[models.Bar, BarDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			case "Bars":
				res = stage.Lane_Bars_reverseMap[inst]
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Gantt:
		tmp := GetInstanceDBFromInstance[models.Gantt, GanttDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Group:
		tmp := GetInstanceDBFromInstance[models.Group, GroupDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Groups":
				res = stage.Gantt_Groups_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.Lane:
		tmp := GetInstanceDBFromInstance[models.Lane, LaneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Lanes":
				res = stage.Gantt_Lanes_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			case "GroupLanes":
				res = stage.Group_GroupLanes_reverseMap[inst]
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	case *models.LaneUse:
		tmp := GetInstanceDBFromInstance[models.LaneUse, LaneUseDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			case "LanesToDisplayMilestoneUse":
				res = stage.Milestone_LanesToDisplayMilestoneUse_reverseMap[inst]
			}
		}

	case *models.Milestone:
		tmp := GetInstanceDBFromInstance[models.Milestone, MilestoneDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Arrow":
			switch reverseField.Fieldname {
			}
		case "Bar":
			switch reverseField.Fieldname {
			}
		case "Gantt":
			switch reverseField.Fieldname {
			case "Milestones":
				res = stage.Gantt_Milestones_reverseMap[inst]
			}
		case "Group":
			switch reverseField.Fieldname {
			}
		case "Lane":
			switch reverseField.Fieldname {
			}
		case "LaneUse":
			switch reverseField.Fieldname {
			}
		case "Milestone":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
