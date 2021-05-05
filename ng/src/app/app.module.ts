import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { GongganttModule } from 'gonggantt'

import { AngularSplitModule } from 'angular-split';

import { GongsvgspecificModule } from 'gongsvgspecific'
import { GongsvgModule } from 'gongsvg'

// mandatory
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    AngularSplitModule,

    GongsvgspecificModule,
    GongsvgModule,

    HttpClientModule,
    GongganttModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
