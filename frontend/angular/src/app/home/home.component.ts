import { Component, OnInit } from '@angular/core';

import { Profile } from '../core/interfaces/user';
import { UserService } from '../core/services/user.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})


export class HomeComponent implements OnInit {

  constructor(
    private userService: UserService
  ) {}


  best: Profile;

  // Get user with more recipes in the app
  getBestuser() {
    this.userService.bestuser()
    .subscribe(
      response => {
        console.log(response);
        this.best = response;
      },
      error => { console.log(error); }
    )
  }

  ngOnInit(): void {
    this.getBestuser();
  }

}
