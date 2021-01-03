import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { CategoryService } from '../../core/services/category.service';

@Component({
  selector: 'app-create-category',
  templateUrl: './create-category.component.html',
  styleUrls: ['./create-category.component.css']
})

export class CreateCategoryComponent {

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private categoryService: CategoryService,
    private toastr: ToastrService
  ) { }


  creatorForm = this.fb.group({
    name: ['', Validators.required],
    image: ['', Validators.required]
  });

  submitCreator() {
    this.categoryService.save(this.creatorForm.value)
    .subscribe(
      response => {
        this.toastr.success('Created successfully');
        setTimeout(() => {this.router.navigate(['/'])}, 1000);
      },
      error => {
        this.toastr.error(error.error);
      }
    )
  }

}
