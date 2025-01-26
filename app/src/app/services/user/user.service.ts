import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/interfaces/user';
import { environment } from 'src/environments/environment.development';
import { GenCsrfService } from '../gen-csrf/gen-csrf.service';
import { UserInfo } from 'src/app/interfaces/user_info';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient, private csrfService: GenCsrfService) {}

  public getUser(id: string): Observable<{data: User}> {
    return this.http.get<{data: User}>(`${environment.API_URL}/user/${id}`, {withCredentials: true})
  }

  public getUserInfo(): Observable<UserInfo> {
    return this.http.get<UserInfo>(`${environment.API_URL}/user/userinfo`, {withCredentials: true})
  }

  public setTimeZone(timezone: string): Observable<any> {
    return this.http.post(`${environment.API_URL}/user/set-timezone`, {timezone}, {withCredentials: true, headers: {"xsrf": this.csrfService.csrfToken!}})
  }

  public userLogout(): Observable<any> {
    return this.http.get(`${environment.API_URL}/user/logout`, {withCredentials: true})
  }
}
