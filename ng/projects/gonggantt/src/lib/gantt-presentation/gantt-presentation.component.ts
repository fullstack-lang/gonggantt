import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GanttDB } from '../gantt-db'
import { GanttService } from '../gantt.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface ganttDummyElement {
}

const ELEMENT_DATA: ganttDummyElement[] = [
];

@Component({
	selector: 'app-gantt-presentation',
	templateUrl: './gantt-presentation.component.html',
	styleUrls: ['./gantt-presentation.component.css'],
})
export class GanttPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	gantt: GanttDB = new (GanttDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private ganttService: GanttService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGantt();

		// observable for changes in 
		this.ganttService.GanttServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGantt()
				}
			}
		)
	}

	getGantt(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.gantt = this.frontRepo.Gantts.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gonggantt_go_presentation: ["github_com_fullstack_lang_gonggantt_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gonggantt_go_editor: ["github_com_fullstack_lang_gonggantt_go-" + "gantt-detail", ID]
			}
		}]);
	}
}
