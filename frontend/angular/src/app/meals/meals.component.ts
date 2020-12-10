import { Component, OnInit } from '@angular/core';

import { Meal } from '../core/interfaces/meal';
import { MealService } from '../core/services/meal.service';


@Component({
  selector: 'app-meals',
  templateUrl: './meals.component.html',
  styleUrls: ['./meals.component.css']
})

export class MealsComponent implements OnInit {

  meals: Meal[];
  
  getMeals(): void {
    this.mealService.getMeals()
        .subscribe(meals => {
          console.log(meals);
          this.meals = meals;
        });
  }

  constructor(private mealService: MealService) {}

  ngOnInit(): void {
    this.getMeals();
  }

}
