import { Component, OnInit } from '@angular/core';
import { ProductService } from './services/product/product.service';
import { Observable } from 'rxjs';
import { LoadingService } from './services/loading/loading.service';
import { CategoryService } from './services/category/category.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent implements OnInit {
  errors: any
  isLoading: Observable<boolean>

  constructor(
    private productService: ProductService,
    private categoryService: CategoryService,
    private loadingService: LoadingService
  ) {
    this.isLoading = this.loadingService.loading
  }

  ngOnInit(): void {
    this.getErrorCategories()
    this.getErrorProducts()
  }

  public getErrorProducts(): void {
    this.productService.getProducts().subscribe(
      data => data,
      error => {
        this.errors = error
      }
    )
  }

  public getErrorCategories(): void {
    this.categoryService.getCategories().subscribe(
      data => data,
      error => {
        this.errors = error
      }
    )
  }
}