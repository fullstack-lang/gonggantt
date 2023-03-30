import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongganttComponent } from './gonggantt.component';

describe('GongganttComponent', () => {
  let component: GongganttComponent;
  let fixture: ComponentFixture<GongganttComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongganttComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongganttComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
