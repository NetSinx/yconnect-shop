import { Component } from '@angular/core';
import { HomeService } from './services/home/home.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  errors: any

  constructor(private homeService: HomeService) {}

  ngOnInit(): void {
    this.getErrorCategories()
    this.getErrorProducts()
  }

  public getErrorCategories() {
    this.homeService.getCategories().subscribe(
      data => {
        if (data.code === 500) {
          this.errors = [data.code, data.status]
        }
      },
      error => {
        this.errors = error
      }
    )
  }

  public getErrorProducts() {
    this.homeService.getProducts().subscribe(
      data => {
        if (data.code === 500) {
          this.errors = [data.code, data.status]
        }
      },
      error => {
        this.errors = error
      }
    )
  }
}