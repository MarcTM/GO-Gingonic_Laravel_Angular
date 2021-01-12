import { Component, OnInit } from '@angular/core';

import { UserService } from '../core/services/user.service';
import { Profile } from '../core/interfaces/user';


@Component({
  selector: 'app-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.css']
})

export class AccountComponent implements OnInit {

  constructor(
    private userService: UserService
  ) {}


  account: Profile;

  // Update profile button
  updateProfile() {
    console.log(this.account.id)
  }

  // Get your profile
  getAccount() {
    this.userService.getAccount()
    .subscribe(
      account => {
        console.log(account)
        this.account = account;
      }
    )
  }

  ngOnInit(): void {
    this.getAccount();
  }

}