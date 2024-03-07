import { NgModule, inject } from '@angular/core';
import { RouterModule, Routes, ActivatedRouteSnapshot, ResolveFn, RouterStateSnapshot } from '@angular/router';
// import { ErrorNotFoundComponent } from './error-not-found/error-not-found.component';
import { HomeComponent } from './components/home/home.component';
import { ProductDetailComponent } from './components/product-detail/product-detail.component';
import { HomeService } from './services/home/home.service';
import { Observable, map } from 'rxjs';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';

const productDetailRoute: ResolveFn<string> = (route: ActivatedRouteSnapshot, state: RouterStateSnapshot): string | Observable<string> | Promise<string> => {
  return inject(HomeService).showDetailProduct(route.paramMap.get("slug")!).pipe(
    map(product => `Y-Connect Shop | ${product.data.name}`)
  )
}

const routes: Routes = [
  {path: '', component: HomeComponent, title: 'Y-Connect Shop'},
  {path: 'products/:slug', component: ProductDetailComponent, title: productDetailRoute},
  {path: 'register', component: RegisterComponent, title: "Register | Y-Connect Shop"},
  {path: 'login', component: LoginComponent, title: "Login | Y-Connect Shop"},
  // {path: '**', component: ErrorNotFoundComponent, title: "404 Not Found"}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule { }