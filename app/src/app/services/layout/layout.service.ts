import { Injectable, signal } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LayoutService {
  sidebarOpen = signal(true);

  toggleSidebar() {
    this.sidebarOpen.update(value => !value);
  }
}
