import { Component, OnInit, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { ProductService } from 'src/app/services/product/product.service';
import { LayoutService } from '../../services/layout/layout.service';
import { Title } from '@angular/platform-browser';
import { WishlistService } from '../../services/wishlist/wishlist.service';
import { CartService } from '../../services/cart/cart.service';

@Component({
  selector: 'app-product-detail',
  templateUrl: './product-detail.component.html',
  styleUrls: ['./product-detail.component.css'],
  standalone: false
})
export class ProductDetailComponent implements OnInit {
  products: any;
  product: any;
  isLoading: Observable<boolean>;
  error: boolean = false;
  layoutService: LayoutService = inject(LayoutService);
  private route = inject(ActivatedRoute);
  private titleService = inject(Title);
  wishlistService: WishlistService = inject(WishlistService);
  cartService: CartService = inject(CartService);

  constructor(
    private productService: ProductService,
    private loadingService: LoadingService
  ) {
    this.isLoading = this.loadingService.loading;
    this.products = [
      {
        id: 1,
        nama: 'Baju Muslim Keren & Kekinian',
        images: 'assets/img/baju_muslim1.jpg',
        slug: 'baju-muslim-keren-kekinian',
        deskripsi: 'Baju muslim keren dan kekinian',
        kategori_id: 1,
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
        kategori_id: 1,
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
        kategori_id: 1,
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
        kategori_id: 1,
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
        kategori_id: 2,
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
        kategori_id: 3,
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
        kategori_id: 3,
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
        kategori_id: 3,
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
        kategori_id: 3,
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
        kategori_id: 3,
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
        kategori_id: 2,
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
        kategori_id: 2,
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
        kategori_id: 2,
        harga: 17000,
        stok: 10,
        rating: 5.0
      }
    ];
  }

  ngOnInit(): void {
    // this.getDetailProduct()
    const slug = this.route.snapshot.paramMap.get('slug');
    this.product = this.products.find((p: any) => p.slug === slug);

    if (this.product) {
      this.titleService.setTitle(`${this.product.nama} | Y-Connect Shop`);
    } else {
      this.titleService.setTitle('Produk Tidak Ditemukan | Y-Connect Shop');
    }
  }

  //public getDetailProduct(): void {
  //  const productSlug = this.route.snapshot.paramMap.get('slug')!;
  // this.productService.getDetailProduct(productSlug).subscribe(
  //    resp => {
  //      this.product = resp.data;
  //    },
  //    () => {
  //      this.error = true;
  //    }
  //  );
  //}

  public addToCart(product: any, qty: string) {
    this.cartService.addToCart(product, Number(qty));
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

  public toggleFavorit(): void {
    if (this.product) {
      this.wishlistService.addToWishlist(this.product);
    }
  }

  get isFavorite(): boolean {
    return this.product ? this.wishlistService.isInWishlist(this.product.id) : false;
  }
}
