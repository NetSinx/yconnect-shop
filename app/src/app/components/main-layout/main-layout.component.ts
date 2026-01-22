import { Component, inject } from '@angular/core';
import { Kategori } from '../../interfaces/category';
import { LayoutService } from '../../services/layout/layout.service';
import { Router, NavigationStart, Event } from '@angular/router';

@Component({
  selector: 'app-main-layout',
  templateUrl: './main-layout.component.html',
  styleUrl: './main-layout.component.css',
  standalone: false
})
export class MainLayoutComponent {
  layoutService: LayoutService = inject(LayoutService);
  categories: Kategori[];
  router: Router = inject(Router);

  constructor() {
    this.router.events.subscribe((event: Event) => {
      if (event instanceof NavigationStart) {
        this.layoutService.sidebarOpen.set(false);
      }
    });

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
}
