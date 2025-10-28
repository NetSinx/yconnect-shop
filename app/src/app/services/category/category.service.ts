import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Kategori } from 'src/app/interfaces/category';
import { environment } from 'src/environments/environment.development';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  constructor(private http: HttpClient) {}

  public getCategories(): Observable<{data: Kategori[]}> {
    return this.http.get<{data: Kategori[]}>(`${environment.API_URL}/category`)
  }
}
