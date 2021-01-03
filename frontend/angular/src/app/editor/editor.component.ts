import { Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { RecipeService } from '../core/services/recipe.service';

@Component({
  selector: 'app-editor',
  templateUrl: './editor.component.html',
  styleUrls: ['./editor.component.css']
})

export class EditorComponent {

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private recipeService: RecipeService,
    private toastr: ToastrService
  ) {}


  // Editor form fields
  editorForm = this.fb.group({
    name: ['', Validators.required],
    description: ['', Validators.required],
    image: ['', Validators.required],
  });

  // Submit form
  submitEditor() {
    this.recipeService.save(this.editorForm.value)
    .subscribe(
      response => {
        this.toastr.success('Created successfully');
        setTimeout(() => {this.router.navigate(['/recipes'])}, 1000);
      },
      error => {
        this.toastr.error(error.error);
      }
    )
  }

}
