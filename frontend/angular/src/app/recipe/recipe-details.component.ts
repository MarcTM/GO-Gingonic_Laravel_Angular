import { Component, OnInit } from '@angular/core';
import { Recipe } from '../core/interfaces/recipe';
import { ActivatedRoute } from '@angular/router';
import { FormBuilder, Validators } from '@angular/forms';
import { Location } from '@angular/common';
import { Router } from "@angular/router"

import { Comment } from '../core/interfaces/comment';
import { RecipeService } from '../core/services/recipe.service';
import { ProfileService } from '../core/services/profile.service';

@Component({
  selector: 'app-recipe-details',
  templateUrl: './recipe-details.component.html',
  styleUrls: ['./recipe-details.component.css']
})

export class RecipeDetailsComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private fb: FormBuilder,
    private recipeService: RecipeService,
    private profileService: ProfileService,
    private location: Location,
    private router: Router
  ) {}


  recipe: Recipe;
  comments: Comment[];
  me: boolean = false;

  // Add comment form
  commentForm = this.fb.group({
    body: ['', Validators.required],
  });
    

  // Submit form to add comment to recipe
  submitForm() {
    this.recipeService.saveComment(this.recipe.id, this.commentForm.value)
    .subscribe(
      response => {
        console.log(response);
        this.getComments();
      },
      error => { console.log(error); }
    )
  }

  // Get one recipe information and if it is mine or not
  getRecipe(): void {
    const id = +this.route.snapshot.paramMap.get('id');

    this.recipeService.getRecipe(id)
    .subscribe(recipe => {
      this.recipe = recipe;

      if(localStorage.getItem('Bearer')) {
        this.profileService.me()
        .subscribe(
          response=>{ (response.username===recipe.author.username) ? this.me = true : this.me = false; }
        );
      }

      // Get comments
      this.getComments();
    })
  }

  // Get recipe comments
  getComments() {
    this.recipeService.getComments(this.recipe.id)
      .subscribe(
        comments => { 
          console.log(comments);
          this.comments = comments;
        },
        error => { console.log(error); }
      )
  }

  // Go back to the previous page
  goBack() {
    this.location.back();
  }

  ngOnInit() {
    this.getRecipe();
  }

}
