import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { catchError, map, of, switchMap } from 'rxjs';
import { AuthService } from 'src/app/services/auth/auth.service';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';

export const authGuard: CanActivateFn = (route, state) => {
  const router: Router = inject(Router)
  const authService: AuthService = inject(AuthService)
  const genCSRFService: GenCsrfService = inject(GenCsrfService)
  return genCSRFService.getCSRF().pipe(
    switchMap(() => {
      return authService.refreshToken().pipe(
        map(() => true),
        catchError(() => {
          router.navigate(["/login"])
          return of(false)
        })
      )
    })
  )
};
