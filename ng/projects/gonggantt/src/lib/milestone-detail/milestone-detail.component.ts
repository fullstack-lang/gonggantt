// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { MilestoneDB } from '../milestone-db'
import { MilestoneService } from '../milestone.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-milestone-detail',
	templateUrl: './milestone-detail.component.html',
	styleUrls: ['./milestone-detail.component.css'],
})
export class MilestoneDetailComponent implements OnInit {

	// insertion point for declarations

	// the MilestoneDB of interest
	milestone: MilestoneDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private milestoneService: MilestoneService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getMilestone()

		// observable for changes in structs
		this.milestoneService.MilestoneServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getMilestone()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getMilestone(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.milestone = frontRepo.Milestones.get(id)
				} else {
					this.milestone = new (MilestoneDB)
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
			if (this.milestone.Gantt_Milestones_reverse != undefined) {
				if (this.milestone.Gantt_MilestonesDBID == undefined) {
					this.milestone.Gantt_MilestonesDBID = new NullInt64
				}
				this.milestone.Gantt_MilestonesDBID.Int64 = this.milestone.Gantt_Milestones_reverse.ID
				this.milestone.Gantt_MilestonesDBID.Valid = true
				if (this.milestone.Gantt_MilestonesDBID_Index == undefined) {
					this.milestone.Gantt_MilestonesDBID_Index = new NullInt64
				}
				this.milestone.Gantt_MilestonesDBID_Index.Valid = true
				this.milestone.Gantt_Milestones_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.milestoneService.updateMilestone(this.milestone)
				.subscribe(milestone => {
					this.milestoneService.MilestoneServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Gantt_Milestones":
					this.milestone.Gantt_MilestonesDBID = new NullInt64
					this.milestone.Gantt_MilestonesDBID.Int64 = id
					this.milestone.Gantt_MilestonesDBID.Valid = true
					this.milestone.Gantt_MilestonesDBID_Index = new NullInt64
					this.milestone.Gantt_MilestonesDBID_Index.Valid = true
					break
			}
			this.milestoneService.postMilestone(this.milestone).subscribe(milestone => {

				this.milestoneService.MilestoneServiceChanged.next("post")

				this.milestone = {} // reset fields
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
			ID: this.milestone.ID,
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
			ID: this.milestone.ID,
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
		if (this.milestone.Name == undefined) {
			this.milestone.Name = event.value.Name		
		}
	}
}
