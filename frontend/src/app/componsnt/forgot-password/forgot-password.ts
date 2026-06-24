import { Component, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { BankService } from '../../services/bank';

@Component({
  selector: 'app-forgot-password',
  standalone:true,
  imports: [FormsModule, CommonModule],
  templateUrl: './forgot-password.html',
  styleUrl: './forgot-password.css',
})
export class ForgotPasswordComponent {
  currentStep=1;
  username='';
  otp='';
  newPassword='';
  confirmPassword='';
  message='';

  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef,
    private router: Router
  ){}

  generateOtp(){
    const request={
      username:this.username
    };

    this.bankService
    .generateOtp(request)
    .subscribe({
      next:(response:any)=>{
        this.message=response.message;
        this.currentStep=2;
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });

  }

  verifyOtp(){
    const request={
      username:this.username,
      otp:this.otp
    };

    this.bankService
    .verifyOtp(request)
    .subscribe({
      next:(response:any)=>{
        this.message=response.message;
        this.currentStep=3;
        this.otp='';
        this.cdr.detectChanges();
      },

      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }

  resetPassword(){
    if (this.newPassword !== this.confirmPassword){
      this.message='Password do not match';
      return;
    }

    const request={
      username:this.username,
      new_password:this.newPassword
    };

    this.bankService.resetPassword(request)
    .subscribe({
      next:(response:any)=>{
        this.message=response.message;
        this.currentStep=4;
        setTimeout(()=>{
          this.router.navigate(['/login']);
        },2000);
        this.cdr.detectChanges();
      },

      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }
}
