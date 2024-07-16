import { Component, OnInit } from '@angular/core';
import { ProductService } from '../../services/product/product.service';
import { Product } from 'src/app/interfaces/product';

@Component({
  selector: 'app-product',
  templateUrl: './product.component.html',
  styleUrls: ['./product.component.css']
})

export class ProductComponent implements OnInit {
  products: Product[] = []

  constructor(private productService: ProductService) {}
  
  ngOnInit(): void {
    this.getProducts()
  }
  
  public getProducts(): void {
    this.productService.getProducts().subscribe(
      resp => {
        this.products = resp.data
      }
    )
  }
}
