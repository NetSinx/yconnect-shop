import { Component, OnInit } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { LoadingService } from './services/loading/loading.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  standalone: false
})

export class AppComponent implements OnInit {
  errors: any
  showNavbar: boolean = true
  isLoading: Observable<boolean> = new Observable<boolean>()

  constructor(private router: Router, private loadingService: LoadingService) {
    this.isLoading = this.loadingService.loading
  }

  ngOnInit(): void {
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        if (event.urlAfterRedirects === '/login' || event.urlAfterRedirects === '/register') {
          this.showNavbar = false
        } else {
          this.showNavbar = true
        }
      }
    })
  }
}
