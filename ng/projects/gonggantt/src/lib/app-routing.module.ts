import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { BarsTableComponent } from './bars-table/bars-table.component'
import { BarDetailComponent } from './bar-detail/bar-detail.component'
import { BarPresentationComponent } from './bar-presentation/bar-presentation.component'

import { GanttsTableComponent } from './gantts-table/gantts-table.component'
import { GanttDetailComponent } from './gantt-detail/gantt-detail.component'
import { GanttPresentationComponent } from './gantt-presentation/gantt-presentation.component'

import { LanesTableComponent } from './lanes-table/lanes-table.component'
import { LaneDetailComponent } from './lane-detail/lane-detail.component'
import { LanePresentationComponent } from './lane-presentation/lane-presentation.component'


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

	{ path: 'lanes', component: LanesTableComponent, outlet: 'table' },
	{ path: 'lane-adder', component: LaneDetailComponent, outlet: 'editor' },
	{ path: 'lane-adder/:id/:association', component: LaneDetailComponent, outlet: 'editor' },
	{ path: 'lane-detail/:id', component: LaneDetailComponent, outlet: 'editor' },
	{ path: 'lane-presentation/:id', component: LanePresentationComponent, outlet: 'presentation' },
	{ path: 'lane-presentation-special/:id', component: LanePresentationComponent, outlet: 'lanepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
