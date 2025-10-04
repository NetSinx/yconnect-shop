import { NgModule } from '@angular/core';
import { BrowserModule, provideClientHydration, withEventReplay } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { provideHttpClient, withFetch, withInterceptors, withInterceptorsFromDi } from '@angular/common/http';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component'
import { ProductDetailComponent } from './components/product-detail/product-detail.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ProductComponent } from './components/product/product.component';
import { NgOptimizedImage } from '@angular/common';
import { loadingInterceptor } from './interceptor/loading/loading.interceptor';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
    declarations: [
        AppComponent,
        ProductComponent,
        NavbarComponent,
        ProductDetailComponent,
        LoginComponent,
        RegisterComponent,
        PageNotFoundComponent,
        DashboardComponent,
    ],
    bootstrap: [AppComponent],
    imports: [
        BrowserModule,
        AppRoutingModule,
        NgOptimizedImage,
        ReactiveFormsModule
    ],
    providers: [
        ProductComponent,
        NavbarComponent,
        LoginComponent,
        RegisterComponent,
        ProductDetailComponent,
        provideHttpClient(withInterceptorsFromDi(), withInterceptors([loadingInterceptor]), withFetch()),
        // provideClientHydration(withEventReplay())
    ]
})

export class AppModule { }