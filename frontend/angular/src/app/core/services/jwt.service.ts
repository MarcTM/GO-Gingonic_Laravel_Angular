import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class JwtService {

  constructor() { }

  getToken(name): String {
    return window.localStorage[name];
  }

  saveToken(name, token: String) {
    window.localStorage[name] = token;
  }

  destroyToken(name) {
    window.localStorage.removeItem(name);
  }

}
