import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/interfaces/user';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})

export class RegisterService {
  constructor(private http: HttpClient) {}

  public registerUser(data: any): Observable<User> {
    return this.http.post<User>(`${environment.apiUrl}/users/sign-up`, data)
  }
}
