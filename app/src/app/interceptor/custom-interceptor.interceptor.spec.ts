import { TestBed } from '@angular/core/testing';

import { CustomInterceptorInterceptor } from './custom-interceptor.interceptor';

describe('CustomInterceptorInterceptor', () => {
  beforeEach(() => TestBed.configureTestingModule({
    providers: [
      CustomInterceptorInterceptor
      ]
  }));

  it('should be created', () => {
    const interceptor: CustomInterceptorInterceptor = TestBed.inject(CustomInterceptorInterceptor);
    expect(interceptor).toBeTruthy();
  });
});
