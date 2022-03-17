import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  
  view = 'Gantt view'
  
  gantt = 'Gantt view'
  data_gantt = 'Data gantt view'

  meta = 'Meta view'
  
  data_svg = 'Data svg view'

  uml_view = 'Uml view'
  
  views: string[] = [this.gantt, this.data_gantt, this.data_svg, this.uml_view, this.meta];


}
