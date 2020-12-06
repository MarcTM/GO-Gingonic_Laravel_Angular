import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormArray, Validators } from '@angular/forms';
import {Router} from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { UserService } from '../../core/services/user.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})


export class LoginComponent implements OnInit {

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
    let data = { "user": this.loginForm.value };

    this.userService.attemptAuth('login', data)
      .subscribe(
        response => {
          this.toastr.success('Logged in')
          console.log(response);
          this.router.navigate(['/']);
        },
        error => {
          this.toastr.error('Invalid credentials', 'Check your email and password', {
          timeOut: 3000
          })
        }
      )
  }


  ngOnInit(): void {
  }

}
