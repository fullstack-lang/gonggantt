import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { AnimatesTableComponent } from './animates-table/animates-table.component'
import { AnimateDetailComponent } from './animate-detail/animate-detail.component'

import { CirclesTableComponent } from './circles-table/circles-table.component'
import { CircleDetailComponent } from './circle-detail/circle-detail.component'

import { EllipsesTableComponent } from './ellipses-table/ellipses-table.component'
import { EllipseDetailComponent } from './ellipse-detail/ellipse-detail.component'

import { LinesTableComponent } from './lines-table/lines-table.component'
import { LineDetailComponent } from './line-detail/line-detail.component'

import { PathsTableComponent } from './paths-table/paths-table.component'
import { PathDetailComponent } from './path-detail/path-detail.component'

import { PolygonesTableComponent } from './polygones-table/polygones-table.component'
import { PolygoneDetailComponent } from './polygone-detail/polygone-detail.component'

import { PolylinesTableComponent } from './polylines-table/polylines-table.component'
import { PolylineDetailComponent } from './polyline-detail/polyline-detail.component'

import { RectsTableComponent } from './rects-table/rects-table.component'
import { RectDetailComponent } from './rect-detail/rect-detail.component'

import { SVGsTableComponent } from './svgs-table/svgs-table.component'
import { SVGDetailComponent } from './svg-detail/svg-detail.component'

