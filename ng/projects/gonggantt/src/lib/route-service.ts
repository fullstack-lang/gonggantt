import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
import { ArrowsTableComponent } from './arrows-table/arrows-table.component'
import { ArrowDetailComponent } from './arrow-detail/arrow-detail.component'

import { BarsTableComponent } from './bars-table/bars-table.component'
import { BarDetailComponent } from './bar-detail/bar-detail.component'

import { GanttsTableComponent } from './gantts-table/gantts-table.component'
import { GanttDetailComponent } from './gantt-detail/gantt-detail.component'

import { GroupsTableComponent } from './groups-table/groups-table.component'
import { GroupDetailComponent } from './group-detail/group-detail.component'

import { LanesTableComponent } from './lanes-table/lanes-table.component'
import { LaneDetailComponent } from './lane-detail/lane-detail.component'

import { LaneUsesTableComponent } from './laneuses-table/laneuses-table.component'
import { LaneUseDetailComponent } from './laneuse-detail/laneuse-detail.component'

import { MilestonesTableComponent } from './milestones-table/milestones-table.component'
import { MilestoneDetailComponent } from './milestone-detail/milestone-detail.component'


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
        return 'github_com_fullstack_lang_gonggantt_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getArrowTablePath(): string {
        return this.getPathRoot() + '-arrows/:GONG__StackPath'
    }
    getArrowTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getArrowTablePath(), component: ArrowsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getArrowAdderPath(): string {
        return this.getPathRoot() + '-arrow-adder/:GONG__StackPath'
    }
    getArrowAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getArrowAdderPath(), component: ArrowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getArrowAdderForUsePath(): string {
        return this.getPathRoot() + '-arrow-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getArrowAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getArrowAdderForUsePath(), component: ArrowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getArrowDetailPath(): string {
        return this.getPathRoot() + '-arrow-detail/:id/:GONG__StackPath'
    }
    getArrowDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getArrowDetailPath(), component: ArrowDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getBarTablePath(): string {
        return this.getPathRoot() + '-bars/:GONG__StackPath'
    }
    getBarTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBarTablePath(), component: BarsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getBarAdderPath(): string {
        return this.getPathRoot() + '-bar-adder/:GONG__StackPath'
    }
    getBarAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBarAdderPath(), component: BarDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getBarAdderForUsePath(): string {
        return this.getPathRoot() + '-bar-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getBarAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBarAdderForUsePath(), component: BarDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getBarDetailPath(): string {
        return this.getPathRoot() + '-bar-detail/:id/:GONG__StackPath'
    }
    getBarDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getBarDetailPath(), component: BarDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGanttTablePath(): string {
        return this.getPathRoot() + '-gantts/:GONG__StackPath'
    }
    getGanttTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGanttTablePath(), component: GanttsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGanttAdderPath(): string {
        return this.getPathRoot() + '-gantt-adder/:GONG__StackPath'
    }
    getGanttAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGanttAdderPath(), component: GanttDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGanttAdderForUsePath(): string {
        return this.getPathRoot() + '-gantt-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGanttAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGanttAdderForUsePath(), component: GanttDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGanttDetailPath(): string {
        return this.getPathRoot() + '-gantt-detail/:id/:GONG__StackPath'
    }
    getGanttDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGanttDetailPath(), component: GanttDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGroupTablePath(): string {
        return this.getPathRoot() + '-groups/:GONG__StackPath'
    }
    getGroupTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGroupTablePath(), component: GroupsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGroupAdderPath(): string {
        return this.getPathRoot() + '-group-adder/:GONG__StackPath'
    }
    getGroupAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGroupAdderPath(), component: GroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGroupAdderForUsePath(): string {
        return this.getPathRoot() + '-group-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGroupAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGroupAdderForUsePath(), component: GroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGroupDetailPath(): string {
        return this.getPathRoot() + '-group-detail/:id/:GONG__StackPath'
    }
    getGroupDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGroupDetailPath(), component: GroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getLaneTablePath(): string {
        return this.getPathRoot() + '-lanes/:GONG__StackPath'
    }
    getLaneTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneTablePath(), component: LanesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getLaneAdderPath(): string {
        return this.getPathRoot() + '-lane-adder/:GONG__StackPath'
    }
    getLaneAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneAdderPath(), component: LaneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLaneAdderForUsePath(): string {
        return this.getPathRoot() + '-lane-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getLaneAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneAdderForUsePath(), component: LaneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLaneDetailPath(): string {
        return this.getPathRoot() + '-lane-detail/:id/:GONG__StackPath'
    }
    getLaneDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneDetailPath(), component: LaneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getLaneUseTablePath(): string {
        return this.getPathRoot() + '-laneuses/:GONG__StackPath'
    }
    getLaneUseTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneUseTablePath(), component: LaneUsesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getLaneUseAdderPath(): string {
        return this.getPathRoot() + '-laneuse-adder/:GONG__StackPath'
    }
    getLaneUseAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneUseAdderPath(), component: LaneUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLaneUseAdderForUsePath(): string {
        return this.getPathRoot() + '-laneuse-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getLaneUseAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneUseAdderForUsePath(), component: LaneUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLaneUseDetailPath(): string {
        return this.getPathRoot() + '-laneuse-detail/:id/:GONG__StackPath'
    }
    getLaneUseDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLaneUseDetailPath(), component: LaneUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getMilestoneTablePath(): string {
        return this.getPathRoot() + '-milestones/:GONG__StackPath'
    }
    getMilestoneTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMilestoneTablePath(), component: MilestonesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getMilestoneAdderPath(): string {
        return this.getPathRoot() + '-milestone-adder/:GONG__StackPath'
    }
    getMilestoneAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMilestoneAdderPath(), component: MilestoneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMilestoneAdderForUsePath(): string {
        return this.getPathRoot() + '-milestone-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getMilestoneAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMilestoneAdderForUsePath(), component: MilestoneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMilestoneDetailPath(): string {
        return this.getPathRoot() + '-milestone-detail/:id/:GONG__StackPath'
    }
    getMilestoneDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMilestoneDetailPath(), component: MilestoneDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getArrowTableRoute(stackPath),
            this.getArrowAdderRoute(stackPath),
            this.getArrowAdderForUseRoute(stackPath),
            this.getArrowDetailRoute(stackPath),

            this.getBarTableRoute(stackPath),
            this.getBarAdderRoute(stackPath),
            this.getBarAdderForUseRoute(stackPath),
            this.getBarDetailRoute(stackPath),

            this.getGanttTableRoute(stackPath),
            this.getGanttAdderRoute(stackPath),
            this.getGanttAdderForUseRoute(stackPath),
            this.getGanttDetailRoute(stackPath),

            this.getGroupTableRoute(stackPath),
            this.getGroupAdderRoute(stackPath),
            this.getGroupAdderForUseRoute(stackPath),
            this.getGroupDetailRoute(stackPath),

            this.getLaneTableRoute(stackPath),
            this.getLaneAdderRoute(stackPath),
            this.getLaneAdderForUseRoute(stackPath),
            this.getLaneDetailRoute(stackPath),

            this.getLaneUseTableRoute(stackPath),
            this.getLaneUseAdderRoute(stackPath),
            this.getLaneUseAdderForUseRoute(stackPath),
            this.getLaneUseDetailRoute(stackPath),

            this.getMilestoneTableRoute(stackPath),
            this.getMilestoneAdderRoute(stackPath),
            this.getMilestoneAdderForUseRoute(stackPath),
            this.getMilestoneDetailRoute(stackPath),

        ])
    }
}
