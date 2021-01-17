import { FormBuilder, FormArray, Validators } from '@angular/forms';
import {Router} from "@angular/router"
import { Component, OnInit } from '@angular/core';
import { ToastrService } from 'ngx-toastr';

import { UserService } from 'src/app/core/services/user.service';


@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})

export class RegisterComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private userService: UserService,
    private toastr: ToastrService
  ) {}


  // Register form fields
  registerForm = this.fb.group({
    username: ['', Validators.required],
    email: ['', Validators.required],
    password: ['', Validators.required],
  });

  
  // Submit register
  submitRegister() {
    this.userService.attemptAuth('register', this.registerForm.value)
      .subscribe(
        response => {
          this.toastr.success('Registration successfull')
          setTimeout(() => {this.router.navigate(['/signin'])}, 1000);
        },
        error => {this.toastr.error(error.error)}
      )
  }

  ngOnInit(): void {
  }

}