import { TextsTableComponent } from './texts-table/texts-table.component'
import { TextDetailComponent } from './text-detail/text-detail.component'


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(newRoutes: Routes): void {
        const existingRoutes = this.router.config
        this.routes = this.router.config

        for (let newRoute of newRoutes) {
            if (!existingRoutes.includes(newRoute)) {
                this.routes.push(newRoute)
            }
        }
        this.router.resetConfig(this.routes)
    }

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongsvg_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getAnimateTablePath(): string {
        return this.getPathRoot() + '-animates/:GONG__StackPath'
    }
    getAnimateTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAnimateTablePath(), component: AnimatesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getAnimateAdderPath(): string {
        return this.getPathRoot() + '-animate-adder/:GONG__StackPath'
    }
    getAnimateAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAnimateAdderPath(), component: AnimateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getAnimateAdderForUsePath(): string {
        return this.getPathRoot() + '-animate-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getAnimateAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAnimateAdderForUsePath(), component: AnimateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getAnimateDetailPath(): string {
        return this.getPathRoot() + '-animate-detail/:id/:GONG__StackPath'
    }
    getAnimateDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getAnimateDetailPath(), component: AnimateDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getCircleTablePath(): string {
        return this.getPathRoot() + '-circles/:GONG__StackPath'
    }
    getCircleTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleTablePath(), component: CirclesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCircleAdderPath(): string {
        return this.getPathRoot() + '-circle-adder/:GONG__StackPath'
    }
    getCircleAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleAdderPath(), component: CircleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCircleAdderForUsePath(): string {
        return this.getPathRoot() + '-circle-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCircleAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleAdderForUsePath(), component: CircleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCircleDetailPath(): string {
        return this.getPathRoot() + '-circle-detail/:id/:GONG__StackPath'
    }
    getCircleDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleDetailPath(), component: CircleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getEllipseTablePath(): string {
        return this.getPathRoot() + '-ellipses/:GONG__StackPath'
    }
    getEllipseTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEllipseTablePath(), component: EllipsesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getEllipseAdderPath(): string {
        return this.getPathRoot() + '-ellipse-adder/:GONG__StackPath'
    }
    getEllipseAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEllipseAdderPath(), component: EllipseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEllipseAdderForUsePath(): string {
        return this.getPathRoot() + '-ellipse-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getEllipseAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEllipseAdderForUsePath(), component: EllipseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEllipseDetailPath(): string {
        return this.getPathRoot() + '-ellipse-detail/:id/:GONG__StackPath'
    }
    getEllipseDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEllipseDetailPath(), component: EllipseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getLineTablePath(): string {
        return this.getPathRoot() + '-lines/:GONG__StackPath'
    }
    getLineTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLineTablePath(), component: LinesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getLineAdderPath(): string {
        return this.getPathRoot() + '-line-adder/:GONG__StackPath'
    }
    getLineAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLineAdderPath(), component: LineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLineAdderForUsePath(): string {
        return this.getPathRoot() + '-line-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getLineAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLineAdderForUsePath(), component: LineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLineDetailPath(): string {
        return this.getPathRoot() + '-line-detail/:id/:GONG__StackPath'
    }
    getLineDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLineDetailPath(), component: LineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getPathTablePath(): string {
        return this.getPathRoot() + '-paths/:GONG__StackPath'
    }
    getPathTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPathTablePath(), component: PathsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getPathAdderPath(): string {
        return this.getPathRoot() + '-path-adder/:GONG__StackPath'
    }
    getPathAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPathAdderPath(), component: PathDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPathAdderForUsePath(): string {
        return this.getPathRoot() + '-path-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getPathAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPathAdderForUsePath(), component: PathDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPathDetailPath(): string {
        return this.getPathRoot() + '-path-detail/:id/:GONG__StackPath'
    }
    getPathDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPathDetailPath(), component: PathDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getPolygoneTablePath(): string {
        return this.getPathRoot() + '-polygones/:GONG__StackPath'
    }
    getPolygoneTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolygoneTablePath(), component: PolygonesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getPolygoneAdderPath(): string {
        return this.getPathRoot() + '-polygone-adder/:GONG__StackPath'
    }
    getPolygoneAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolygoneAdderPath(), component: PolygoneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPolygoneAdderForUsePath(): string {
        return this.getPathRoot() + '-polygone-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getPolygoneAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolygoneAdderForUsePath(), component: PolygoneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPolygoneDetailPath(): string {
        return this.getPathRoot() + '-polygone-detail/:id/:GONG__StackPath'
    }
    getPolygoneDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolygoneDetailPath(), component: PolygoneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getPolylineTablePath(): string {
        return this.getPathRoot() + '-polylines/:GONG__StackPath'
    }
    getPolylineTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolylineTablePath(), component: PolylinesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getPolylineAdderPath(): string {
        return this.getPathRoot() + '-polyline-adder/:GONG__StackPath'
    }
    getPolylineAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolylineAdderPath(), component: PolylineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPolylineAdderForUsePath(): string {
        return this.getPathRoot() + '-polyline-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getPolylineAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolylineAdderForUsePath(), component: PolylineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getPolylineDetailPath(): string {
        return this.getPathRoot() + '-polyline-detail/:id/:GONG__StackPath'
    }
    getPolylineDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getPolylineDetailPath(), component: PolylineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getRectTablePath(): string {
        return this.getPathRoot() + '-rects/:GONG__StackPath'
    }
    getRectTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRectTablePath(), component: RectsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getRectAdderPath(): string {
        return this.getPathRoot() + '-rect-adder/:GONG__StackPath'
    }
    getRectAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRectAdderPath(), component: RectDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRectAdderForUsePath(): string {
        return this.getPathRoot() + '-rect-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getRectAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRectAdderForUsePath(), component: RectDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRectDetailPath(): string {
        return this.getPathRoot() + '-rect-detail/:id/:GONG__StackPath'
    }
    getRectDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRectDetailPath(), component: RectDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getSVGTablePath(): string {
        return this.getPathRoot() + '-svgs/:GONG__StackPath'
    }
    getSVGTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSVGTablePath(), component: SVGsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getSVGAdderPath(): string {
        return this.getPathRoot() + '-svg-adder/:GONG__StackPath'
    }
    getSVGAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSVGAdderPath(), component: SVGDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getSVGAdderForUsePath(): string {
        return this.getPathRoot() + '-svg-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getSVGAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSVGAdderForUsePath(), component: SVGDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getSVGDetailPath(): string {
        return this.getPathRoot() + '-svg-detail/:id/:GONG__StackPath'
    }
    getSVGDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSVGDetailPath(), component: SVGDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getTextTablePath(): string {
        return this.getPathRoot() + '-texts/:GONG__StackPath'
    }
    getTextTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextTablePath(), component: TextsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getTextAdderPath(): string {
        return this.getPathRoot() + '-text-adder/:GONG__StackPath'
    }
    getTextAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextAdderPath(), component: TextDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTextAdderForUsePath(): string {
        return this.getPathRoot() + '-text-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getTextAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextAdderForUsePath(), component: TextDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getTextDetailPath(): string {
        return this.getPathRoot() + '-text-detail/:id/:GONG__StackPath'
    }
    getTextDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getTextDetailPath(), component: TextDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getAnimateTableRoute(stackPath),
            this.getAnimateAdderRoute(stackPath),
            this.getAnimateAdderForUseRoute(stackPath),
            this.getAnimateDetailRoute(stackPath),

            this.getCircleTableRoute(stackPath),
            this.getCircleAdderRoute(stackPath),
            this.getCircleAdderForUseRoute(stackPath),
            this.getCircleDetailRoute(stackPath),

            this.getEllipseTableRoute(stackPath),
            this.getEllipseAdderRoute(stackPath),
            this.getEllipseAdderForUseRoute(stackPath),
            this.getEllipseDetailRoute(stackPath),

            this.getLineTableRoute(stackPath),
            this.getLineAdderRoute(stackPath),
            this.getLineAdderForUseRoute(stackPath),
            this.getLineDetailRoute(stackPath),

            this.getPathTableRoute(stackPath),
            this.getPathAdderRoute(stackPath),
            this.getPathAdderForUseRoute(stackPath),
            this.getPathDetailRoute(stackPath),

            this.getPolygoneTableRoute(stackPath),
            this.getPolygoneAdderRoute(stackPath),
            this.getPolygoneAdderForUseRoute(stackPath),
            this.getPolygoneDetailRoute(stackPath),

            this.getPolylineTableRoute(stackPath),
            this.getPolylineAdderRoute(stackPath),
            this.getPolylineAdderForUseRoute(stackPath),
            this.getPolylineDetailRoute(stackPath),

            this.getRectTableRoute(stackPath),
            this.getRectAdderRoute(stackPath),
            this.getRectAdderForUseRoute(stackPath),
            this.getRectDetailRoute(stackPath),

            this.getSVGTableRoute(stackPath),
            this.getSVGAdderRoute(stackPath),
            this.getSVGAdderForUseRoute(stackPath),
            this.getSVGDetailRoute(stackPath),

            this.getTextTableRoute(stackPath),
            this.getTextAdderRoute(stackPath),
            this.getTextAdderForUseRoute(stackPath),
            this.getTextDetailRoute(stackPath),

        ])
    }
}