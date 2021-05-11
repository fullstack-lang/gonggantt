// insertion point sub template for components imports 
  import { BarsTableComponent } from './bars-table/bars-table.component'
  import { BarSortingComponent } from './bar-sorting/bar-sorting.component'
  import { GanttsTableComponent } from './gantts-table/gantts-table.component'
  import { GanttSortingComponent } from './gantt-sorting/gantt-sorting.component'
  import { GroupsTableComponent } from './groups-table/groups-table.component'
  import { GroupSortingComponent } from './group-sorting/group-sorting.component'
  import { LanesTableComponent } from './lanes-table/lanes-table.component'
  import { LaneSortingComponent } from './lane-sorting/lane-sorting.component'
  import { MilestonesTableComponent } from './milestones-table/milestones-table.component'
  import { MilestoneSortingComponent } from './milestone-sorting/milestone-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfBarsComponents: Map<string, any> = new Map([["BarsTableComponent", BarsTableComponent],])
  export const MapOfBarSortingComponents: Map<string, any> = new Map([["BarSortingComponent", BarSortingComponent],])
  export const MapOfGanttsComponents: Map<string, any> = new Map([["GanttsTableComponent", GanttsTableComponent],])
  export const MapOfGanttSortingComponents: Map<string, any> = new Map([["GanttSortingComponent", GanttSortingComponent],])
  export const MapOfGroupsComponents: Map<string, any> = new Map([["GroupsTableComponent", GroupsTableComponent],])
  export const MapOfGroupSortingComponents: Map<string, any> = new Map([["GroupSortingComponent", GroupSortingComponent],])
  export const MapOfLanesComponents: Map<string, any> = new Map([["LanesTableComponent", LanesTableComponent],])
  export const MapOfLaneSortingComponents: Map<string, any> = new Map([["LaneSortingComponent", LaneSortingComponent],])
  export const MapOfMilestonesComponents: Map<string, any> = new Map([["MilestonesTableComponent", MilestonesTableComponent],])
  export const MapOfMilestoneSortingComponents: Map<string, any> = new Map([["MilestoneSortingComponent", MilestoneSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Bar", MapOfBarsComponents],
      ["Gantt", MapOfGanttsComponents],
      ["Group", MapOfGroupsComponents],
      ["Lane", MapOfLanesComponents],
      ["Milestone", MapOfMilestonesComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Bar", MapOfBarSortingComponents],
      ["Gantt", MapOfGanttSortingComponents],
      ["Group", MapOfGroupSortingComponents],
      ["Lane", MapOfLaneSortingComponents],
      ["Milestone", MapOfMilestoneSortingComponents],
    ]
  )
