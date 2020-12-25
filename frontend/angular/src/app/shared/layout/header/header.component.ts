import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { UserService } from 'src/app/core/services/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})

export class HeaderComponent{

    // Admin variables
    admin: boolean = false;
    adminText: string = 'SHOW PANNEL ADMIN'

    constructor(private userService: UserService,
      private router: Router) { }

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

}
