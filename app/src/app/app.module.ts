import { NgModule } from '@angular/core';
import { BrowserModule, provideClientHydration, withEventReplay } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { HTTP_INTERCEPTORS, provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { AppComponent } from './app.component';
import { CategoryComponent } from './components/category/category.component';
import { ProductDetailComponent } from './components/product-detail/product-detail.component';
import { LoginComponent } from './components/login/login.component';
import { ReactiveFormsModule } from '@angular/forms';
import { RegisterComponent } from './components/register/register.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ProductComponent } from './components/product/product.component';
import { CustomInterceptor } from './interceptor/auth/custom.interceptor';
import { NgOptimizedImage } from '@angular/common';

@NgModule({ declarations: [
        AppComponent,
        ProductComponent,
        CategoryComponent,
        ProductDetailComponent,
        LoginComponent,
        RegisterComponent,
        PageNotFoundComponent,
        DashboardComponent,
    ],
    bootstrap: [AppComponent], imports: [BrowserModule,
        AppRoutingModule,
        ReactiveFormsModule,
        NgOptimizedImage], providers: [
        ProductComponent,
        CategoryComponent,
        LoginComponent,
        RegisterComponent,
        ProductDetailComponent,
        { provide: HTTP_INTERCEPTORS, useClass: CustomInterceptor, multi: true },
        provideHttpClient(withInterceptorsFromDi()),
        provideClientHydration(withEventReplay())
    ] })

export class AppModule { }