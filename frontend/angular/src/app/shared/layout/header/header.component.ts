import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { UserService } from 'src/app/core/services/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})

export class HeaderComponent implements OnInit{

    constructor(
      private userService: UserService,
      private router: Router
    ) { }


    // Admin variables
    admin: boolean = false;
    adminText: string = 'SHOW PANNEL ADMIN'

    // Logout
    logout() {
      this.userService.purgeAuth('Bearer');
      this.userService.purgeAuth('Bearer_lar');
      this.router.navigate(['/']);
      this.admin = false;
    }

    // Activate and desactivate pannel admin
    showAdmin() {
      if (this.admin === true) {
        this.admin=false;
        this.adminText="SHOW PANNEL ADMIN";
      } else {
        this.admin=true;
        this.adminText="HIDE PANNEL ADMIN";
      }
    }

    // See if token is still valid or not
    validateToken() {
      if(localStorage.getItem('Bearer')) {
        this.userService.validateToken()
        .subscribe(
          response=>{console.log("User validated")},
          err=>{this.logout()}
        )
      }
    }

    ngOnInit(): void {
      this.validateToken();
    }

}
