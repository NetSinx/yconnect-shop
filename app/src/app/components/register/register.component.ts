import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { matchingPasswordValidator } from './confirmpassword.validator';
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
  userProfileForm: FormGroup
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
    this.userProfileForm = new FormGroup({
      nama_lengkap: new FormControl('', [Validators.required]),
      username: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      alamat: new FormGroup({
        nama_jalan: new FormControl('', [Validators.required]),
        rt: new FormControl(0, [Validators.required, Validators.min(1)]),
        rw: new FormControl(0, [Validators.required, Validators.min(1)]),
        kelurahan: new FormControl('', [Validators.required]),
        kecamatan: new FormControl('', [Validators.required]),
        kota: new FormControl('', [Validators.required]),
        kode_pos: new FormControl(0, [Validators.required, Validators.minLength(5)])
      }),
      no_telp: new FormControl('', [Validators.required, Validators.minLength(8)]),
      password: new FormControl('', [Validators.required, Validators.minLength(5)]),
      confirmpassword: new FormControl('', [Validators.required, Validators.minLength(5)])
    }, { validators: matchingPasswordValidator })
  }

  get nama_lengkap() {return this.userProfileForm.get("nama_lengkap")}
  get username() {return this.userProfileForm.get("username")}
  get email() {return this.userProfileForm.get("email")}
  get alamat() {return this.userProfileForm.get("alamat")}
  get no_telp() {return this.userProfileForm.get("no_telp")}
  get password() {return this.userProfileForm.get("password")}
  get confirmpassword() {return this.userProfileForm.get("confirmpassword")}

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

    dataUser.name = this.userProfileForm.value.name
    dataUser.username = this.userProfileForm.value.username
    dataUser.email = this.userProfileForm.value.email
    dataUser.alamat = this.userProfileForm.value.alamat
    dataUser.no_telp = this.userProfileForm.value.no_tlp
    dataUser.password = this.userProfileForm.value.password

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
