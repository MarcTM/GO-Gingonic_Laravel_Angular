import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';

import { JwtService } from '../jwt.service';


@Injectable({
  providedIn: 'root'
})


export class NoAuthGuardService {

  constructor(private jwtService: JwtService) { }

  canActivate(): boolean {
    if(this.jwtService.getToken('Bearer')){
      return false;
    } else {
      return true;
    }
  }

}
