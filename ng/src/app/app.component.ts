import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongsvg from 'gongsvg'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {


  GONG__StackPath="github.com/fullstack-lang/gonggantt/go/models"

  view = 'View'
  default = 'View'
  data = 'Data'
  model = 'Model'

  loading = true

  views: string[] = [this.default, this.data, this.model];


  constructor(
  ) {

  }

  ngOnInit(): void {
    this.loading = false
  }
}
