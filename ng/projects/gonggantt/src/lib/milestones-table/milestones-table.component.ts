// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { MilestoneDB } from '../milestone-db'
import { MilestoneService } from '../milestone.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-milestonestable',
  templateUrl: './milestones-table.component.html',
  styleUrls: ['./milestones-table.component.css'],
})
export class MilestonesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Milestone instances
  selection: SelectionModel<MilestoneDB> = new (SelectionModel)
  initialSelection = new Array<MilestoneDB>()

  // the data source for the table
  milestones: MilestoneDB[] = []
  matTableDataSource: MatTableDataSource<MilestoneDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.milestones
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (milestoneDB: MilestoneDB, property: string) => {
      switch (property) {
        case 'ID':
          return milestoneDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return milestoneDB.Name;

        case 'Date':
          return milestoneDB.Date.getDate();

        case 'Gantt_Milestones':
          return this.frontRepo.Gantts.get(milestoneDB.Gantt_MilestonesDBID.Int64)!.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (milestoneDB: MilestoneDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the milestoneDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += milestoneDB.Name.toLowerCase()
      if (milestoneDB.Gantt_MilestonesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Gantts.get(milestoneDB.Gantt_MilestonesDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private milestoneService: MilestoneService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of milestone instances
    public dialogRef: MatDialogRef<MilestonesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.milestoneService.MilestoneServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getMilestones()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Date",
        "Gantt_Milestones",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Date",
        "Gantt_Milestones",
      ]
      this.selection = new SelectionModel<MilestoneDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getMilestones()
    this.matTableDataSource = new MatTableDataSource(this.milestones)
  }

  getMilestones(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.milestones = this.frontRepo.Milestones_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let milestone of this.milestones) {
            let ID = this.dialogData.ID
            let revPointer = milestone[this.dialogData.ReversePointer as keyof MilestoneDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(milestone)
            }
            this.selection = new SelectionModel<MilestoneDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, MilestoneDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as MilestoneDB[]
          for (let associationInstance of sourceField) {
            let milestone = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as MilestoneDB
            this.initialSelection.push(milestone)
          }

          this.selection = new SelectionModel<MilestoneDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.milestones
      }
    )
  }

  // newMilestone initiate a new milestone
  // create a new Milestone objet
  newMilestone() {
  }

  deleteMilestone(milestoneID: number, milestone: MilestoneDB) {
    // list of milestones is truncated of milestone before the delete
    this.milestones = this.milestones.filter(h => h !== milestone);

    this.milestoneService.deleteMilestone(milestoneID).subscribe(
      milestone => {
        this.milestoneService.MilestoneServiceChanged.next("delete")
      }
    );
  }

  editMilestone(milestoneID: number, milestone: MilestoneDB) {

  }

  // display milestone in router
  displayMilestoneInRouter(milestoneID: number) {
    this.router.navigate(["github_com_fullstack_lang_gonggantt_go-" + "milestone-display", milestoneID])
  }

  // set editor outlet
  setEditorRouterOutlet(milestoneID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gonggantt_go_editor: ["github_com_fullstack_lang_gonggantt_go-" + "milestone-detail", milestoneID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(milestoneID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gonggantt_go_presentation: ["github_com_fullstack_lang_gonggantt_go-" + "milestone-presentation", milestoneID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.milestones.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.milestones.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<MilestoneDB>()

      // reset all initial selection of milestone that belong to milestone
      for (let milestone of this.initialSelection) {
        let index = milestone[this.dialogData.ReversePointer as keyof MilestoneDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(milestone)

      }

      // from selection, set milestone that belong to milestone
      for (let milestone of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = milestone[this.dialogData.ReversePointer as keyof MilestoneDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(milestone)
      }


      // update all milestone (only update selection & initial selection)
      for (let milestone of toUpdate) {
        this.milestoneService.updateMilestone(milestone)
          .subscribe(milestone => {
            this.milestoneService.MilestoneServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, MilestoneDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedMilestone = new Set<number>()
      for (let milestone of this.initialSelection) {
        if (this.selection.selected.includes(milestone)) {
          // console.log("milestone " + milestone.Name + " is still selected")
        } else {
          console.log("milestone " + milestone.Name + " has been unselected")
          unselectedMilestone.add(milestone.ID)
          console.log("is unselected " + unselectedMilestone.has(milestone.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let milestone = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as MilestoneDB
      if (unselectedMilestone.has(milestone.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<MilestoneDB>) = new Array<MilestoneDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          milestone => {
            if (!this.initialSelection.includes(milestone)) {
              // console.log("milestone " + milestone.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + milestone.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = milestone.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = milestone.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("milestone " + milestone.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<MilestoneDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}