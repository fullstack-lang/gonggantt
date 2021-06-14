import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { BarsTableComponent } from './bars-table/bars-table.component'
import { BarDetailComponent } from './bar-detail/bar-detail.component'
import { BarPresentationComponent } from './bar-presentation/bar-presentation.component'

import { GanttsTableComponent } from './gantts-table/gantts-table.component'
import { GanttDetailComponent } from './gantt-detail/gantt-detail.component'
import { GanttPresentationComponent } from './gantt-presentation/gantt-presentation.component'

import { GroupsTableComponent } from './groups-table/groups-table.component'
import { GroupDetailComponent } from './group-detail/group-detail.component'
import { GroupPresentationComponent } from './group-presentation/group-presentation.component'

import { LanesTableComponent } from './lanes-table/lanes-table.component'
import { LaneDetailComponent } from './lane-detail/lane-detail.component'
import { LanePresentationComponent } from './lane-presentation/lane-presentation.component'

import { MilestonesTableComponent } from './milestones-table/milestones-table.component'
import { MilestoneDetailComponent } from './milestone-detail/milestone-detail.component'
import { MilestonePresentationComponent } from './milestone-presentation/milestone-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gonggantt_go-bars', component: BarsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-adder', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-adder/:id/:association', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-detail/:id', component: BarDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-presentation/:id', component: BarPresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_presentation' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-bar-presentation-special/:id', component: BarPresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_gobarpres' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-gantts', component: GanttsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-adder', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-adder/:id/:association', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-detail/:id', component: GanttDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-presentation/:id', component: GanttPresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_presentation' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-gantt-presentation-special/:id', component: GanttPresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_goganttpres' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-groups', component: GroupsTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-adder', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-adder/:id/:association', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-detail/:id', component: GroupDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-presentation/:id', component: GroupPresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_presentation' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-group-presentation-special/:id', component: GroupPresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_gogrouppres' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-lanes', component: LanesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-adder', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-adder/:id/:association', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-detail/:id', component: LaneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-presentation/:id', component: LanePresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_presentation' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-lane-presentation-special/:id', component: LanePresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_golanepres' },

	{ path: 'github_com_fullstack_lang_gonggantt_go-milestones', component: MilestonesTableComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_table' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-adder', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-adder/:id/:association', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-detail/:id', component: MilestoneDetailComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_editor' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-presentation/:id', component: MilestonePresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_go_presentation' },
	{ path: 'github_com_fullstack_lang_gonggantt_go-milestone-presentation-special/:id', component: MilestonePresentationComponent, outlet: 'github_com_fullstack_lang_gonggantt_gomilestonepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
