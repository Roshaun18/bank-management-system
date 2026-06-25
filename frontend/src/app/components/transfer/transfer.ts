import { Component, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BankService } from '../../services/bank'; 
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-transfer',
  standalone:true,
  imports: [FormsModule, CommonModule],
  templateUrl: './transfer.html',
  styleUrl: './transfer.css',
})
export class TransferComponent {
  transfer={
    from_account:'',
    to_account:'',
    amount:0
  };
  message='';
  selectedAccount='';
  beneficiaries:any[]=[];
  showConfirmation=false;
  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}
  ngOnInit(){
    this.selectedAccount=localStorage.getItem('accountId') || '';
    this.loadBeneficiaries();
  }

  loadBeneficiaries(){
    const customerId = localStorage.getItem('customerId');

    if (!customerId) return;

    this.bankService
    .getBeneficiaries(customerId)
    .subscribe({
      next:(response:any)=>{
        this.beneficiaries=response;
        this.cdr.detectChanges();
      },
      error:(err)=>{
        console.log(err);
        this.cdr.detectChanges();
      }
    });
  }

  confirmTransfer(){
    this.bankService
    .transferMoney(this.transfer)
    .subscribe({
      next:(response)=>{
        this.message=response.message;
        this.transfer.to_account='';
        this.transfer.amount=0;
        this.showConfirmation=false;
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error?.message || err.error || "Transfer Failed";
        this.showConfirmation=false;
        this.cdr.detectChanges();
      }
    });
  }

  cancelTransfer(){
    this.showConfirmation=false;
  }

  transferMoney(){
    const accountId=localStorage.getItem('accountId');
    if(!accountId){
      this.message="No account selected";
      return;
    }

    this.transfer.from_account=accountId;
    this.showConfirmation=true;
  }

}
