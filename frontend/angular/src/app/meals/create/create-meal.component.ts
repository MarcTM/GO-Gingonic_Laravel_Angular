import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { MealService } from '../../core/services/meal.service';


@Component({
  selector: 'app-create-meal',
  templateUrl: './create-meal.component.html',
  styleUrls: ['./create-meal.component.css']
})

export class CreateMealComponent implements OnInit {

  constructor(private fb: FormBuilder,
    private router: Router,
    private mealService: MealService,
    private toastr: ToastrService) { }


  // Create meal form fields
  creatorForm = this.fb.group({
    name: ['', Validators.required],
    description: ['', Validators.required],
    price: ['', Validators.required]
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
          this.toastr.error(error.error);
        }
      )
  }


  ngOnInit(): void {
  }

}
