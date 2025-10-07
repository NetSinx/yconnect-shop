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
  formGroup: FormGroup
  formBuilder: FormBuilder = new FormBuilder()
  errorLogin: string | undefined
  successMessage: string = ""

  constructor(
    private loginService: LoginService,
    private router: Router,
    private loadingService: LoadingService,
    private userService: UserService,
    private csrfService: GenCsrfService
  ) {
    this.formGroup = this.formBuilder.group({
      email: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required]),
    })
  }

  ngOnInit(): void {
    this.csrfService.getCSRF().subscribe()
    if (history.state && history.state.success) {
      this.successMessage = history.state.success
    }

    if (this.successMessage) {
      history.replaceState({}, "")
    }
  }

  get email() { return this.formGroup.get("email")! }
  get password() { return this.formGroup.get("password")! }

  public userLogin(): void {
    let dataLogin = {
      email: "",
      password: ""
    }

    dataLogin.email = this.formGroup.value.email
    dataLogin.password = this.formGroup.value.password

    this.loginService.userLogin(dataLogin).subscribe(
      () => {
        this.router.navigate(["/dashboard/"])
      }, () => {
        this.loadingService.setLoading(false)
        this.errorLogin = "Email / password Anda salah!"
      }
    )
  }
}
