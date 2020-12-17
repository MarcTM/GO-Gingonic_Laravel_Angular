import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';

import { User } from '../interfaces/user';
import { JwtService } from './jwt.service';
import { environment } from '../../../environments/environment';
import { nextTick } from 'process';


@Injectable({
  providedIn: 'root'
})


export class UserService {

  constructor(private jwtService: JwtService,
              private http: HttpClient) { }

  
  // Save JWT in localstorage
  setAuth(name, token) {
    this.jwtService.saveToken(name, token);
  }


  // Remove JWT from localstorage
  purgeAuth(name) {
    this.jwtService.destroyToken(name);
  }


  // Returns a token or not, to know if user is authenticated
  isAuthenticated(){
    return this.jwtService.getToken('Bearer');
  }

  isAdmin(){
    return this.jwtService.getToken('Bearer_lar');
  }


  // Attempt to login or register
  attemptAuth(type, credentials) {
    const attemptType = (type === 'login') ? 'login' : 'register';
    return this.http.post<any>(environment.api_url + '/users/' + attemptType, {user: credentials})
    .pipe(map(
        data => {
          if (type === 'login') {
            this.setAuth('Bearer', data.user.bearer)
          };
          return data;
        }
    ));
  }

  attemptAuthAdmin(credentials) {
    return this.http.post<any>(environment.laravel_url+'/auth/login', credentials)
    .pipe(map(
        data => {
          this.setAuth('Bearer_lar', data.token);
          return data;
        }
    ));
  }

}
