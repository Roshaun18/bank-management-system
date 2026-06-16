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
  selectedAccount='';
  constructor(
    private bankService:BankService,
  private cdr: ChangeDetectorRef
){}

ngOnInit(){
  this.selectedAccount=localStorage.getItem('accountId') || '';
}

  withdrawMoney(){
    const accountId=localStorage.getItem('accountId');
      if(!accountId){
        this.message='No account selected';
        return;
      }
      this.withdraw.account_id=accountId;
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
