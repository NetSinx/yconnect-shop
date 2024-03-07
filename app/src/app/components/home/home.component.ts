import { Component, OnInit } from '@angular/core';
import { HomeService } from '../../services/home/home.service';
import { Product } from 'src/app/interfaces/product';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {
  products: Product[] = []

  ngOnInit(): void {
    this.getProducts()
  }

  constructor(private homeService: HomeService) {}

  public getProducts(): void {
    this.homeService.getProducts().subscribe(
      data => this.products = data.data
    )
  }

  public getCarts(): void {
    this.homeService.getCarts()
  }
}
