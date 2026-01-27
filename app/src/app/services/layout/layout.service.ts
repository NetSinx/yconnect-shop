import { Injectable, signal } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LayoutService {
  sidebarOpen = signal(false);

  toggleSidebar() {
    this.sidebarOpen.update(value => !value);
  }
}
