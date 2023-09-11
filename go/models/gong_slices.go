// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[T PointerToGongstruct, TF PointerToGongstruct](
	stage *StageStruct,
	owningInstance T,
	sliceField []TF,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[TF]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Arrow:
		// insertion point per field

	case *Bar:
		// insertion point per field

	case *Gantt:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Lanes" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Gantt)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Lanes).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Lanes = make([]*Lane, 0)
				if any(_instance).(*Gantt) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Milestones" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Gantt)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Milestones).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Milestones = make([]*Milestone, 0)
				if any(_instance).(*Gantt) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Groups" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Gantt)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Groups).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Groups = make([]*Group, 0)
				if any(_instance).(*Gantt) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Arrows" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Gantt)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Arrows).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Arrows = make([]*Arrow, 0)
				if any(_instance).(*Gantt) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Group:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "GroupLanes" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Group)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.GroupLanes).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.GroupLanes = make([]*Lane, 0)
				if any(_instance).(*Group) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Lane:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Bars" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Lane)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Bars).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Bars = make([]*Bar, 0)
				if any(_instance).(*Lane) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *LaneUse:
		// insertion point per field

	case *Milestone:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "LanesToDisplayMilestoneUse" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Milestone)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.LanesToDisplayMilestoneUse).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.LanesToDisplayMilestoneUse = make([]*LaneUse, 0)
				if any(_instance).(*Milestone) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
