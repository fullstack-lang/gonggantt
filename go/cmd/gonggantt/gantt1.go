package main

import (
	"time"

	"github.com/fullstack-lang/gonggantt/go/models"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["gantt1"] = gantt1Injection
}

// gantt1Injection will stage objects of database "gantt1"
func gantt1Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Arrow
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1 := (&models.Arrow{Name: "Lane 1 - Bar 1 to Lan 2 - Bar 1"}).Stage()

	// Declarations of staged instances of Bar
	__Bar__000000_Lane_1_Bar_1 := (&models.Bar{Name: "Lane 1 - Bar 1"}).Stage()
	__Bar__000001_Lane_1_Bar_2 := (&models.Bar{Name: "Lane 1 - Bar 2"}).Stage()
	__Bar__000002_Lane_2_Bar_1 := (&models.Bar{Name: "Lane 2 - Bar 1"}).Stage()
	__Bar__000003_Lane_3_Bar_1 := (&models.Bar{Name: "Lane-3-Bar-1"}).Stage()

	// Declarations of staged instances of Gantt
	__Gantt__000000_Test := (&models.Gantt{Name: "Test"}).Stage()

	// Declarations of staged instances of Group
	__Group__000000_Group_of_lanes_3_4 := (&models.Group{Name: "Group of lanes 3&4"}).Stage()

	// Declarations of staged instances of Lane
	__Lane__000000_Lane_1 := (&models.Lane{Name: "Lane-1"}).Stage()
	__Lane__000001_Lane_2 := (&models.Lane{Name: "Lane-2"}).Stage()
	__Lane__000002_Lane_3 := (&models.Lane{Name: "Lane-3"}).Stage()
	__Lane__000003_Lane_4 := (&models.Lane{Name: "Lane-4"}).Stage()

	// Declarations of staged instances of LaneUse
	__LaneUse__000000_Lane_1 := (&models.LaneUse{Name: "Lane-1"}).Stage()
	__LaneUse__000001_Lane_3 := (&models.LaneUse{Name: "Lane-3"}).Stage()

	// Declarations of staged instances of Milestone
	__Milestone__000000_june_2023_ := (&models.Milestone{Name: "june 2023 !"}).Stage()

	// Setup of values

	// Arrow Lane 1 - Bar 1 to Lan 2 - Bar 1 values setup
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.Name = `Lane 1 - Bar 1 to Lan 2 - Bar 1`
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalColor = ``
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalStroke = ``

	// Bar Lane 1 - Bar 1 values setup
	__Bar__000000_Lane_1_Bar_1.Name = `Lane 1 - Bar 1`
	__Bar__000000_Lane_1_Bar_1.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-04-01 00:00:00 +0000 +0000")
	__Bar__000000_Lane_1_Bar_1.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-04-01 00:00:00 +0000 +0000")
	__Bar__000000_Lane_1_Bar_1.OptionnalColor = `red`
	__Bar__000000_Lane_1_Bar_1.OptionnalStroke = `red`
	__Bar__000000_Lane_1_Bar_1.FillOpacity = 0.000000
	__Bar__000000_Lane_1_Bar_1.StrokeWidth = 3.000000
	__Bar__000000_Lane_1_Bar_1.StrokeDashArray = `5 5`

	// Bar Lane 1 - Bar 2 values setup
	__Bar__000001_Lane_1_Bar_2.Name = `Lane 1 - Bar 2`
	__Bar__000001_Lane_1_Bar_2.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-01-01 00:00:00 +0000 +0000")
	__Bar__000001_Lane_1_Bar_2.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-01 00:00:00 +0000 +0000")
	__Bar__000001_Lane_1_Bar_2.OptionnalColor = `green`
	__Bar__000001_Lane_1_Bar_2.OptionnalStroke = ``
	__Bar__000001_Lane_1_Bar_2.FillOpacity = 0.000000
	__Bar__000001_Lane_1_Bar_2.StrokeWidth = 0.000000
	__Bar__000001_Lane_1_Bar_2.StrokeDashArray = ``

	// Bar Lane 2 - Bar 1 values setup
	__Bar__000002_Lane_2_Bar_1.Name = `Lane 2 - Bar 1`
	__Bar__000002_Lane_2_Bar_1.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-01-01 00:00:00 +0000 +0000")
	__Bar__000002_Lane_2_Bar_1.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-01-01 00:00:00 +0000 +0000")
	__Bar__000002_Lane_2_Bar_1.OptionnalColor = ``
	__Bar__000002_Lane_2_Bar_1.OptionnalStroke = ``
	__Bar__000002_Lane_2_Bar_1.FillOpacity = 0.000000
	__Bar__000002_Lane_2_Bar_1.StrokeWidth = 0.000000
	__Bar__000002_Lane_2_Bar_1.StrokeDashArray = ``

	// Bar Lane-3-Bar-1 values setup
	__Bar__000003_Lane_3_Bar_1.Name = `Lane-3-Bar-1`
	__Bar__000003_Lane_3_Bar_1.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 +0000")
	__Bar__000003_Lane_3_Bar_1.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-01 00:00:00 +0000 +0000")
	__Bar__000003_Lane_3_Bar_1.OptionnalColor = ``
	__Bar__000003_Lane_3_Bar_1.OptionnalStroke = ``
	__Bar__000003_Lane_3_Bar_1.FillOpacity = 0.000000
	__Bar__000003_Lane_3_Bar_1.StrokeWidth = 0.000000
	__Bar__000003_Lane_3_Bar_1.StrokeDashArray = ``

	// Gantt Test values setup
	__Gantt__000000_Test.Name = `Test`
	__Gantt__000000_Test.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-01-01 00:00:00 +0000 +0000")
	__Gantt__000000_Test.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-12-31 00:00:00 +0000 +0000")
	__Gantt__000000_Test.UseManualStartAndEndDates = false
	__Gantt__000000_Test.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2017-02-12 00:00:00 +0000 +0000")
	__Gantt__000000_Test.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-02-12 00:00:00 +0000 +0000")
	__Gantt__000000_Test.LaneHeight = 80.000000
	__Gantt__000000_Test.RatioBarToLaneHeight = 0.700000
	__Gantt__000000_Test.YTopMargin = 10.000000
	__Gantt__000000_Test.XLeftText = 8.000000
	__Gantt__000000_Test.TextHeight = 16.000000
	__Gantt__000000_Test.XLeftLanes = 150.000000
	__Gantt__000000_Test.XRightMargin = 800.000000
	__Gantt__000000_Test.ArrowLengthToTheRightOfStartBar = 50.000000
	__Gantt__000000_Test.ArrowTipLenght = 15.000000
	__Gantt__000000_Test.TimeLine_Color = `"black"`
	__Gantt__000000_Test.TimeLine_FillOpacity = 1.000000
	__Gantt__000000_Test.TimeLine_Stroke = `"black"`
	__Gantt__000000_Test.TimeLine_StrokeWidth = 1.000000
	__Gantt__000000_Test.Group_Stroke = `blue`
	__Gantt__000000_Test.Group_StrokeWidth = 0.300000
	__Gantt__000000_Test.Group_StrokeDashArray = ``
	__Gantt__000000_Test.DateYOffset = 20.000000
	__Gantt__000000_Test.AlignOnStartEndOnYearStart = true

	// Group Group of lanes 3&4 values setup
	__Group__000000_Group_of_lanes_3_4.Name = `Group of lanes 3&4`

	// Lane Lane-1 values setup
	__Lane__000000_Lane_1.Name = `Lane-1`
	__Lane__000000_Lane_1.Order = 61

	// Lane Lane-2 values setup
	__Lane__000001_Lane_2.Name = `Lane-2`
	__Lane__000001_Lane_2.Order = 20

	// Lane Lane-3 values setup
	__Lane__000002_Lane_3.Name = `Lane-3`
	__Lane__000002_Lane_3.Order = 50

	// Lane Lane-4 values setup
	__Lane__000003_Lane_4.Name = `Lane-4`
	__Lane__000003_Lane_4.Order = 23

	// LaneUse Lane-1 values setup
	__LaneUse__000000_Lane_1.Name = `Lane-1`

	// LaneUse Lane-3 values setup
	__LaneUse__000001_Lane_3.Name = `Lane-3`

	// Milestone june 2023 ! values setup
	__Milestone__000000_june_2023_.Name = `june 2023 !`
	__Milestone__000000_june_2023_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-06-01 00:00:00 +0000 +0000")

	// Setup of pointers
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.From = __Bar__000003_Lane_3_Bar_1
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.To = __Bar__000002_Lane_2_Bar_1
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000001_Lane_2)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000003_Lane_4)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000002_Lane_3)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000000_Lane_1)
	__Gantt__000000_Test.Milestones = append(__Gantt__000000_Test.Milestones, __Milestone__000000_june_2023_)
	__Gantt__000000_Test.Groups = append(__Gantt__000000_Test.Groups, __Group__000000_Group_of_lanes_3_4)
	__Gantt__000000_Test.Arrows = append(__Gantt__000000_Test.Arrows, __Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1)
	__Group__000000_Group_of_lanes_3_4.GroupLanes = append(__Group__000000_Group_of_lanes_3_4.GroupLanes, __Lane__000002_Lane_3)
	__Group__000000_Group_of_lanes_3_4.GroupLanes = append(__Group__000000_Group_of_lanes_3_4.GroupLanes, __Lane__000003_Lane_4)
	__Lane__000000_Lane_1.Bars = append(__Lane__000000_Lane_1.Bars, __Bar__000000_Lane_1_Bar_1)
	__Lane__000000_Lane_1.Bars = append(__Lane__000000_Lane_1.Bars, __Bar__000001_Lane_1_Bar_2)
	__Lane__000001_Lane_2.Bars = append(__Lane__000001_Lane_2.Bars, __Bar__000002_Lane_2_Bar_1)
	__Lane__000002_Lane_3.Bars = append(__Lane__000002_Lane_3.Bars, __Bar__000003_Lane_3_Bar_1)
	__LaneUse__000000_Lane_1.Lane = __Lane__000000_Lane_1
	__LaneUse__000001_Lane_3.Lane = __Lane__000002_Lane_3
	__Milestone__000000_june_2023_.LanesToDisplayMilestone = append(__Milestone__000000_june_2023_.LanesToDisplayMilestone, __LaneUse__000000_Lane_1)
	__Milestone__000000_june_2023_.LanesToDisplayMilestone = append(__Milestone__000000_june_2023_.LanesToDisplayMilestone, __LaneUse__000001_Lane_3)
}


