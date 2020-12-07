import { Component, OnInit } from '@angular/core';

import { Recipe } from '../../core/interfaces/recipe';
import { RecipeService } from '../../core/services/recipe.service';

@Component({
  selector: 'app-recipe-list',
  templateUrl: './recipe-list.component.html',
  styleUrls: ['./recipe-list.component.css']
})


export class RecipeListComponent implements OnInit {

  constructor(private recipeService: RecipeService) {}


  recipes: Recipe[];

  getRecipes(): void {
    this.recipeService.getRecipes().subscribe(
      recipes => {
        console.log(recipes)
        this.recipes = recipes;
      }
    );
  }

  ngOnInit(): void {
    this.getRecipes();
  }

}
