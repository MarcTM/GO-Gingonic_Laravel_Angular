import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { environment } from '../../../../environments/environment';
import { JwtService } from '../jwt.service';

@Injectable({
  providedIn: 'root'
})

export class AdminGuardService implements CanActivate {

  constructor(private jwtService: JwtService,
    private router: Router,
    private http: HttpClient) { }

    canActivate(): boolean {
      if(this.jwtService.getToken('Bearer_lar')){
        this.http.post(environment.laravel_url+'/auth/validate', null, {headers: new HttpHeaders({'Authorization': `Bearer ${localStorage.getItem('Bearer_lar')}`})})
        .subscribe(
          valid => {},
          error => {
            localStorage.removeItem('Bearer_lar')
            this.router.navigate(['']);
          }
        );
        return true;

      } else {
        (this.jwtService.getToken('Bearer')) ? this.router.navigate(['']) : this.router.navigate(['signin']);
        return false;
      }
    }
    
}