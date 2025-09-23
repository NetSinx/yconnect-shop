import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { matchPassword } from './confirmpassword.validator';
import { RegisterService } from 'src/app/services/register/register.service';
import { User } from 'src/app/interfaces/user';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';
import { LoadingService } from 'src/app/services/loading/loading.service';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css'],
    standalone: false
})

export class RegisterComponent implements OnInit {
  formGroup: FormGroup = new FormGroup({})
  formBuilder: FormBuilder = new FormBuilder
  errorRegister: string | undefined
  successRegister: string | undefined

  ngOnInit(): void {
    this.csrfService.getCSRF().subscribe()
  }

  constructor(
    private registerService: RegisterService,
    private csrfService: GenCsrfService,
    private loadingService: LoadingService
  ) {
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

  public registerUser(): void {
    this.csrfService.getCSRF().subscribe()
    
    let dataUser: User = {
      name: "",
      username: "",
      avatar: "",
      email: "",
      alamat: "",
      no_telp: "",
      password: ""
    }

    dataUser.name = this.formGroup.value.name
    dataUser.username = this.formGroup.value.username
    dataUser.email = this.formGroup.value.email
    dataUser.alamat = this.formGroup.value.alamat
    dataUser.no_telp = this.formGroup.value.no_tlp
    dataUser.password = this.formGroup.value.password

    this.registerService.registerUser(dataUser).subscribe(
      () => {
        this.loadingService.setLoading(false)
        this.successRegister = "Registrasi akun berhasil!"
        this.errorRegister = ""
      },
      error => {
        this.loadingService.setLoading(false)
        
        if (error.status === 409) {
          this.errorRegister = "Registrasi akun gagal! User sudah pernah dibuat."
          this.successRegister = ""
        } else {
          this.errorRegister = "Registrasi akun gagal!"
          this.successRegister = ""
        }
      }
    )
  }
}
