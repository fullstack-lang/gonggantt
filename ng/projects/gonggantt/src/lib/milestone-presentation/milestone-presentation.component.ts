import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { MilestoneDB } from '../milestone-db'
import { MilestoneService } from '../milestone.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface milestoneDummyElement {
}

const ELEMENT_DATA: milestoneDummyElement[] = [
];

@Component({
	selector: 'app-milestone-presentation',
	templateUrl: './milestone-presentation.component.html',
	styleUrls: ['./milestone-presentation.component.css'],
})
export class MilestonePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	milestone: MilestoneDB;
 
	constructor(
		private milestoneService: MilestoneService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getMilestone();

		// observable for changes in 
		this.milestoneService.MilestoneServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMilestone()
				}
			}
		)
	}

	getMilestone(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.milestoneService.getMilestone(id)
			.subscribe(
				milestone => {
					this.milestone = milestone

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
				editor: ["milestone-detail", ID]
			}
		}]);
	}
}