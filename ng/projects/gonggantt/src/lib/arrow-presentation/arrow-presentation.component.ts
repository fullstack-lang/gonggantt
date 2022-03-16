import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ArrowDB } from '../arrow-db'
import { ArrowService } from '../arrow.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface arrowDummyElement {
}

const ELEMENT_DATA: arrowDummyElement[] = [
];

@Component({
	selector: 'app-arrow-presentation',
	templateUrl: './arrow-presentation.component.html',
	styleUrls: ['./arrow-presentation.component.css'],
})
export class ArrowPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	arrow: ArrowDB = new (ArrowDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private arrowService: ArrowService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getArrow();

		// observable for changes in 
		this.arrowService.ArrowServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getArrow()
				}
			}
		)
	}

	getArrow(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.arrow = this.frontRepo.Arrows.get(id)!

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
				github_com_fullstack_lang_gonggantt_go_editor: ["github_com_fullstack_lang_gonggantt_go-" + "arrow-detail", ID]
			}
		}]);
	}
}
