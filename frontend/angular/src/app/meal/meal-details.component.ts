import { Component, OnInit, Input } from '@angular/core';
import { Meal } from '../core/interfaces/meal';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';

import { MealService } from '../core/services/meal.service';


@Component({
  selector: 'app-meal-details',
  templateUrl: './meal-details.component.html',
  styleUrls: ['./meal-details.component.css']
})

export class MealDetailsComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private mealService: MealService,
    private location: Location
  ) {}


  @Input() meal: Meal;

  // Get one meal information
  getMeal(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    
    this.mealService.getMeal(id)
    .subscribe(meal => {
      console.log(meal);
      this.meal = meal;
    })
  }

  // Go back to the previous page
  goBack(): void {
    this.location.back();
  }

  ngOnInit(): void {
    this.getMeal();
  }

}
