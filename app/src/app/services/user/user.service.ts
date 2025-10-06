import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/interfaces/user';
import { environment } from 'src/environments/environment.development';
import { GenCsrfService } from '../gen-csrf/gen-csrf.service';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient, private csrfService: GenCsrfService) {}

  public getUser(id: string): Observable<{data: User}> {
    return this.http.get<{data: User}>(`http://localhost:8082/api/user/${id}`, {withCredentials: true})
  }

  public userLogout(): Observable<any> {
    return this.http.get(`${environment.API_URL}/user/logout`, {withCredentials: true})
  }
}
