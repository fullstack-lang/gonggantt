// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { BarDB } from '../bar-db'
import { BarService } from '../bar.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-bar-detail',
	templateUrl: './bar-detail.component.html',
	styleUrls: ['./bar-detail.component.css'],
})
export class BarDetailComponent implements OnInit {

	// insertion point for declarations

	// the BarDB of interest
	bar: BarDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private barService: BarService,
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
		this.getBar()

		// observable for changes in structs
		this.barService.BarServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getBar()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getBar(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				console.log("front repo BarPull returned")

				if (id != 0 && association == undefined) {
					this.bar = frontRepo.Bars.get(id)
				} else {
					this.bar = new (BarDB)
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
			if (this.bar.Lane_Bars_reverse != undefined) {
				this.bar.Lane_BarsDBID = new NullInt64
				this.bar.Lane_BarsDBID.Int64 = this.bar.Lane_Bars_reverse.ID
				this.bar.Lane_BarsDBID.Valid = true
				this.bar.Lane_Bars_reverse = undefined // very important, otherwise, circular JSON
			}
		}

		if (id != 0 && association == undefined) {

			this.barService.updateBar(this.bar)
				.subscribe(bar => {
					this.barService.BarServiceChanged.next("update")

					console.log("bar saved")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
				case "Lane_Bars":
					this.bar.Lane_BarsDBID = new NullInt64
					this.bar.Lane_BarsDBID.Int64 = id
					this.bar.Lane_BarsDBID.Valid = true
					break
			}
			this.barService.postBar(this.bar).subscribe(bar => {

				this.barService.BarServiceChanged.next("post")

				this.bar = {} // reset fields
				console.log("bar added")
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
			ID: this.bar.ID,
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
