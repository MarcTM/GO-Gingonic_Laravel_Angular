import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { UserService } from 'src/app/core/services/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent{

    constructor(private userService: UserService,
      private router: Router) { }

    // Logout
    logout() {
      this.userService.purgeAuth('Bearer');
      this.userService.purgeAuth('Bearer_lar');
      this.router.navigate(['/']); 
    }

    currentUser: boolean;
}
