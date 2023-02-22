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
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrows/:GONG__StackPath', component: ArrowsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrow-adder/:GONG__StackPath', component: ArrowDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrow-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: ArrowDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-arrow-detail/:id/:GONG__StackPath', component: ArrowDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-bars/:GONG__StackPath', component: BarsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-adder/:GONG__StackPath', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-detail/:id/:GONG__StackPath', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-gantts/:GONG__StackPath', component: GanttsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-adder/:GONG__StackPath', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-detail/:id/:GONG__StackPath', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-groups/:GONG__StackPath', component: GroupsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-adder/:GONG__StackPath', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-detail/:id/:GONG__StackPath', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-lanes/:GONG__StackPath', component: LanesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-adder/:GONG__StackPath', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-detail/:id/:GONG__StackPath', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuses/:GONG__StackPath', component: LaneUsesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuse-adder/:GONG__StackPath', component: LaneUseDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuse-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: LaneUseDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-laneuse-detail/:id/:GONG__StackPath', component: LaneUseDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-milestones/:GONG__StackPath', component: MilestonesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-adder/:GONG__StackPath', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-detail/:id/:GONG__StackPath', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
