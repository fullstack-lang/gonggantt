// insertion point sub template for components imports 
import { BarsTableComponent } from './bars-table/bars-table.component'
import { GanttsTableComponent } from './gantts-table/gantts-table.component'
import { LanesTableComponent } from './lanes-table/lanes-table.component'
import { MilestonesTableComponent } from './milestones-table/milestones-table.component'

// insertion point sub template for map of components per struct 
export const MapOfBarsComponents: Map<string, any> = new Map([["BarsTableComponent", BarsTableComponent],])
export const MapOfGanttsComponents: Map<string, any> = new Map([["GanttsTableComponent", GanttsTableComponent],])
export const MapOfLanesComponents: Map<string, any> = new Map([["LanesTableComponent", LanesTableComponent],])
export const MapOfMilestonesComponents: Map<string, any> = new Map([["MilestonesTableComponent", MilestonesTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Bar", MapOfBarsComponents],
      ["Gantt", MapOfGanttsComponents],
      ["Lane", MapOfLanesComponents],
      ["Milestone", MapOfMilestonesComponents],
    ]
  )