import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute } from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { RecipeService } from '../../core/services/recipe.service';
import { Recipe } from '../../core/interfaces/recipe';

@Component({
  selector: 'app-editor-update',
  templateUrl: './editor-update.component.html',
  styleUrls: ['./editor-update.component.css']
})

export class EditorUpdateComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private recipeService: RecipeService,
    private toastr: ToastrService
  ) {}

  
  recipe: Recipe;

  editorForm;

  // Submit form
  submitEditor() {
    let updateRecipe = {'id':this.recipe.id, 'name':this.editorForm.value.name, 'description':this.editorForm.value.description};

    this.recipeService.update(updateRecipe)
    .subscribe(
      response => {
        this.toastr.success('Updated successfully');
        setTimeout(() => {this.router.navigate(['/recipe/'+updateRecipe.id])}, 1000);
      },
      error => {
        this.toastr.error(error.error);
      }
    )
  }

  // Get recipe
  validateRecipe(): void {
    const id = +this.route.snapshot.paramMap.get('id');
    
    this.recipeService.owns(id)
    .subscribe(
      recipe => {
        this.recipe = recipe

        this.editorForm = this.fb.group({
          name: [this.recipe.name, Validators.required],
          description: [this.recipe.description, Validators.required],
        });
      },
      error => {this.router.navigate(['recipes'])}
    )
  }

  
  ngOnInit(): void {
    this.validateRecipe();
  }

}
