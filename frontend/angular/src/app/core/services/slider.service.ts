import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { environment } from '../../../environments/environment';
import { Slider } from '../interfaces/slider';

@Injectable({
  providedIn: 'root'
})

export class SliderService {

  constructor(private http: HttpClient) {}


  // Get slides
  getSlides(): Observable<Slider[]> {
    return this.http.get<Slider[]>(environment.laravel_url+'/slider/');
  }
}
