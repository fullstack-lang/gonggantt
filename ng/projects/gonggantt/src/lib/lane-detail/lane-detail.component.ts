// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { LaneDB } from '../lane-db'
import { LaneService } from '../lane.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { GanttDB } from '../gantt-db'
import { GroupDB } from '../group-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// LaneDetailComponent is initilizaed from different routes
// LaneDetailComponentState detail different cases 
enum LaneDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Gantt_Lanes_SET,
	CREATE_INSTANCE_WITH_ASSOCIATION_Group_GroupLanes_SET,
}

@Component({
	selector: 'app-lane-detail',
	templateUrl: './lane-detail.component.html',
	styleUrls: ['./lane-detail.component.css'],
})
export class LaneDetailComponent implements OnInit {

	// insertion point for declarations

	// the LaneDB of interest
	lane: LaneDB = new LaneDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: LaneDetailComponentState = LaneDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	GONG__StackPath: string = ""

	constructor(
		private laneService: LaneService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		this.GONG__StackPath = this.activatedRoute.snapshot.paramMap.get('GONG__StackPath')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = LaneDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = LaneDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Lanes":
						// console.log("Lane" + " is instanciated with back pointer to instance " + this.id + " Gantt association Lanes")
						this.state = LaneDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Gantt_Lanes_SET
						break;
					case "GroupLanes":
						// console.log("Lane" + " is instanciated with back pointer to instance " + this.id + " Group association GroupLanes")
						this.state = LaneDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Group_GroupLanes_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getLane()

		// observable for changes in structs
		this.laneService.LaneServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLane()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getLane(): void {

		this.frontRepoService.pull(this.GONG__StackPath).subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case LaneDetailComponentState.CREATE_INSTANCE:
						this.lane = new (LaneDB)
						break;
					case LaneDetailComponentState.UPDATE_INSTANCE:
						let lane = frontRepo.Lanes.get(this.id)
						console.assert(lane != undefined, "missing lane with id:" + this.id)
						this.lane = lane!
						break;
					// insertion point for init of association field
					case LaneDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Gantt_Lanes_SET:
						this.lane = new (LaneDB)
						this.lane.Gantt_Lanes_reverse = frontRepo.Gantts.get(this.id)!
						break;
					case LaneDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Group_GroupLanes_SET:
						this.lane = new (LaneDB)
						this.lane.Group_GroupLanes_reverse = frontRepo.Groups.get(this.id)!
						break;
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.lane.Gantt_Lanes_reverse != undefined) {
			if (this.lane.Gantt_LanesDBID == undefined) {
				this.lane.Gantt_LanesDBID = new NullInt64
			}
			this.lane.Gantt_LanesDBID.Int64 = this.lane.Gantt_Lanes_reverse.ID
			this.lane.Gantt_LanesDBID.Valid = true
			if (this.lane.Gantt_LanesDBID_Index == undefined) {
				this.lane.Gantt_LanesDBID_Index = new NullInt64
			}
			this.lane.Gantt_LanesDBID_Index.Valid = true
			this.lane.Gantt_Lanes_reverse = new GanttDB // very important, otherwise, circular JSON
		}
		if (this.lane.Group_GroupLanes_reverse != undefined) {
			if (this.lane.Group_GroupLanesDBID == undefined) {
				this.lane.Group_GroupLanesDBID = new NullInt64
			}
			this.lane.Group_GroupLanesDBID.Int64 = this.lane.Group_GroupLanes_reverse.ID
			this.lane.Group_GroupLanesDBID.Valid = true
			if (this.lane.Group_GroupLanesDBID_Index == undefined) {
				this.lane.Group_GroupLanesDBID_Index = new NullInt64
			}
			this.lane.Group_GroupLanesDBID_Index.Valid = true
			this.lane.Group_GroupLanes_reverse = new GroupDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case LaneDetailComponentState.UPDATE_INSTANCE:
				this.laneService.updateLane(this.lane, this.GONG__StackPath)
					.subscribe(lane => {
						this.laneService.LaneServiceChanged.next("update")
					});
				break;
			default:
				this.laneService.postLane(this.lane, this.GONG__StackPath).subscribe(lane => {
					this.laneService.LaneServiceChanged.next("post")
					this.lane = new (LaneDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.lane.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.lane.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode
			dialogData.GONG__StackPath = this.GONG__StackPath

			// set up the source
			dialogData.SourceStruct = "Lane"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.lane.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
			GONG__StackPath: this.GONG__StackPath,
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

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.lane.Name == "") {
			this.lane.Name = event.value.Name
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
			return this.mapFields_displayAsTextArea.get(fieldName)!
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
