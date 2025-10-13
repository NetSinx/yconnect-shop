import { Component, OnInit } from '@angular/core';
import { ProductService } from '../../services/product/product.service';
import { Product } from 'src/app/interfaces/product';
import { LoadingService } from 'src/app/services/loading/loading.service';

@Component({
    selector: 'app-product',
    templateUrl: './product.component.html',
    styleUrls: ['./product.component.css'],
    standalone: false
})

export class ProductComponent implements OnInit {
  products: Product[] = []
  activeSidebar: string = ""

  constructor(private productService: ProductService, private loadingService: LoadingService) {
    this.products = [
      {
        id: 1,
        nama: "Baju Muslim Keren & Kekinian",
        images: "assets/img/baju_muslim1.jpg",
        slug: 'baju-muslim-keren-kekinian',
        deskripsi: 'Baju Muslim Keren & Kekinian',
        kategori_id: 1,
        harga: 17000,
        stok: 25,
        rating: 5.0
      },
      {
        id: 2,
        nama: "Baju Koko Pria Beragam Ukuran S/M/L/XL/XXL",
        images: "assets/img/baju_muslim2.jpg",
        slug: 'baju-koko-pria-beragam-ukuran',
        deskripsi: 'Baju Koko Pria Beragam Ukuran',
        kategori_id: 1,
        harga: 17000,
        stok: 12,
        rating: 5.0
      }
    ]
  }
  
  ngOnInit(): void {

  }
  
  public getProducts(): void {
    this.productService.getProducts().subscribe(
      resp => {
        this.loadingService.setLoading(false)
        this.products = resp.data
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

  public setActiveSidebar(sidebar: string): void {
    this.activeSidebar = sidebar
  }
}
