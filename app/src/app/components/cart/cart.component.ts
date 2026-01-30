import { Component, inject, ViewChild } from '@angular/core';
import { CartService } from '../../services/cart/cart.service';
import { SwalComponent } from '@sweetalert2/ngx-sweetalert2';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrl: './cart.component.css',
  standalone: false
})
export class CartComponent {
  cartService: CartService = inject(CartService);

  @ViewChild('successSwal') public readonly successSwal!: SwalComponent;

  onToggleAll(event: any) {
    const isChecked = event.target.checked;
    this.cartService.toggleAllSelection(isChecked);
  }

  onToggleItem(id: number) {
    this.cartService.toggleItemSelection(id);
  }

  handleDeleteItemCart(id: number) {
    this.cartService.removeFromCart(id);
    this.successSwal.fire();
  }
}
