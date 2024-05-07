import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gonggantt from 'gonggantt'
import * as gongsvg from 'gongsvg'

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { GongtableModule } from 'gongtable'
import { GongtablespecificModule } from 'gongtablespecific'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  gantt_stack_probe = "Gantt Stack Probe"
  svg_stack_probe = 'SVG Stack Probe'
  gantt = "Gantt diagram"
  view = this.gantt

  views: string[] = [this.gantt, this.gantt_stack_probe, this.svg_stack_probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  GanttStackType = gonggantt.StackType
  SVGStackType = gongsvg.StackType

  GanttStacksNames = gonggantt.GanttStacksNames

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
