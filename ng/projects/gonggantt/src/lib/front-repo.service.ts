import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { ArrowDB } from './arrow-db'
import { ArrowService } from './arrow.service'

import { BarDB } from './bar-db'
import { BarService } from './bar.service'

import { GanttDB } from './gantt-db'
import { GanttService } from './gantt.service'

import { GroupDB } from './group-db'
import { GroupService } from './group.service'

import { LaneDB } from './lane-db'
import { LaneService } from './lane.service'

import { LaneUseDB } from './laneuse-db'
import { LaneUseService } from './laneuse.service'

import { MilestoneDB } from './milestone-db'
import { MilestoneService } from './milestone.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Arrows_array = new Array<ArrowDB>(); // array of repo instances
  Arrows = new Map<number, ArrowDB>(); // map of repo instances
  Arrows_batch = new Map<number, ArrowDB>(); // same but only in last GET (for finding repo instances to delete)
  Bars_array = new Array<BarDB>(); // array of repo instances
  Bars = new Map<number, BarDB>(); // map of repo instances
  Bars_batch = new Map<number, BarDB>(); // same but only in last GET (for finding repo instances to delete)
  Gantts_array = new Array<GanttDB>(); // array of repo instances
  Gantts = new Map<number, GanttDB>(); // map of repo instances
  Gantts_batch = new Map<number, GanttDB>(); // same but only in last GET (for finding repo instances to delete)
  Groups_array = new Array<GroupDB>(); // array of repo instances
  Groups = new Map<number, GroupDB>(); // map of repo instances
  Groups_batch = new Map<number, GroupDB>(); // same but only in last GET (for finding repo instances to delete)
  Lanes_array = new Array<LaneDB>(); // array of repo instances
  Lanes = new Map<number, LaneDB>(); // map of repo instances
  Lanes_batch = new Map<number, LaneDB>(); // same but only in last GET (for finding repo instances to delete)
  LaneUses_array = new Array<LaneUseDB>(); // array of repo instances
  LaneUses = new Map<number, LaneUseDB>(); // map of repo instances
  LaneUses_batch = new Map<number, LaneUseDB>(); // same but only in last GET (for finding repo instances to delete)
  Milestones_array = new Array<MilestoneDB>(); // array of repo instances
  Milestones = new Map<number, MilestoneDB>(); // map of repo instances
  Milestones_batch = new Map<number, MilestoneDB>(); // same but only in last GET (for finding repo instances to delete)
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"

  GONG__StackPath: string = ""
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  GONG__StackPath: string = ""

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  //
  // Store of all instances of the stack
  //
  frontRepo = new (FrontRepo)

  constructor(
    private http: HttpClient, // insertion point sub template 
    private arrowService: ArrowService,
    private barService: BarService,
    private ganttService: GanttService,
    private groupService: GroupService,
    private laneService: LaneService,
    private laneuseService: LaneUseService,
    private milestoneService: MilestoneService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<ArrowDB[]>,
    Observable<BarDB[]>,
    Observable<GanttDB[]>,
    Observable<GroupDB[]>,
    Observable<LaneDB[]>,
    Observable<LaneUseDB[]>,
    Observable<MilestoneDB[]>,
  ] = [ // insertion point sub template
      this.arrowService.getArrows(this.GONG__StackPath),
      this.barService.getBars(this.GONG__StackPath),
      this.ganttService.getGantts(this.GONG__StackPath),
      this.groupService.getGroups(this.GONG__StackPath),
      this.laneService.getLanes(this.GONG__StackPath),
      this.laneuseService.getLaneUses(this.GONG__StackPath),
      this.milestoneService.getMilestones(this.GONG__StackPath),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

    this.GONG__StackPath = GONG__StackPath

    this.observableFrontRepo = [ // insertion point sub template
      this.arrowService.getArrows(this.GONG__StackPath),
      this.barService.getBars(this.GONG__StackPath),
      this.ganttService.getGantts(this.GONG__StackPath),
      this.groupService.getGroups(this.GONG__StackPath),
      this.laneService.getLanes(this.GONG__StackPath),
      this.laneuseService.getLaneUses(this.GONG__StackPath),
      this.milestoneService.getMilestones(this.GONG__StackPath),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            arrows_,
            bars_,
            gantts_,
            groups_,
            lanes_,
            laneuses_,
            milestones_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var arrows: ArrowDB[]
            arrows = arrows_ as ArrowDB[]
            var bars: BarDB[]
            bars = bars_ as BarDB[]
            var gantts: GanttDB[]
            gantts = gantts_ as GanttDB[]
            var groups: GroupDB[]
            groups = groups_ as GroupDB[]
            var lanes: LaneDB[]
            lanes = lanes_ as LaneDB[]
            var laneuses: LaneUseDB[]
            laneuses = laneuses_ as LaneUseDB[]
            var milestones: MilestoneDB[]
            milestones = milestones_ as MilestoneDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.Arrows_array = arrows

            // clear the map that counts Arrow in the GET
            this.frontRepo.Arrows_batch.clear()

            arrows.forEach(
              arrow => {
                this.frontRepo.Arrows.set(arrow.ID, arrow)
                this.frontRepo.Arrows_batch.set(arrow.ID, arrow)
              }
            )

            // clear arrows that are absent from the batch
            this.frontRepo.Arrows.forEach(
              arrow => {
                if (this.frontRepo.Arrows_batch.get(arrow.ID) == undefined) {
                  this.frontRepo.Arrows.delete(arrow.ID)
                }
              }
            )

            // sort Arrows_array array
            this.frontRepo.Arrows_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Bars_array = bars

            // clear the map that counts Bar in the GET
            this.frontRepo.Bars_batch.clear()

            bars.forEach(
              bar => {
                this.frontRepo.Bars.set(bar.ID, bar)
                this.frontRepo.Bars_batch.set(bar.ID, bar)
              }
            )

            // clear bars that are absent from the batch
            this.frontRepo.Bars.forEach(
              bar => {
                if (this.frontRepo.Bars_batch.get(bar.ID) == undefined) {
                  this.frontRepo.Bars.delete(bar.ID)
                }
              }
            )

            // sort Bars_array array
            this.frontRepo.Bars_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Gantts_array = gantts

            // clear the map that counts Gantt in the GET
            this.frontRepo.Gantts_batch.clear()

            gantts.forEach(
              gantt => {
                this.frontRepo.Gantts.set(gantt.ID, gantt)
                this.frontRepo.Gantts_batch.set(gantt.ID, gantt)
              }
            )

            // clear gantts that are absent from the batch
            this.frontRepo.Gantts.forEach(
              gantt => {
                if (this.frontRepo.Gantts_batch.get(gantt.ID) == undefined) {
                  this.frontRepo.Gantts.delete(gantt.ID)
                }
              }
            )

            // sort Gantts_array array
            this.frontRepo.Gantts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Groups_array = groups

            // clear the map that counts Group in the GET
            this.frontRepo.Groups_batch.clear()

            groups.forEach(
              group => {
                this.frontRepo.Groups.set(group.ID, group)
                this.frontRepo.Groups_batch.set(group.ID, group)
              }
            )

            // clear groups that are absent from the batch
            this.frontRepo.Groups.forEach(
              group => {
                if (this.frontRepo.Groups_batch.get(group.ID) == undefined) {
                  this.frontRepo.Groups.delete(group.ID)
                }
              }
            )

            // sort Groups_array array
            this.frontRepo.Groups_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Lanes_array = lanes

            // clear the map that counts Lane in the GET
            this.frontRepo.Lanes_batch.clear()

            lanes.forEach(
              lane => {
                this.frontRepo.Lanes.set(lane.ID, lane)
                this.frontRepo.Lanes_batch.set(lane.ID, lane)
              }
            )

            // clear lanes that are absent from the batch
            this.frontRepo.Lanes.forEach(
              lane => {
                if (this.frontRepo.Lanes_batch.get(lane.ID) == undefined) {
                  this.frontRepo.Lanes.delete(lane.ID)
                }
              }
            )

            // sort Lanes_array array
            this.frontRepo.Lanes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.LaneUses_array = laneuses

            // clear the map that counts LaneUse in the GET
            this.frontRepo.LaneUses_batch.clear()

            laneuses.forEach(
              laneuse => {
                this.frontRepo.LaneUses.set(laneuse.ID, laneuse)
                this.frontRepo.LaneUses_batch.set(laneuse.ID, laneuse)
              }
            )

            // clear laneuses that are absent from the batch
            this.frontRepo.LaneUses.forEach(
              laneuse => {
                if (this.frontRepo.LaneUses_batch.get(laneuse.ID) == undefined) {
                  this.frontRepo.LaneUses.delete(laneuse.ID)
                }
              }
            )

            // sort LaneUses_array array
            this.frontRepo.LaneUses_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Milestones_array = milestones

            // clear the map that counts Milestone in the GET
            this.frontRepo.Milestones_batch.clear()

            milestones.forEach(
              milestone => {
                this.frontRepo.Milestones.set(milestone.ID, milestone)
                this.frontRepo.Milestones_batch.set(milestone.ID, milestone)
              }
            )

            // clear milestones that are absent from the batch
            this.frontRepo.Milestones.forEach(
              milestone => {
                if (this.frontRepo.Milestones_batch.get(milestone.ID) == undefined) {
                  this.frontRepo.Milestones.delete(milestone.ID)
                }
              }
            )

            // sort Milestones_array array
            this.frontRepo.Milestones_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            arrows.forEach(
              arrow => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field From redeeming
                {
                  let _bar = this.frontRepo.Bars.get(arrow.FromID.Int64)
                  if (_bar) {
                    arrow.From = _bar
                  }
                }
                // insertion point for pointer field To redeeming
                {
                  let _bar = this.frontRepo.Bars.get(arrow.ToID.Int64)
                  if (_bar) {
                    arrow.To = _bar
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Arrows redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(arrow.Gantt_ArrowsDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Arrows == undefined) {
                      _gantt.Arrows = new Array<ArrowDB>()
                    }
                    _gantt.Arrows.push(arrow)
                    if (arrow.Gantt_Arrows_reverse == undefined) {
                      arrow.Gantt_Arrows_reverse = _gantt
                    }
                  }
                }
              }
            )
            bars.forEach(
              bar => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Lane.Bars redeeming
                {
                  let _lane = this.frontRepo.Lanes.get(bar.Lane_BarsDBID.Int64)
                  if (_lane) {
                    if (_lane.Bars == undefined) {
                      _lane.Bars = new Array<BarDB>()
                    }
                    _lane.Bars.push(bar)
                    if (bar.Lane_Bars_reverse == undefined) {
                      bar.Lane_Bars_reverse = _lane
                    }
                  }
                }
              }
            )
            gantts.forEach(
              gantt => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            groups.forEach(
              group => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Groups redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(group.Gantt_GroupsDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Groups == undefined) {
                      _gantt.Groups = new Array<GroupDB>()
                    }
                    _gantt.Groups.push(group)
                    if (group.Gantt_Groups_reverse == undefined) {
                      group.Gantt_Groups_reverse = _gantt
                    }
                  }
                }
              }
            )
            lanes.forEach(
              lane => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Lanes redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(lane.Gantt_LanesDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Lanes == undefined) {
                      _gantt.Lanes = new Array<LaneDB>()
                    }
                    _gantt.Lanes.push(lane)
                    if (lane.Gantt_Lanes_reverse == undefined) {
                      lane.Gantt_Lanes_reverse = _gantt
                    }
                  }
                }
                // insertion point for slice of pointer field Group.GroupLanes redeeming
                {
                  let _group = this.frontRepo.Groups.get(lane.Group_GroupLanesDBID.Int64)
                  if (_group) {
                    if (_group.GroupLanes == undefined) {
                      _group.GroupLanes = new Array<LaneDB>()
                    }
                    _group.GroupLanes.push(lane)
                    if (lane.Group_GroupLanes_reverse == undefined) {
                      lane.Group_GroupLanes_reverse = _group
                    }
                  }
                }
              }
            )
            laneuses.forEach(
              laneuse => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Lane redeeming
                {
                  let _lane = this.frontRepo.Lanes.get(laneuse.LaneID.Int64)
                  if (_lane) {
                    laneuse.Lane = _lane
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Milestone.LanesToDisplayMilestoneUse redeeming
                {
                  let _milestone = this.frontRepo.Milestones.get(laneuse.Milestone_LanesToDisplayMilestoneUseDBID.Int64)
                  if (_milestone) {
                    if (_milestone.LanesToDisplayMilestoneUse == undefined) {
                      _milestone.LanesToDisplayMilestoneUse = new Array<LaneUseDB>()
                    }
                    _milestone.LanesToDisplayMilestoneUse.push(laneuse)
                    if (laneuse.Milestone_LanesToDisplayMilestoneUse_reverse == undefined) {
                      laneuse.Milestone_LanesToDisplayMilestoneUse_reverse = _milestone
                    }
                  }
                }
              }
            )
            milestones.forEach(
              milestone => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Milestones redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(milestone.Gantt_MilestonesDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Milestones == undefined) {
                      _gantt.Milestones = new Array<MilestoneDB>()
                    }
                    _gantt.Milestones.push(milestone)
                    if (milestone.Gantt_Milestones_reverse == undefined) {
                      milestone.Gantt_Milestones_reverse = _gantt
                    }
                  }
                }
              }
            )

            // 
            // Third Step: sort arrays (slices in go) according to their index
            // insertion point sub template for redeem 
            arrows.forEach(
              arrow => {
                // insertion point for sorting
              }
            )
            bars.forEach(
              bar => {
                // insertion point for sorting
              }
            )
            gantts.forEach(
              gantt => {
                // insertion point for sorting
                gantt.Lanes?.sort((t1, t2) => {
                  if (t1.Gantt_LanesDBID_Index.Int64 > t2.Gantt_LanesDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Gantt_LanesDBID_Index.Int64 < t2.Gantt_LanesDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                gantt.Milestones?.sort((t1, t2) => {
                  if (t1.Gantt_MilestonesDBID_Index.Int64 > t2.Gantt_MilestonesDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Gantt_MilestonesDBID_Index.Int64 < t2.Gantt_MilestonesDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                gantt.Groups?.sort((t1, t2) => {
                  if (t1.Gantt_GroupsDBID_Index.Int64 > t2.Gantt_GroupsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Gantt_GroupsDBID_Index.Int64 < t2.Gantt_GroupsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

                gantt.Arrows?.sort((t1, t2) => {
                  if (t1.Gantt_ArrowsDBID_Index.Int64 > t2.Gantt_ArrowsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Gantt_ArrowsDBID_Index.Int64 < t2.Gantt_ArrowsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            groups.forEach(
              group => {
                // insertion point for sorting
                group.GroupLanes?.sort((t1, t2) => {
                  if (t1.Group_GroupLanesDBID_Index.Int64 > t2.Group_GroupLanesDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Group_GroupLanesDBID_Index.Int64 < t2.Group_GroupLanesDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            lanes.forEach(
              lane => {
                // insertion point for sorting
                lane.Bars?.sort((t1, t2) => {
                  if (t1.Lane_BarsDBID_Index.Int64 > t2.Lane_BarsDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Lane_BarsDBID_Index.Int64 < t2.Lane_BarsDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )
            laneuses.forEach(
              laneuse => {
                // insertion point for sorting
              }
            )
            milestones.forEach(
              milestone => {
                // insertion point for sorting
                milestone.LanesToDisplayMilestoneUse?.sort((t1, t2) => {
                  if (t1.Milestone_LanesToDisplayMilestoneUseDBID_Index.Int64 > t2.Milestone_LanesToDisplayMilestoneUseDBID_Index.Int64) {
                    return 1;
                  }
                  if (t1.Milestone_LanesToDisplayMilestoneUseDBID_Index.Int64 < t2.Milestone_LanesToDisplayMilestoneUseDBID_Index.Int64) {
                    return -1;
                  }
                  return 0;
                })

              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // ArrowPull performs a GET on Arrow of the stack and redeem association pointers 
  ArrowPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.arrowService.getArrows(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            arrows,
          ]) => {
            // init the array
            this.frontRepo.Arrows_array = arrows

            // clear the map that counts Arrow in the GET
            this.frontRepo.Arrows_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            arrows.forEach(
              arrow => {
                this.frontRepo.Arrows.set(arrow.ID, arrow)
                this.frontRepo.Arrows_batch.set(arrow.ID, arrow)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field From redeeming
                {
                  let _bar = this.frontRepo.Bars.get(arrow.FromID.Int64)
                  if (_bar) {
                    arrow.From = _bar
                  }
                }
                // insertion point for pointer field To redeeming
                {
                  let _bar = this.frontRepo.Bars.get(arrow.ToID.Int64)
                  if (_bar) {
                    arrow.To = _bar
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Arrows redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(arrow.Gantt_ArrowsDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Arrows == undefined) {
                      _gantt.Arrows = new Array<ArrowDB>()
                    }
                    _gantt.Arrows.push(arrow)
                    if (arrow.Gantt_Arrows_reverse == undefined) {
                      arrow.Gantt_Arrows_reverse = _gantt
                    }
                  }
                }
              }
            )

            // clear arrows that are absent from the GET
            this.frontRepo.Arrows.forEach(
              arrow => {
                if (this.frontRepo.Arrows_batch.get(arrow.ID) == undefined) {
                  this.frontRepo.Arrows.delete(arrow.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // BarPull performs a GET on Bar of the stack and redeem association pointers 
  BarPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.barService.getBars(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            bars,
          ]) => {
            // init the array
            this.frontRepo.Bars_array = bars

            // clear the map that counts Bar in the GET
            this.frontRepo.Bars_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bars.forEach(
              bar => {
                this.frontRepo.Bars.set(bar.ID, bar)
                this.frontRepo.Bars_batch.set(bar.ID, bar)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Lane.Bars redeeming
                {
                  let _lane = this.frontRepo.Lanes.get(bar.Lane_BarsDBID.Int64)
                  if (_lane) {
                    if (_lane.Bars == undefined) {
                      _lane.Bars = new Array<BarDB>()
                    }
                    _lane.Bars.push(bar)
                    if (bar.Lane_Bars_reverse == undefined) {
                      bar.Lane_Bars_reverse = _lane
                    }
                  }
                }
              }
            )

            // clear bars that are absent from the GET
            this.frontRepo.Bars.forEach(
              bar => {
                if (this.frontRepo.Bars_batch.get(bar.ID) == undefined) {
                  this.frontRepo.Bars.delete(bar.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // GanttPull performs a GET on Gantt of the stack and redeem association pointers 
  GanttPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.ganttService.getGantts(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            gantts,
          ]) => {
            // init the array
            this.frontRepo.Gantts_array = gantts

            // clear the map that counts Gantt in the GET
            this.frontRepo.Gantts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gantts.forEach(
              gantt => {
                this.frontRepo.Gantts.set(gantt.ID, gantt)
                this.frontRepo.Gantts_batch.set(gantt.ID, gantt)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear gantts that are absent from the GET
            this.frontRepo.Gantts.forEach(
              gantt => {
                if (this.frontRepo.Gantts_batch.get(gantt.ID) == undefined) {
                  this.frontRepo.Gantts.delete(gantt.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // GroupPull performs a GET on Group of the stack and redeem association pointers 
  GroupPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.groupService.getGroups(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            groups,
          ]) => {
            // init the array
            this.frontRepo.Groups_array = groups

            // clear the map that counts Group in the GET
            this.frontRepo.Groups_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            groups.forEach(
              group => {
                this.frontRepo.Groups.set(group.ID, group)
                this.frontRepo.Groups_batch.set(group.ID, group)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Groups redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(group.Gantt_GroupsDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Groups == undefined) {
                      _gantt.Groups = new Array<GroupDB>()
                    }
                    _gantt.Groups.push(group)
                    if (group.Gantt_Groups_reverse == undefined) {
                      group.Gantt_Groups_reverse = _gantt
                    }
                  }
                }
              }
            )

            // clear groups that are absent from the GET
            this.frontRepo.Groups.forEach(
              group => {
                if (this.frontRepo.Groups_batch.get(group.ID) == undefined) {
                  this.frontRepo.Groups.delete(group.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LanePull performs a GET on Lane of the stack and redeem association pointers 
  LanePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.laneService.getLanes(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            lanes,
          ]) => {
            // init the array
            this.frontRepo.Lanes_array = lanes

            // clear the map that counts Lane in the GET
            this.frontRepo.Lanes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            lanes.forEach(
              lane => {
                this.frontRepo.Lanes.set(lane.ID, lane)
                this.frontRepo.Lanes_batch.set(lane.ID, lane)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Lanes redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(lane.Gantt_LanesDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Lanes == undefined) {
                      _gantt.Lanes = new Array<LaneDB>()
                    }
                    _gantt.Lanes.push(lane)
                    if (lane.Gantt_Lanes_reverse == undefined) {
                      lane.Gantt_Lanes_reverse = _gantt
                    }
                  }
                }
                // insertion point for slice of pointer field Group.GroupLanes redeeming
                {
                  let _group = this.frontRepo.Groups.get(lane.Group_GroupLanesDBID.Int64)
                  if (_group) {
                    if (_group.GroupLanes == undefined) {
                      _group.GroupLanes = new Array<LaneDB>()
                    }
                    _group.GroupLanes.push(lane)
                    if (lane.Group_GroupLanes_reverse == undefined) {
                      lane.Group_GroupLanes_reverse = _group
                    }
                  }
                }
              }
            )

            // clear lanes that are absent from the GET
            this.frontRepo.Lanes.forEach(
              lane => {
                if (this.frontRepo.Lanes_batch.get(lane.ID) == undefined) {
                  this.frontRepo.Lanes.delete(lane.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LaneUsePull performs a GET on LaneUse of the stack and redeem association pointers 
  LaneUsePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.laneuseService.getLaneUses(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            laneuses,
          ]) => {
            // init the array
            this.frontRepo.LaneUses_array = laneuses

            // clear the map that counts LaneUse in the GET
            this.frontRepo.LaneUses_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            laneuses.forEach(
              laneuse => {
                this.frontRepo.LaneUses.set(laneuse.ID, laneuse)
                this.frontRepo.LaneUses_batch.set(laneuse.ID, laneuse)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Lane redeeming
                {
                  let _lane = this.frontRepo.Lanes.get(laneuse.LaneID.Int64)
                  if (_lane) {
                    laneuse.Lane = _lane
                  }
                }

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Milestone.LanesToDisplayMilestoneUse redeeming
                {
                  let _milestone = this.frontRepo.Milestones.get(laneuse.Milestone_LanesToDisplayMilestoneUseDBID.Int64)
                  if (_milestone) {
                    if (_milestone.LanesToDisplayMilestoneUse == undefined) {
                      _milestone.LanesToDisplayMilestoneUse = new Array<LaneUseDB>()
                    }
                    _milestone.LanesToDisplayMilestoneUse.push(laneuse)
                    if (laneuse.Milestone_LanesToDisplayMilestoneUse_reverse == undefined) {
                      laneuse.Milestone_LanesToDisplayMilestoneUse_reverse = _milestone
                    }
                  }
                }
              }
            )

            // clear laneuses that are absent from the GET
            this.frontRepo.LaneUses.forEach(
              laneuse => {
                if (this.frontRepo.LaneUses_batch.get(laneuse.ID) == undefined) {
                  this.frontRepo.LaneUses.delete(laneuse.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // MilestonePull performs a GET on Milestone of the stack and redeem association pointers 
  MilestonePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.milestoneService.getMilestones(this.GONG__StackPath)
        ]).subscribe(
          ([ // insertion point sub template 
            milestones,
          ]) => {
            // init the array
            this.frontRepo.Milestones_array = milestones

            // clear the map that counts Milestone in the GET
            this.frontRepo.Milestones_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            milestones.forEach(
              milestone => {
                this.frontRepo.Milestones.set(milestone.ID, milestone)
                this.frontRepo.Milestones_batch.set(milestone.ID, milestone)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Gantt.Milestones redeeming
                {
                  let _gantt = this.frontRepo.Gantts.get(milestone.Gantt_MilestonesDBID.Int64)
                  if (_gantt) {
                    if (_gantt.Milestones == undefined) {
                      _gantt.Milestones = new Array<MilestoneDB>()
                    }
                    _gantt.Milestones.push(milestone)
                    if (milestone.Gantt_Milestones_reverse == undefined) {
                      milestone.Gantt_Milestones_reverse = _gantt
                    }
                  }
                }
              }
            )

            // clear milestones that are absent from the GET
            this.frontRepo.Milestones.forEach(
              milestone => {
                if (this.frontRepo.Milestones_batch.get(milestone.ID) == undefined) {
                  this.frontRepo.Milestones.delete(milestone.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getArrowUniqueID(id: number): number {
  return 31 * id
}
export function getBarUniqueID(id: number): number {
  return 37 * id
}
export function getGanttUniqueID(id: number): number {
  return 41 * id
}
export function getGroupUniqueID(id: number): number {
  return 43 * id
}
export function getLaneUniqueID(id: number): number {
  return 47 * id
}
export function getLaneUseUniqueID(id: number): number {
  return 53 * id
}
export function getMilestoneUniqueID(id: number): number {
  return 59 * id
}
