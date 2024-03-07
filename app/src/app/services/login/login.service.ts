import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private httpClient: HttpClient) {}

  public csrfToken(): Observable<any> {
    return this.httpClient.get(`${environment.apiUrl}/user/csrf`)
  }

  public userLogin(): void {
    // this.httpClient.post<User>(`${environment.apiUrl}/users/sign-up`, )
  }
}
