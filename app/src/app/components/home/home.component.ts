import { Component, OnInit } from '@angular/core';
import { Kategori } from 'src/app/interfaces/category';
import { CategoryService } from 'src/app/services/category/category.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
  standalone: false
})
export class HomeComponent implements OnInit {
  categories: Kategori[] = []

  constructor(private categoryService: CategoryService) {
    this.categories = [
      {
        id: 1, nama: "Pakaian", slug: "pakaian", total: 14
      },
      {
        id: 2, nama: "Makanan", slug: "makanan", total: 25
      },
      {
        id: 3, nama: "Minuman", slug: "minuman", total: 10
      },
      {
        id: 4, nama: "Detergen", slug: "detergen", total: 8
      }
    ]
  }

  ngOnInit(): void {
    // this.categoryService.getCategories().subscribe(
    //   value => this.categories = value.data
    // )
  }
}
