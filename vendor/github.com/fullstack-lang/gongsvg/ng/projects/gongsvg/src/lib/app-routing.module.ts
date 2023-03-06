import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

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


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongsvg_go-animates/:GONG__StackPath', component: AnimatesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-adder/:GONG__StackPath', component: AnimateDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: AnimateDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-animate-detail/:id/:GONG__StackPath', component: AnimateDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-circles/:GONG__StackPath', component: CirclesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-adder/:GONG__StackPath', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-circle-detail/:id/:GONG__StackPath', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipses/:GONG__StackPath', component: EllipsesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-adder/:GONG__StackPath', component: EllipseDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: EllipseDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-ellipse-detail/:id/:GONG__StackPath', component: EllipseDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-lines/:GONG__StackPath', component: LinesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-adder/:GONG__StackPath', component: LineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: LineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-line-detail/:id/:GONG__StackPath', component: LineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-paths/:GONG__StackPath', component: PathsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-adder/:GONG__StackPath', component: PathDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: PathDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-path-detail/:id/:GONG__StackPath', component: PathDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-polygones/:GONG__StackPath', component: PolygonesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-adder/:GONG__StackPath', component: PolygoneDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: PolygoneDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polygone-detail/:id/:GONG__StackPath', component: PolygoneDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-polylines/:GONG__StackPath', component: PolylinesTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-adder/:GONG__StackPath', component: PolylineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: PolylineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-polyline-detail/:id/:GONG__StackPath', component: PolylineDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-rects/:GONG__StackPath', component: RectsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-adder/:GONG__StackPath', component: RectDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: RectDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-rect-detail/:id/:GONG__StackPath', component: RectDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-svgs/:GONG__StackPath', component: SVGsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-adder/:GONG__StackPath', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-svg-detail/:id/:GONG__StackPath', component: SVGDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsvg_go-texts/:GONG__StackPath', component: TextsTableComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_table' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-adder/:GONG__StackPath', component: TextDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: TextDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsvg_go-text-detail/:id/:GONG__StackPath', component: TextDetailComponent, outlet: 'github_com_fullstack_lang_gongsvg_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
