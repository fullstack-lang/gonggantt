import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { BarService } from '../bar.service'
import { getBarUniqueID } from '../front-repo.service'
import { GanttService } from '../gantt.service'
import { getGanttUniqueID } from '../front-repo.service'
import { GroupService } from '../group.service'
import { getGroupUniqueID } from '../front-repo.service'
import { LaneService } from '../lane.service'
import { getLaneUniqueID } from '../front-repo.service'
import { MilestoneService } from '../milestone.service'
import { getMilestoneUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children?: GongNode[];
  type: GongNodeType;
  structName: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gonggantt-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo
  commitNb: number

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,

    // insertion point for per struct service declaration
    private barService: BarService,
    private ganttService: GanttService,
    private groupService: GroupService,
    private laneService: LaneService,
    private milestoneService: MilestoneService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.barService.BarServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.ganttService.GanttServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.groupService.GroupServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.laneService.LaneServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.milestoneService.MilestoneServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (this.treeControl.isExpanded(node)) {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = true
            } else {
              memoryOfExpandedNodes[node.uniqueIdPerStack] = false
            }
          }
        )
      }

      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Bar part of the mat tree
      */
      let barGongNodeStruct: GongNode = {
        name: "Bar",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Bar",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(barGongNodeStruct)

      this.frontRepo.Bars_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Bars_array.forEach(
        barDB => {
          let barGongNodeInstance: GongNode = {
            name: barDB.Name,
            type: GongNodeType.INSTANCE,
            id: barDB.ID,
            uniqueIdPerStack: getBarUniqueID(barDB.ID),
            structName: "Bar",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          barGongNodeStruct.children.push(barGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Gantt part of the mat tree
      */
      let ganttGongNodeStruct: GongNode = {
        name: "Gantt",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Gantt",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(ganttGongNodeStruct)

      this.frontRepo.Gantts_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Gantts_array.forEach(
        ganttDB => {
          let ganttGongNodeInstance: GongNode = {
            name: ganttDB.Name,
            type: GongNodeType.INSTANCE,
            id: ganttDB.ID,
            uniqueIdPerStack: getGanttUniqueID(ganttDB.ID),
            structName: "Gantt",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          ganttGongNodeStruct.children.push(ganttGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Lanes
          */
          let LanesGongNodeAssociation: GongNode = {
            name: "(Lane) Lanes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: ganttDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Gantt",
            associatedStructName: "Lane",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          ganttGongNodeInstance.children.push(LanesGongNodeAssociation)

          ganttDB.Lanes?.forEach(laneDB => {
            let laneNode: GongNode = {
              name: laneDB.Name,
              type: GongNodeType.INSTANCE,
              id: laneDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGanttUniqueID(ganttDB.ID)
                + 11 * getLaneUniqueID(laneDB.ID),
              structName: "Lane",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LanesGongNodeAssociation.children.push(laneNode)
          })

          /**
          * let append a node for the slide of pointer Milestones
          */
          let MilestonesGongNodeAssociation: GongNode = {
            name: "(Milestone) Milestones",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: ganttDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Gantt",
            associatedStructName: "Milestone",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          ganttGongNodeInstance.children.push(MilestonesGongNodeAssociation)

          ganttDB.Milestones?.forEach(milestoneDB => {
            let milestoneNode: GongNode = {
              name: milestoneDB.Name,
              type: GongNodeType.INSTANCE,
              id: milestoneDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGanttUniqueID(ganttDB.ID)
                + 11 * getMilestoneUniqueID(milestoneDB.ID),
              structName: "Milestone",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            MilestonesGongNodeAssociation.children.push(milestoneNode)
          })

          /**
          * let append a node for the slide of pointer Groups
          */
          let GroupsGongNodeAssociation: GongNode = {
            name: "(Group) Groups",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: ganttDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Gantt",
            associatedStructName: "Group",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          ganttGongNodeInstance.children.push(GroupsGongNodeAssociation)

          ganttDB.Groups?.forEach(groupDB => {
            let groupNode: GongNode = {
              name: groupDB.Name,
              type: GongNodeType.INSTANCE,
              id: groupDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGanttUniqueID(ganttDB.ID)
                + 11 * getGroupUniqueID(groupDB.ID),
              structName: "Group",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GroupsGongNodeAssociation.children.push(groupNode)
          })

        }
      )

      /**
      * fill up the Group part of the mat tree
      */
      let groupGongNodeStruct: GongNode = {
        name: "Group",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Group",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(groupGongNodeStruct)

      this.frontRepo.Groups_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Groups_array.forEach(
        groupDB => {
          let groupGongNodeInstance: GongNode = {
            name: groupDB.Name,
            type: GongNodeType.INSTANCE,
            id: groupDB.ID,
            uniqueIdPerStack: getGroupUniqueID(groupDB.ID),
            structName: "Group",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          groupGongNodeStruct.children.push(groupGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer GroupLanes
          */
          let GroupLanesGongNodeAssociation: GongNode = {
            name: "(Lane) GroupLanes",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: groupDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Group",
            associatedStructName: "Lane",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          groupGongNodeInstance.children.push(GroupLanesGongNodeAssociation)

          groupDB.GroupLanes?.forEach(laneDB => {
            let laneNode: GongNode = {
              name: laneDB.Name,
              type: GongNodeType.INSTANCE,
              id: laneDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGroupUniqueID(groupDB.ID)
                + 11 * getLaneUniqueID(laneDB.ID),
              structName: "Lane",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GroupLanesGongNodeAssociation.children.push(laneNode)
          })

        }
      )

      /**
      * fill up the Lane part of the mat tree
      */
      let laneGongNodeStruct: GongNode = {
        name: "Lane",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Lane",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(laneGongNodeStruct)

      this.frontRepo.Lanes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Lanes_array.forEach(
        laneDB => {
          let laneGongNodeInstance: GongNode = {
            name: laneDB.Name,
            type: GongNodeType.INSTANCE,
            id: laneDB.ID,
            uniqueIdPerStack: getLaneUniqueID(laneDB.ID),
            structName: "Lane",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          laneGongNodeStruct.children.push(laneGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Bars
          */
          let BarsGongNodeAssociation: GongNode = {
            name: "(Bar) Bars",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: laneDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Lane",
            associatedStructName: "Bar",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          laneGongNodeInstance.children.push(BarsGongNodeAssociation)

          laneDB.Bars?.forEach(barDB => {
            let barNode: GongNode = {
              name: barDB.Name,
              type: GongNodeType.INSTANCE,
              id: barDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getLaneUniqueID(laneDB.ID)
                + 11 * getBarUniqueID(barDB.ID),
              structName: "Bar",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            BarsGongNodeAssociation.children.push(barNode)
          })

        }
      )

      /**
      * fill up the Milestone part of the mat tree
      */
      let milestoneGongNodeStruct: GongNode = {
        name: "Milestone",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Milestone",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(milestoneGongNodeStruct)

      this.frontRepo.Milestones_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Milestones_array.forEach(
        milestoneDB => {
          let milestoneGongNodeInstance: GongNode = {
            name: milestoneDB.Name,
            type: GongNodeType.INSTANCE,
            id: milestoneDB.ID,
            uniqueIdPerStack: getMilestoneUniqueID(milestoneDB.ID),
            structName: "Milestone",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          milestoneGongNodeStruct.children.push(milestoneGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer DiamonfAndTextAnchors
          */
          let DiamonfAndTextAnchorsGongNodeAssociation: GongNode = {
            name: "(Lane) DiamonfAndTextAnchors",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: milestoneDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Milestone",
            associatedStructName: "Lane",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          milestoneGongNodeInstance.children.push(DiamonfAndTextAnchorsGongNodeAssociation)

          milestoneDB.DiamonfAndTextAnchors?.forEach(laneDB => {
            let laneNode: GongNode = {
              name: laneDB.Name,
              type: GongNodeType.INSTANCE,
              id: laneDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getMilestoneUniqueID(milestoneDB.ID)
                + 11 * getLaneUniqueID(laneDB.ID),
              structName: "Lane",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            DiamonfAndTextAnchorsGongNodeAssociation.children.push(laneNode)
          })

        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      if (this.treeControl.dataNodes != undefined) {
        this.treeControl.dataNodes.forEach(
          node => {
            if (memoryOfExpandedNodes[node.uniqueIdPerStack] != undefined) {
              if (memoryOfExpandedNodes[node.uniqueIdPerStack]) {
                this.treeControl.expand(node)
              }
            }
          }
        )
      }
    });

    // fetch the number of commits
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        table: [path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          table: [path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          presentation: [structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path) {
    this.router.navigate([{
      outlets: {
        editor: [path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet( node: GongFlatNode) {
    console.log("setEditorSpecialRouterOutlet " + node)
    this.router.navigate([{
      outlets: {
        editor: [node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName + "_" + node.name]
      }
    }]);
  }
}
