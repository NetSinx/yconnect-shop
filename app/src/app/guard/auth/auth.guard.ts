import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';

export const authGuard: CanActivateFn = (route, state) => {
  const router: Router = inject(Router)
  const authService: AuthService = inject(AuthService)
  authService.refreshToken()
  if (authService.accessToken) {
    console.log(authService.accessToken)
    return true
  } else {
    console.log(authService.accessToken)
    return false
  }
};
