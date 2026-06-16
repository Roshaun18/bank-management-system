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
  role=localStorage.getItem('role');
  accountId: string | null='' ;
  transactions:any[]=[];
  message='';


  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}
  ngOnInit(){
    if(this.role==='customer'){
      this.accountId=localStorage.getItem('accountId') || '';
    this.loadTransactions();
    }
    
  }

  loadTransactions(){
  
      if(!this.accountId){
        this.message='No account selected';
        return;
      }
    this.bankService
    .getTransactions(this.accountId)
    .subscribe({
      next:(response)=>{
        this.transactions=response || [];
        if(this.transactions.length===0){
          this.message='Not Transactions Found';
          
        }
        else{
          this.message='';
        }
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
