import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../../../environments/environment';
import { User } from '../interfaces/user';


@Injectable({
  providedIn: 'root'
})

export class ProfileService {

  constructor(private http: HttpClient) { }

  
  // Return my user
  me(){
    return this.http.get<User>(environment.api_users_url+'/users/me', {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

  // Already following
  isFollowing(username) {
    return this.http.post(environment.api_users_url+'/users/following/'+username, null, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

  // Follow user
  follow(id) {
    return this.http.put(environment.api_users_url+'/users/follow/'+id, null, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

  // Unfollow user
  unfollow(id) {
    return this.http.put(environment.api_users_url+'/users/unfollow/'+id, null, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

}
