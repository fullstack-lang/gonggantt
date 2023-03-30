import { TestBed } from '@angular/core/testing';

import { GongganttService } from './gonggantt.service';

describe('GongganttService', () => {
  let service: GongganttService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongganttService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
