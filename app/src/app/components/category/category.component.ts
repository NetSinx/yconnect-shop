import { Component, ElementRef, HostListener, OnInit, ViewChild } from '@angular/core';
import { NavigationEnd, Router } from '@angular/router';
import { Category } from 'src/app/interfaces/category';
import { CategoryService } from 'src/app/services/category/category.service';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { LoginService } from 'src/app/services/login/login.service';
import { UserService } from 'src/app/services/user/user.service';

@Component({
    selector: 'app-category',
    templateUrl: './category.component.html',
    styleUrls: ['./category.component.css'],
    standalone: false
})
export class CategoryComponent implements OnInit {
  categories: Category[] = []
  currentRoute: string = ""
  isLoggedIn: boolean | null = null
  username: string | null = null
  @ViewChild('dropdownMenu') dropdownMenu: ElementRef | null = null
  isDropDownOpen: boolean = false

  constructor(
    private categoryService: CategoryService,
    private router: Router,
    private loginService: LoginService,
    private loadingService: LoadingService,
    private userService: UserService,
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
    this.getCategories()
    this.csrfService.getCSRF().subscribe()
    this.router.events.subscribe(
      nav => {
        if (nav instanceof NavigationEnd) {
          this.loginService.verifyUser().subscribe(
            () => {
              this.isLoggedIn = true
              this.userService.getUserInfo().subscribe(
                resp => {
                  this.loadingService.setLoading(false)
                  this.username = resp.data.user_id
                }
              )
            },
            () => {
              this.loadingService.setLoading(false)
              this.isLoggedIn = false
              this.username = null
            }
            )
        }
      }
    )
      
    const timezone: string = Intl.DateTimeFormat().resolvedOptions().timeZone
    this.userService.setTimeZone(timezone).subscribe()
    this.loginService.verifyUser().subscribe(
      () => {
        this.isLoggedIn = true
        this.userService.getUserInfo().subscribe(
          resp => {
            this.loadingService.setLoading(false)
            this.username = resp.data.user_id
          }
        )
      },
      () => {
        this.loadingService.setLoading(false)
        this.isLoggedIn = false
        this.username = null
      }
    )
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
  }

  @HostListener('document:click', ['$event'])
  public popMenuProfile(event: MouseEvent): void {
    if (!this.dropdownMenu?.nativeElement.contains(event.target)) {
      this.isDropDownOpen = false
    }
  }

  public userLogout(): void {
    this.userService.userLogout().subscribe(
      () => {
        this.isLoggedIn = false
        this.username = null
        this.router.navigate(['/'])
        this.loadingService.setLoading(false)
      }
    )
  }

  public hamburgerMenu(el1: HTMLElement, el2: HTMLElement): void {
    el1.classList.toggle('active')
    el2.classList.toggle('active')
  }
}
