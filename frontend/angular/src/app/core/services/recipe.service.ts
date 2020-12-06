import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { Recipe } from '../interfaces/recipe';


@Injectable({
  providedIn: 'root'
})
export class RecipeService {

  constructor(private http: HttpClient) { }


  // Get recipes
  getRecipes(): Observable<Recipe[]> {
    return this.http.get<Recipe[]>(environment.api_url+'/recipes');
  }


  // Get recipe by Id
  getRecipe(id: number): Observable<Recipe> {
    return this.http.get<Recipe>(`${environment.api_url}/recipes/${id}`)
  }


  // Create recipe
  save(data) {
    return this.http.post<any>(environment.api_url+'/recipes', data)
  }

}
