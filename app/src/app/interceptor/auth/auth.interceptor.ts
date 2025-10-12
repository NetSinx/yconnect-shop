import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { AuthService } from 'src/app/services/auth/auth.service';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const authService: AuthService = inject(AuthService)
  const newReq = req.clone({
    setHeaders: {
      "Authorization": `Bearer ${authService.accessToken}`
    },
    withCredentials: true
  })

  return next(newReq)
};
