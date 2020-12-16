import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormArray, Validators } from '@angular/forms';
import { Router } from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { UserService } from '../../core/services/user.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})


export class LoginComponent {

  constructor(private fb: FormBuilder,
    private router: Router,
    private userService: UserService,
    private toastr: ToastrService) { }


  // Login form
  loginForm = this.fb.group({
    email: ['', Validators.required],
    password: ['', Validators.required],
  });


  // Submit login
  submitLogin() {
      this.userService.attemptAuth('login', this.loginForm.value)
      .subscribe(
        response => {
          // Type user
            if (response.user.type==="user"){

              this.toastr.success('Logged in')
              console.log(response);
              setTimeout(() => {this.router.navigate(['/'])}, 1000);
        
          // Type admin
            } else if (response.user.type==="admin") {

              this.userService.attemptAuthAdmin(this.loginForm.value)
              .subscribe(
                response => {
                  console.log(response)
                  this.toastr.success('Logged in')
                  setTimeout(() => {this.router.navigate(['/'])}, 1000);
                }
              )
            }
        },

        error => {this.toastr.error(error.error, 'Invalid credentials')}
      )
  }
}
