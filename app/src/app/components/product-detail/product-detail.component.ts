import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { HomeService } from 'src/app/services/home/home.service';

@Component({
  selector: 'app-product-detail',
  templateUrl: './product-detail.component.html',
  styleUrls: ['./product-detail.component.css']
})

export class ProductDetailComponent implements OnInit {
  product: any
  
  constructor(
    private homeService: HomeService,
    private route: ActivatedRoute
  ) {}

  ngOnInit(): void {
    this.getDetailProduct()
  }

  public getDetailProduct(): void {
    const productSlug = this.route.snapshot.paramMap.get("slug")
    this.homeService.showDetailProduct(productSlug!).subscribe(
      data => this.product = data.data
    )
  }
}