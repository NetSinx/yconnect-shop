import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { LoginService } from 'src/app/services/login/login.service';

export const loginGuard: CanActivateFn = (route, state) => {
  const router: Router = inject(Router)
  const loginService: LoginService = inject(LoginService)
  return loginService.verifyUser().toPromise()
  .then(() => {
    let cookie_userId: string = document.cookie.split(';').find(
      row => row.startsWith("user_id=")
    )!.split('=')[1]
    
    router.navigate([`/dashboard/${cookie_userId}`])
    return false
  }, () => true)
};
