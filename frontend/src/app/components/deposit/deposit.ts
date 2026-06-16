import { Component, ChangeDetectorRef} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {BankService} from '../../services/bank';

@Component({
  selector: 'app-deposit',
  standalone:true,
  imports: [FormsModule],
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


  constructor(
    private bankService: BankService,
  private cdr: ChangeDetectorRef){}

    ngOnInit(){
    this.selectedAccount=localStorage.getItem('accountId') || '';
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
        next:(response)=>{
          this.message=`${response.message} | Balance: ${response.new_balance}`;
          this.deposit.amount=0;
          this.cdr.detectChanges();
        },
        error:(err)=>{
          this.message=err.error;
          this.cdr.detectChanges();
        }
      });
    }
  }

