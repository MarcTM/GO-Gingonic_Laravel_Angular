import { Component, OnInit, Input } from '@angular/core';
import { Profile } from '../core/interfaces/user';
import { ActivatedRoute, Router } from '@angular/router';

import { UserService } from '../core/services/user.service';

@Component({
  selector: 'app-profiles',
  templateUrl: './profiles.component.html',
  styleUrls: ['./profiles.component.css']
})

export class ProfilesComponent implements OnInit {

  profile: Profile;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private userService: UserService,
  ) {}


  ngOnInit(): void {
    this.getProfile()
  }


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
  }


}
