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
    return this.http.get<Recipe>(`${environment.api_recipes_url}/recipes/recipe/${id}`)
  }


  // Already favorited
  isFavorited(id) {
    return this.http.post(environment.api_recipes_url+'/recipes/favorited/'+id, {}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

  // Favorite recipe
  favorite(id) {
    return this.http.put(environment.api_recipes_url+'/recipes/favorite/'+id, {}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

  // Unfavorite recipe
  unfavorite(id) {
    return this.http.put(environment.api_recipes_url+'/recipes/unfavorite/'+id, {}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

}
