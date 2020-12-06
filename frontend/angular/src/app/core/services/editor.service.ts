import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { Recipe } from '../interfaces/recipe';


@Injectable({
  providedIn: 'root'
})


export class EditorService {

  constructor(private http: HttpClient) { }


  uploadRecipe(data) {
    return this.http.post<any>(environment.api_url+'/recipes', data)
  }

}
