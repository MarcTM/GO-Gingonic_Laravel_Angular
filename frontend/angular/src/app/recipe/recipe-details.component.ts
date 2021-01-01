import { Component, OnInit } from '@angular/core';
import { Recipe } from '../core/interfaces/recipe';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';

import { RecipeService } from '../core/services/recipe.service';
import { ProfileService } from '../core/services/profile.service';

@Component({
  selector: 'app-recipe-details',
  templateUrl: './recipe-details.component.html',
  styleUrls: ['./recipe-details.component.css']
})

export class RecipeDetailsComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private recipeService: RecipeService,
    private profileService: ProfileService,
    private location: Location
  ) {}


  recipe: Recipe;
  me: boolean = false;

  // Get one recipe information and if it is mine or not
  getRecipe(): void {
    const id = +this.route.snapshot.paramMap.get('id');

    this.recipeService.getRecipe(id)
    .subscribe(recipe => {
      this.recipe = recipe;

      if(localStorage.getItem('Bearer')) {
        this.profileService.me()
        .subscribe(
          response=>{
            (response.username===recipe.author.username) ? this.me = true : this.me = false;
          }
        );
      }
      
    })
  }

  // Go back to the previous page
  goBack(): void {
    this.location.back();
  }

  ngOnInit(): void {
    this.getRecipe();
  }

}
