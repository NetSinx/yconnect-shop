import { NgModule, inject } from '@angular/core';
import { RouterModule, Routes, ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { ProductDetailComponent } from './components/product-detail/product-detail.component';
import { ProductService } from './services/product/product.service';
import { Observable, map } from 'rxjs';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { authGuard } from './guard/auth/auth.guard';
import { PageForbiddenComponent } from './components/page-forbidden/page-forbidden.component';
import { guestGuard } from './guard/guest/guest.guard';
import { HomeComponent } from './components/home/home.component';

const productDetailRoute: ResolveFn<string> = (route: ActivatedRouteSnapshot, state: RouterStateSnapshot): string | Observable<string> | Promise<string> => {
  return inject(ProductService).getDetailProduct(route.paramMap.get("slug")!).pipe(
    map(product => `${product.data.nama} | Y-Connect Shop`)
  )
}

const routes: Routes = [
  {path: '', component: HomeComponent, title: 'Y-Connect Shop'},
  {path: 'product/:slug', component: ProductDetailComponent},
  {path: 'register', component: RegisterComponent, title: "Register | Y-Connect Shop", canActivate: [guestGuard]},
  {path: 'login', component: LoginComponent, title: "Login | Y-Connect Shop", canActivate: [guestGuard]},
  {path: 'dashboard', component: DashboardComponent, title: "Dashboard | Y-Connect Shop", canActivate: [authGuard]},
  {path: 'forbidden', component: PageForbiddenComponent, title: "Forbidden | Y-Connect Shop"},
  {path: '**', component: PageNotFoundComponent, title: "Upss... | Y-Connect Shop"}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule { }