import { Component, OnInit, Input } from '@angular/core';

import * as gongsvg from 'gongsvg'

@Component({
  selector: 'lib-rect',
  templateUrl: './rect.component.svg',
  styleUrls: ['./rect.component.css']
})
export class RectComponent implements OnInit {

  @Input() Rect?: gongsvg.RectDB
  @Input() GONG__StackPath: string = ""

  // In your component
  anchorRadius = 8; // Adjust this value according to your desired anchor size
  anchorFillColor = 'blue'; // Choose your desired anchor fill color
  draggingAnchorFillColor = 'red'; // Change this to the desired color when dragging

  constructor(
    private rectService: gongsvg.RectService) { }

  ngOnInit(): void {
  }

  onClick(event: MouseEvent) {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    console.log("rect, onClick() : ", this.Rect?.Name)
    if (this.Rect?.IsSelectable) {
      this.Rect!.IsSelected = !this.Rect!.IsSelected
      this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()
    }
  }

  dragging: boolean = false;
  activeAnchor: 'left' | 'right' | null = null;

  startDrag(event: MouseEvent, anchor: 'left' | 'right'): void {
    event.preventDefault();
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    this.dragging = true;
    this.activeAnchor = anchor;
  }

  drag(event: MouseEvent, anchor: 'left' | 'right'): void {
    event.stopPropagation(); // Prevent the event from bubbling up to the SVG element

    if (!this.dragging) {
      return;
    }

    const newX = event.clientX;


    if (anchor === 'left') {
      const originalRightEdge = this.Rect!.X + this.Rect!.Width;
      this.Rect!.X = newX;
      this.Rect!.Width = originalRightEdge - newX;
    } else if (anchor === 'right') {
      this.Rect!.Width = newX - this.Rect!.X;
    }
  }

  endDrag(): void {
    this.dragging = false;
    this.activeAnchor = null;
    this.rectService.updateRect(this.Rect!, this.GONG__StackPath).subscribe()

  }
}
