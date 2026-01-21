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
import { NgOptimizedImage, registerLocaleData } from '@angular/common';
import { loadingInterceptor } from './interceptor/loading/loading.interceptor';
import { ReactiveFormsModule } from '@angular/forms';
import { RegisterComponent } from './components/register/register.component';
import { authInterceptor } from './interceptor/auth/auth.interceptor';
import { HomeComponent } from './components/home/home.component';
import { FooterComponent } from './components/footer/footer.component';
import localeId from '@angular/common/locales/id';

registerLocaleData(localeId, 'id');

@NgModule({
  declarations: [
    AppComponent,
    ProductComponent,
    NavbarComponent,
    ProductDetailComponent,
    RegisterComponent,
    LoginComponent,
    PageNotFoundComponent,
    DashboardComponent,
    HomeComponent,
    FooterComponent
  ],
  bootstrap: [AppComponent],
  imports: [BrowserModule, AppRoutingModule, NgOptimizedImage, ReactiveFormsModule],
  providers: [
    ProductComponent,
    NavbarComponent,
    LoginComponent,
    ProductDetailComponent,
    HomeComponent,
    FooterComponent,
    provideHttpClient(withInterceptorsFromDi(), withInterceptors([loadingInterceptor, authInterceptor]), withFetch()),
    provideClientHydration(withEventReplay()),
    { provide: LOCALE_ID, useValue: 'id-ID' }
  ]
})
export class AppModule {}
