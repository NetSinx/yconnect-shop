import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { finalize } from 'rxjs';
import { LoadingService } from 'src/app/services/loading/loading.service';

export const loadingInterceptor: HttpInterceptorFn = (req, next) => {
  const loadingService: LoadingService = inject(LoadingService)
  loadingService.setLoading(true)

  return next(req).pipe(
    finalize(() => {
      loadingService.setLoading(false)
    })
  );
};
