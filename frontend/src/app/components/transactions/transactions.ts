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
  selectedType='All';


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

  downloadCSV(){
    const headers=[
      'Date',
      'Type',
      'Amount',
      'From Account',
      'To Account',
      'Balance'
    ];

    const rows=this.transactions.map(tx=>[
      new Date(tx.created_at)
      .toLocaleString(),
      tx.type || '',
      tx.amount || '',
      tx.from_account || '',
      tx.to_account || '',
      tx.balance || ''
    ]);

    const csvContent=[
      headers,
      ...rows
    ]
    .map(row=>row.join(','))
    .join('/n');
    const blob=new Blob(
      [csvContent],
      {type: 'text/csv;charset=utf-8;'}
    );

    const url=window.URL.createObjectURL(blob);
    const link=document.createElement('a');
    link.href=url;
    link.download='statement.csv';
    link.click();
    window.URL.revokeObjectURL(url);
  }

  get filteredTransactions(){
    if(this.selectedType==='ALL'){
      return this.transactions;
    }
    return this.transactions.filter(
      tx=>tx.type===this.selectedType
    );
}
}
