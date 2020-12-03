import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { Recipe } from '../interfaces/recipe';
import { RECIPES } from '../mocks/mock-recipes';


@Injectable({
  providedIn: 'root'
})
export class RecipeService {

  getRecipes(): Observable<Recipe[]> {
    return this.http.get<Recipe[]>(environment.api_url+'/recipes');
  }

  getRecipe(id: number): Observable<Recipe> {
    return this.http.get<Recipe>(`${environment.api_url}/recipes/${id}`)
  }

  constructor(private http: HttpClient) { }
}
