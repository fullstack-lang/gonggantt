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
				Y: 280.000000,
			},
			Width:  240.000000,
			Heigth: 93.000000,
			Links: []*uml.Link{
				{
					Field: models.Arrow{}.From,
					Middlevertice: &uml.Vertice{
						X: 1010.000000,
						Y: 386.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
				{
					Field: models.Arrow{}.To,
					Middlevertice: &uml.Vertice{
						X: 1190.000000,
						Y: 356.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Arrow{}.Name,
				},
				{
					Field: models.Arrow{}.OptionnalColor,
				},
				{
					Field: models.Arrow{}.OptionnalStroke,
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
			Heigth: 123.000000,
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
			Heigth: 333.000000,
			Links: []*uml.Link{
				{
					Field: models.Gantt{}.Arrows,
					Middlevertice: &uml.Vertice{
						X: 445.000000,
						Y: 261.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Gantt{}.Groups,
					Middlevertice: &uml.Vertice{
						X: 428.500000,
						Y: 107.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Gantt{}.Lanes,
					Middlevertice: &uml.Vertice{
						X: 411.500000,
						Y: 175.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
				{
					Field: models.Gantt{}.Milestones,
					Middlevertice: &uml.Vertice{
						X: 435.000000,
						Y: 361.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
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
					Field: models.Gantt{}.ComputedStart,
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
						X: 1080.000000,
						Y: 76.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Group{}.Name,
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
			Heigth: 78.000000,
			Links: []*uml.Link{
				{
					Field: models.Lane{}.Bars,
					Middlevertice: &uml.Vertice{
						X: 870.000000,
						Y: 171.500000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Lane{}.Name,
				},
				{
					Field: models.Lane{}.Order,
				},
			},
		},
		{
			Struct: &(models.LaneUse{}),
			Position: &uml.Position{
				X: 930.000000,
				Y: 520.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.LaneUse{}.Lane,
					Middlevertice: &uml.Vertice{
						X: 885.000000,
						Y: 306.500000,
					},
					TargetMultiplicity: "0..1",
					SourceMultiplicity: "*",
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
			Heigth: 93.000000,
			Links: []*uml.Link{
				{
					Field: models.Milestone{}.LanesToDisplayMilestoneUse,
					Middlevertice: &uml.Vertice{
						X: 510.000000,
						Y: 559.000000,
					},
					TargetMultiplicity: "*",
					SourceMultiplicity: "0..1",
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Milestone{}.Date,
				},
				{
					Field: models.Milestone{}.DisplayVerticalBar,
				},
				{
					Field: models.Milestone{}.Name,
				},
			},
		},
	},
}
