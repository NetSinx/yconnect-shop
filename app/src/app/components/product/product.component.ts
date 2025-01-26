import { Component, HostListener, OnInit } from '@angular/core';
import { ProductService } from '../../services/product/product.service';
import { Product } from 'src/app/interfaces/product';
import { LoadingService } from 'src/app/services/loading/loading.service';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css']
})

export class ProductComponent implements OnInit {
  products: Product[] = []

  constructor(private productService: ProductService, private loadingService: LoadingService) {}
  
  ngOnInit(): void {
    this.getProducts()
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
}
