import { Component, ChangeDetectorRef } from '@angular/core';
import {FormsModule} from '@angular/forms';
import { BankService } from '../../services/bank';  
import { ToastComponent } from '../toast/toast';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-withdraw',
  standalone:true,
  imports: [FormsModule, ToastComponent, CommonModule],
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
  toastMessage='';
  showToast=false;
  toastType: 'success' | 'error'='success';
  private toastTimer:any;
  constructor(
    private bankService:BankService,
  private cdr: ChangeDetectorRef
){}

ngOnInit(){
  this.selectedAccount=localStorage.getItem('accountId') || '';
}

 showNotification(
    message: string,
    type:'success' | 'error'
  ){
    console.log("Resevied message:",message);
    this.toastMessage=message;
    this.toastType=type;
    this.showToast=true;

    if(this.toastTimer){
      clearTimeout(this.toastTimer);
    }
    this.toastTimer=setTimeout(()=>{
      this.showToast=false;
      this.cdr.detectChanges();
    },3000);
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
      next:(response:any)=>{
        this.showNotification('Withdraw Successful','success')
        this.message=`${response.message} | Balance: ${response.new_balance}`;
        this.withdraw.account_id='';
        this.withdraw.amount=0;
        this.cdr.detectChanges();
      },
      error:(err:any)=>{
        this.showNotification('Withdraw Failed','error')
        this.message=err.error?.message || err.error;
        this.cdr.detectChanges();
      }
    });
  }
}
