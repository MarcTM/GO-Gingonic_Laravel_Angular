import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import { Category, CategoryMeals } from '../core/interfaces/category';
import { Meal } from '../core/interfaces/meal';
import { CategoryService } from '../core/services/category.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})

export class CategoryComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private categoryService: CategoryService
  ) {}


  category: CategoryMeals;

  // Get category meals
  getCategory(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    
    this.categoryService.getMeals(id)
    .subscribe(response => {
      console.log(response);
      this.category = response;
    })
  }

  ngOnInit(): void {
    this.getCategory();
  }

}
