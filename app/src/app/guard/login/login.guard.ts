import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { LoginService } from 'src/app/services/login/login.service';
import { UserService } from 'src/app/services/user/user.service';

export const loginGuard: CanActivateFn = (route, state) => {
  const router: Router = inject(Router)
  const loginService: LoginService = inject(LoginService)
  const userService: UserService = inject(UserService)
  return loginService.verifyUser().toPromise()
  .then(() => 
    userService.getUserInfo().toPromise().then(
      resp => {
        router.navigate([`/dashboard/${resp!.data.user_id}`])
        return false
      }
    ), () => true)
};
