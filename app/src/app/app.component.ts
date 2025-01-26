import { Component, OnInit } from '@angular/core';
import { ProductService } from './services/product/product.service';
import { filter, Observable } from 'rxjs';
import { LoadingService } from './services/loading/loading.service';
import { CategoryService } from './services/category/category.service';
import { NavigationEnd, Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})

export class AppComponent implements OnInit {
  errors: any
  isLoading: Observable<boolean>
  showNavbar: boolean = true

  constructor(
    private productService: ProductService,
    private categoryService: CategoryService,
    private loadingService: LoadingService,
    private router: Router
  ) {
    this.isLoading = this.loadingService.loading
  }

  ngOnInit(): void {
    this.loadingService.setLoading(true)
    this.getErrorCategories()
    this.getErrorProducts()
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        if (event.urlAfterRedirects === '/login' || event.urlAfterRedirects === '/register') {
          this.showNavbar = false
        } else {
          this.showNavbar = true
        }
      }
    })
  }

  public getErrorProducts(): void {
    this.productService.getProducts().subscribe(
      () => this.loadingService.setLoading(false),
      error => {
        this.loadingService.setLoading(false)
        this.errors = error
      }
    )
  }

  public getErrorCategories(): void {
    this.categoryService.getCategories().subscribe(
      () => this.loadingService.setLoading(false),
      error => {
        this.loadingService.setLoading(false)
        this.errors = error
      }
    )
  }
}