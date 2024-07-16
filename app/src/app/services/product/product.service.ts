import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Product } from 'src/app/interfaces/product';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class ProductService {
  constructor(private http: HttpClient) {}

  public getProducts(): Observable<{data: Product[]}> {
    return this.http.get<{data: Product[]}>(`${environment.apiUrl}/product`)
  }

  public getDetailProduct(slug: string): Observable<{data: Product}> {
    return this.http.get<{data: Product}>(`${environment.apiUrl}/product/${slug}`)
  }
}
