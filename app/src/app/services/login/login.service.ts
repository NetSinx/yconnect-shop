import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.development';
import { GenCsrfService } from '../gen-csrf/gen-csrf.service';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private httpClient: HttpClient, private csrfService: GenCsrfService) {}

  public userLogin(data: any): Observable<any> {
    return this.httpClient.post<{token: string}>(`${environment.apiUrl}/user/sign-in`, data, {headers: {'xsrf': this.csrfService.csrfToken!}, withCredentials: true})
  }

  public verifyUser(): Observable<any> {
    return this.httpClient.get(`${environment.apiUrl}/user/verify`, {withCredentials: true})
  }
}
