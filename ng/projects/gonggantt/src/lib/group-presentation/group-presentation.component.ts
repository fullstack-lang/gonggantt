import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { GroupDB } from '../group-db'
import { GroupService } from '../group.service'

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

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	group: GroupDB;
 
	constructor(
		private groupService: GroupService,
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
		const id = +this.route.snapshot.paramMap.get('id');
		this.groupService.getGroup(id)
			.subscribe(
				group => {
					this.group = group

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
				editor: ["group-detail", ID]
			}
		}]);
	}
}
