import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { switchMap, map, catchError, of } from 'rxjs';
import { AuthService } from 'src/app/services/auth/auth.service';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';

export const guestGuard: CanActivateFn = (route, state) => {
  const router: Router = inject(Router)
    const authService: AuthService = inject(AuthService)
    const genCSRFService: GenCsrfService = inject(GenCsrfService)
    return genCSRFService.getCSRF().pipe(
      switchMap(() => {
        return authService.refreshToken().pipe(
          map(() => {
            router.navigate(["/dashboard"])
            return false
          }),
          catchError(() => of(true))
        )
      }),
      catchError(() => of(true))
    )
};
