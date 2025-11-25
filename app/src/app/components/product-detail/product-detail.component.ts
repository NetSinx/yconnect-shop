import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { Kategori } from 'src/app/interfaces/category';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { ProductService } from 'src/app/services/product/product.service';

@Component({
  selector: 'app-product-detail',
  templateUrl: './product-detail.component.html',
  styleUrls: ['./product-detail.component.css'],
  standalone: false
})
export class ProductDetailComponent implements OnInit {
  product: any;
  categories: Kategori[] = [];
  isLoading: Observable<boolean>;
  error: boolean = false;

  constructor(
    private productService: ProductService,
    private route: ActivatedRoute,
    private loadingService: LoadingService
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
    // this.getDetailProduct()
  }

  public getDetailProduct(): void {
    const productSlug = this.route.snapshot.paramMap.get('slug')!;
    this.productService.getDetailProduct(productSlug).subscribe(
      resp => {
        this.product = resp.data;
      },
      () => {
        this.error = true;
      }
    );
  }

  public increaseQuantity(el: HTMLInputElement): void {
    let val = parseInt(el.value);
    el.value = (val + 1).toString();
  }

  public decreaseQuantity(el: HTMLInputElement): void {
    let val = parseInt(el.value);
    if (val > 1) {
      el.value = (val - 1).toString();
    } else {
      el.value = '1';
    }
  }

  public qtyValidation(el: HTMLInputElement): void {
    if (el.value === '0') {
      el.value = '1';
    }
    if (parseInt(el.value) < 0) {
      el.value = '1';
    }
  }
}
