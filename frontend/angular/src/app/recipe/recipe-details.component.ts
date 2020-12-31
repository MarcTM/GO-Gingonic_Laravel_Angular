import { Component, OnInit, Input } from '@angular/core';
import { Recipe } from '../core/interfaces/recipe';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';

import { RecipeService } from '../core/services/recipe.service';


@Component({
  selector: 'app-recipe-details',
  templateUrl: './recipe-details.component.html',
  styleUrls: ['./recipe-details.component.css']
})

export class RecipeDetailsComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private recipeService: RecipeService,
    private location: Location
  ) {}


  recipe: Recipe;

  // Get one recipe information
  getRecipe(): void {
    const id = +this.route.snapshot.paramMap.get('id');

    this.recipeService.getRecipe(id)
    .subscribe(recipe => {
      console.log(recipe);
      this.recipe = recipe;
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
