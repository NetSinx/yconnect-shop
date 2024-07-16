import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { LoginService } from 'src/app/services/login/login.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  formGroup: FormGroup = new FormGroup({})
  formBuilder: FormBuilder = new FormBuilder
  errorLogin: string | undefined
  isLoading: Observable<boolean>
  dataLogin: {
    UsernameorEmail: string,
    password: string
  } = {
    UsernameorEmail: "",
    password: ""
  }

  constructor(
    private loginService: LoginService,
    private csrfService: GenCsrfService,
    private router: Router,
    private loadingService: LoadingService,
    private userService: UserService
  ) {
    this.formGroup = this.formBuilder.group({
      UsernameorEmail: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required]),
    })
    
    this.isLoading = this.loadingService.loading
  }

  ngOnInit(): void {
    this.csrfService.getCSRF().subscribe()
  }

  get f() {return this.formGroup.controls}

  public userLogin(): void {
    this.dataLogin.UsernameorEmail = this.formGroup.value.UsernameorEmail
    this.dataLogin.password = this.formGroup.value.password

    this.loginService.userLogin(this.dataLogin).subscribe(() => {
      this.userService.getUser(this.dataLogin.UsernameorEmail).subscribe(
        resp => {
          this.router.navigate([`/dashboard/${resp.data.username}`])
          document.cookie = `user_id=${resp.data.username}`
        }
      )
    }, () => {
      this.errorLogin = "Email / password Anda salah!"
    })
  }
}
