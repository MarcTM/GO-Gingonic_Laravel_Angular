import { Component, OnInit } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router, ActivatedRoute } from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { UserService } from '../../core/services/user.service';
import { Profile } from '../../core/interfaces/user';

@Component({
  selector: 'app-update-account',
  templateUrl: './update-account.component.html',
  styleUrls: ['./update-account.component.css']
})

export class UpdateAccountComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private userService: UserService,
    private toastr: ToastrService
  ) {}

  
  account: Profile;
  updateForm;

  // Submit form
  submitUpdate() {
    let updateAccount = {'id':this.account.id, 'username':this.account.username, 'bio':this.updateForm.value.bio, 'image':this.updateForm.value.image};

    this.userService.update(updateAccount)
    .subscribe(
      response => {
        this.toastr.success('Updated successfully');
        setTimeout(() => {this.router.navigate(['/account'])}, 1000);
      },
      error => {
        this.toastr.error(error.error);
      }
    )
  }

  // Get account profile
  getAccount() {
    this.userService.getAccount()
    .subscribe(
      account => {
        this.account = account;

        this.updateForm = this.fb.group({
          image: [this.account.image, Validators.required],
          bio: [this.account.bio, Validators.required],
        });
      },
      error => {this.router.navigate(['account'])}
    )
  }
  
  ngOnInit(): void {
    this.getAccount();
  }

}
