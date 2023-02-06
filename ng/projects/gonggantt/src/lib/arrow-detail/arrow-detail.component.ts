// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { UntypedFormControl } from '@angular/forms';

import { ArrowDB } from '../arrow-db'
import { ArrowService } from '../arrow.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { GanttDB } from '../gantt-db'

import { Router, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// ArrowDetailComponent is initilizaed from different routes
// ArrowDetailComponentState detail different cases 
enum ArrowDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
	CREATE_INSTANCE_WITH_ASSOCIATION_Gantt_Arrows_SET,
}

@Component({
	selector: 'app-arrow-detail',
	templateUrl: './arrow-detail.component.html',
	styleUrls: ['./arrow-detail.component.css'],
})
export class ArrowDetailComponent implements OnInit {

	// insertion point for declarations

	// the ArrowDB of interest
	arrow: ArrowDB = new ArrowDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: ArrowDetailComponentState = ArrowDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private arrowService: ArrowService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private activatedRoute: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.activatedRoute.params.subscribe(params => {
			this.onChangedActivatedRoute()
		});
	}
	onChangedActivatedRoute(): void {

		// compute state
		this.id = +this.activatedRoute.snapshot.paramMap.get('id')!;
		this.originStruct = this.activatedRoute.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.activatedRoute.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.activatedRoute.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = ArrowDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = ArrowDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					case "Arrows":
						// console.log("Arrow" + " is instanciated with back pointer to instance " + this.id + " Gantt association Arrows")
						this.state = ArrowDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Gantt_Arrows_SET
						break;
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getArrow()

		// observable for changes in structs
		this.arrowService.ArrowServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getArrow()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getArrow(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case ArrowDetailComponentState.CREATE_INSTANCE:
						this.arrow = new (ArrowDB)
						break;
					case ArrowDetailComponentState.UPDATE_INSTANCE:
						let arrow = frontRepo.Arrows.get(this.id)
						console.assert(arrow != undefined, "missing arrow with id:" + this.id)
						this.arrow = arrow!
						break;
					// insertion point for init of association field
					case ArrowDetailComponentState.CREATE_INSTANCE_WITH_ASSOCIATION_Gantt_Arrows_SET:
						this.arrow = new (ArrowDB)
						this.arrow.Gantt_Arrows_reverse = frontRepo.Gantts.get(this.id)!
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
		if (this.arrow.FromID == undefined) {
			this.arrow.FromID = new NullInt64
		}
		if (this.arrow.From != undefined) {
			this.arrow.FromID.Int64 = this.arrow.From.ID
			this.arrow.FromID.Valid = true
		} else {
			this.arrow.FromID.Int64 = 0
			this.arrow.FromID.Valid = true
		}
		if (this.arrow.ToID == undefined) {
			this.arrow.ToID = new NullInt64
		}
		if (this.arrow.To != undefined) {
			this.arrow.ToID.Int64 = this.arrow.To.ID
			this.arrow.ToID.Valid = true
		} else {
			this.arrow.ToID.Int64 = 0
			this.arrow.ToID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers
		if (this.arrow.Gantt_Arrows_reverse != undefined) {
			if (this.arrow.Gantt_ArrowsDBID == undefined) {
				this.arrow.Gantt_ArrowsDBID = new NullInt64
			}
			this.arrow.Gantt_ArrowsDBID.Int64 = this.arrow.Gantt_Arrows_reverse.ID
			this.arrow.Gantt_ArrowsDBID.Valid = true
			if (this.arrow.Gantt_ArrowsDBID_Index == undefined) {
				this.arrow.Gantt_ArrowsDBID_Index = new NullInt64
			}
			this.arrow.Gantt_ArrowsDBID_Index.Valid = true
			this.arrow.Gantt_Arrows_reverse = new GanttDB // very important, otherwise, circular JSON
		}

		switch (this.state) {
			case ArrowDetailComponentState.UPDATE_INSTANCE:
				this.arrowService.updateArrow(this.arrow)
					.subscribe(arrow => {
						this.arrowService.ArrowServiceChanged.next("update")
					});
				break;
			default:
				this.arrowService.postArrow(this.arrow).subscribe(arrow => {
					this.arrowService.ArrowServiceChanged.next("post")
					this.arrow = new (ArrowDB) // reset fields
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

			dialogData.ID = this.arrow.ID!
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
			dialogData.ID = this.arrow.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Arrow"
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
			ID: this.arrow.ID,
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
		if (this.arrow.Name == "") {
			this.arrow.Name = event.value.Name
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
