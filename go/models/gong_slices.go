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
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
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
		if fieldName == "Lanes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Gantt) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Gantt)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Lanes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Lanes = _inferedTypeInstance.Lanes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Lanes =
								append(_inferedTypeInstance.Lanes, any(fieldInstance).(*Lane))
						}
					}
				}
			}
		}
		if fieldName == "Milestones" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Gantt) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Gantt)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Milestones).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Milestones = _inferedTypeInstance.Milestones[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Milestones =
								append(_inferedTypeInstance.Milestones, any(fieldInstance).(*Milestone))
						}
					}
				}
			}
		}
		if fieldName == "Groups" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Gantt) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Gantt)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Groups).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Groups = _inferedTypeInstance.Groups[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Groups =
								append(_inferedTypeInstance.Groups, any(fieldInstance).(*Group))
						}
					}
				}
			}
		}
		if fieldName == "Arrows" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Gantt) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Gantt)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Arrows).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Arrows = _inferedTypeInstance.Arrows[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Arrows =
								append(_inferedTypeInstance.Arrows, any(fieldInstance).(*Arrow))
						}
					}
				}
			}
		}

	case *Group:
		// insertion point per field
		if fieldName == "GroupLanes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Group) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Group)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GroupLanes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GroupLanes = _inferedTypeInstance.GroupLanes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GroupLanes =
								append(_inferedTypeInstance.GroupLanes, any(fieldInstance).(*Lane))
						}
					}
				}
			}
		}

	case *Lane:
		// insertion point per field
		if fieldName == "Bars" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Lane) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Lane)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Bars).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Bars = _inferedTypeInstance.Bars[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Bars =
								append(_inferedTypeInstance.Bars, any(fieldInstance).(*Bar))
						}
					}
				}
			}
		}

	case *LaneUse:
		// insertion point per field

	case *Milestone:
		// insertion point per field
		if fieldName == "LanesToDisplayMilestoneUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Milestone) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Milestone)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.LanesToDisplayMilestoneUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.LanesToDisplayMilestoneUse = _inferedTypeInstance.LanesToDisplayMilestoneUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.LanesToDisplayMilestoneUse =
								append(_inferedTypeInstance.LanesToDisplayMilestoneUse, any(fieldInstance).(*LaneUse))
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
