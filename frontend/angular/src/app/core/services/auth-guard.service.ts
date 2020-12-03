import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';

import { JwtService } from './jwt.service';


@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate {

  constructor(private jwtService: JwtService,
    private router: Router) { }


    canActivate(): boolean {
      if(this.jwtService.getToken()){
        return true;
      } else {
        this.router.navigate(['signin']);
        return false;
      }
    }
    
}
