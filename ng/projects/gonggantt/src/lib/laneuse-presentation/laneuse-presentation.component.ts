import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LaneUseDB } from '../laneuse-db'
import { LaneUseService } from '../laneuse.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface laneuseDummyElement {
}

const ELEMENT_DATA: laneuseDummyElement[] = [
];

@Component({
	selector: 'app-laneuse-presentation',
	templateUrl: './laneuse-presentation.component.html',
	styleUrls: ['./laneuse-presentation.component.css'],
})
export class LaneUsePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	laneuse: LaneUseDB = new (LaneUseDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private laneuseService: LaneUseService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLaneUse();

		// observable for changes in 
		this.laneuseService.LaneUseServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLaneUse()
				}
			}
		)
	}

	getLaneUse(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.laneuse = this.frontRepo.LaneUses.get(id)!

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
				github_com_fullstack_lang_gonggantt_go_editor: ["github_com_fullstack_lang_gonggantt_go-" + "laneuse-detail", ID]
			}
		}]);
	}
}
