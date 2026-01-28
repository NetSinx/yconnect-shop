import {
  Component,
  computed,
  ElementRef,
  HostListener,
  Inject,
  inject,
  input,
  OnInit,
  PLATFORM_ID,
  ViewChild
} from '@angular/core';
import { Router } from '@angular/router';
import { Kategori } from 'src/app/interfaces/category';
import { AuthService } from 'src/app/services/auth/auth.service';
import { CategoryService } from 'src/app/services/category/category.service';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { TransferState, makeStateKey } from '@angular/core';
import { isPlatformServer } from '@angular/common';
import { LayoutService } from 'src/app/services/layout/layout.service';
import { CartService } from '../../services/cart/cart.service';

const IS_LOGGED_IN_KEY = makeStateKey<boolean>('isLoggedIn');

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'],
  standalone: false
})
export class NavbarComponent implements OnInit {
  categories = input.required<Kategori[]>();
  currentRoute: string = '';
  isLoggedIn: boolean | null = null;
  username: string | null = null;
  @ViewChild('dropdownMenu') dropdownMenu: ElementRef | null = null;
  isDropDownOpen: boolean = false;
  layoutService: LayoutService = inject(LayoutService);
  cartService: CartService = inject(CartService);

  constructor(
    private authService: AuthService,
    private csrfService: GenCsrfService,
    private state: TransferState,
    @Inject(PLATFORM_ID) private platformId: Object
  ) {
    this.isLoggedIn = true;
  }

  ngOnInit(): void {
    this.categories;

    const savedAPIResp = this.state.get(IS_LOGGED_IN_KEY, null);
    if (savedAPIResp) {
      this.isLoggedIn = savedAPIResp;
    } else {
      this.csrfService.getCSRF().subscribe(() => {
        this.authService.refreshToken().subscribe(() => {
          if (isPlatformServer(this.platformId)) {
            this.isLoggedIn = true;
            this.state.set(IS_LOGGED_IN_KEY, this.isLoggedIn);
          }
        });
      });
    }

    // this.getCategories()
  }

  // public getCategories(): void {
  //   this.categoryService.getCategories().subscribe(
  //     resp => {
  //       this.loadingService.setLoading(false)
  //       this.categories = resp.data
  //     }
  //   )
  // }

  public toggleDropdown(event: MouseEvent): void {
    event.stopPropagation();
    this.isDropDownOpen = !this.isDropDownOpen;
    console.log(this.isDropDownOpen);
  }

  @HostListener('document:click', ['$event'])
  public popMenuProfile(event: MouseEvent): void {
    if (!this.dropdownMenu?.nativeElement.contains(event.target)) {
      this.isDropDownOpen = false;
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
    el1.classList.toggle('active');
    el2.classList.toggle('active');
    el3.classList.toggle('active');
    el4.classList.toggle('active');
  }
}
