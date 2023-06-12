import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongganttdatamodelComponent } from './gongganttdatamodel.component';

describe('GongganttdatamodelComponent', () => {
  let component: GongganttdatamodelComponent;
  let fixture: ComponentFixture<GongganttdatamodelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongganttdatamodelComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongganttdatamodelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
