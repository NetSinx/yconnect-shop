import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, NavigationEnd, Router } from '@angular/router';
import { Category } from 'src/app/interfaces/category';
import { CategoryService } from 'src/app/services/category/category.service';
import { LoginService } from 'src/app/services/login/login.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {
  categories: Category[] = []
  currentRoute: any
  isLoggedIn: boolean = false

  constructor(
    private categoryService: CategoryService,
    private router: Router,
    private loginService: LoginService,
    private userService: UserService
  ) {
    this.router.events.subscribe(
      nav => {
        if (nav instanceof NavigationEnd) {
          this.currentRoute = nav.url
        }
      }
    )
  }

  ngOnInit(): void {
    this.getCategories()
    this.loginService.verifyUser().toPromise().then(
      () => this.isLoggedIn = true,
      () => this.isLoggedIn = false
    )

    // this.userService.getUser()
  }
  
  public getCategories(): void {
    this.categoryService.getCategories().subscribe(
      resp => {
        this.categories = resp.data
      }
    )
  }
}
