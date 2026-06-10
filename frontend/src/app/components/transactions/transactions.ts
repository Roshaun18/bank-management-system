import { Component, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BankService } from '../../services/bank';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-transactions',
  standalone:true,
  imports: [FormsModule, CommonModule],
  templateUrl: './transactions.html',
  styleUrl: './transactions.css',
})
export class TransactionsComponent {
  accountId='';
  transactions:any[]=[];
  message='';

  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}

  loadTransactions(){
    if(!this.accountId){
      this.message="Enter Account ID";
      this.cdr.detectChanges();
      return;
    }
    this.bankService.getTransactions(this.accountId)
    .subscribe({
      next:(response)=>{
        this.transactions=response;
        this.message='';
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.transactions=[];
        this.cdr.detectChanges();
      }
    });
  }
}
