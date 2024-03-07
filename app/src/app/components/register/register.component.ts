import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { matchPassword } from './confirmpassword.validator';
import { RegisterService } from 'src/app/services/register/register.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})

export class RegisterComponent {
  formGroup: FormGroup = new FormGroup({})
  formBuilder: FormBuilder = new FormBuilder

  constructor(private registerService: RegisterService) {
    this.formGroup = this.formBuilder.group({
      name: new FormControl('', [Validators.required, Validators.minLength(3)]),
      username: new FormControl('', [Validators.required, Validators.minLength(3)]),
      email: new FormControl('', [Validators.required, Validators.email]),
      alamat: new FormControl('', [Validators.required, Validators.minLength(5)]),
      no_tlp: new FormControl('', [Validators.required, Validators.minLength(8)]),
      password: new FormControl('', [Validators.required, Validators.minLength(5)]),
      confirmpassword: new FormControl('', [Validators.required])
    }, {
      validator: matchPassword('password', 'confirmpassword')
    })
  }

  get f() {return this.formGroup.controls}

  public registUser() {
    console.warn(this.formGroup.value.password)
  }
}
