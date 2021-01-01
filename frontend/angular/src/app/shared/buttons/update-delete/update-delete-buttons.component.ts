import { Component, OnInit, Input } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';

import { ProfileService } from '../../../core/services/profile.service';
import { RecipeService } from '../../../core/services/recipe.service';

@Component({
  selector: 'app-update-delete-buttons',
  templateUrl: './update-delete-buttons.component.html',
  styleUrls: ['./update-delete-buttons.component.css']
})

export class UpdateDeleteButtonsComponent {

  constructor(
    private router: Router,
    private recipeService: RecipeService,
  ) {}


  @Input() recipe: number;

  // Delete recipe
  delete(id) {
    this.recipeService.delete(id)
    .subscribe(
      response=>{this.router.navigate(['/recipes'])}
    )
  }

  // Update recipe
  update(id) {
    this.router.navigate(['editor/'+id]);
  }


}
