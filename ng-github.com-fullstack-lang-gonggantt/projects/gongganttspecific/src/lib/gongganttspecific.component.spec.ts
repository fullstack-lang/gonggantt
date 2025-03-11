import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongganttspecificComponent } from './gongganttspecific.component';

describe('GongganttspecificComponent', () => {
  let component: GongganttspecificComponent;
  let fixture: ComponentFixture<GongganttspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongganttspecificComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongganttspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
