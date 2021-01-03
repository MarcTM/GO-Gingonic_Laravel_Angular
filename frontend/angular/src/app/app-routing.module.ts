import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { HomeComponent } from './home/home.component';
import { RecipesComponent } from './recipes/recipes.component';
import { RecipeDetailsComponent } from './recipe/recipe-details.component';
import { ProfilesComponent } from './profiles/profiles.component';
import { EditorComponent } from './editor/editor.component';
import { LoginComponent } from './auth/login/login.component';
import { RegisterComponent } from './auth/register/register.component';
import { RecoverPasswordComponent } from './auth/recover-password/recover-password.component';
import { MealsComponent } from './meals/meals.component';
import { MealDetailsComponent } from './meal/meal-details.component';
import { CreateMealComponent } from './meals/create/create-meal.component';
import { AuthGuardService } from './core/services/guards/auth-guard.service';
import { AdminGuardService } from './core/services/guards/admin-guard.service';
import { NoAuthGuardService } from './core/services/guards/no-auth-guard.service';
import { EditorUpdateComponent } from './editor/update/editor-update.component';
import { CreateCategoryComponent } from './categories/create/create-category.component';


const routes: Routes = [
  { path: '', component: HomeComponent },

  { path: 'recipes', component: RecipesComponent },

  { path: 'recipe/:id', component: RecipeDetailsComponent },

  { path: 'profiles/:username', component: ProfilesComponent },

  { path: 'editor', component: EditorComponent,
    canActivate: [AuthGuardService] },

  { path: 'editor/:id', component: EditorUpdateComponent,
    canActivate: [AuthGuardService] },

  { path: 'signin', component: LoginComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'signup', component: RegisterComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'recover', component: RecoverPasswordComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'meals', component: MealsComponent },

  { path: 'meal/:id', component: MealDetailsComponent },

  { path: 'create_meal', component: CreateMealComponent,
    canActivate: [AdminGuardService] },

  { path: 'create_category', component: CreateCategoryComponent,
    canActivate: [AdminGuardService] },

  { path: '**', redirectTo: '', pathMatch: 'full' }
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})

export class AppRoutingModule { }