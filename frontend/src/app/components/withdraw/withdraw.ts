import { Component, ChangeDetectorRef } from '@angular/core';
import {FormsModule} from '@angular/forms';
import { BankService } from '../../services/bank';  

@Component({
  selector: 'app-withdraw',
  standalone:true,
  imports: [FormsModule],
  templateUrl: './withdraw.html',
  styleUrl: './withdraw.css',
})
export class WithdrawComponent {
  withdraw={
    account_id:'',
    amount:0
  };

  message='';
  constructor(
    private bankService:BankService,
  private cdr: ChangeDetectorRef){}
  withdrawMoney(){
    this.bankService.withdrawMoney(this.withdraw)
    .subscribe({
      next:(response)=>{
        this.message=`${response.message} | Balance: ${response.new_balance}`;
        this.withdraw.account_id='';
        this.withdraw.amount=0;
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }
}
