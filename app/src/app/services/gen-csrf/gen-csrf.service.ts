import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap, throwError } from 'rxjs';
import { environment } from 'src/environments/environment.development';
import { LoadingService } from '../loading/loading.service';

@Injectable({
  providedIn: 'root'
})
export class GenCsrfService {
  public csrfToken: string | undefined

  constructor(private http: HttpClient, private loadingService: LoadingService) {}

  private handleError(err: HttpErrorResponse) {
    if (err.status === 0) {
      console.error("Server sedang tidak terhubung")
    } else {
      console.error(`Server mengalami error: ${err.error}`)
    }

    return throwError(() => new Error("Ada kesalahan pada server. Mohon refresh halaman."))
  }

  public getCSRF(): Observable<{csrf_token: string}> {
    return this.http.get<{csrf_token: string}>(`${environment.apiUrl}/gencsrf`, {withCredentials: true}).pipe(
      catchError(this.handleError),
      tap(resp => {
        this.csrfToken = resp.csrf_token
      })
    )
  }
}
