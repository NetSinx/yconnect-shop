import { inject } from '@angular/core';
import { CanActivateFn } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';

export const loginGuard: CanActivateFn = (route, state) => {
  const authService: AuthService = inject(AuthService)
  return authService.isAuthenticated
};
