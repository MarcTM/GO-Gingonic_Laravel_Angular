import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RecipesComponent } from './recipes/recipes.component';
import { RecipeDetailsComponent } from './recipe/recipe-details.component';
import { EditorComponent } from './editor/editor.component';
import { LoginComponent } from './auth/login/login.component';
import { RegisterComponent } from './auth/register/register.component';
import { MealsComponent } from './meals/meals.component';
import { AuthGuardService } from './core/services/auth-guard.service';
import { NoAuthGuardService } from './core/services/no-auth-guard.service';


const routes: Routes = [
  { path: '', redirectTo: '/recipes', pathMatch: 'full' },

  { path: 'recipes', component: RecipesComponent },

  { path: 'recipe/:id', component: RecipeDetailsComponent },

  { path: 'editor', component: EditorComponent,
    canActivate: [AuthGuardService] },

  { path: 'signin', component: LoginComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'signup', component: RegisterComponent,
    canActivate: [NoAuthGuardService] },

  { path: 'meals', component: MealsComponent,
    canActivate: [AuthGuardService] }
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})


export class AppRoutingModule { }