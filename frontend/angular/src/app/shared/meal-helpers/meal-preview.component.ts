import { Component, Input } from '@angular/core';

import { Meal } from '../../core/interfaces/meal';


@Component({
  selector: 'app-meal-preview',
  templateUrl: './meal-preview.component.html',
  styleUrls: ['./meal-preview.component.css']
})

export class MealPreviewComponent {

  constructor() { }


  @Input() meal: Meal;

}
