import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { ArrowsTableComponent } from './arrows-table/arrows-table.component'
import { ArrowSortingComponent } from './arrow-sorting/arrow-sorting.component'
import { ArrowDetailComponent } from './arrow-detail/arrow-detail.component'

import { BarsTableComponent } from './bars-table/bars-table.component'
import { BarSortingComponent } from './bar-sorting/bar-sorting.component'
import { BarDetailComponent } from './bar-detail/bar-detail.component'

import { GanttsTableComponent } from './gantts-table/gantts-table.component'
import { GanttSortingComponent } from './gantt-sorting/gantt-sorting.component'
import { GanttDetailComponent } from './gantt-detail/gantt-detail.component'

import { GroupsTableComponent } from './groups-table/groups-table.component'
import { GroupSortingComponent } from './group-sorting/group-sorting.component'
import { GroupDetailComponent } from './group-detail/group-detail.component'

import { LanesTableComponent } from './lanes-table/lanes-table.component'
import { LaneSortingComponent } from './lane-sorting/lane-sorting.component'
import { LaneDetailComponent } from './lane-detail/lane-detail.component'

import { LaneUsesTableComponent } from './laneuses-table/laneuses-table.component'
import { LaneUseSortingComponent } from './laneuse-sorting/laneuse-sorting.component'
import { LaneUseDetailComponent } from './laneuse-detail/laneuse-detail.component'

import { MilestonesTableComponent } from './milestones-table/milestones-table.component'
import { MilestoneSortingComponent } from './milestone-sorting/milestone-sorting.component'
import { MilestoneDetailComponent } from './milestone-detail/milestone-detail.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		ArrowsTableComponent,
		ArrowSortingComponent,
		ArrowDetailComponent,

		BarsTableComponent,
		BarSortingComponent,
		BarDetailComponent,

		GanttsTableComponent,
		GanttSortingComponent,
		GanttDetailComponent,

		GroupsTableComponent,
		GroupSortingComponent,
		GroupDetailComponent,

		LanesTableComponent,
		LaneSortingComponent,
		LaneDetailComponent,

		LaneUsesTableComponent,
		LaneUseSortingComponent,
		LaneUseDetailComponent,

		MilestonesTableComponent,
		MilestoneSortingComponent,
		MilestoneDetailComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		ArrowsTableComponent,
		ArrowSortingComponent,
		ArrowDetailComponent,

		BarsTableComponent,
		BarSortingComponent,
		BarDetailComponent,

		GanttsTableComponent,
		GanttSortingComponent,
		GanttDetailComponent,

		GroupsTableComponent,
		GroupSortingComponent,
		GroupDetailComponent,

		LanesTableComponent,
		LaneSortingComponent,
		LaneDetailComponent,

		LaneUsesTableComponent,
		LaneUseSortingComponent,
		LaneUseDetailComponent,

		MilestonesTableComponent,
		MilestoneSortingComponent,
		MilestoneDetailComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongganttModule { }
