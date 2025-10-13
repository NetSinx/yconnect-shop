import { Component, ElementRef, HostListener, OnInit, ViewChild } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { Category } from 'src/app/interfaces/category';
import { AuthService } from 'src/app/services/auth/auth.service';
import { CategoryService } from 'src/app/services/category/category.service';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';
import { LoadingService } from 'src/app/services/loading/loading.service';

@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html',
    styleUrls: ['./navbar.component.css'],
    standalone: false
})
export class NavbarComponent implements OnInit {
  categories: Category[] = []
  currentRoute: string = ""
  isLoggedIn: boolean | null = null
  username: string | null = null
  @ViewChild('dropdownMenu') dropdownMenu: ElementRef | null = null
  isDropDownOpen: boolean = false

  constructor(
    private categoryService: CategoryService,
    private router: Router,
    private loadingService: LoadingService,
    private authService: AuthService,
    private csrfService: GenCsrfService
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
    this.csrfService.getCSRF().subscribe(
      () => {
        this.authService.refreshToken().subscribe(
          () => this.isLoggedIn = true
        )
      }
    )
    // this.getCategories()
  }
  
  public getCategories(): void {
    this.categoryService.getCategories().subscribe(
      resp => {
        this.loadingService.setLoading(false)
        this.categories = resp.data
      }
    )
  }

  public toggleDropdown(event: MouseEvent): void {
    event.stopPropagation()
    this.isDropDownOpen = !this.isDropDownOpen
    console.log(this.isDropDownOpen)
  }

  @HostListener('document:click', ['$event'])
  public popMenuProfile(event: MouseEvent): void {
    if (!this.dropdownMenu?.nativeElement.contains(event.target)) {
      this.isDropDownOpen = false
    }
  }

  // public userLogout(): void {
  //   this.userService.userLogout().subscribe(
  //     () => {
  //       this.isLoggedIn = false
  //       this.username = null
  //       this.router.navigate(['/'])
  //       this.loadingService.setLoading(false)
  //     }
  //   )
  // }

  public hamburgerMenu(el1: HTMLElement, el2: HTMLElement, el3: HTMLElement, el4: HTMLElement): void {
    el1.classList.toggle('active')
    el2.classList.toggle('active')
    el3.classList.toggle('active')
    el4.classList.toggle('active')
  }
}
