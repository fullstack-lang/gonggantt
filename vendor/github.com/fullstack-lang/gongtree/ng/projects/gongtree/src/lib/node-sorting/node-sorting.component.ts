// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { NodeDB } from '../node-db'
import { NodeService } from '../node.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-node-sorting',
  templateUrl: './node-sorting.component.html',
  styleUrls: ['./node-sorting.component.css']
})
export class NodeSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Node instances that are in the association
  associatedNodes = new Array<NodeDB>();

  constructor(
    private nodeService: NodeService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of node instances
    public dialogRef: MatDialogRef<NodeSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getNodes()
  }

  getNodes(): void {
    this.frontRepoService.pull(this.dialogData.GONG__StackPath).subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let node of this.frontRepo.Nodes_array) {
          let ID = this.dialogData.ID
          let revPointerID = node[this.dialogData.ReversePointer as keyof NodeDB] as unknown as NullInt64
          let revPointerID_Index = node[this.dialogData.ReversePointer + "_Index" as keyof NodeDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedNodes.push(node)
          }
        }

        // sort associated node according to order
        this.associatedNodes.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedNodes, event.previousIndex, event.currentIndex);

    // set the order of Node instances
    let index = 0

    for (let node of this.associatedNodes) {
      let revPointerID_Index = node[this.dialogData.ReversePointer + "_Index" as keyof NodeDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedNodes.forEach(
      node => {
        this.nodeService.updateNode(node, this.dialogData.GONG__StackPath)
          .subscribe(node => {
            this.nodeService.NodeServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer + ' done');
  }
}