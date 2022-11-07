import { Component, OnInit, ViewChild } from '@angular/core';
import { Observable, Subscription, timer } from 'rxjs';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-svg',
  templateUrl: './svg.component.html',
  styleUrls: ['./svg.component.css']
})
export class SvgComponent implements OnInit {

  SVG: string = ""

  rect: number = 0

  public gongsvgFrontRepo?: gongsvg.FrontRepo

  public Rects = new Array<gongsvg.RectDB>()
  public Texts = new Array<gongsvg.TextDB>()
  public Lines = new Array<gongsvg.LineDB>()
  public Circles = new Array<gongsvg.CircleDB>()
  public Ellipses = new Array<gongsvg.EllipseDB>()
  public Paths = new Array<gongsvg.PathDB>()
  public Polygones = new Array<gongsvg.PolygoneDB>()
  public Polylines = new Array<gongsvg.PolylineDB>()

  /**
 * the component is refreshed when modification are performed in the back repo 
 * 
 * the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
 * if the commit number has increased, it pulls the front repo and redraw the diagram
 */
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommiNbFromBagetCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  svgSingloton?: gongsvg.SVGDB

  constructor(
    private gongsvgFrontRepoService: gongsvg.FrontRepoService,
    private gongsvgNbFromBackService: gongsvg.CommitNbFromBackService,
    private gongsvgPushFromFrontNbService: gongsvg.PushFromFrontNbService,
  ) { }

  ngOnInit(): void {

    this.checkCommiNbFromBagetCommitNbFromBackTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.gongsvgNbFromBackService.getCommitNbFromBack().subscribe(
          commiNbFromBagetCommitNbFromBack => {
            if (this.lastCommiNbFromBagetCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

              console.log("last commit nb " + this.lastCommiNbFromBagetCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
              this.refresh()
              this.lastCommiNbFromBagetCommitNbFromBack = commiNbFromBagetCommitNbFromBack
            }
          }
        )

        // see above for the explanation
        this.gongsvgPushFromFrontNbService.getPushFromFrontNb().subscribe(
          pushFromFrontNb => {
            if (this.lastPushFromFrontNb < pushFromFrontNb) {

              console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
              this.refresh()
              this.lastPushFromFrontNb = pushFromFrontNb
            }
          }
        )
      }
    )
  }

  refresh(): void {

    this.gongsvgFrontRepoService.pull().subscribe(
      gongsvgsFrontRepo => {
        this.gongsvgFrontRepo = gongsvgsFrontRepo

        this.svgSingloton = this.gongsvgFrontRepo.SVGs_array[0]
        console.log("svgSingloton " + this.svgSingloton?.Name)

      }

    )
  }

  fillColor = 'rgb(255, 0, 0)';
  Text = "Toto"
  width = 150
  height = 200
  changeColor() {
    const r = Math.floor(Math.random() * 256);
    const g = Math.floor(Math.random() * 256);
    const b = Math.floor(Math.random() * 256);
    this.fillColor = `rgb(${r}, ${g}, ${b})`;
  }

}
