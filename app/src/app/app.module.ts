import { NgModule, LOCALE_ID } from '@angular/core';
import { BrowserModule, provideClientHydration, withEventReplay } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { provideHttpClient, withFetch, withInterceptors, withInterceptorsFromDi } from '@angular/common/http';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { ProductDetailComponent } from './components/product-detail/product-detail.component';
import { LoginComponent } from './components/login/login.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ProductComponent } from './components/product/product.component';
import { NgOptimizedImage, registerLocaleData, DecimalPipe } from '@angular/common';
import { loadingInterceptor } from './interceptor/loading/loading.interceptor';
import { ReactiveFormsModule } from '@angular/forms';
import { RegisterComponent } from './components/register/register.component';
import { authInterceptor } from './interceptor/auth/auth.interceptor';
import { FooterComponent } from './components/footer/footer.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import localeId from '@angular/common/locales/id';
import { MainLayoutComponent } from './components/main-layout/main-layout.component';
import { CartComponent } from './components/cart/cart.component';
import { WishlistComponent } from './components/wishlist/wishlist.component';
import { provideSweetAlert2 } from '@sweetalert2/ngx-sweetalert2';
import { SwalComponent, SwalDirective } from '@sweetalert2/ngx-sweetalert2';

registerLocaleData(localeId, 'id');

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    ProductComponent,
    ProductDetailComponent,
    RegisterComponent,
    LoginComponent,
    SidebarComponent,
    PageNotFoundComponent,
    DashboardComponent,
    FooterComponent,
    MainLayoutComponent,
    CartComponent,
    WishlistComponent
  ],
  bootstrap: [AppComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgOptimizedImage,
    ReactiveFormsModule,
    DecimalPipe,
    SwalComponent,
    SwalDirective
  ],
  providers: [
    provideHttpClient(withInterceptorsFromDi(), withInterceptors([loadingInterceptor, authInterceptor]), withFetch()),
    provideSweetAlert2({
      fireOnInit: false,
      dismissOnDestroy: true
    }),
    provideClientHydration(withEventReplay()),
    { provide: LOCALE_ID, useValue: 'id-ID' }
  ]
})
export class AppModule {}
