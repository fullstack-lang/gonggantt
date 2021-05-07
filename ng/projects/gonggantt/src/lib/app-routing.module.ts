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
	{ path: 'bars', component: BarsTableComponent, outlet: 'table' },
	{ path: 'bar-adder', component: BarDetailComponent, outlet: 'editor' },
	{ path: 'bar-adder/:id/:association', component: BarDetailComponent, outlet: 'editor' },
	{ path: 'bar-detail/:id', component: BarDetailComponent, outlet: 'editor' },
	{ path: 'bar-presentation/:id', component: BarPresentationComponent, outlet: 'presentation' },
	{ path: 'bar-presentation-special/:id', component: BarPresentationComponent, outlet: 'barpres' },

	{ path: 'gantts', component: GanttsTableComponent, outlet: 'table' },
	{ path: 'gantt-adder', component: GanttDetailComponent, outlet: 'editor' },
	{ path: 'gantt-adder/:id/:association', component: GanttDetailComponent, outlet: 'editor' },
	{ path: 'gantt-detail/:id', component: GanttDetailComponent, outlet: 'editor' },
	{ path: 'gantt-presentation/:id', component: GanttPresentationComponent, outlet: 'presentation' },
	{ path: 'gantt-presentation-special/:id', component: GanttPresentationComponent, outlet: 'ganttpres' },

	{ path: 'groups', component: GroupsTableComponent, outlet: 'table' },
	{ path: 'group-adder', component: GroupDetailComponent, outlet: 'editor' },
	{ path: 'group-adder/:id/:association', component: GroupDetailComponent, outlet: 'editor' },
	{ path: 'group-detail/:id', component: GroupDetailComponent, outlet: 'editor' },
	{ path: 'group-presentation/:id', component: GroupPresentationComponent, outlet: 'presentation' },
	{ path: 'group-presentation-special/:id', component: GroupPresentationComponent, outlet: 'grouppres' },

	{ path: 'lanes', component: LanesTableComponent, outlet: 'table' },
	{ path: 'lane-adder', component: LaneDetailComponent, outlet: 'editor' },
	{ path: 'lane-adder/:id/:association', component: LaneDetailComponent, outlet: 'editor' },
	{ path: 'lane-detail/:id', component: LaneDetailComponent, outlet: 'editor' },
	{ path: 'lane-presentation/:id', component: LanePresentationComponent, outlet: 'presentation' },
	{ path: 'lane-presentation-special/:id', component: LanePresentationComponent, outlet: 'lanepres' },

	{ path: 'milestones', component: MilestonesTableComponent, outlet: 'table' },
	{ path: 'milestone-adder', component: MilestoneDetailComponent, outlet: 'editor' },
	{ path: 'milestone-adder/:id/:association', component: MilestoneDetailComponent, outlet: 'editor' },
	{ path: 'milestone-detail/:id', component: MilestoneDetailComponent, outlet: 'editor' },
	{ path: 'milestone-presentation/:id', component: MilestonePresentationComponent, outlet: 'presentation' },
	{ path: 'milestone-presentation-special/:id', component: MilestonePresentationComponent, outlet: 'milestonepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
