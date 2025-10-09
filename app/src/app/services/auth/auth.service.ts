import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, map, Observable, of } from 'rxjs';
import { UserLogin } from 'src/app/interfaces/user-login';
import { GenCsrfService } from '../gen-csrf/gen-csrf.service';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  isAuthenticated: boolean = false
  accessToken: string | undefined

  constructor(private http: HttpClient, private genCSRFService: GenCsrfService) { }

  public loginUser(requestLogin: UserLogin): Observable<any> {
    return this.http.post<UserLogin>("http://localhost:8086/api/auth/login", requestLogin, {
      headers: {
        "X-CSRF-Token": this.genCSRFService.csrfToken
      },
      withCredentials: true
    })
  }

  public verifyUser(access_token: string): Observable<any> {
    return this.http.get("http://localhost:8086/api/auth/verify", {
      headers: { "Authorization": `Bearer ${access_token}` }
    })
  }

  public refreshToken(): Observable<boolean> {
    return this.http.post<{ access_token: string }>(
      "http://localhost:8086/api/auth/refresh",
      null,
      {
        headers: { "X-CSRF-Token": this.genCSRFService.csrfToken },
        withCredentials: true,
      }
    ).pipe(
      map((resp) => {
        this.isAuthenticated = true;
        this.accessToken = resp.access_token;
        return true;
      }),
      catchError((err) => {
        this.isAuthenticated = false;
        this.accessToken = '';
        return of(false);
      })
    );
  }
}
