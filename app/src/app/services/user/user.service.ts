import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/interfaces/user';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) {}

  public getUser(id: string): Observable<{data: User}> {
    return this.http.get<{data: User}>(`${environment.apiUrl}/user/${id}`, {withCredentials: true})
  }
}
