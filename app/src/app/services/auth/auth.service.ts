import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/interfaces/user';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  isAuthenticated: boolean = false

  constructor(private http: HttpClient) {}

  public verifyUser(): Observable<User> {
    return this.http.get<User>("http://localhost:8086/api/auth/verify")
  }
}
