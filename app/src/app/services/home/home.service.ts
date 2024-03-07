import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})

export class HomeService {
  constructor(private http: HttpClient) {}

  public getProducts(): Observable<any> {
    return this.http.get(`${environment.apiUrl}/product`)
  }

  public showDetailProduct(slug: string): Observable<any> {
    return this.http.get(`${environment.apiUrl}/product/${slug}`)
  }

  public getCategories(): Observable<any> {
    return this.http.get(`${environment.apiUrl}/category`)
  }

  public getImages(name: string): Promise<Object | undefined> {
    return this.http.get(`${environment.apiUrl}${name}`).toPromise()
  }

  public getCarts(): Object {
    return this.http.get(`${environment.apiUrl}/cart`).subscribe(data => {
      let obj: Object = data

      return obj
    })
  }
}
