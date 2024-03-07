import { Component, OnInit } from '@angular/core';
import { LoginService } from 'src/app/services/login/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  currentRoute = '/login'
  csrf: any

  ngOnInit(): void {
    this.getCSRF()
  }

  constructor(private loginService: LoginService) {}

  public getCSRF(): void {
    this.loginService.csrfToken().subscribe(
      data => this.csrf = data.csrf_token
    )
  }
}
