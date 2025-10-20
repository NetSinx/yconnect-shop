import { Component, input, OnInit, signal } from '@angular/core';
import { ProductService } from '../../services/product/product.service';
import { Product } from 'src/app/interfaces/product';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { Kategori } from 'src/app/interfaces/category';

@Component({
    selector: 'app-product',
    templateUrl: './product.component.html',
    styleUrls: ['./product.component.css'],
    standalone: false
})

export class ProductComponent implements OnInit {
  promoProducts: Product[] = []
  flashSaleProducts: Product[] = []
  activeSidebar = signal<string>('')
  categorySidebar = input.required<Kategori[]>()

  constructor(private productService: ProductService, private loadingService: LoadingService) {
    this.promoProducts = [
      {
        id: 1,
        nama: "Baju Muslim Keren & Kekinian",
        images: "assets/img/baju_muslim1.jpg",
        slug: 'baju-muslim-keren-kekinian',
        deskripsi: 'Baju muslim keren dan kekinian',
        kategori_id: 1,
        harga: 205000,
        stok: 25,
        rating: 5.0
      },
      {
        id: 2,
        nama: "Baju Koko Pria Beragam Ukuran S/M/L/XL/XXL",
        images: "assets/img/baju_muslim2.jpg",
        slug: 'baju-koko-pria-beragam-ukuran',
        deskripsi: 'Baju koko pria beragam ukuran',
        kategori_id: 1,
        harga: 170000,
        stok: 12,
        rating: 5.0
      },
      {
        id: 3,
        nama: "Koko Pria Murah Berkualitas",
        images: "assets/img/baju_muslim3.jpg",
        slug: 'koko-pria-murah-berkualitas',
        deskripsi: 'Baju koko pria murah dan berkualitas',
        kategori_id: 1,
        harga: 250000,
        stok: 8,
        rating: 5.0
      }
    ]

    this.flashSaleProducts = [
      {
        id: 1,
        nama: "Sepatu Adidas Samba",
        images: "assets/img/sepatu_adidas_samba.jpeg",
        slug: 'sepatu-adidas-samba',
        deskripsi: 'Sepatu Adidas dengan kualitas ori',
        kategori_id: 2,
        harga: 250000,
        stok: 8,
        rating: 5.0
      },
      {
        id: 2,
        nama: "Keyboard RGB Apex Pro Mini",
        images: "assets/img/keyboard_rgb_apex_pro_mini.jpeg",
        slug: 'keyboard-rgb-apex-pro-mini',
        deskripsi: 'Keyboard RGB kualitas mewah dan elegan',
        kategori_id: 3,
        harga: 370000,
        stok: 24,
        rating: 5.0
      },
      {
        id: 3,
        nama: "Air Pods Max Nouveau",
        images: "assets/img/Air pods max nouveau.jpeg",
        slug: 'air-pods-max-nouveau',
        deskripsi: 'Air Pods Max Nouveau',
        kategori_id: 3,
        harga: 32000,
        stok: 55,
        rating: 4.8
      },
      {
        id: 4,
        nama: "Air Pods Max Nouveau",
        images: "assets/img/Air pods max nouveau.jpeg",
        slug: 'air-pods-max-nouveau',
        deskripsi: 'Air Pods Max Nouveau',
        kategori_id: 3,
        harga: 32000,
        stok: 55,
        rating: 4.8
      },
      {
        id: 5,
        nama: "Air Pods Max Nouveau",
        images: "assets/img/Air pods max nouveau.jpeg",
        slug: 'air-pods-max-nouveau',
        deskripsi: 'Air Pods Max Nouveau',
        kategori_id: 3,
        harga: 32000,
        stok: 55,
        rating: 4.8
      },
      {
        id: 6,
        nama: "Air Pods Max Nouveau",
        images: "assets/img/Air pods max nouveau.jpeg",
        slug: 'air-pods-max-nouveau',
        deskripsi: 'Air Pods Max Nouveau',
        kategori_id: 3,
        harga: 32000,
        stok: 55,
        rating: 4.8
      },
    ]
  }
  
  ngOnInit(): void {

  }
  
  public getProducts(): void {
    this.productService.getProducts().subscribe(
      resp => {
        this.loadingService.setLoading(false)
        this.promoProducts = resp.data
      }
    )
  }

  public scrollToProduct(event: Event): void {
    event.target?.addEventListener('click', () => {
      window.scrollTo({top: 890, behavior: 'smooth'})
    })
  }

  public scrollToFlashSale(event: Event): void {
    event.target?.addEventListener('click', () => {
      window.scrollTo({top: 1490, behavior: 'smooth'})
    })
  }

  public scrollToDiskon(event: Event): void {
    event.target?.addEventListener('click', () => {
      window.scrollTo({top: 2520, behavior: 'smooth'})
    })
  }

  public setActiveSidebar(event: Event, sidebar: string): void {
    console.log(`clicked is: ${sidebar}`)
    event.preventDefault()
    this.activeSidebar.set(sidebar)
  }

  public ratingNumber(n: number): number[] {
    return Array.from({ length: n }, (_, i) => i)
  }
}
