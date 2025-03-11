import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongganttSpecificComponent } from './gonggantt-specific.component';

describe('GongganttSpecificComponent', () => {
  let component: GongganttSpecificComponent;
  let fixture: ComponentFixture<GongganttSpecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongganttSpecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongganttSpecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
