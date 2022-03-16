import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LaneDB } from '../lane-db'
import { LaneService } from '../lane.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface laneDummyElement {
}

const ELEMENT_DATA: laneDummyElement[] = [
];

@Component({
	selector: 'app-lane-presentation',
	templateUrl: './lane-presentation.component.html',
	styleUrls: ['./lane-presentation.component.css'],
})
export class LanePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	lane: LaneDB = new (LaneDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private laneService: LaneService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLane();

		// observable for changes in 
		this.laneService.LaneServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLane()
				}
			}
		)
	}

	getLane(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.lane = this.frontRepo.Lanes.get(id)!

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
				github_com_fullstack_lang_gonggantt_go_editor: ["github_com_fullstack_lang_gonggantt_go-" + "lane-detail", ID]
			}
		}]);
	}
}
