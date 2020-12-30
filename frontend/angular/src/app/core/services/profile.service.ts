import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../../../environments/environment';

@Injectable({
  providedIn: 'root'
})


export class ProfileService {

  constructor(private http: HttpClient) { }


  // Already following
  isFollowing(username) {
    return this.http.post(environment.api_users_url+'/users/following/'+username, {}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

  // Follow user
  follow(id) {
    return this.http.put(environment.api_users_url+'/users/follow/'+id, {}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

  // Unfollow user
  unfollow(id) {
    return this.http.put(environment.api_users_url+'/users/unfollow/'+id, {}, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})});
  }

}
