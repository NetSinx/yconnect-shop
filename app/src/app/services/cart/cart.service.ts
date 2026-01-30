import { Injectable, signal, computed, effect } from '@angular/core';
import { CartItem } from '../../interfaces/cart-item';

@Injectable({
  providedIn: 'root'
})
export class CartService {
  cartItems = signal<CartItem[]>([]);
  totalItems: any = computed(() => this.cartItems().reduce((acc, item) => acc + item.quantity, 0));
  grandTotal: any = computed(() =>
    this.cartItems()
      .filter(item => item.selected)
      .reduce((acc, item) => acc + item.harga * item.quantity, 0)
  );
  isAllSelected: any = computed(() => {
    const items = this.cartItems();
    return items.length > 0 && items.every(item => item.selected);
  });

  constructor() {
    if (typeof localStorage !== 'undefined') {
      const savedCart = localStorage.getItem('my_cart');
      if (savedCart) {
        this.cartItems.set(JSON.parse(savedCart));
      }

      effect(() => localStorage.setItem('my_cart', JSON.stringify(this.cartItems())));
    }
  }

  addToCart(product: any, qty: number) {
    const currentItems = this.cartItems();
    const existingItem = currentItems.find(item => item.id === product.id);

    if (existingItem) {
      this.cartItems.update(items =>
        items.map(item => (item.id === product.id ? { ...item, quantity: item.quantity + qty } : item))
      );
    } else {
      const newItem: CartItem = {
        id: product.id,
        nama: product.nama,
        harga: product.harga,
        images: product.images,
        quantity: qty,
        selected: true
      };
      this.cartItems.update(items => [...items, newItem]);
    }
  }

  updateQuantity(id: number, quantity: number) {
    if (quantity <= 0) return;
    this.cartItems.update(items => items.map(item => (item.id === id ? { ...item, quantity } : item)));
  }

  removeFromCart(id: number) {
    this.cartItems.update(items => items.filter(item => item.id !== id));
  }

  toggleItemSelection(id: number) {
    this.cartItems.update(items => items.map(item => (item.id === id ? { ...item, selected: !item.selected } : item)));
  }

  toggleAllSelection(isChecked: boolean) {
    this.cartItems.update(items => items.map(item => ({ ...item, selected: isChecked })));
  }
}
