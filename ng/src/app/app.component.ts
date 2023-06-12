import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
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

  GONG__GANTT_DATA_MODEL__StackPath = "github.com/fullstack-lang/gonggantt/go/models"
  GONG__SVG_DATA_MODEL__StackPath = "github.com/fullstack-lang/gongsvg/go/models"
  GONG__GANTT_SVG__StackPath = "gonggantt+github.com/fullstack-lang/gongsvg/go/models"
  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
