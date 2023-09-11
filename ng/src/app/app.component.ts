import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gonggantt from 'gonggantt'


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {


  default = 'Gantt diagram Data/Model'
  svg = 'SVG diagram Data/Model'
  gantt = "Gantt diagram"
  view = this.gantt

  views: string[] = [this.gantt, this.default, this.svg];

  GanttStacksNames = gonggantt.GanttStacksNames
  
  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
