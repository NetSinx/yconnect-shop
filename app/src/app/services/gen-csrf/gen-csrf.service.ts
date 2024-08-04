import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, tap } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class GenCsrfService {
  csrfToken: string = ""

  constructor(private http: HttpClient) {}

  public getCSRF(): Observable<{csrf_token: string}> {
    return this.http.get<{csrf_token: string}>(`${environment.apiUrl}/gencsrf`, {withCredentials: true}).pipe(
      tap(
        resp => this.csrfToken = resp.csrf_token
      )
    )
  }
}
