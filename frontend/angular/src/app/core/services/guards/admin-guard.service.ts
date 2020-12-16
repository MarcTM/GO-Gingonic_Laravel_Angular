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
        return true;
      } else {
        this.router.navigate(['signin']);
        return false;
      }
    }
    
}