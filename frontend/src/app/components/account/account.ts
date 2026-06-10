import { Component, ChangeDetectorRef} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {BankService} from '../../services/bank';

@Component({
  selector: 'app-account',
  standalone:true,
  imports: [FormsModule],
  templateUrl: './account.html',
  styleUrl: './account.css',
})
export class AccountComponent {
  account={
    customer_id:'',
    balance:0
  };
  message='';

  constructor(
    private bankService:BankService,
    private cdr: ChangeDetectorRef
  ){}

  createAccount(){
    this.bankService.createAccount(this.account)
    .subscribe({
      next:(response)=>{
        console.log("Response:",response);
        this.message=`${response.message}(ACCOUNT_ID: ${response.account_id})`;
        this.account.customer_id='';
        this.account.balance=0;
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }
}
