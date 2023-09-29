// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for GanttStacksNames
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (ganttstacksnames GanttStacksNames) ToString() (res string) {

	// migration of former implementation of enum
	switch ganttstacksnames {
	// insertion code per enum code
	case SvgStackName:
		res = "svg"
	case GanttStackName:
		res = "gantt"
	case GanttProbeStacksPrefix:
		res = "gantt-probe"
	case SVGProbeStacksPrefix:
		res = "svg-probe"
	}
	return
}

func (ganttstacksnames *GanttStacksNames) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "svg":
		*ganttstacksnames = SvgStackName
	case "gantt":
		*ganttstacksnames = GanttStackName
	case "gantt-probe":
		*ganttstacksnames = GanttProbeStacksPrefix
	case "svg-probe":
		*ganttstacksnames = SVGProbeStacksPrefix
	default:
		return errUnkownEnum
	}
	return
}

func (ganttstacksnames *GanttStacksNames) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "SvgStackName":
		*ganttstacksnames = SvgStackName
	case "GanttStackName":
		*ganttstacksnames = GanttStackName
	case "GanttProbeStacksPrefix":
		*ganttstacksnames = GanttProbeStacksPrefix
	case "SVGProbeStacksPrefix":
		*ganttstacksnames = SVGProbeStacksPrefix
	default:
		return errUnkownEnum
	}
	return
}

func (ganttstacksnames *GanttStacksNames) ToCodeString() (res string) {

	switch *ganttstacksnames {
	// insertion code per enum code
	case SvgStackName:
		res = "SvgStackName"
	case GanttStackName:
		res = "GanttStackName"
	case GanttProbeStacksPrefix:
		res = "GanttProbeStacksPrefix"
	case SVGProbeStacksPrefix:
		res = "SVGProbeStacksPrefix"
	}
	return
}

func (ganttstacksnames GanttStacksNames) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "SvgStackName")
	res = append(res, "GanttStackName")
	res = append(res, "GanttProbeStacksPrefix")
	res = append(res, "SVGProbeStacksPrefix")

	return
}

func (ganttstacksnames GanttStacksNames) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "svg")
	res = append(res, "gantt")
	res = append(res, "gantt-probe")
	res = append(res, "svg-probe")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | GanttStacksNames
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*GanttStacksNames
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
