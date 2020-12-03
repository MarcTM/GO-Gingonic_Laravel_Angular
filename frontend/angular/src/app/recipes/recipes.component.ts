import { Component, OnInit } from '@angular/core';

import { Recipe } from '../core/interfaces/recipe';
import { RecipeService } from '../core/services/recipe.service';


@Component({
  selector: 'app-recipes',
  templateUrl: './recipes.component.html',
  styleUrls: ['./recipes.component.css']
})
export class RecipesComponent implements OnInit {

  constructor(private recipeService: RecipeService) {}

  recipes: Recipe[];
  
  getRecipes(): void {
    this.recipeService.getRecipes().subscribe(
      recipes => {
        this.recipes = recipes;
      }
    );
  }


  ngOnInit(): void {
    this.getRecipes();
  }

}