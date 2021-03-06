// insertion point sub template for components imports 
  import { AnimatesTableComponent } from './animates-table/animates-table.component'
  import { AnimateSortingComponent } from './animate-sorting/animate-sorting.component'
  import { CirclesTableComponent } from './circles-table/circles-table.component'
  import { CircleSortingComponent } from './circle-sorting/circle-sorting.component'
  import { EllipsesTableComponent } from './ellipses-table/ellipses-table.component'
  import { EllipseSortingComponent } from './ellipse-sorting/ellipse-sorting.component'
  import { LinesTableComponent } from './lines-table/lines-table.component'
  import { LineSortingComponent } from './line-sorting/line-sorting.component'
  import { PathsTableComponent } from './paths-table/paths-table.component'
  import { PathSortingComponent } from './path-sorting/path-sorting.component'
  import { PolygonesTableComponent } from './polygones-table/polygones-table.component'
  import { PolygoneSortingComponent } from './polygone-sorting/polygone-sorting.component'
  import { PolylinesTableComponent } from './polylines-table/polylines-table.component'
  import { PolylineSortingComponent } from './polyline-sorting/polyline-sorting.component'
  import { RectsTableComponent } from './rects-table/rects-table.component'
  import { RectSortingComponent } from './rect-sorting/rect-sorting.component'
  import { SVGsTableComponent } from './svgs-table/svgs-table.component'
  import { SVGSortingComponent } from './svg-sorting/svg-sorting.component'
  import { TextsTableComponent } from './texts-table/texts-table.component'
  import { TextSortingComponent } from './text-sorting/text-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfAnimatesComponents: Map<string, any> = new Map([["AnimatesTableComponent", AnimatesTableComponent],])
  export const MapOfAnimateSortingComponents: Map<string, any> = new Map([["AnimateSortingComponent", AnimateSortingComponent],])
  export const MapOfCirclesComponents: Map<string, any> = new Map([["CirclesTableComponent", CirclesTableComponent],])
  export const MapOfCircleSortingComponents: Map<string, any> = new Map([["CircleSortingComponent", CircleSortingComponent],])
  export const MapOfEllipsesComponents: Map<string, any> = new Map([["EllipsesTableComponent", EllipsesTableComponent],])
  export const MapOfEllipseSortingComponents: Map<string, any> = new Map([["EllipseSortingComponent", EllipseSortingComponent],])
  export const MapOfLinesComponents: Map<string, any> = new Map([["LinesTableComponent", LinesTableComponent],])
  export const MapOfLineSortingComponents: Map<string, any> = new Map([["LineSortingComponent", LineSortingComponent],])
  export const MapOfPathsComponents: Map<string, any> = new Map([["PathsTableComponent", PathsTableComponent],])
  export const MapOfPathSortingComponents: Map<string, any> = new Map([["PathSortingComponent", PathSortingComponent],])
  export const MapOfPolygonesComponents: Map<string, any> = new Map([["PolygonesTableComponent", PolygonesTableComponent],])
  export const MapOfPolygoneSortingComponents: Map<string, any> = new Map([["PolygoneSortingComponent", PolygoneSortingComponent],])
  export const MapOfPolylinesComponents: Map<string, any> = new Map([["PolylinesTableComponent", PolylinesTableComponent],])
  export const MapOfPolylineSortingComponents: Map<string, any> = new Map([["PolylineSortingComponent", PolylineSortingComponent],])
  export const MapOfRectsComponents: Map<string, any> = new Map([["RectsTableComponent", RectsTableComponent],])
  export const MapOfRectSortingComponents: Map<string, any> = new Map([["RectSortingComponent", RectSortingComponent],])
  export const MapOfSVGsComponents: Map<string, any> = new Map([["SVGsTableComponent", SVGsTableComponent],])
  export const MapOfSVGSortingComponents: Map<string, any> = new Map([["SVGSortingComponent", SVGSortingComponent],])
  export const MapOfTextsComponents: Map<string, any> = new Map([["TextsTableComponent", TextsTableComponent],])
  export const MapOfTextSortingComponents: Map<string, any> = new Map([["TextSortingComponent", TextSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["Animate", MapOfAnimatesComponents],
      ["Circle", MapOfCirclesComponents],
      ["Ellipse", MapOfEllipsesComponents],
      ["Line", MapOfLinesComponents],
      ["Path", MapOfPathsComponents],
      ["Polygone", MapOfPolygonesComponents],
      ["Polyline", MapOfPolylinesComponents],
      ["Rect", MapOfRectsComponents],
      ["SVG", MapOfSVGsComponents],
      ["Text", MapOfTextsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["Animate", MapOfAnimateSortingComponents],
      ["Circle", MapOfCircleSortingComponents],
      ["Ellipse", MapOfEllipseSortingComponents],
      ["Line", MapOfLineSortingComponents],
      ["Path", MapOfPathSortingComponents],
      ["Polygone", MapOfPolygoneSortingComponents],
      ["Polyline", MapOfPolylineSortingComponents],
      ["Rect", MapOfRectSortingComponents],
      ["SVG", MapOfSVGSortingComponents],
      ["Text", MapOfTextSortingComponents],
    ]
  )
