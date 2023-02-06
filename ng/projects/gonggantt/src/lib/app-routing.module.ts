import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ArrowsTableComponent } from './arrows-table/arrows-table.component'
import { ArrowDetailComponent } from './arrow-detail/arrow-detail.component'

import { BarsTableComponent } from './bars-table/bars-table.component'
import { BarDetailComponent } from './bar-detail/bar-detail.component'

import { GanttsTableComponent } from './gantts-table/gantts-table.component'
import { GanttDetailComponent } from './gantt-detail/gantt-detail.component'

import { GroupsTableComponent } from './groups-table/groups-table.component'
import { GroupDetailComponent } from './group-detail/group-detail.component'

import { LanesTableComponent } from './lanes-table/lanes-table.component'
import { LaneDetailComponent } from './lane-detail/lane-detail.component'

import { LaneUsesTableComponent } from './laneuses-table/laneuses-table.component'
import { LaneUseDetailComponent } from './laneuse-detail/laneuse-detail.component'

import { MilestonesTableComponent } from './milestones-table/milestones-table.component'
import { MilestoneDetailComponent } from './milestone-detail/milestone-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrows', component: ArrowsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrow-adder', component: ArrowDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrow-adder/:id/:originStruct/:originStructFieldName', component: ArrowDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrow-detail/:id', component: ArrowDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-bars', component: BarsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-adder', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-adder/:id/:originStruct/:originStructFieldName', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-detail/:id', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-gantts', component: GanttsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-adder', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-adder/:id/:originStruct/:originStructFieldName', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-detail/:id', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-groups', component: GroupsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-adder', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-adder/:id/:originStruct/:originStructFieldName', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-detail/:id', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-lanes', component: LanesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-adder', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-adder/:id/:originStruct/:originStructFieldName', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-detail/:id', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuses', component: LaneUsesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuse-adder', component: LaneUseDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuse-adder/:id/:originStruct/:originStructFieldName', component: LaneUseDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuse-detail/:id', component: LaneUseDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-milestones', component: MilestonesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-adder', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-adder/:id/:originStruct/:originStructFieldName', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-detail/:id', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
