import { Component, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as gonggantt from '../../projects/gonggantt/src/public-api'

import { GongganttSpecificComponent } from '../../projects/gongganttspecific/src/lib/gonggantt-specific/gonggantt-specific.component'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component'


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    AngularSplitModule,
    SplitSpecificComponent,
    GongganttSpecificComponent
  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  StackName = gonggantt.GanttStacksNames.GanttStackName

  constructor(
  ) {
  }

  ngOnInit(): void {
  }
}
