import { NgModule, inject } from '@angular/core';
import { RouterModule, Routes, ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
import { ProductComponent } from './components/product/product.component';
import { ProductDetailComponent } from './components/product-detail/product-detail.component';
import { ProductService } from './services/product/product.service';
import { Observable, map } from 'rxjs';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { authGuardGuard } from './guard/auth-guard/auth-guard.guard';
import { loginGuard } from './guard/login/login.guard';

const productDetailRoute: ResolveFn<string> = (route: ActivatedRouteSnapshot, state: RouterStateSnapshot): string | Observable<string> | Promise<string> => {
  return inject(ProductService).getDetailProduct(route.paramMap.get("slug")!).pipe(
    map(product => `Y-Connect Shop | ${product.data.name}`)
  )
}

const routes: Routes = [
  {path: '', component: ProductComponent, title: 'Y-Connect Shop'},
  {path: 'product', component: ProductDetailComponent, title: productDetailRoute},
  {path: 'register', component: RegisterComponent, title: "Register | Y-Connect Shop", canActivate: [loginGuard]},
  {path: 'login', component: LoginComponent, title: "Login | Y-Connect Shop", canActivate: [loginGuard]},
  {path: 'dashboard/:userId', component: DashboardComponent, title: "Dashboard | Y-Connect Shop", canActivate: [authGuardGuard], pathMatch: 'full'},
  {path: '**', component: PageNotFoundComponent, title: "404 Not Found"}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule { }