package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
	"github.com/fullstack-lang/gonggantt/go/models"
)

var Default uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Arrow{}),
			Position: &uml.Position{
				X: 570.000000,
				Y: 310.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Arrow{}.From,
					Middlevertice: &uml.Vertice{
						X: 1010.000000,
						Y: 386.500000,
					},
					Multiplicity: "0..1",
				},
				{
					Field: models.Arrow{}.To,
					Middlevertice: &uml.Vertice{
						X: 1190.000000,
						Y: 356.500000,
					},
					Multiplicity: "0..1",
				},
			},
		},
		{
			Struct: &(models.Bar{}),
			Position: &uml.Position{
				X: 950.000000,
				Y: 140.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
			Fields: []*uml.Field{
				{
					Field: models.Bar{}.End,
				},
				{
					Field: models.Bar{}.Name,
				},
				{
					Field: models.Bar{}.OptionnalColor,
				},
				{
					Field: models.Bar{}.OptionnalStroke,
				},
				{
					Field: models.Bar{}.Start,
				},
			},
		},
		{
			Struct: &(models.Gantt{}),
			Position: &uml.Position{
				X: 60.000000,
				Y: 40.000000,
			},
			Width:  240.000000,
			Heigth: 318.000000,
			Links: []*uml.Link{
				{
					Field: models.Gantt{}.Arrows,
					Middlevertice: &uml.Vertice{
						X: 445.000000,
						Y: 261.500000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Gantt{}.Groups,
					Middlevertice: &uml.Vertice{
						X: 428.500000,
						Y: 107.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Gantt{}.Lanes,
					Middlevertice: &uml.Vertice{
						X: 411.500000,
						Y: 175.000000,
					},
					Multiplicity: "*",
				},
				{
					Field: models.Gantt{}.Milestones,
					Middlevertice: &uml.Vertice{
						X: 435.000000,
						Y: 361.500000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Gantt{}.AlignOnStartEndOnYearStart,
				},
				{
					Field: models.Gantt{}.ArrowLengthToTheRightOfStartBar,
				},
				{
					Field: models.Gantt{}.ArrowTipLenght,
				},
				{
					Field: models.Gantt{}.Group_Stroke,
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
					Field: models.Gantt{}.Start,
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
				X: 570.000000,
				Y: 50.000000,
			},
			Width:  240.000000,
			Heigth: 63.000000,
			Links: []*uml.Link{
				{
					Field: models.Group{}.GroupLanes,
					Middlevertice: &uml.Vertice{
						X: 930.000000,
						Y: 126.500000,
					},
					Multiplicity: "*",
				},
			},
		},
		{
			Struct: &(models.Lane{}),
			Position: &uml.Position{
				X: 570.000000,
				Y: 140.000000,
			},
			Width:  240.000000,
			Heigth: 138.000000,
			Links: []*uml.Link{
				{
					Field: models.Lane{}.Bars,
					Middlevertice: &uml.Vertice{
						X: 870.000000,
						Y: 171.500000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Lane{}.Name,
				},
				{
					Field: models.Lane{}.Name,
				},
				{
					Field: models.Lane{}.Order,
				},
				{
					Field: models.Lane{}.Order,
				},
				{
					Field: models.Lane{}.Order,
				},
			},
		},
		{
			Struct: &(models.Milestone{}),
			Position: &uml.Position{
				X: 570.000000,
				Y: 400.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Milestone{}.LanesToDisplayMilestone,
					Middlevertice: &uml.Vertice{
						X: 1010.000000,
						Y: 429.000000,
					},
					Multiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Milestone{}.Date,
				},
			},
		},
	},
}
