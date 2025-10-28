import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { ProductService } from 'src/app/services/product/product.service';

@Component({
    selector: 'app-product-detail',
    templateUrl: './product-detail.component.html',
    styleUrls: ['./product-detail.component.css'],
    standalone: false
})

export class ProductDetailComponent implements OnInit {
  product: any
  isLoading: Observable<boolean>
  error: boolean = false
  
  constructor(
    private productService: ProductService,
    private route: ActivatedRoute,
    private loadingService: LoadingService
  ) {
    this.isLoading = this.loadingService.loading
  }

  ngOnInit(): void {
    this.getDetailProduct()
  }

  public getDetailProduct(): void {
    const productSlug = this.route.snapshot.paramMap.get("slug")!
    this.productService.getDetailProduct(productSlug).subscribe(
      resp => {
        this.product = resp.data
      },
      () => {
        this.error = true
      }
    )
  }
}