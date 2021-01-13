import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";

import { UserService } from '../core/services/user.service';
import { Profile } from '../core/interfaces/user';


@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})

export class AccountComponent implements OnInit {

  constructor(
    private router: Router,
    private userService: UserService
  ) {}


  account: Profile;

  // Update profile button
  updateProfile() {
    this.router.navigate(['/account/update'])
  }

  // Get your profile
  getAccount() {
    this.userService.getAccount()
    .subscribe(
      account => {
        this.account = account;
      }
    )
  }

  ngOnInit(): void {
    this.getAccount();
  }

}