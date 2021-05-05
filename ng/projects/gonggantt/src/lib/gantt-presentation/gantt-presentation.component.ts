import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GanttDB } from '../gantt-db'
import { GanttService } from '../gantt.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

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

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	gantt: GanttDB;
 
	constructor(
		private ganttService: GanttService,
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
		const id = +this.route.snapshot.paramMap.get('id');
		this.ganttService.getGantt(id)
			.subscribe(
				gantt => {
					this.gantt = gantt

					// insertion point for recovery of durations

				}
			);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				presentation: [structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				editor: ["gantt-detail", ID]
			}
		}]);
	}
}
