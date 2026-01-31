import { Component, inject, ViewChild } from '@angular/core';
import { WishlistService } from '../../services/wishlist/wishlist.service';
import { SwalComponent } from '@sweetalert2/ngx-sweetalert2';

@Component({
  selector: 'app-wishlist',
  templateUrl: './wishlist.component.html',
  styleUrl: './wishlist.component.css',
  standalone: false
})
export class WishlistComponent {
  wishlistService: WishlistService = inject(WishlistService);
  @ViewChild('successSwal') public readonly successSwal!: SwalComponent;
  @ViewChild('successMoveSwal') public readonly successMoveSwal!: SwalComponent;

  handleDeleteItemWishlist(id: number) {
    this.wishlistService.removeFromWishlist(id);
    this.successSwal.fire();
  }

  handleMoveItemToCart() {
    this.wishlistService.moveSelectedToCart();
    this.successMoveSwal.fire();
  }

  onToggleAll(event: any) {
    this.wishlistService.toggleAllSelection(event.target.checked);
  }

  onToggleItem(id: number) {
    this.wishlistService.toggleItemSelection(id);
  }
}
