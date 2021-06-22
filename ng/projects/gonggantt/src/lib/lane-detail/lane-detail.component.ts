// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { LaneDB } from '../lane-db'
import { LaneService } from '../lane.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-lane-detail',
	templateUrl: './lane-detail.component.html',
	styleUrls: ['./lane-detail.component.css'],
})
export class LaneDetailComponent implements OnInit {

	// insertion point for declarations

	// the LaneDB of interest
	lane: LaneDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private laneService: LaneService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
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
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.lane = frontRepo.Lanes.get(id)
				} else {
					this.lane = new (LaneDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
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
				this.lane.Gantt_Lanes_reverse = undefined // very important, otherwise, circular JSON
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
				this.lane.Group_GroupLanes_reverse = undefined // very important, otherwise, circular JSON
			}
			if (this.lane.Milestone_DiamonfAndTextAnchors_reverse != undefined) {
				if (this.lane.Milestone_DiamonfAndTextAnchorsDBID == undefined) {
					this.lane.Milestone_DiamonfAndTextAnchorsDBID = new NullInt64
				}
				this.lane.Milestone_DiamonfAndTextAnchorsDBID.Int64 = this.lane.Milestone_DiamonfAndTextAnchors_reverse.ID
				this.lane.Milestone_DiamonfAndTextAnchorsDBID.Valid = true
				if (this.lane.Milestone_DiamonfAndTextAnchorsDBID_Index == undefined) {
					this.lane.Milestone_DiamonfAndTextAnchorsDBID_Index = new NullInt64
				}
				this.lane.Milestone_DiamonfAndTextAnchorsDBID_Index.Valid = true
				this.lane.Milestone_DiamonfAndTextAnchors_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.laneService.updateLane(this.lane)
				.subscribe(lane => {
					this.laneService.LaneServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Gantt_Lanes":
					this.lane.Gantt_LanesDBID = new NullInt64
					this.lane.Gantt_LanesDBID.Int64 = id
					this.lane.Gantt_LanesDBID.Valid = true
					this.lane.Gantt_LanesDBID_Index = new NullInt64
					this.lane.Gantt_LanesDBID_Index.Valid = true
					break
				case "Group_GroupLanes":
					this.lane.Group_GroupLanesDBID = new NullInt64
					this.lane.Group_GroupLanesDBID.Int64 = id
					this.lane.Group_GroupLanesDBID.Valid = true
					this.lane.Group_GroupLanesDBID_Index = new NullInt64
					this.lane.Group_GroupLanesDBID_Index.Valid = true
					break
				case "Milestone_DiamonfAndTextAnchors":
					this.lane.Milestone_DiamonfAndTextAnchorsDBID = new NullInt64
					this.lane.Milestone_DiamonfAndTextAnchorsDBID.Int64 = id
					this.lane.Milestone_DiamonfAndTextAnchorsDBID.Valid = true
					this.lane.Milestone_DiamonfAndTextAnchorsDBID_Index = new NullInt64
					this.lane.Milestone_DiamonfAndTextAnchorsDBID_Index.Valid = true
					break
			}
			this.laneService.postLane(this.lane).subscribe(lane => {

				this.laneService.LaneServiceChanged.next("post")

				this.lane = {} // reset fields
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
			ID: this.lane.ID,
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
			ID: this.lane.ID,
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
		if (this.lane.Name == undefined) {
			this.lane.Name = event.value.Name		
		}
	}
}
