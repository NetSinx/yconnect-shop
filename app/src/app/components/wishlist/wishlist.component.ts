import { Component, inject } from '@angular/core';
import { WishlistService } from '../../services/wishlist/wishlist.service';

@Component({
  selector: 'app-wishlist',
  templateUrl: './wishlist.component.html',
  styleUrl: './wishlist.component.css',
  standalone: false
})
export class WishlistComponent {
  wishlistService: WishlistService = inject(WishlistService);

  onToggleAll(event: any) {
    this.wishlistService.toggleAllSelection(event.target.checked);
  }

  onToggleItem(id: number) {
    this.wishlistService.toggleItemSelection(id);
  }

  addToCartSelected() {
    this.wishlistService.moveSelectedToCart();
    alert('Produk terpilih berhasil dipindahkan ke Keranjang!');
  }
}
