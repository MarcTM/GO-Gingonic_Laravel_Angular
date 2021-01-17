import { Component, OnInit } from '@angular/core';

import { Meal } from '../../core/interfaces/meal';
import { MealService } from '../../core/services/meal.service';

@Component({
  selector: 'app-meal-list',
  templateUrl: './meal-list.component.html',
  styleUrls: ['./meal-list.component.css']
})

export class MealListComponent implements OnInit {

  constructor(private mealService: MealService) {}


  meals: Meal[];
  
  actualPage: number = 1;

  getMeals(): void {
    this.mealService.getMeals()
    .subscribe(meals => { this.meals = meals; });
  }

  ngOnInit(): void {
    this.getMeals();
  }

}
