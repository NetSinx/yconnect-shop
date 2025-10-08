import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from 'src/app/services/auth/auth.service';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const authToken: Observable<any> = inject(AuthService).refreshToken()
  const reqWithHeaders = req.clone({
    headers: req.headers.set("Authorization", `Bearer ${authToken}`)
  })
  
  return next(reqWithHeaders);
};
