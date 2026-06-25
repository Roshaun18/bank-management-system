import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { BankService } from '../../services/bank';

@Component({
  selector: 'app-beneficiaries',
  standalone:true,
  imports: [FormsModule, CommonModule],
  templateUrl: './beneficiaries.html',
  styleUrl: './beneficiaries.css',
})
export class BeneficiariesComponent {
  beneficiaryName='';
  accountId='';
  message='';
  beneficiaries:any[]=[];

  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}

  ngOnInit(): void{
    this.loadBeneficiaries();
  }

  addBeneficiary(){
    const customerId=localStorage.getItem('customerId');

    const request={
      customer_id:customerId,
      beneficiary_name:this.beneficiaryName,
      account_id: this.accountId
    };

    console.log(request);

    this.bankService
    .addBeneficiary(request)
    .subscribe({
      next:(response:any)=>{
        this.message=response.message;
        this.loadBeneficiaries();
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }

  loadBeneficiaries(){
    const customerId=localStorage.getItem('customerId');
    if (!customerId)return;

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

  deleteBeneficiary(accountId:string){
    const customerId=localStorage.getItem('customerId');
    console.log("Customer ID:",customerId);
    const request={
      customer_id:customerId,
      account_id:accountId
    };
    
    this.bankService
    .deleteBeneficiary(request)
    .subscribe({

      next:(response:any)=>{
      this.message=response.message;
      this.loadBeneficiaries();
      this.cdr.detectChanges();
    },
    error:(err)=>{
      this.message=err.error;
      this.cdr.detectChanges();
    }
    });
    

  }

}
