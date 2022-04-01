// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { LaneUseDB } from '../laneuse-db'
import { LaneUseService } from '../laneuse.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { MilestoneDB } from '../milestone-db'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// LaneUseDetailComponent is initilizaed from different routes
// LaneUseDetailComponentState detail different cases 
enum LaneUseDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Milestone_LanesToDisplayMilestone_SET,
}

@Component({
	selector: 'app-laneuse-detail',
	templateUrl: './laneuse-detail.component.html',
	styleUrls: ['./laneuse-detail.component.css'],
})
export class LaneUseDetailComponent implements OnInit {

	// insertion point for declarations

	// the LaneUseDB of interest
	laneuse: LaneUseDB = new LaneUseDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: LaneUseDetailComponentState = LaneUseDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private laneuseService: LaneUseService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id')!;
		this.originStruct = this.route.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = LaneUseDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = LaneUseDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "LanesToDisplayMilestone":
						// console.log("LaneUse" + " is instanciated with back pointer to instance " + this.id + " Milestone association LanesToDisplayMilestone")
						this.state = LaneUseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Milestone_LanesToDisplayMilestone_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getLaneUse()

		// observable for changes in structs
		this.laneuseService.LaneUseServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getLaneUse()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getLaneUse(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case LaneUseDetailComponentState.CREATE_INSTANCE:
						this.laneuse = new (LaneUseDB)
						break;
					case LaneUseDetailComponentState.UPDATE_INSTANCE:
						let laneuse = frontRepo.LaneUses.get(this.id)
						console.assert(laneuse != undefined, "missing laneuse with id:" + this.id)
						this.laneuse = laneuse!
						break;
					// insertion point for init of association field
					case LaneUseDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Milestone_LanesToDisplayMilestone_SET:
						this.laneuse = new (LaneUseDB)
						this.laneuse.Milestone_LanesToDisplayMilestone_reverse = frontRepo.Milestones.get(this.id)!
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
		if (this.laneuse.LaneID == undefined) {
			this.laneuse.LaneID = new NullInt64
		}
		if (this.laneuse.Lane != undefined) {
			this.laneuse.LaneID.Int64 = this.laneuse.Lane.ID
			this.laneuse.LaneID.Valid = true
		} else {
			this.laneuse.LaneID.Int64 = 0
			this.laneuse.LaneID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.laneuse.Milestone_LanesToDisplayMilestone_reverse != undefined) {
			if (this.laneuse.Milestone_LanesToDisplayMilestoneDBID == undefined) {
				this.laneuse.Milestone_LanesToDisplayMilestoneDBID = new NullInt64
			}
			this.laneuse.Milestone_LanesToDisplayMilestoneDBID.Int64 = this.laneuse.Milestone_LanesToDisplayMilestone_reverse.ID
			this.laneuse.Milestone_LanesToDisplayMilestoneDBID.Valid = true
			if (this.laneuse.Milestone_LanesToDisplayMilestoneDBID_Index == undefined) {
				this.laneuse.Milestone_LanesToDisplayMilestoneDBID_Index = new NullInt64
			}
			this.laneuse.Milestone_LanesToDisplayMilestoneDBID_Index.Valid = true
			this.laneuse.Milestone_LanesToDisplayMilestone_reverse = new MilestoneDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case LaneUseDetailComponentState.UPDATE_INSTANCE:
				this.laneuseService.updateLaneUse(this.laneuse)
					.subscribe(laneuse => {
						this.laneuseService.LaneUseServiceChanged.next("update")
					});
				break;
			default:
				this.laneuseService.postLaneUse(this.laneuse).subscribe(laneuse => {
					this.laneuseService.LaneUseServiceChanged.next("post")
					this.laneuse = new (LaneUseDB) // reset fields
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

			dialogData.ID = this.laneuse.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

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
			dialogData.ID = this.laneuse.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "LaneUse"
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
			ID: this.laneuse.ID,
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

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.laneuse.Name == "") {
			this.laneuse.Name = event.value.Name
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
