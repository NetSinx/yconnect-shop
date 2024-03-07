import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor,
  HttpXsrfTokenExtractor
} from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class CustomInterceptorInterceptor implements HttpInterceptor {

  constructor(private tokenExtractor: HttpXsrfTokenExtractor) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const cookieHeaderName = "XSRF-Token";
    let csrfToken = this.tokenExtractor.getToken()!

    if (csrfToken !== null && !request.headers.has(cookieHeaderName)) {
      request = request.clone({headers: request.headers.set(cookieHeaderName, csrfToken)})
    }

    return next.handle(request);
  }
}
