import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormArray, Validators } from '@angular/forms';
import { Router } from "@angular/router"
import { ToastrService } from 'ngx-toastr';

import { UserService } from '../../core/services/user.service';

@Component({
  selector: 'app-recover-password',
  templateUrl: './recover-password.component.html',
  styleUrls: ['./recover-password.component.css']
})

export class RecoverPasswordComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private userService: UserService,
    private toastr: ToastrService
  ) {}


  // Recover password form
  recoverForm = this.fb.group({
    email: ['', Validators.required],
  });

  // Submit recover password
  submitRecover() {
    console.log(this.recoverForm.value)
  }

  ngOnInit(): void {
  }

}
