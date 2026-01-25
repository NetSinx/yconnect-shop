import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth/auth.service';
import { LoadingService } from 'src/app/services/loading/loading.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  standalone: false
})
export class LoginComponent implements OnInit {
  formGroup: FormGroup;
  formBuilder: FormBuilder = new FormBuilder();
  errorLogin: string | undefined;
  successMessage: string = '';

  constructor(
    private router: Router,
    private loadingService: LoadingService,
    private authService: AuthService
  ) {
    this.formGroup = this.formBuilder.group({
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    });
  }

  ngOnInit(): void {
    if (history.state && history.state.success) {
      this.successMessage = history.state.success;
    }

    if (this.successMessage) {
      history.replaceState({}, '');
    }
  }

  get email() {
    return this.formGroup.get('email')!;
  }
  get password() {
    return this.formGroup.get('password')!;
  }

  public userLogin(): void {
    let dataLogin = {
      email: '',
      password: ''
    };

    dataLogin.email = this.formGroup.value.email;
    dataLogin.password = this.formGroup.value.password;

    this.authService.loginUser(dataLogin).subscribe(
      () => {
        this.authService.verifyUser().subscribe(
          () => {
            this.loadingService.setLoading(false);
            this.router.navigate(['/dashboard']);
          },
          () => {
            this.loadingService.setLoading(false);
            this.router.navigate(['/forbidden']);
          }
        );
      },
      () => {
        this.loadingService.setLoading(false);
        this.errorLogin = 'Email / password Anda salah!';
      }
    );
  }
}
