import { Component, ChangeDetectorRef} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {BankService} from '../../services/bank';
import { ToastComponent } from '../toast/toast';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-deposit',
  standalone:true,
  imports: [FormsModule, CommonModule, ToastComponent],
  templateUrl: './deposit.html',
  styleUrl: './deposit.css',
})
export class DepositComponent {
  deposit={
    account_id:'',
    amount:0
  };

  message='';
  selectedAccount='';
  toastMessage='';
  toastType: 'success' | 'error' = 'success';
  showToast=false;
  private toastTimer:any;

  constructor(
    private bankService: BankService,
  private cdr: ChangeDetectorRef){}

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
    depositMoney(){
      const accountId=localStorage.getItem('accountId');
      if(!accountId){
        this.message='No account selected';
        return;
      }
      this.deposit.account_id=accountId;
      this.bankService.depositMoney(this.deposit)
      .subscribe({
        next:(response:any)=>{
          console.log(response);
          this.showNotification('Deposit Successful','success');
          this.message=`${response.message} | Balance: ${response.new_balance}`;
          this.deposit.amount=0;
          this.cdr.detectChanges();
        },
        error:(err:any)=>{
          this.showNotification('Deposit Failed','error');
          this.message=err.error?.message || err.error;
          this.cdr.detectChanges();
        }
      });
    }
  }

