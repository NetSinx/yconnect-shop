import { Component, OnInit } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { LoadingService } from './services/loading/loading.service';
import { Observable } from 'rxjs';
import { CategoryService } from './services/category/category.service';
import { NavbarComponent } from './components/navbar/navbar.component';
import { Kategori } from './interfaces/category';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  standalone: false
})
export class AppComponent implements OnInit {
  errors: any;
  showNavbar: boolean = true;
  isLoading: Observable<boolean> = new Observable<boolean>();
  categories: Kategori[] = [];

  constructor(
    private router: Router,
    private loadingService: LoadingService,
    private categoryService: CategoryService
  ) {
    this.isLoading = this.loadingService.loading;
    this.categories = [
      {
        id: 1,
        nama: 'Pakaian',
        slug: 'pakaian',
        gambar: 'assets/img/clothes-rack.png'
      },
      {
        id: 2,
        nama: 'Makanan',
        slug: 'makanan',
        gambar: 'assets/img/food-category.png'
      },
      {
        id: 3,
        nama: 'Minuman',
        slug: 'minuman',
        gambar: 'assets/img/soft-drink.png'
      },
      {
        id: 4,
        nama: 'Detergen',
        slug: 'detergen',
        gambar: 'assets/img/laundry-detergent.png'
      }
    ];
  }

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        if (event.urlAfterRedirects === '/login' || event.urlAfterRedirects === '/register') {
          this.showNavbar = false;
        } else {
          this.showNavbar = true;
        }
      }
    });
  }
}
