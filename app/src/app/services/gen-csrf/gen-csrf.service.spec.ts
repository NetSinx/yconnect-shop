import { TestBed } from '@angular/core/testing';

import { GenCsrfService } from './gen-csrf.service';

describe('GenCsrfService', () => {
  let service: GenCsrfService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GenCsrfService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
