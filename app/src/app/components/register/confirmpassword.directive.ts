import { AbstractControl, ValidationErrors, ValidatorFn } from "@angular/forms"


export function matchingPasswordValidator(): ValidatorFn {
  return (control: AbstractControl): ValidationErrors | null => {
    const password = control.get("password")
    const confirmPassword = control.get("confirmpassword")

    return password && confirmPassword && password.value !== confirmPassword.value ? { passwordMismatch: true } : null
  }
}