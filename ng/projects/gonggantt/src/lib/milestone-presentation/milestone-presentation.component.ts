import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { MilestoneDB } from '../milestone-db'
import { MilestoneService } from '../milestone.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

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

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	milestone: MilestoneDB = new (MilestoneDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private milestoneService: MilestoneService,
		private frontRepoService: FrontRepoService,
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
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.milestone = this.frontRepo.Milestones.get(id)!

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
				github_com_fullstack_lang_gonggantt_go_editor: ["github_com_fullstack_lang_gonggantt_go-" + "milestone-detail", ID]
			}
		}]);
	}
}
