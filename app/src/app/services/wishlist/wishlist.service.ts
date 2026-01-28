import { Injectable } from '@angular/core';
import { WishlistItem } from '../../interfaces/wishlist-item';
import { Product } from '../../interfaces/product';

@Injectable({
  providedIn: 'root'
})
export class WishlistService {
  wishlistItem: WishlistItem[] = [];

  constructor() {}

  addToWishlist(product: Product) {
    this.wishlistItem().find(item);
  }
}
