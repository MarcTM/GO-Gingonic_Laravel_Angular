import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from "@angular/router"
import { ToastrService } from 'ngx-toastr';
import { CategoryService } from 'src/app/core/services/category.service';

import { MealService } from '../../core/services/meal.service';
import { Category } from '../../core/interfaces/category';


@Component({
  selector: 'app-create-meal',
  templateUrl: './create-meal.component.html',
  styleUrls: ['./create-meal.component.css']
})

export class CreateMealComponent implements OnInit {

  constructor(private fb: FormBuilder,
    private router: Router,
    private mealService: MealService,
    private toastr: ToastrService,
    private categoryService: CategoryService
  ) {}


  categories: Category[];

  // Get cateogires
  getCategories() {
    this.categoryService.getCategories()
    .subscribe(
      categories => {this.categories = categories}
    )
  }

  creatorForm = this.fb.group({
    name: ['', Validators.required],
    description: ['', Validators.required],
    price: ['', Validators.required],
    category: ['', Validators.required]
  });

  // Submit form
  submitCreator() {
    this.mealService.save(this.creatorForm.value)
    .subscribe(
      response => {
        this.toastr.success('Created successfully');
        setTimeout(() => {this.router.navigate(['/meals'])}, 1000);
      },
      error => {
        console.log(error.error);
      }
    )
  }

  ngOnInit(): void {
    this.getCategories();
  }

}
