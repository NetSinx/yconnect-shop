import { Component, inject } from '@angular/core';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrl: './cart.component.css',
  standalone: false
})
export class CartComponent {
  cartService = inject(CartService);

  onToggleAll(event: any) {
    const isChecked = event.target.checked;
    this.cartService.toggleAllSelection(isChecked);
  }

  onToggleItem(id: number) {
    this.cartService.toggleItemSelection(id);
  }
}
