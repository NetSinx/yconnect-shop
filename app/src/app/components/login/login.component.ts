import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { LoginService } from 'src/app/services/login/login.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  standalone: false
})
export class LoginComponent implements OnInit {
  formGroup: FormGroup = new FormGroup({})
  formBuilder: FormBuilder = new FormBuilder
  errorLogin: string | undefined
  successMessage: string = ""
  dataLogin: {
    UsernameorEmail: string,
    password: string
  } = {
      UsernameorEmail: "",
      password: ""
    }

  constructor(
    private loginService: LoginService,
    private router: Router,
    private loadingService: LoadingService,
    private userService: UserService,
    private csrfService: GenCsrfService
  ) {
    this.formGroup = this.formBuilder.group({
      UsernameorEmail: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required]),
    })
  }

  ngOnInit(): void {
    this.csrfService.getCSRF().subscribe()
    if (history.state.success) {
      this.successMessage = history.state.success
    }
  }

  get f() { return this.formGroup.controls }

  public userLogin(): void {
    this.dataLogin.UsernameorEmail = this.formGroup.value.UsernameorEmail
    this.dataLogin.password = this.formGroup.value.password
    const timezone: string = Intl.DateTimeFormat().resolvedOptions().timeZone

    this.loginService.userLogin(this.dataLogin).subscribe(() => {
      this.userService.setTimeZone(timezone).subscribe(
        () => {
          this.userService.getUserInfo().subscribe(resp => {
            this.loadingService.setLoading(false)
            this.router.navigate([`/dashboard/${resp.data.user_id}`])
          })
        }
      )
    }, () => {
      this.loadingService.setLoading(false)
      this.errorLogin = "Email / password Anda salah!"
    })
  }
}
