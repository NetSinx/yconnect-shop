import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { HTTP_INTERCEPTORS, HttpClientModule, HttpClientXsrfModule, HttpXsrfTokenExtractor } from '@angular/common/http';
import { AppComponent } from './app.component';
import { HomeComponent } from './components/home/home.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { ProductDetailComponent } from './components/product-detail/product-detail.component';
import { LoginComponent } from './components/login/login.component';
import { ReactiveFormsModule } from '@angular/forms';
import { RegisterComponent } from './components/register/register.component';
import { CustomInterceptorInterceptor } from './interceptor/custom-interceptor.interceptor';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    NavbarComponent,
    ProductDetailComponent,
    LoginComponent,
    RegisterComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule,
    HttpClientXsrfModule.withOptions({
      cookieName: "XSRF-Token",
      headerName: "XSRF-Token"
    })
  ],
  providers: [
    HomeComponent,
    NavbarComponent,
    LoginComponent,
    RegisterComponent,
    ProductDetailComponent,
    {provide: HTTP_INTERCEPTORS, useClass: CustomInterceptorInterceptor, multi: true},
    {provide: HttpXsrfTokenExtractor, useClass: HttpClientXsrfModule}
  ],
  bootstrap: [AppComponent]

})

export class AppModule { }