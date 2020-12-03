import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class JwtService {

  constructor() { }

  getToken(): String {
    return window.localStorage['Bearer'];
  }

  saveToken(token: String) {
    window.localStorage['Bearer'] = token;
  }

  destroyToken() {
    window.localStorage.removeItem('Bearer');
  }

}
