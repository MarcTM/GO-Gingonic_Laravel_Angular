import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { Meal } from '../interfaces/meal';


@Injectable({
  providedIn: 'root'
})


export class MealService {

  constructor(private http: HttpClient) { }


  // Create meal
  save(data) {
    return this.http.post<any>(environment.laravel_url+'/meals/', {"meal":data}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer_lar')}`})})
  }

  // Get meals
  getMeals(): Observable<Meal[]> {
    return this.http.get<Meal[]>(environment.laravel_url+'/meals/');
  }

  // Get meal by Id
  getMeal(id: number): Observable<Meal> {
    return this.http.get<Meal>(`${environment.laravel_url}/meals/${id}`)
  }
  
}
