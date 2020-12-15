import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';

import { JwtService } from './jwt.service';
import { environment } from '../../../environments/environment';


@Injectable({
  providedIn: 'root'
})


export class UserService {

  constructor(private jwtService: JwtService,
              private http: HttpClient) { }

  
  // Save JWT in localstorage
  setAuth(token) {
    this.jwtService.saveToken(token);
  }


  // Remove JWT from localstorage
  purgeAuth() {
    this.jwtService.destroyToken();
  }


  // Returns a token or not, to know if user is authenticated
  isAuthenticated(){
    return this.jwtService.getToken();
  }


  // Attempt to login or register
  attemptAuth(type, credentials) {
    const attemptType = (type === 'login') ? 'login' : 'register';
    return this.http.post<any>(environment.api_url + '/users/' + attemptType, {user: credentials})
    .pipe(map(
        data => {
          if (type === 'login') this.setAuth(data.user.bearer);
          return data;
        }
    ));
  }

  attemptAuthLaravel(credentials) {
    return this.http.post<any>(environment.laravel_url+'/auth/login', credentials)
    .pipe(map(
        data => {
          return data;
        }
    ));
  }

}
