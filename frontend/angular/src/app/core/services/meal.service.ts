import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { Meal } from '../interfaces/meal';


@Injectable({
  providedIn: 'root'
})


export class MealService {

  constructor(private http: HttpClient) { }


  getMeals(): Observable<Meal[]> {
    return this.http.get<Meal[]>(environment.laravel_url+'/meals/');
  }
  
}
