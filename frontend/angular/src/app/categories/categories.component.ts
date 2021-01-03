import { Component, OnInit } from '@angular/core';
import { Category } from '../core/interfaces/category';
import { ActivatedRoute, Router } from '@angular/router';

import { CategoryService } from '../core/services/category.service';

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.css']
})

export class CategoriesComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private categoryService: CategoryService
  ) {}


  categories: Category[];

  // Get categories
  getCategories() {
    this.categoryService.getCategories()
    .subscribe(
      categories => {
          this.categories = categories;
      }
    )
  }

  ngOnInit(): void {
    this.getCategories();
  }

}
