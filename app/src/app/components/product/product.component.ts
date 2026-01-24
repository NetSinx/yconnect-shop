import { Component, input, OnInit, signal, inject } from '@angular/core';
import { NgOptimizedImage } from '@angular/common';
import { ProductService } from '../../services/product/product.service';
import { Product } from 'src/app/interfaces/product';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { Kategori } from 'src/app/interfaces/category';
import { LayoutService } from 'src/app/services/layout/layout.service';
import { RouterModule, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css'],
  standalone: false
})
export class ProductComponent implements OnInit {
  products: Product[] = [];
  activeSidebar = signal<string>('');
  categorySidebar = input.required<Kategori[]>();
  layoutService = inject(LayoutService);
  route = inject(ActivatedRoute);
  displayProducts: any[] = [];

  constructor(
    private productService: ProductService,
    private loadingService: LoadingService
  ) {
    this.products = [
      {
        id: 1,
        nama: 'Baju Muslim Keren & Kekinian',
        images: 'assets/img/baju_muslim1.jpg',
        slug: 'baju-muslim-keren-kekinian',
        deskripsi: 'Baju muslim keren dan kekinian',
        kategori: 'pakaian',
        harga: 205000,
        stok: 25,
        rating: 5.0
      },
      {
        id: 2,
        nama: 'Baju Koko Pria Beragam Ukuran S/M/L/XL/XXL',
        images: 'assets/img/baju_muslim2.jpg',
        slug: 'baju-koko-pria-beragam-ukuran',
        deskripsi: 'Baju koko pria beragam ukuran',
        kategori: 'pakaian',
        harga: 170000,
        stok: 12,
        rating: 5.0
      },
      {
        id: 3,
        nama: 'Koko Pria Murah Berkualitas',
        images: 'assets/img/baju_muslim3.jpg',
        slug: 'koko-pria-murah-berkualitas',
        deskripsi: 'Baju koko pria murah dan berkualitas',
        kategori: 'pakaian',
        harga: 250000,
        stok: 8,
        rating: 5.0
      },
      {
        id: 4,
        nama: 'Koko Pria Berbagai Ukuran',
        images: 'assets/img/baju_muslim4.jpg',
        slug: 'koko-pria-murah-berkualitas',
        deskripsi: 'Baju koko pria murah dan berkualitas',
        kategori: 'pakaian',
        harga: 150000,
        stok: 15,
        rating: 5.0
      },
      {
        id: 5,
        nama: 'Sepatu Adidas Samba',
        images: 'assets/img/sepatu_adidas_samba.jpeg',
        slug: 'sepatu-adidas-samba',
        deskripsi: 'Sepatu Adidas dengan kualitas ori',
        kategori: 'makanan',
        harga: 250000,
        stok: 8,
        rating: 5.0
      },
      {
        id: 6,
        nama: 'Keyboard RGB Apex Pro Mini',
        images: 'assets/img/keyboard_rgb_apex_pro_mini.jpeg',
        slug: 'keyboard-rgb-apex-pro-mini',
        deskripsi: 'Keyboard RGB kualitas mewah dan elegan',
        kategori: 'minuman',
        harga: 370000,
        stok: 24,
        rating: 5.0
      },
      {
        id: 7,
        nama: 'Air Pods Max Nouveau',
        images: 'assets/img/Air pods max nouveau.jpeg',
        slug: 'air-pods-max-nouveau',
        deskripsi: 'Air Pods Max Nouveau',
        kategori: 'makanan',
        harga: 32000,
        stok: 55,
        rating: 4.8
      },
      {
        id: 8,
        nama: 'Nikon Camera',
        images: 'assets/img/nikon-camera.jpeg',
        slug: 'nikon-camera',
        deskripsi: 'Nikon camera',
        kategori: 'minuman',
        harga: 6700000,
        stok: 55,
        rating: 4.8
      },
      {
        id: 9,
        nama: 'Silicone Case Apple Iphone 13',
        images: 'assets/img/silicone-case-apple-iphone-13.jpeg',
        slug: 'silicon-case-apple-iphone-13',
        deskripsi: 'Silicone Case Apple Iphone 13',
        kategori: 'pakaian',
        harga: 25000,
        stok: 55,
        rating: 4.8
      },
      {
        id: 10,
        nama: 'Nike Sneaker Dunk',
        images: 'assets/img/nike-sneaker-dunk.jpeg',
        slug: 'nike-sneaker-dunk',
        deskripsi: 'Nike Sneaker Dunk',
        kategori: 'pakaian',
        harga: 350000,
        stok: 100,
        rating: 4.8
      },
      {
        id: 11,
        nama: 'Wired Mouse',
        images: 'assets/img/wired-mouse.jpeg',
        slug: 'wired-mouse',
        deskripsi: 'Wired Mouse',
        kategori: 'pakaian',
        harga: 17000,
        stok: 10,
        rating: 5.0
      },
      {
        id: 12,
        nama: 'Lenovo Thinkpad X270 Intel Core i5-7300U',
        images: 'assets/img/lenovo-thinkpax-x270.jpeg',
        slug: 'lenovo-thinkpad-x270-intel-core-i5-7300u',
        deskripsi: 'Lenovo Thinkpad X270',
        kategori: 'pakaian',
        harga: 1500000,
        stok: 8,
        rating: 5.0
      },
      {
        id: 13,
        nama: 'Kompor Portable',
        images: 'assets/img/kompor-portable.jpeg',
        slug: 'kompor-portable',
        deskripsi: 'Kompor Portable',
        kategori: 'makanan',
        harga: 17000,
        stok: 10,
        rating: 5.0
      }
    ];
  }

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      const categorySlug = params['kategori']; // Ambil ?category=...

      if (categorySlug) {
        this.displayProducts = this.products.filter(
          product => product.kategori.toLowerCase() === categorySlug.toLowerCase()
        );
      } else {
        this.displayProducts = this.products;
      }
    });
  }

  public getProducts(): void {
    this.productService.getProducts().subscribe(resp => {
      this.loadingService.setLoading(false);
      this.products = resp.data;
    });
  }

  public scrollToProduct(event: Event): void {
    event.target?.addEventListener('click', () => {
      window.scrollTo({ top: 890, behavior: 'smooth' });
    });
  }

  public scrollToFlashSale(event: Event): void {
    event.target?.addEventListener('click', () => {
      window.scrollTo({ top: 1490, behavior: 'smooth' });
    });
  }

  public scrollToDiskon(event: Event): void {
    event.target?.addEventListener('click', () => {
      window.scrollTo({ top: 2520, behavior: 'smooth' });
    });
  }

  public setActiveSidebar(event: Event, sidebar: string): void {
    console.log(`clicked is: ${sidebar}`);
    event.preventDefault();
    this.activeSidebar.set(sidebar);
  }

  public ratingNumber(n: number): number[] {
    return Array.from({ length: n }, (_, i) => i);
  }
}
