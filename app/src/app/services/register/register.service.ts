import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/interfaces/user';
import { environment } from 'src/environments/environment.development';
import { GenCsrfService } from '../gen-csrf/gen-csrf.service';

@Injectable({
  providedIn: 'root'
})

export class RegisterService {

  constructor(private http: HttpClient, private csrfService: GenCsrfService) {}

  public registerUser(data: User): Observable<any> {
    if (!this.csrfService.csrfToken) {
      throw new Error('CSRF token is not available.');
    }

    return this.http.post(`${environment.apiUrl}/user/sign-up`, data, {headers: {'xsrf': this.csrfService.csrfToken}, withCredentials: true})
  }
}