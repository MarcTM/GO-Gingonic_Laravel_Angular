import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../../../../environments/environment';
import { JwtService } from '../jwt.service';


@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate {

  constructor(private jwtService: JwtService,
    private router: Router,
    private http: HttpClient) { }

    canActivate(): boolean {
      if(this.jwtService.getToken('Bearer')){
        // Check if token is valid in the server
        this.http.post(environment.api_url+'/auth/validate', null, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer')}`})})
        .subscribe(
          valid => {},
          error => {
            localStorage.removeItem('Bearer')
            localStorage.removeItem('Bearer_lar')
          }
        );
        if (this.jwtService.getToken('Bearer')){
          return true;
        } else {
          this.router.navigate(['signin']); 
          return false;
        }
      } else {
        this.router.navigate(['signin']);
        return false;
      }
    }
    
}
