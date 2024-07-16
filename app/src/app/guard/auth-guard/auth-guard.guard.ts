import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { tap } from 'rxjs';
import { LoginService } from 'src/app/services/login/login.service';

export const authGuardGuard: CanActivateFn = (route, state) => {
  const router: Router = inject(Router)
  const loginService: LoginService = inject(LoginService)
  return loginService.verifyUser().toPromise().then(
    () => {
      return true;
    }, () => {
      router.navigate(['/login']);
      return false;
    }
  )
};
