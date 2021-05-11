// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GroupDB } from '../group-db'
import { GroupService } from '../group.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-group-detail',
	templateUrl: './group-detail.component.html',
	styleUrls: ['./group-detail.component.css'],
})
export class GroupDetailComponent implements OnInit {

	// insertion point for declarations

	// the GroupDB of interest
	group: GroupDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private groupService: GroupService,
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
		this.getGroup()

		// observable for changes in structs
		this.groupService.GroupServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGroup()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getGroup(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo GroupPull returned")

				if (id != 0 && association == undefined) {
					this.group = frontRepo.Groups.get(id)
				} else {
					this.group = new (GroupDB)
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
			if (this.group.Gantt_Groups_reverse != undefined) {
				this.group.Gantt_GroupsDBID = new NullInt64
				this.group.Gantt_GroupsDBID.Int64 = this.group.Gantt_Groups_reverse.ID
				this.group.Gantt_GroupsDBID.Valid = true
				this.group.Gantt_GroupsDBID_Index.Valid = true
				this.group.Gantt_Groups_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.groupService.updateGroup(this.group)
				.subscribe(group => {
					this.groupService.GroupServiceChanged.next("update")

					console.log("group saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Gantt_Groups":
					this.group.Gantt_GroupsDBID = new NullInt64
					this.group.Gantt_GroupsDBID.Int64 = id
					this.group.Gantt_GroupsDBID.Valid = true
					this.group.Gantt_GroupsDBID_Index.Valid = true
					break
			}
			this.groupService.postGroup(this.group).subscribe(group => {

				this.groupService.GroupServiceChanged.next("post")

				this.group = {} // reset fields
				console.log("group added")
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
			ID: this.group.ID,
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
			console.log('The dialog was closed');
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.group.ID,
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
			console.log('The dialog was closed');
		});
	}
}
