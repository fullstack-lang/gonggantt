import { TestBed } from '@angular/core/testing';

import { GongganttdatamodelService } from './gongganttdatamodel.service';

describe('GongganttdatamodelService', () => {
  let service: GongganttdatamodelService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongganttdatamodelService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
