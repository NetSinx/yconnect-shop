import { Component, Inject, makeStateKey, OnInit, PLATFORM_ID, TransferState } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { matchingPasswordValidator } from './confirmpassword.directive';
import { RegisterService } from 'src/app/services/register/register.service';
import { User } from 'src/app/interfaces/user';
import { GenCsrfService } from 'src/app/services/gen-csrf/gen-csrf.service';
import { LoadingService } from 'src/app/services/loading/loading.service';
import { isPlatformBrowser, isPlatformServer } from '@angular/common';

const CSRF_KEY = makeStateKey<string>("csrf-token")

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css'],
    standalone: false
})

export class RegisterComponent implements OnInit {
  userProfileForm: FormGroup = new FormGroup({
      nama_lengkap: new FormControl('', [Validators.required]),
      username: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      no_hp: new FormControl('', [Validators.required, Validators.minLength(8)]),
      password: new FormControl('', [Validators.required, Validators.minLength(5)]),
      confirmpassword: new FormControl('', [Validators.required, Validators.minLength(5)])
    }, { validators: matchingPasswordValidator() })
  errorRegister: string | undefined
  successRegister: string | undefined
  csrfToken: string = ""

  ngOnInit(): void {
    const saved = this.state.get(CSRF_KEY, null)
    if (saved) {
      this.csrfToken = saved
      console.log(saved)
    } else {
      this.csrfService.getCSRF().subscribe(data => {
        this.csrfToken = data.csrf_token

        if (isPlatformBrowser(this.platformId)) {
          this.state.set(CSRF_KEY, data.csrf_token)
        }
      })
    }
  }

  constructor(
    private registerService: RegisterService,
    private csrfService: GenCsrfService,
    private state: TransferState,
    @Inject(PLATFORM_ID) private platformId: Object,
    private loadingService: LoadingService
  ) {}

  get nama_lengkap() {return this.userProfileForm.get("nama_lengkap")!}
  get username() {return this.userProfileForm.get("username")!}
  get email() {return this.userProfileForm.get("email")!}
  get no_hp() {return this.userProfileForm.get("no_hp")!}
  get password() {return this.userProfileForm.get("password")!}
  get confirmpassword() {return this.userProfileForm.get("confirmpassword")!}

  public registerUser(): void {
    console.log(this.userProfileForm.value)
    
    // let dataUser: User = {
    //   nama_lengkap: "",
    //   username: "",
    //   email: "",
    //   no_hp: "",
    //   password: ""
    // }

    // dataUser.nama_lengkap = this.userProfileForm.value.nama_lengkap
    // dataUser.username = this.userProfileForm.value.username
    // dataUser.email = this.userProfileForm.value.email
    // dataUser.no_hp = this.userProfileForm.value.no_tlp
    // dataUser.password = this.userProfileForm.value.password

    // this.registerService.registerUser(dataUser).subscribe(
    //   () => {
    //     this.loadingService.setLoading(false)
    //     this.successRegister = "Registrasi akun berhasil!"
    //     this.errorRegister = ""
    //   },
    //   error => {
    //     this.loadingService.setLoading(false)
        
    //     if (error.status === 409) {
    //       this.errorRegister = "Registrasi akun gagal! User sudah pernah dibuat."
    //       this.successRegister = ""
    //     } else {
    //       this.errorRegister = "Registrasi akun gagal!"
    //       this.successRegister = ""
    //     }
    //   }
    // )
  }
}
