import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { catchError, map, of } from 'rxjs';
import { AuthService } from 'src/app/services/auth/auth.service';

export const authGuard: CanActivateFn = (route, state) => {
  const router: Router = inject(Router)
  const authService: AuthService = inject(AuthService)
  return authService.refreshToken().pipe(
    map((isAuth) => {
      if (isAuth) return true;
      router.navigate(['/forbidden']);
      return false;
    }),
    catchError(() => {
      router.navigate(['/forbidden']);
      return of(false);
    })
  );
};
