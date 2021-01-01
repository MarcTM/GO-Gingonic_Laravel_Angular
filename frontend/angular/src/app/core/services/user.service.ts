import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { User, Profile } from '../interfaces/user';
import { JwtService } from './jwt.service';
import { environment } from '../../../environments/environment';
import { nextTick } from 'process';


@Injectable({
  providedIn: 'root'
})


export class UserService {

  constructor(
    private jwtService: JwtService,
    private http: HttpClient
  ) {}

  
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


  // Returns an admin token or not, to know if user is authenticated as admin
  isAdmin(){
    return this.jwtService.getToken('Bearer_lar');
  }


  // See if token is still valid or not
  validateToken() {
    return this.http.post(environment.api_users_url+'/auth/validate', null, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})})
  }


  // Attempt to login or register
  attemptAuth(type, credentials) {
    const attemptType = (type === 'login') ? 'login' : 'register';
    return this.http.post<any>(environment.api_users_url + '/users/' + attemptType, {user: credentials})
    .pipe(map(
        data => {
          if (type === 'login') {
            this.setAuth('Bearer', data.user.bearer)
          };
          return data;
        }
    ));
  }


  // Attempt to login when user is admin
  attemptAuthAdmin(credentials) {
    return this.http.post<any>(environment.laravel_url+'/auth/login', credentials)
    .pipe(map(
        data => {
          this.setAuth('Bearer_lar', data.token);
          return data;
        }
    ));
  }


  // Get user profile by username
  getProfile(username) {
    return this.http.get<Profile>(environment.api_users_url+'/profiles/'+username)
    .pipe(map(
      data => {
        return data;
      }
    ))
  }

}
