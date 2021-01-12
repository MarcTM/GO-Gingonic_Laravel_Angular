import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToastrModule } from 'ngx-toastr';
import { CarouselModule } from 'ngx-owl-carousel-o';
import { NgxPaginationModule } from 'ngx-pagination';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RecipesComponent } from './recipes/recipes.component';
import { HeaderComponent } from './shared/layout/header/header.component';
import { FooterComponent } from './shared/layout/footer/footer.component';
import { RecipeDetailsComponent } from './recipe/recipe-details.component';
import { LoginComponent } from './auth/login/login.component';
import { RegisterComponent } from './auth/register/register.component';
import { RecoverPasswordComponent } from './auth/recover-password/recover-password.component';
import { MealsComponent } from './meals/meals.component';
import { ShowAuthedDirective } from './core/directives/show-authed.directive';
import { ShowAdminDirective } from './core/directives/show-admin.directive';
import { EditorComponent } from './editor/editor.component';
import { RecipeListComponent } from './shared/recipe-helpers/recipe-list.component';
import { RecipePreviewComponent } from './shared/recipe-helpers/recipe-preview.component';
import { CreateMealComponent } from './meals/create/create-meal.component';
import { ProfilesComponent } from './profiles/profiles.component';
import { MealListComponent } from './shared/meal-helpers/meal-list.component';
import { MealPreviewComponent } from './shared/meal-helpers/meal-preview.component';
import { MealDetailsComponent } from './meal/meal-details.component';
import { HomeComponent } from './home/home.component';
import { SliderComponent } from './shared/slider/slider.component';
import { FollowFavButtonsComponent } from './shared/buttons/fav_follow/follow-fav-buttons.component';
import { UpdateDeleteButtonsComponent } from './shared/buttons/update-delete/update-delete-buttons.component';
import { EditorUpdateComponent } from './editor/update/editor-update.component';
import { CategoriesComponent } from './categories/categories.component';
import { CreateCategoryComponent } from './categories/create/create-category.component';
import { CategoryComponent } from './category/category.component';
import { AccountComponent } from './account/account.component';
import { UpdateAccountComponent } from './account/update/update-account.component';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    RecipesComponent,
    RecipeDetailsComponent,
    ProfilesComponent,
    FooterComponent,
    LoginComponent,
    RegisterComponent,
    MealsComponent,
    ShowAuthedDirective,
    ShowAdminDirective,
    EditorComponent,
    RecipeListComponent,
    RecipePreviewComponent,
    RecoverPasswordComponent,
    CreateMealComponent,
    MealListComponent,
    MealPreviewComponent,
    MealDetailsComponent,
    HomeComponent,
    SliderComponent,
    FollowFavButtonsComponent,
    UpdateDeleteButtonsComponent,
    EditorUpdateComponent,
    CategoriesComponent,
    CreateCategoryComponent,
    CategoryComponent,
    AccountComponent,
    UpdateAccountComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    ReactiveFormsModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    CarouselModule,
    ToastrModule.forRoot(),
    NgxPaginationModule
  ],
  providers: [],
  bootstrap: [
    AppComponent,
  ]
})

export class AppModule { }
