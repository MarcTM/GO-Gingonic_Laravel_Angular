import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { Recipe } from '../interfaces/recipe';


@Injectable({
  providedIn: 'root'
})
export class RecipeService {

  constructor(private http: HttpClient) { }


  // Create recipe
  save(data) {
    return this.http.post<any>(environment.api_recipes_url+'/recipes/', { recipe: data }, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})})
  }

  // Get recipes
  getRecipes(): Observable<Recipe[]> {
    return this.http.get<Recipe[]>(environment.api_recipes_url+'/recipes/');
  }


  // Get recipe by Id
  getRecipe(id: number): Observable<Recipe> {
    return this.http.get<Recipe>(`${environment.api_recipes_url}/recipes/${id}`)
  }

}
