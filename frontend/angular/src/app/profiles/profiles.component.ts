import { Component, OnInit } from '@angular/core';
import { Profile } from '../core/interfaces/user';
import { ActivatedRoute, Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';

import { UserService } from '../core/services/user.service';
import { ProfileService } from '../core/services/profile.service';

@Component({
  selector: 'app-profiles',
  templateUrl: './profiles.component.html',
  styleUrls: ['./profiles.component.css']
})

export class ProfilesComponent implements OnInit {

  profile: Profile;
  me: boolean = false;
  following: boolean = false;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private userService: UserService,
    private profileService: ProfileService,
    private toastr: ToastrService) {}


  // Get one user profile
  getProfile(): void {
    const username = this.route.snapshot.paramMap.get('username');
    this.userService.getProfile(username)
    .subscribe(
      profile => {
          console.log(profile);
          this.profile = profile;
      },
      error => {this.router.navigate(['/recipes'])}
    )

    if(localStorage.getItem('Bearer')) {
      this.profileService.me()
      .subscribe(
        response=>{
          (response.username===username) ? this.me = true : this.me = false;

          if(!this.me) {
            this.profileService.isFollowing(username)
            .subscribe(
              res => {this.following = true},
              error => {this.following = false}
            )
          }
          
        }
      );
    }
  }

  // Follow user
  follow(id) {
    this.profileService.follow(id)
    .subscribe(
      response => {this.following = true},
      error => {this.toastr.error('You have to login to do this')}
    )
  }

  // Follow user
  unfollow(id) {
    this.profileService.unfollow(id)
    .subscribe(
      response => {this.following = false},
      error => {this.toastr.error('You have to login to do this')}
    )
  }

  ngOnInit(): void {
    this.getProfile();
  }

}
