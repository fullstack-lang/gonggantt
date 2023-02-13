package main

import (
	"time"

	"github.com/fullstack-lang/gonggantt/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_gantt1 models.StageStruct
var ___dummy__Time_gantt1 time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_gantt1 map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["gantt1"] = gantt1Injection
// }

// gantt1Injection will stage objects of database "gantt1"
func gantt1Injection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Arrow
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1 := (&models.Arrow{Name: `Lane 1 - Bar 1 to Lan 2 - Bar 1`}).Stage()

	// Declarations of staged instances of Bar
	__Bar__000000_Lane_1_Bar_1 := (&models.Bar{Name: `Lane 1 - Bar 1`}).Stage()
	__Bar__000001_Lane_1_Bar_2 := (&models.Bar{Name: `Lane 1 - Bar 2`}).Stage()
	__Bar__000002_Lane_2_Bar_1 := (&models.Bar{Name: `Lane 2 - Bar 1`}).Stage()
	__Bar__000003_Lane_3_Bar_1 := (&models.Bar{Name: `Lane-3-Bar-1`}).Stage()

	// Declarations of staged instances of Gantt
	__Gantt__000000_Test := (&models.Gantt{Name: `Test`}).Stage()

	// Declarations of staged instances of Group
	__Group__000000_Group_of_lanes_3_4 := (&models.Group{Name: `Group of lanes 3&4`}).Stage()

	// Declarations of staged instances of Lane
	__Lane__000000_Lane_1 := (&models.Lane{Name: `Lane-1`}).Stage()
	__Lane__000001_Lane_2 := (&models.Lane{Name: `Lane-2`}).Stage()
	__Lane__000002_Lane_3 := (&models.Lane{Name: `Lane-3`}).Stage()
	__Lane__000003_Lane_4 := (&models.Lane{Name: `Lane-4`}).Stage()

	// Declarations of staged instances of LaneUse
	__LaneUse__000000_Lane_1 := (&models.LaneUse{Name: `Lane-1`}).Stage()
	__LaneUse__000001_Lane_2 := (&models.LaneUse{Name: `Lane-2`}).Stage()
	__LaneUse__000002_Lane_3 := (&models.LaneUse{Name: `Lane-3`}).Stage()
	__LaneUse__000003_Lane_3 := (&models.LaneUse{Name: `Lane-3`}).Stage()
	__LaneUse__000004_Lane_4 := (&models.LaneUse{Name: `Lane-4`}).Stage()

	// Declarations of staged instances of Milestone
	__Milestone__000000_Oct_2024 := (&models.Milestone{Name: `Oct 2024`}).Stage()
	__Milestone__000001_june_2023_ := (&models.Milestone{Name: `june 2023 !`}).Stage()

	// Setup of values

	// Arrow values setup
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.Name = `Lane 1 - Bar 1 to Lan 2 - Bar 1`
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalColor = ``
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.OptionnalStroke = ``

	// Bar values setup
	__Bar__000000_Lane_1_Bar_1.Name = `Lane 1 - Bar 1`
	__Bar__000000_Lane_1_Bar_1.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-04-01 00:00:00 +0000 +0000")
	__Bar__000000_Lane_1_Bar_1.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-04-01 00:00:00 +0000 +0000")
	__Bar__000000_Lane_1_Bar_1.OptionnalColor = `red`
	__Bar__000000_Lane_1_Bar_1.OptionnalStroke = `red`
	__Bar__000000_Lane_1_Bar_1.FillOpacity = 0.000000
	__Bar__000000_Lane_1_Bar_1.StrokeWidth = 3.000000
	__Bar__000000_Lane_1_Bar_1.StrokeDashArray = `5 5`

	// Bar values setup
	__Bar__000001_Lane_1_Bar_2.Name = `Lane 1 - Bar 2`
	__Bar__000001_Lane_1_Bar_2.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-01-01 00:00:00 +0000 +0000")
	__Bar__000001_Lane_1_Bar_2.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-01 00:00:00 +0000 +0000")
	__Bar__000001_Lane_1_Bar_2.OptionnalColor = `green`
	__Bar__000001_Lane_1_Bar_2.OptionnalStroke = ``
	__Bar__000001_Lane_1_Bar_2.FillOpacity = 0.000000
	__Bar__000001_Lane_1_Bar_2.StrokeWidth = 0.000000
	__Bar__000001_Lane_1_Bar_2.StrokeDashArray = ``

	// Bar values setup
	__Bar__000002_Lane_2_Bar_1.Name = `Lane 2 - Bar 1`
	__Bar__000002_Lane_2_Bar_1.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-01-01 00:00:00 +0000 +0000")
	__Bar__000002_Lane_2_Bar_1.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-01-01 00:00:00 +0000 +0000")
	__Bar__000002_Lane_2_Bar_1.OptionnalColor = ``
	__Bar__000002_Lane_2_Bar_1.OptionnalStroke = ``
	__Bar__000002_Lane_2_Bar_1.FillOpacity = 0.000000
	__Bar__000002_Lane_2_Bar_1.StrokeWidth = 0.000000
	__Bar__000002_Lane_2_Bar_1.StrokeDashArray = ``

	// Bar values setup
	__Bar__000003_Lane_3_Bar_1.Name = `Lane-3-Bar-1`
	__Bar__000003_Lane_3_Bar_1.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 +0000")
	__Bar__000003_Lane_3_Bar_1.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-01-01 00:00:00 +0000 +0000")
	__Bar__000003_Lane_3_Bar_1.OptionnalColor = ``
	__Bar__000003_Lane_3_Bar_1.OptionnalStroke = ``
	__Bar__000003_Lane_3_Bar_1.FillOpacity = 0.000000
	__Bar__000003_Lane_3_Bar_1.StrokeWidth = 0.000000
	__Bar__000003_Lane_3_Bar_1.StrokeDashArray = ``

	// Gantt values setup
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

	// Group values setup
	__Group__000000_Group_of_lanes_3_4.Name = `Group of lanes 3&4`

	// Lane values setup
	__Lane__000000_Lane_1.Name = `Lane-1`
	__Lane__000000_Lane_1.Order = 61

	// Lane values setup
	__Lane__000001_Lane_2.Name = `Lane-2`
	__Lane__000001_Lane_2.Order = 20

	// Lane values setup
	__Lane__000002_Lane_3.Name = `Lane-3`
	__Lane__000002_Lane_3.Order = 50

	// Lane values setup
	__Lane__000003_Lane_4.Name = `Lane-4`
	__Lane__000003_Lane_4.Order = 23

	// LaneUse values setup
	__LaneUse__000000_Lane_1.Name = `Lane-1`

	// LaneUse values setup
	__LaneUse__000001_Lane_2.Name = `Lane-2`

	// LaneUse values setup
	__LaneUse__000002_Lane_3.Name = `Lane-3`

	// LaneUse values setup
	__LaneUse__000003_Lane_3.Name = `Lane-3`

	// LaneUse values setup
	__LaneUse__000004_Lane_4.Name = `Lane-4`

	// Milestone values setup
	__Milestone__000000_Oct_2024.Name = `Oct 2024`
	__Milestone__000000_Oct_2024.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2024-09-30 22:00:00 +0000 +0000")
	__Milestone__000000_Oct_2024.DisplayVerticalBar = false

	// Milestone values setup
	__Milestone__000001_june_2023_.Name = `june 2023 !`
	__Milestone__000001_june_2023_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-06-01 00:00:00 +0000 +0000")
	__Milestone__000001_june_2023_.DisplayVerticalBar = true

	// Setup of pointers
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.From = __Bar__000003_Lane_3_Bar_1
	__Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1.To = __Bar__000002_Lane_2_Bar_1
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000001_Lane_2)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000003_Lane_4)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000002_Lane_3)
	__Gantt__000000_Test.Lanes = append(__Gantt__000000_Test.Lanes, __Lane__000000_Lane_1)
	__Gantt__000000_Test.Milestones = append(__Gantt__000000_Test.Milestones, __Milestone__000001_june_2023_)
	__Gantt__000000_Test.Milestones = append(__Gantt__000000_Test.Milestones, __Milestone__000000_Oct_2024)
	__Gantt__000000_Test.Groups = append(__Gantt__000000_Test.Groups, __Group__000000_Group_of_lanes_3_4)
	__Gantt__000000_Test.Arrows = append(__Gantt__000000_Test.Arrows, __Arrow__000000_Lane_1_Bar_1_to_Lan_2_Bar_1)
	__Group__000000_Group_of_lanes_3_4.GroupLanes = append(__Group__000000_Group_of_lanes_3_4.GroupLanes, __Lane__000002_Lane_3)
	__Group__000000_Group_of_lanes_3_4.GroupLanes = append(__Group__000000_Group_of_lanes_3_4.GroupLanes, __Lane__000003_Lane_4)
	__Lane__000000_Lane_1.Bars = append(__Lane__000000_Lane_1.Bars, __Bar__000000_Lane_1_Bar_1)
	__Lane__000000_Lane_1.Bars = append(__Lane__000000_Lane_1.Bars, __Bar__000001_Lane_1_Bar_2)
	__Lane__000001_Lane_2.Bars = append(__Lane__000001_Lane_2.Bars, __Bar__000002_Lane_2_Bar_1)
	__Lane__000002_Lane_3.Bars = append(__Lane__000002_Lane_3.Bars, __Bar__000003_Lane_3_Bar_1)
	__LaneUse__000000_Lane_1.Lane = __Lane__000000_Lane_1
	__LaneUse__000001_Lane_2.Lane = __Lane__000001_Lane_2
	__LaneUse__000002_Lane_3.Lane = __Lane__000002_Lane_3
	__LaneUse__000003_Lane_3.Lane = __Lane__000002_Lane_3
	__LaneUse__000004_Lane_4.Lane = __Lane__000003_Lane_4
	__Milestone__000000_Oct_2024.LanesToDisplayMilestoneUse = append(__Milestone__000000_Oct_2024.LanesToDisplayMilestoneUse, __LaneUse__000001_Lane_2)
	__Milestone__000000_Oct_2024.LanesToDisplayMilestoneUse = append(__Milestone__000000_Oct_2024.LanesToDisplayMilestoneUse, __LaneUse__000004_Lane_4)
	__Milestone__000000_Oct_2024.LanesToDisplayMilestoneUse = append(__Milestone__000000_Oct_2024.LanesToDisplayMilestoneUse, __LaneUse__000002_Lane_3)
	__Milestone__000001_june_2023_.LanesToDisplayMilestoneUse = append(__Milestone__000001_june_2023_.LanesToDisplayMilestoneUse, __LaneUse__000000_Lane_1)
	__Milestone__000001_june_2023_.LanesToDisplayMilestoneUse = append(__Milestone__000001_june_2023_.LanesToDisplayMilestoneUse, __LaneUse__000003_Lane_3)
}


