// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GanttDB } from '../gantt-db'
import { GanttService } from '../gantt.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

// GanttDetailComponent is initilizaed from different routes
// GanttDetailComponentState detail different cases 
enum GanttDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-gantt-detail',
	templateUrl: './gantt-detail.component.html',
	styleUrls: ['./gantt-detail.component.css'],
})
export class GanttDetailComponent implements OnInit {

	// insertion point for declarations
	AlignOnStartEndOnYearStartFormControl = new FormControl(false);

	// the GanttDB of interest
	gantt: GanttDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: GanttDetailComponentState

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string
	originStructFieldName: string

	constructor(
		private ganttService: GanttService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id');
		this.originStruct = this.route.snapshot.paramMap.get('originStruct');
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName');

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = GanttDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = GanttDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getGantt()

		// observable for changes in structs
		this.ganttService.GanttServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGantt()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGantt(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case GanttDetailComponentState.CREATE_INSTANCE:
						this.gantt = new (GanttDB)
						break;
					case GanttDetailComponentState.UPDATE_INSTANCE:
						this.gantt = frontRepo.Gantts.get(this.id)
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.AlignOnStartEndOnYearStartFormControl.setValue(this.gantt.AlignOnStartEndOnYearStart)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.gantt.AlignOnStartEndOnYearStart = this.AlignOnStartEndOnYearStartFormControl.value

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case GanttDetailComponentState.UPDATE_INSTANCE:
				this.ganttService.updateGantt(this.gantt)
					.subscribe(gantt => {
						this.ganttService.GanttServiceChanged.next("update")
					});
				break;
			default:
				this.ganttService.postGantt(this.gantt).subscribe(gantt => {
					this.ganttService.GanttServiceChanged.next("post")
					this.gantt = {} // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.gantt.ID,
			ReversePointer: reverseField,
			OrderingMode: false,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.gantt.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event) {
		if (this.gantt.Name == undefined) {
			this.gantt.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
