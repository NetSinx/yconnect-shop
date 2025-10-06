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
    return this.httpClient.post<{token: string}>(`http://localhost:8086/api/auth/login`, data, {headers: {'X-CSRF-Token': this.csrfService.csrfToken}, withCredentials: true})
  }

  public verifyUser(): Observable<any> {
    return this.httpClient.get(`${environment.API_URL}/user/verify`, {withCredentials: true})
  }
}
