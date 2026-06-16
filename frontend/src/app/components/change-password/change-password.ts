import { Component, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BankService } from '../../services/bank';

@Component({
  selector: 'app-change-password',
  standalone:true,
  imports: [FormsModule],
  templateUrl: './change-password.html',
  styleUrl: './change-password.css',
})
export class ChangePasswordComponent {
  passwordData={
    old_password:'',
    new_password:'',
    confirm_password:''
  };
  message='';

  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}

  changePassword(){
    if( this.passwordData.new_password!==this.passwordData.confirm_password){
      this.message='Password do not match';
      this.cdr.detectChanges();
      return;
    }

    const customerId=localStorage.getItem('customerId');
    console.log(customerId);
    const request={
      customer_id: customerId,
      old_password: this.passwordData.old_password,
      new_password: this.passwordData.new_password
    };
console.log('Customer ID:', customerId);
console.log(request);
    this.bankService.changePassword(request)
    .subscribe({
      next:(response:any)=>{
        this.message=response.message;
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });

    console.log(request);
    this.cdr.detectChanges();
  }
}
