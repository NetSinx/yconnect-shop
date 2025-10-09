import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { UserLogin } from 'src/app/interfaces/user-login';
import { GenCsrfService } from '../gen-csrf/gen-csrf.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  accessToken: string | undefined

  constructor(private http: HttpClient, private genCSRFService: GenCsrfService) {}

  public loginUser(requestLogin: UserLogin): Observable<any> {
    return this.http.post<UserLogin>("http://localhost:8086/api/auth/login", requestLogin, {
      headers: {
        "X-CSRF-Token": this.genCSRFService.csrfToken
      },
      withCredentials: true
    })
  }

  public verifyUser(): Observable<any> {
    return this.http.get("http://localhost:8086/api/auth/verify", {
      headers: { "Authorization": `Bearer ${this.accessToken}` }
    })
  }

  public refreshToken(): any {
    this.http.post<{access_token: string}>("http://localhost:8086/api/auth/refresh", null).subscribe(
      resp => this.accessToken = resp.access_token
    )
  }
}
