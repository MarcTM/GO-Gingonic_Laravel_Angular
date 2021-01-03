import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { Category } from '../interfaces/category';

@Injectable({
  providedIn: 'root'
})

export class CategoryService {

  constructor(private http: HttpClient) {} 


  // Create category
  save(data) {
    return this.http.post<any>(environment.laravel_url+'/categories/', {"category":data}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer_lar')}`})})
  }

  // Get categories
  getCategories(): Observable<Category[]> {
    return this.http.get<Category[]>(environment.laravel_url+'/categories/');
  }

}