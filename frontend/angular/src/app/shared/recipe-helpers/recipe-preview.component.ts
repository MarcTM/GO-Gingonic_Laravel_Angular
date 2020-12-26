import { Component, Input } from '@angular/core';

import { Recipe } from '../../core/interfaces/recipe';


@Component({
  selector: 'app-recipe-preview',
  templateUrl: './recipe-preview.component.html',
  styleUrls: ['./recipe-preview.component.css']
})

export class RecipePreviewComponent {

  constructor() { }


  @Input() recipe: Recipe;

}
