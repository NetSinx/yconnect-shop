import { Component, inject } from '@angular/core';
import { LayoutService } from '../../services/layout/layout.service';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrl: './footer.component.css',
  standalone: false
})
export class FooterComponent {
  layoutService: LayoutService = inject(LayoutService);

  constructor() {}
}
