// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { MilestoneDB } from '../milestone-db'
import { MilestoneService } from '../milestone.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

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
		// https://stackoverflow.com/questions/54627478/angular-7-routing-to-same-component-but-different-param-not-working
		// this is for routerLink on same component when only queryParameter changes
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
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
				console.log("front repo MilestonePull returned")

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
				this.milestone.Gantt_MilestonesDBID = new NullInt64
				this.milestone.Gantt_MilestonesDBID.Int64 = this.milestone.Gantt_Milestones_reverse.ID
				this.milestone.Gantt_MilestonesDBID.Valid = true
				this.milestone.Gantt_Milestones_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.milestoneService.updateMilestone(this.milestone)
				.subscribe(milestone => {
					this.milestoneService.MilestoneServiceChanged.next("update")

					console.log("milestone saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Gantt_Milestones":
					this.milestone.Gantt_MilestonesDBID = new NullInt64
					this.milestone.Gantt_MilestonesDBID.Int64 = id
					this.milestone.Gantt_MilestonesDBID.Valid = true
					break
			}
			this.milestoneService.postMilestone(this.milestone).subscribe(milestone => {

				this.milestoneService.MilestoneServiceChanged.next("post")

				this.milestone = {} // reset fields
				console.log("milestone added")
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
		dialogConfig.data = {
			ID: this.milestone.ID,
			ReversePointer: reverseField,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
			console.log('The dialog was closed');
		});
	}
}