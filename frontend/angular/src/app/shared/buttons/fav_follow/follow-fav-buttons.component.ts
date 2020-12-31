import { Component, OnInit, Input } from '@angular/core';
import { ToastrService } from 'ngx-toastr';

import { ProfileService } from '../../../core/services/profile.service';
import { RecipeService } from '../../../core/services/recipe.service';

@Component({
  selector: 'app-follow-fav-buttons',
  templateUrl: './follow-fav-buttons.component.html',
  styleUrls: ['./follow-fav-buttons.component.css']
})

export class FollowFavButtonsComponent implements OnInit {

  constructor(
    private profileService: ProfileService,
    private recipeService: RecipeService,
    private toastr: ToastrService
  ) {}


  @Input() author_id: number;
  @Input() author_username: string;
  @Input() recipe: number;


  following: boolean = false;
  favorited: boolean = false;

  // Follow user
  follow($event) {
    $event.preventDefault();

    this.profileService.follow(this.author_id)
    .subscribe(
      response => {this.following = true},
      error => {this.toastr.error('You have to login to do this')}
    )
  }

  // Unfollow user
  unfollow($event) {
    $event.preventDefault();

    this.profileService.unfollow(this.author_id)
    .subscribe(
      response => {this.following = false},
      error => {this.toastr.error('You have to login to do this')}
    )
  }

  // Favorite recipe
  favorite($event) {
    $event.preventDefault();

    this.recipeService.favorite(this.recipe)
    .subscribe(
      response => {this.favorited = true},
      error => {this.toastr.error('You have to login to do this')}
    )
  }

  // Unfollow user
  unfavorite($event) {
    $event.preventDefault();

    this.recipeService.unfavorite(this.recipe)
    .subscribe(
      response => {this.favorited = false},
      error => {this.toastr.error('You have to login to do this')}
    )
  }


  // Already following?
  isFollowing() {
    this.profileService.isFollowing(this.author_username)
      .subscribe(
        res => {this.following = true},
        error => {this.following = false}
      )
  }

  // Already favorited?
  isFavorited() {
    this.recipeService.isFavorited(this.recipe)
      .subscribe(
        res => {this.favorited = true},
        error => {this.favorited = false}
      )
  }

  
  ngOnInit(): void {
    if(localStorage.getItem('Bearer')) {
      this.isFollowing();
      this.isFavorited();
    }
  }

}
