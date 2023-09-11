package main

import (
	"time"

	"github.com/fullstack-lang/gonggantt/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_migration models.StageStruct
var ___dummy__Time_migration time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_migration map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["migration"] = migrationInjection
// }

// migrationInjection will stage objects of database "migration"
func migrationInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Arrow

	// Declarations of staged instances of Bar
	__Bar__000000_Conversion_mod_le := (&models.Bar{Name: `Conversion modèle`}).Stage(stage)
	__Bar__000001_configuration_MMP := (&models.Bar{Name: `configuration MMP`}).Stage(stage)

	// Declarations of staged instances of Gantt
	__Gantt__000000_migration := (&models.Gantt{Name: `migration`}).Stage(stage)

	// Declarations of staged instances of Group
	__Group__000000_Sprint_1 := (&models.Group{Name: `Sprint 1`}).Stage(stage)
	__Group__000001_sprint_2 := (&models.Group{Name: `sprint 2`}).Stage(stage)
	__Group__000002_sprint_3 := (&models.Group{Name: `sprint 3`}).Stage(stage)

	// Declarations of staged instances of Lane
	__Lane__000000_sprint_2_1 := (&models.Lane{Name: `sprint 2.1`}).Stage(stage)
	__Lane__000001_sprint_3_1 := (&models.Lane{Name: `sprint 3.1`}).Stage(stage)

	// Declarations of staged instances of LaneUse

	// Declarations of staged instances of Milestone

	// Setup of values

	// Bar values setup
	__Bar__000000_Conversion_mod_le.Name = `Conversion modèle`
	__Bar__000000_Conversion_mod_le.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Bar__000000_Conversion_mod_le.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Bar__000000_Conversion_mod_le.OptionnalColor = ``
	__Bar__000000_Conversion_mod_le.OptionnalStroke = ``
	__Bar__000000_Conversion_mod_le.FillOpacity = 0.000000
	__Bar__000000_Conversion_mod_le.StrokeWidth = 0.000000
	__Bar__000000_Conversion_mod_le.StrokeDashArray = ``

	// Bar values setup
	__Bar__000001_configuration_MMP.Name = `configuration MMP`
	__Bar__000001_configuration_MMP.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-09-11 14:12:22.214 +0000 UTC")
	__Bar__000001_configuration_MMP.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2023-09-11 14:12:22.214 +0000 UTC")
	__Bar__000001_configuration_MMP.OptionnalColor = ``
	__Bar__000001_configuration_MMP.OptionnalStroke = ``
	__Bar__000001_configuration_MMP.FillOpacity = 0.000000
	__Bar__000001_configuration_MMP.StrokeWidth = 0.000000
	__Bar__000001_configuration_MMP.StrokeDashArray = ``

	// Gantt values setup
	__Gantt__000000_migration.Name = `migration`
	__Gantt__000000_migration.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Gantt__000000_migration.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Gantt__000000_migration.UseManualStartAndEndDates = false
	__Gantt__000000_migration.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Gantt__000000_migration.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Gantt__000000_migration.LaneHeight = 0.000000
	__Gantt__000000_migration.RatioBarToLaneHeight = 0.000000
	__Gantt__000000_migration.YTopMargin = 0.000000
	__Gantt__000000_migration.XLeftText = 0.000000
	__Gantt__000000_migration.TextHeight = 0.000000
	__Gantt__000000_migration.XLeftLanes = 0.000000
	__Gantt__000000_migration.XRightMargin = 0.000000
	__Gantt__000000_migration.ArrowLengthToTheRightOfStartBar = 0.000000
	__Gantt__000000_migration.ArrowTipLenght = 0.000000
	__Gantt__000000_migration.TimeLine_Color = ``
	__Gantt__000000_migration.TimeLine_FillOpacity = 0.000000
	__Gantt__000000_migration.TimeLine_Stroke = ``
	__Gantt__000000_migration.TimeLine_StrokeWidth = 0.000000
	__Gantt__000000_migration.Group_Stroke = ``
	__Gantt__000000_migration.Group_StrokeWidth = 0.000000
	__Gantt__000000_migration.Group_StrokeDashArray = ``
	__Gantt__000000_migration.DateYOffset = 0.000000
	__Gantt__000000_migration.AlignOnStartEndOnYearStart = false

	// Group values setup
	__Group__000000_Sprint_1.Name = `Sprint 1`

	// Group values setup
	__Group__000001_sprint_2.Name = `sprint 2`

	// Group values setup
	__Group__000002_sprint_3.Name = `sprint 3`

	// Lane values setup
	__Lane__000000_sprint_2_1.Name = `sprint 2.1`
	__Lane__000000_sprint_2_1.Order = 0

	// Lane values setup
	__Lane__000001_sprint_3_1.Name = `sprint 3.1`
	__Lane__000001_sprint_3_1.Order = 0

	// Setup of pointers
	__Gantt__000000_migration.Lanes = append(__Gantt__000000_migration.Lanes, __Lane__000000_sprint_2_1)
	__Gantt__000000_migration.Lanes = append(__Gantt__000000_migration.Lanes, __Lane__000001_sprint_3_1)
	__Gantt__000000_migration.Groups = append(__Gantt__000000_migration.Groups, __Group__000000_Sprint_1)
	__Gantt__000000_migration.Groups = append(__Gantt__000000_migration.Groups, __Group__000001_sprint_2)
	__Gantt__000000_migration.Groups = append(__Gantt__000000_migration.Groups, __Group__000002_sprint_3)
	__Group__000001_sprint_2.GroupLanes = append(__Group__000001_sprint_2.GroupLanes, __Lane__000000_sprint_2_1)
	__Group__000002_sprint_3.GroupLanes = append(__Group__000002_sprint_3.GroupLanes, __Lane__000001_sprint_3_1)
	__Lane__000000_sprint_2_1.Bars = append(__Lane__000000_sprint_2_1.Bars, __Bar__000001_configuration_MMP)
	__Lane__000001_sprint_3_1.Bars = append(__Lane__000001_sprint_3_1.Bars, __Bar__000000_Conversion_mod_le)
}


