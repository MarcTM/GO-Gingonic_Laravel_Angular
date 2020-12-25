import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RecipesComponent } from './recipes/recipes.component';
import { RecipeDetailsComponent } from './recipe/recipe-details.component';
import { ProfilesComponent } from './profiles/profiles.component';
import { EditorComponent } from './editor/editor.component';
import { LoginComponent } from './auth/login/login.component';
import { RegisterComponent } from './auth/register/register.component';
import { RecoverPasswordComponent } from './auth/recover-password/recover-password.component';
import { MealsComponent } from './meals/meals.component';
import { CreateMealComponent } from './meals/create/create-meal.component';
import { AuthGuardService } from './core/services/guards/auth-guard.service';
import { AdminGuardService } from './core/services/guards/admin-guard.service';
import { NoAuthGuardService } from './core/services/guards/no-auth-guard.service';


const routes: Routes = [
  { path: '', redirectTo: '/recipes', pathMatch: 'full' },

  { path: 'recipes', component: RecipesComponent },

  { path: 'recipe/:id', component: RecipeDetailsComponent },

  { path: 'profiles/:username', component: ProfilesComponent },

  { path: 'editor', component: EditorComponent,
    canActivate: [AuthGuardService] },

  { path: 'signin', component: LoginComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'signup', component: RegisterComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'recover', component: RecoverPasswordComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'meals', component: MealsComponent,
    canActivate: [AuthGuardService] },

  { path: 'create', component: CreateMealComponent,
    canActivate: [AdminGuardService] },

  { path: '**', redirectTo: '/recipes', pathMatch: 'full' }
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})


export class AppRoutingModule { }