import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { BarDB } from '../bar-db'
import { BarService } from '../bar.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface barDummyElement {
}

const ELEMENT_DATA: barDummyElement[] = [
];

@Component({
	selector: 'app-bar-presentation',
	templateUrl: './bar-presentation.component.html',
	styleUrls: ['./bar-presentation.component.css'],
})
export class BarPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	bar: BarDB;
 
	constructor(
		private barService: BarService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getBar();

		// observable for changes in 
		this.barService.BarServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getBar()
				}
			}
		)
	}

	getBar(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.barService.getBar(id)
			.subscribe(
				bar => {
					this.bar = bar

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
				editor: ["bar-detail", ID]
			}
		}]);
	}
}
