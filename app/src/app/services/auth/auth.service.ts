import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { map, Observable, tap } from 'rxjs';
import { UserLogin } from 'src/app/interfaces/user-login';
import { GenCsrfService } from '../gen-csrf/gen-csrf.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  isAuthenticated: boolean = false
  accessToken: string = ""

  constructor(private http: HttpClient, private genCSRFService: GenCsrfService) { }

  public loginUser(requestLogin: UserLogin): Observable<any> {
    return this.http.post<{ access_token: string }>("http://localhost:8086/api/auth/login", requestLogin, {
      headers: {
        "X-CSRF-Token": this.genCSRFService.csrfToken
      },
      withCredentials: true
    }).pipe(
      tap(resp => {
        this.accessToken = resp.access_token
      })
    )
  }

  public verifyUser(): Observable<boolean> {
    return this.http.get("http://localhost:8086/api/auth/verify", {
      headers: {"Authorization": `Bearer ${this.accessToken}`}
    }).pipe(
      map(() => true),
      tap(
        () => this.isAuthenticated = true,
        () => this.isAuthenticated = false
      )
    )
  }

  public refreshToken(): Observable<string> {
    return this.http.post<{ auth_token: string }>("http://localhost:8086/api/auth/refresh", null, {
      headers: {
        "X-CSRF-Token": this.genCSRFService.csrfToken
      },
      withCredentials: true
    }).pipe(
      map(
        resp => resp.auth_token
      ),
      tap(token => this.accessToken = token)
    )
  }
}
