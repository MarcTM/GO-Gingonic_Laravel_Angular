import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { CategoryMeals } from '../core/interfaces/category';
import { CategoryService } from '../core/services/category.service';

@Component({
  selector: 'app-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})

export class CategoryComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private categoryService: CategoryService
  ) {}


  category: CategoryMeals;

  // Get category meals
  getCategory(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    
    this.categoryService.getMeals(id)
    .subscribe(response => { this.category = response; })
  }

  ngOnInit(): void {
    this.getCategory();
  }

}
