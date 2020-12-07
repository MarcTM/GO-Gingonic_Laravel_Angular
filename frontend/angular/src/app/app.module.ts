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
import { MealsComponent } from './meals/meals.component';
import { ShowAuthedDirective } from './core/directives/show-authed.directive';
import { EditorComponent } from './editor/editor.component';
import { RecipeListComponent } from './shared/recipe-helpers/recipe-list.component';
import { RecipePreviewComponent } from './shared/recipe-helpers/recipe-preview.component';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    RecipesComponent,
    RecipeDetailsComponent,
    FooterComponent,
    LoginComponent,
    RegisterComponent,
    MealsComponent,
    ShowAuthedDirective,
    EditorComponent,
    RecipeListComponent,
    RecipePreviewComponent
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
