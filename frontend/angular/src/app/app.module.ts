import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToastrModule } from 'ngx-toastr';

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
    CreateMealComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    ReactiveFormsModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot()
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
