import { NgModule } from '@angular/core';

import { GongganttspecificComponent } from './gongganttspecific.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongganttModule } from 'gonggantt'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

@NgModule({
  declarations: [
    GongganttspecificComponent,
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongganttModule,
  ],
  exports: [
    GongganttspecificComponent,
    DataModelPanelComponent
  ]
})
export class GongganttspecificModule { }
