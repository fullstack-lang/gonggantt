import { TestBed } from '@angular/core/testing';

import { GongganttspecificService } from './gongganttspecific.service';

describe('GongganttspecificService', () => {
  let service: GongganttspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongganttspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
