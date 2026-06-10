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
  constructor(
    private bankService: BankService,
  private cdr: ChangeDetectorRef){}
    depositMoney(){
      this.bankService.depositMoney(this.deposit)
      .subscribe({
        next:(response)=>{
          this.message=`${response.message} | Balance: ${response.new_balance}`;
          this.deposit.account_id='';
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

