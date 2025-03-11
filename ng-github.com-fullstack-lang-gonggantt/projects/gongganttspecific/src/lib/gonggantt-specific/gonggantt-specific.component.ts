import { Component, Input } from '@angular/core';

import * as gonggant from '../../../../gonggantt/src/public-api'

import {SvgSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'


@Component({
  selector: 'lib-gonggantt-specific',
  imports: [
    SvgSpecificComponent
  ],
  templateUrl: './gonggantt-specific.component.html',
  styleUrl: './gonggantt-specific.component.css'
})
export class GongganttSpecificComponent {

    @Input() GONG__StackPath: string = ""

    SvgStackName = gonggant.GanttStacksNames.SvgStackName
}
