package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	""
)

var Diagram uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Bar{}),
			Position: &uml.Position{
				X: 450.000000,
				Y: 410.000000,
			},
			Width:  240.000000,
			Heigth: 108.000000,
			Fields: []*uml.Field{
				{
					Field: models.Bar{}.Name,
				},
				{
					Field: models.Bar{}.OptionnalColor,
				},
				{
					Field: models.Bar{}.OptionnalStroke,
				},
			},
		},
		{
			Struct: &(models.Gantt{}),
			Position: &uml.Position{
				X: 70.000000,
				Y: 30.000000,
			},
			Width:  240.000000,
			Heigth: 318.000000,
			Links: []*uml.Link{
				{
					Field: models.Gantt{}.Groups,
					Middlevertice: &uml.Vertice{
						X: 188.000000,
						Y: 387.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Gantt{}.Lanes,
					Middlevertice: &uml.Vertice{
						X: 346.500000,
						Y: 225.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Gantt{}.Milestones,
					Middlevertice: &uml.Vertice{
						X: 334.500000,
						Y: 87.000000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Gantt{}.AlignOnStartEndOnYearStart,
				},
				{
					Field: models.Gantt{}.DateYOffset,
				},
				{
					Field: models.Gantt{}.Group_Stroke,
				},
				{
					Field: models.Gantt{}.Group_StrokeDashArray,
				},
				{
					Field: models.Gantt{}.Group_StrokeWidth,
				},
				{
					Field: models.Gantt{}.LaneHeight,
				},
				{
					Field: models.Gantt{}.Name,
				},
				{
					Field: models.Gantt{}.RatioBarToLaneHeight,
				},
				{
					Field: models.Gantt{}.TextHeight,
				},
				{
					Field: models.Gantt{}.TimeLine_Color,
				},
				{
					Field: models.Gantt{}.TimeLine_FillOpacity,
				},
				{
					Field: models.Gantt{}.TimeLine_Stroke,
				},
				{
					Field: models.Gantt{}.TimeLine_StrokeWidth,
				},
				{
					Field: models.Gantt{}.XLeftLanes,
				},
				{
					Field: models.Gantt{}.XLeftText,
				},
				{
					Field: models.Gantt{}.XRightMargin,
				},
				{
					Field: models.Gantt{}.YTopMargin,
				},
			},
		},
		{
			Struct: &(models.Group{}),
			Position: &uml.Position{
				X: 90.000000,
				Y: 450.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Group{}.GroupLanes,
					Middlevertice: &uml.Vertice{
						X: 317.500000,
						Y: 391.500000,
					},
					Multiplicity: "*",
				},
			},
		},
		{
			Struct: &(models.Lane{}),
			Position: &uml.Position{
				X: 460.000000,
				Y: 200.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Lane{}.Bars,
					Middlevertice: &uml.Vertice{
						X: 586.000000,
						Y: 343.000000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Lane{}.Name,
				},
			},
		},
		{
			Struct: &(models.Milestone{}),
			Position: &uml.Position{
				X: 460.000000,
				Y: 50.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Milestone{}.DiamonfAndTextAnchors,
					Middlevertice: &uml.Vertice{
						X: 584.000000,
						Y: 139.000000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Milestone{}.Name,
				},
			},
		},
	},
}
