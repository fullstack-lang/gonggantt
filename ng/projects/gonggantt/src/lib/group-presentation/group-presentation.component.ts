import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { GroupDB } from '../group-db'
import { GroupService } from '../group.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface groupDummyElement {
}

const ELEMENT_DATA: groupDummyElement[] = [
];

@Component({
	selector: 'app-group-presentation',
	templateUrl: './group-presentation.component.html',
	styleUrls: ['./group-presentation.component.css'],
})
export class GroupPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	group: GroupDB = new (GroupDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private groupService: GroupService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getGroup();

		// observable for changes in 
		this.groupService.GroupServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getGroup()
				}
			}
		)
	}

	getGroup(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.group = this.frontRepo.Groups.get(id)!

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_gonggantt_go_editor: ["github_com_fullstack_lang_gonggantt_go-" + "group-detail", ID]
			}
		}]);
	}
}
