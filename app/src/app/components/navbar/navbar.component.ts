import { Component, OnInit } from '@angular/core';
import { Category } from 'src/app/interfaces/category';
import { HomeService } from 'src/app/services/home/home.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  categories: Category[] = [];

  constructor(private homeService: HomeService) {}

  ngOnInit(): void {
    this.getCategories()
  }
  
  public getCategories(): void {
    this.homeService.getCategories().subscribe(
      data => this.categories = data.data
    )
  }
}
