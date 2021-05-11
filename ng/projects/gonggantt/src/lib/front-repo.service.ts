import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest } from 'rxjs';

// insertion point sub template for services imports 
import { BarDB } from './bar-db'
import { BarService } from './bar.service'

import { GanttDB } from './gantt-db'
import { GanttService } from './gantt.service'

import { GroupDB } from './group-db'
import { GroupService } from './group.service'

import { LaneDB } from './lane-db'
import { LaneService } from './lane.service'

import { MilestoneDB } from './milestone-db'
import { MilestoneService } from './milestone.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
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
  Milestones_array = new Array<MilestoneDB>(); // array of repo instances
  Milestones = new Map<number, MilestoneDB>(); // map of repo instances
  Milestones_batch = new Map<number, MilestoneDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// define the type of nullable Int64 in order to support back pointers IDs
export class NullInt64 {
    Int64: number
    Valid: boolean
}

// define the interface for information that is forwarded from the calling instance to 
// the select table
export interface DialogData {
  ID: number; // ID of the calling instance
  ReversePointer: string; // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean; // if true, this is for ordering items
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private barService: BarService,
    private ganttService: GanttService,
    private groupService: GroupService,
    private laneService: LaneService,
    private milestoneService: MilestoneService,
  ) { }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<BarDB[]>,
    Observable<GanttDB[]>,
    Observable<GroupDB[]>,
    Observable<LaneDB[]>,
    Observable<MilestoneDB[]>,
  ] = [ // insertion point sub template 
      this.barService.getBars(),
      this.ganttService.getGantts(),
      this.groupService.getGroups(),
      this.laneService.getLanes(),
      this.milestoneService.getMilestones(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            bars_,
            gantts_,
            groups_,
            lanes_,
            milestones_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var bars: BarDB[]
            bars = bars_
            var gantts: GanttDB[]
            gantts = gantts_
            var groups: GroupDB[]
            groups = groups_
            var lanes: LaneDB[]
            lanes = lanes_
            var milestones: MilestoneDB[]
            milestones = milestones_

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Bars_array = bars

            // clear the map that counts Bar in the GET
            FrontRepoSingloton.Bars_batch.clear()
            
            bars.forEach(
              bar => {
                FrontRepoSingloton.Bars.set(bar.ID, bar)
                FrontRepoSingloton.Bars_batch.set(bar.ID, bar)
              }
            )
            
            // clear bars that are absent from the batch
            FrontRepoSingloton.Bars.forEach(
              bar => {
                if (FrontRepoSingloton.Bars_batch.get(bar.ID) == undefined) {
                  FrontRepoSingloton.Bars.delete(bar.ID)
                }
              }
            )
            
            // sort Bars_array array
            FrontRepoSingloton.Bars_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });
            
            // init the array
            FrontRepoSingloton.Gantts_array = gantts

            // clear the map that counts Gantt in the GET
            FrontRepoSingloton.Gantts_batch.clear()
            
            gantts.forEach(
              gantt => {
                FrontRepoSingloton.Gantts.set(gantt.ID, gantt)
                FrontRepoSingloton.Gantts_batch.set(gantt.ID, gantt)
              }
            )
            
            // clear gantts that are absent from the batch
            FrontRepoSingloton.Gantts.forEach(
              gantt => {
                if (FrontRepoSingloton.Gantts_batch.get(gantt.ID) == undefined) {
                  FrontRepoSingloton.Gantts.delete(gantt.ID)
                }
              }
            )
            
            // sort Gantts_array array
            FrontRepoSingloton.Gantts_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });
            
            // init the array
            FrontRepoSingloton.Groups_array = groups

            // clear the map that counts Group in the GET
            FrontRepoSingloton.Groups_batch.clear()
            
            groups.forEach(
              group => {
                FrontRepoSingloton.Groups.set(group.ID, group)
                FrontRepoSingloton.Groups_batch.set(group.ID, group)
              }
            )
            
            // clear groups that are absent from the batch
            FrontRepoSingloton.Groups.forEach(
              group => {
                if (FrontRepoSingloton.Groups_batch.get(group.ID) == undefined) {
                  FrontRepoSingloton.Groups.delete(group.ID)
                }
              }
            )
            
            // sort Groups_array array
            FrontRepoSingloton.Groups_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });
            
            // init the array
            FrontRepoSingloton.Lanes_array = lanes

            // clear the map that counts Lane in the GET
            FrontRepoSingloton.Lanes_batch.clear()
            
            lanes.forEach(
              lane => {
                FrontRepoSingloton.Lanes.set(lane.ID, lane)
                FrontRepoSingloton.Lanes_batch.set(lane.ID, lane)
              }
            )
            
            // clear lanes that are absent from the batch
            FrontRepoSingloton.Lanes.forEach(
              lane => {
                if (FrontRepoSingloton.Lanes_batch.get(lane.ID) == undefined) {
                  FrontRepoSingloton.Lanes.delete(lane.ID)
                }
              }
            )
            
            // sort Lanes_array array
            FrontRepoSingloton.Lanes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });
            
            // init the array
            FrontRepoSingloton.Milestones_array = milestones

            // clear the map that counts Milestone in the GET
            FrontRepoSingloton.Milestones_batch.clear()
            
            milestones.forEach(
              milestone => {
                FrontRepoSingloton.Milestones.set(milestone.ID, milestone)
                FrontRepoSingloton.Milestones_batch.set(milestone.ID, milestone)
              }
            )
            
            // clear milestones that are absent from the batch
            FrontRepoSingloton.Milestones.forEach(
              milestone => {
                if (FrontRepoSingloton.Milestones_batch.get(milestone.ID) == undefined) {
                  FrontRepoSingloton.Milestones.delete(milestone.ID)
                }
              }
            )
            
            // sort Milestones_array array
            FrontRepoSingloton.Milestones_array.sort((t1, t2) => {
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
            bars.forEach(
              bar => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Lane.Bars redeeming
                {
                  let _lane = FrontRepoSingloton.Lanes.get(bar.Lane_BarsDBID.Int64)
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
                  let _gantt = FrontRepoSingloton.Gantts.get(group.Gantt_GroupsDBID.Int64)
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
                  let _gantt = FrontRepoSingloton.Gantts.get(lane.Gantt_LanesDBID.Int64)
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
                  let _group = FrontRepoSingloton.Groups.get(lane.Group_GroupLanesDBID.Int64)
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
                // insertion point for slice of pointer field Milestone.DiamonfAndTextAnchors redeeming
                {
                  let _milestone = FrontRepoSingloton.Milestones.get(lane.Milestone_DiamonfAndTextAnchorsDBID.Int64)
                  if (_milestone) {
                    if (_milestone.DiamonfAndTextAnchors == undefined) {
                      _milestone.DiamonfAndTextAnchors = new Array<LaneDB>()
                    }
                    _milestone.DiamonfAndTextAnchors.push(lane)
                    if (lane.Milestone_DiamonfAndTextAnchors_reverse == undefined) {
                      lane.Milestone_DiamonfAndTextAnchors_reverse = _milestone
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
                  let _gantt = FrontRepoSingloton.Gantts.get(milestone.Gantt_MilestonesDBID.Int64)
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

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // BarPull performs a GET on Bar of the stack and redeem association pointers 
  BarPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.barService.getBars()
        ]).subscribe(
          ([ // insertion point sub template 
            bars,
          ]) => {
            // init the array
            FrontRepoSingloton.Bars_array = bars

            // clear the map that counts Bar in the GET
            FrontRepoSingloton.Bars_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            bars.forEach(
              bar => {
                FrontRepoSingloton.Bars.set(bar.ID, bar)
                FrontRepoSingloton.Bars_batch.set(bar.ID, bar)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Lane.Bars redeeming
                {
                  let _lane = FrontRepoSingloton.Lanes.get(bar.Lane_BarsDBID.Int64)
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
            FrontRepoSingloton.Bars.forEach(
              bar => {
                if (FrontRepoSingloton.Bars_batch.get(bar.ID) == undefined) {
                  FrontRepoSingloton.Bars.delete(bar.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.ganttService.getGantts()
        ]).subscribe(
          ([ // insertion point sub template 
            gantts,
          ]) => {
            // init the array
            FrontRepoSingloton.Gantts_array = gantts

            // clear the map that counts Gantt in the GET
            FrontRepoSingloton.Gantts_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            gantts.forEach(
              gantt => {
                FrontRepoSingloton.Gantts.set(gantt.ID, gantt)
                FrontRepoSingloton.Gantts_batch.set(gantt.ID, gantt)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
              }
            )

            // clear gantts that are absent from the GET
            FrontRepoSingloton.Gantts.forEach(
              gantt => {
                if (FrontRepoSingloton.Gantts_batch.get(gantt.ID) == undefined) {
                  FrontRepoSingloton.Gantts.delete(gantt.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.groupService.getGroups()
        ]).subscribe(
          ([ // insertion point sub template 
            groups,
          ]) => {
            // init the array
            FrontRepoSingloton.Groups_array = groups

            // clear the map that counts Group in the GET
            FrontRepoSingloton.Groups_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            groups.forEach(
              group => {
                FrontRepoSingloton.Groups.set(group.ID, group)
                FrontRepoSingloton.Groups_batch.set(group.ID, group)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Gantt.Groups redeeming
                {
                  let _gantt = FrontRepoSingloton.Gantts.get(group.Gantt_GroupsDBID.Int64)
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
            FrontRepoSingloton.Groups.forEach(
              group => {
                if (FrontRepoSingloton.Groups_batch.get(group.ID) == undefined) {
                  FrontRepoSingloton.Groups.delete(group.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.laneService.getLanes()
        ]).subscribe(
          ([ // insertion point sub template 
            lanes,
          ]) => {
            // init the array
            FrontRepoSingloton.Lanes_array = lanes

            // clear the map that counts Lane in the GET
            FrontRepoSingloton.Lanes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            lanes.forEach(
              lane => {
                FrontRepoSingloton.Lanes.set(lane.ID, lane)
                FrontRepoSingloton.Lanes_batch.set(lane.ID, lane)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Gantt.Lanes redeeming
                {
                  let _gantt = FrontRepoSingloton.Gantts.get(lane.Gantt_LanesDBID.Int64)
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
                  let _group = FrontRepoSingloton.Groups.get(lane.Group_GroupLanesDBID.Int64)
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
                // insertion point for slice of pointer field Milestone.DiamonfAndTextAnchors redeeming
                {
                  let _milestone = FrontRepoSingloton.Milestones.get(lane.Milestone_DiamonfAndTextAnchorsDBID.Int64)
                  if (_milestone) {
                    if (_milestone.DiamonfAndTextAnchors == undefined) {
                      _milestone.DiamonfAndTextAnchors = new Array<LaneDB>()
                    }
                    _milestone.DiamonfAndTextAnchors.push(lane)
                    if (lane.Milestone_DiamonfAndTextAnchors_reverse == undefined) {
                      lane.Milestone_DiamonfAndTextAnchors_reverse = _milestone
                    }
                  }
                }
              }
            )

            // clear lanes that are absent from the GET
            FrontRepoSingloton.Lanes.forEach(
              lane => {
                if (FrontRepoSingloton.Lanes_batch.get(lane.ID) == undefined) {
                  FrontRepoSingloton.Lanes.delete(lane.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.milestoneService.getMilestones()
        ]).subscribe(
          ([ // insertion point sub template 
            milestones,
          ]) => {
            // init the array
            FrontRepoSingloton.Milestones_array = milestones

            // clear the map that counts Milestone in the GET
            FrontRepoSingloton.Milestones_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            milestones.forEach(
              milestone => {
                FrontRepoSingloton.Milestones.set(milestone.ID, milestone)
                FrontRepoSingloton.Milestones_batch.set(milestone.ID, milestone)

                // insertion point for redeeming ONE/ZERO-ONE associations 

                // insertion point for redeeming ONE-MANY associations 
                // insertion point for slice of pointer field Gantt.Milestones redeeming
                {
                  let _gantt = FrontRepoSingloton.Gantts.get(milestone.Gantt_MilestonesDBID.Int64)
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
            FrontRepoSingloton.Milestones.forEach(
              milestone => {
                if (FrontRepoSingloton.Milestones_batch.get(milestone.ID) == undefined) {
                  FrontRepoSingloton.Milestones.delete(milestone.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getBarUniqueID(id: number): number {
  return 31 * id
}
export function getGanttUniqueID(id: number): number {
  return 37 * id
}
export function getGroupUniqueID(id: number): number {
  return 41 * id
}
export function getLaneUniqueID(id: number): number {
  return 43 * id
}
export function getMilestoneUniqueID(id: number): number {
  return 47 * id
}
