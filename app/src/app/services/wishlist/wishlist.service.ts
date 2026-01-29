import { Injectable, inject, computed, signal } from '@angular/core';
import { WishlistItem } from '../../interfaces/wishlist-item';
import { Product } from '../../interfaces/product';
import { CartService } from '../../services/cart/cart.service';

@Injectable({
  providedIn: 'root'
})
export class WishlistService {
  cartService: CartService = inject(CartService);
  wishlistItems = signal<WishlistItem[]>([]);

  hasAnySelection = computed(() => {
    return this.wishlistItems().some(item => item.selected);
  });

  totalItems = computed(() => this.wishlistItems().length);

  isAllSelected = computed(() => {
    const items = this.wishlistItems();
    return items.length > 0 && items.every((item: any) => item.selected);
  });

  addToWishlist(product: Product) {
    const existingItem = this.wishlistItems().find((item: any) => item.id === product.id);

    if (existingItem) {
      this.removeFromWishlist(product.id);
    } else {
      const newItem: WishlistItem = {
        id: product.id,
        nama: product.nama,
        harga: product.harga,
        images: product.images,
        stok: true,
        selected: false
      };

      this.wishlistItems.update((items: any) => [...items, newItem]);
    }
  }

  isInWishlist(productId: number): boolean {
    return this.wishlistItems().some((item: any) => item.id === productId);
  }

  removeFromWishlist(id: number) {
    this.wishlistItems.update((items: any) => items.filter((item: any) => item.id !== id));
  }

  toggleItemSelection(id: number) {
    this.wishlistItems.update((items: any) =>
      items.map((item: any) => (item.id === id ? { ...item, selected: !item.selected } : item))
    );
  }

  toggleAllSelection(isChecked: boolean) {
    this.wishlistItems.update((items: any) => items.map((item: any) => ({ ...item, selected: isChecked })));
  }

  moveSelectedToCart() {
    const selectedItems = this.wishlistItems().filter((item: any) => item.selected);

    selectedItems.forEach((item: any) => {
      this.cartService.addToCart(item, 1);
      this.removeFromWishlist(item.id);
    });
  }
}
