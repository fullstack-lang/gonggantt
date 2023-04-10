package gantt2svg

import (
	gonggantt_models "github.com/fullstack-lang/gonggantt/go/models"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

// type of the singloton for interception of gantt commit in order to generate
// the svg
type GanttToSVGMapper struct {
	mapBar_BarSVG map[*gonggantt_models.Bar]*gongsvg_models.Rect
}

var GanttToSVGMapperSingloton GanttToSVGMapper
