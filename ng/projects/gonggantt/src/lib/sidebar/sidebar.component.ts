import { Component, Input, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { ArrowService } from '../arrow.service'
import { getArrowUniqueID } from '../front-repo.service'
import { BarService } from '../bar.service'
import { getBarUniqueID } from '../front-repo.service'
import { GanttService } from '../gantt.service'
import { getGanttUniqueID } from '../front-repo.service'
import { GroupService } from '../group.service'
import { getGroupUniqueID } from '../front-repo.service'
import { LaneService } from '../lane.service'
import { getLaneUniqueID } from '../front-repo.service'
import { LaneUseService } from '../laneuse.service'
import { getLaneUseUniqueID } from '../front-repo.service'
import { MilestoneService } from '../milestone.service'
import { getMilestoneUniqueID } from '../front-repo.service'

import { RouteService } from '../route-service';

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
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
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
  associationField: string;
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
      associationField: node.associationField,
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
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  @Input() GONG__StackPath: string = ""

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private arrowService: ArrowService,
    private barService: BarService,
    private ganttService: GanttService,
    private groupService: GroupService,
    private laneService: LaneService,
    private laneuseService: LaneUseService,
    private milestoneService: MilestoneService,

    private routeService: RouteService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    console.log("Sidebar init: " + this.GONG__StackPath)

    // add the routes that will used by this side panel component and
    // by the component that are called from this component
    this.routeService.addDataPanelRoutes(this.GONG__StackPath)

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.arrowService.ArrowServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
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
    this.laneuseService.LaneUseServiceChanged.subscribe(
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
    this.frontRepoService.pull(this.GONG__StackPath).subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Arrow part of the mat tree
      */
      let arrowGongNodeStruct: GongNode = {
        name: "Arrow",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Arrow",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(arrowGongNodeStruct)

      this.frontRepo.Arrows_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Arrows_array.forEach(
        arrowDB => {
          let arrowGongNodeInstance: GongNode = {
            name: arrowDB.Name,
            type: GongNodeType.INSTANCE,
            id: arrowDB.ID,
            uniqueIdPerStack: getArrowUniqueID(arrowDB.ID),
            structName: "Arrow",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          arrowGongNodeStruct.children!.push(arrowGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association From
          */
          let FromGongNodeAssociation: GongNode = {
            name: "(Bar) From",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: arrowDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Arrow",
            associationField: "From",
            associatedStructName: "Bar",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          arrowGongNodeInstance.children!.push(FromGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation From
            */
          if (arrowDB.From != undefined) {
            let arrowGongNodeInstance_From: GongNode = {
              name: arrowDB.From.Name,
              type: GongNodeType.INSTANCE,
              id: arrowDB.From.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getArrowUniqueID(arrowDB.ID)
                + 5 * getBarUniqueID(arrowDB.From.ID),
              structName: "Bar",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            FromGongNodeAssociation.children.push(arrowGongNodeInstance_From)
          }

          /**
          * let append a node for the association To
          */
          let ToGongNodeAssociation: GongNode = {
            name: "(Bar) To",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: arrowDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Arrow",
            associationField: "To",
            associatedStructName: "Bar",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          arrowGongNodeInstance.children!.push(ToGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation To
            */
          if (arrowDB.To != undefined) {
            let arrowGongNodeInstance_To: GongNode = {
              name: arrowDB.To.Name,
              type: GongNodeType.INSTANCE,
              id: arrowDB.To.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getArrowUniqueID(arrowDB.ID)
                + 5 * getBarUniqueID(arrowDB.To.ID),
              structName: "Bar",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ToGongNodeAssociation.children.push(arrowGongNodeInstance_To)
          }

        }
      )

      /**
      * fill up the Bar part of the mat tree
      */
      let barGongNodeStruct: GongNode = {
        name: "Bar",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Bar",
        associationField: "",
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
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          barGongNodeStruct.children!.push(barGongNodeInstance)

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
        associationField: "",
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
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          ganttGongNodeStruct.children!.push(ganttGongNodeInstance)

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
            associationField: "Lanes",
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
              associationField: "",
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
            associationField: "Milestones",
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
              associationField: "",
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
            associationField: "Groups",
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
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            GroupsGongNodeAssociation.children.push(groupNode)
          })

          /**
          * let append a node for the slide of pointer Arrows
          */
          let ArrowsGongNodeAssociation: GongNode = {
            name: "(Arrow) Arrows",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: ganttDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Gantt",
            associationField: "Arrows",
            associatedStructName: "Arrow",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          ganttGongNodeInstance.children.push(ArrowsGongNodeAssociation)

          ganttDB.Arrows?.forEach(arrowDB => {
            let arrowNode: GongNode = {
              name: arrowDB.Name,
              type: GongNodeType.INSTANCE,
              id: arrowDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getGanttUniqueID(ganttDB.ID)
                + 11 * getArrowUniqueID(arrowDB.ID),
              structName: "Arrow",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ArrowsGongNodeAssociation.children.push(arrowNode)
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
        associationField: "",
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
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          groupGongNodeStruct.children!.push(groupGongNodeInstance)

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
            associationField: "GroupLanes",
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
              associationField: "",
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
        associationField: "",
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
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          laneGongNodeStruct.children!.push(laneGongNodeInstance)

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
            associationField: "Bars",
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
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            BarsGongNodeAssociation.children.push(barNode)
          })

        }
      )

      /**
      * fill up the LaneUse part of the mat tree
      */
      let laneuseGongNodeStruct: GongNode = {
        name: "LaneUse",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "LaneUse",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(laneuseGongNodeStruct)

      this.frontRepo.LaneUses_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.LaneUses_array.forEach(
        laneuseDB => {
          let laneuseGongNodeInstance: GongNode = {
            name: laneuseDB.Name,
            type: GongNodeType.INSTANCE,
            id: laneuseDB.ID,
            uniqueIdPerStack: getLaneUseUniqueID(laneuseDB.ID),
            structName: "LaneUse",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          laneuseGongNodeStruct.children!.push(laneuseGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Lane
          */
          let LaneGongNodeAssociation: GongNode = {
            name: "(Lane) Lane",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: laneuseDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "LaneUse",
            associationField: "Lane",
            associatedStructName: "Lane",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          laneuseGongNodeInstance.children!.push(LaneGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Lane
            */
          if (laneuseDB.Lane != undefined) {
            let laneuseGongNodeInstance_Lane: GongNode = {
              name: laneuseDB.Lane.Name,
              type: GongNodeType.INSTANCE,
              id: laneuseDB.Lane.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getLaneUseUniqueID(laneuseDB.ID)
                + 5 * getLaneUniqueID(laneuseDB.Lane.ID),
              structName: "Lane",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LaneGongNodeAssociation.children.push(laneuseGongNodeInstance_Lane)
          }

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
        associationField: "",
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
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          milestoneGongNodeStruct.children!.push(milestoneGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer LanesToDisplayMilestoneUse
          */
          let LanesToDisplayMilestoneUseGongNodeAssociation: GongNode = {
            name: "(LaneUse) LanesToDisplayMilestoneUse",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: milestoneDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Milestone",
            associationField: "LanesToDisplayMilestoneUse",
            associatedStructName: "LaneUse",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          milestoneGongNodeInstance.children.push(LanesToDisplayMilestoneUseGongNodeAssociation)

          milestoneDB.LanesToDisplayMilestoneUse?.forEach(laneuseDB => {
            let laneuseNode: GongNode = {
              name: laneuseDB.Name,
              type: GongNodeType.INSTANCE,
              id: laneuseDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getMilestoneUniqueID(milestoneDB.ID)
                + 11 * getLaneUseUniqueID(laneuseDB.ID),
              structName: "LaneUse",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            LanesToDisplayMilestoneUseGongNodeAssociation.children.push(laneuseNode)
          })

        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    })
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path
    this.router.navigate([{
      outlets: {
        outletName: [fullPath]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      let outletName = this.routeService.getTableOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
      let outletConf: any = {}
      outletConf[outletName] = [fullPath, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }

    if (type == GongNodeType.INSTANCE) {
      let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
      let fullPath = this.routeService.getPathRoot() + "-" + structName.toLowerCase() + "-detail"

      let outletConf: any = {}
      outletConf[outletName] = [fullPath, id, this.GONG__StackPath]

      this.router.navigate([{ outlets: outletConf }])
    }
  }

  setEditorRouterOutlet(path: string) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + path.toLowerCase()
    let outletConf : any = {}
    outletConf[outletName] = [fullPath, this.GONG__StackPath]
    this.router.navigate([{ outlets: outletConf }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    let outletName = this.routeService.getEditorOutlet(this.GONG__StackPath)
    let fullPath = this.routeService.getPathRoot() + "-" + node.associatedStructName.toLowerCase() + "-adder"
    this.router.navigate([{
      outlets: {
        outletName: [fullPath, node.id, node.structName, node.associationField, this.GONG__StackPath]
      }
    }]);
  }
}
