import { Component, ChangeDetectorRef } from '@angular/core';
import { Router } from '@angular/router';
import { BankService } from '../../services/bank';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-account-selection',
  standalone:true,
  imports:[CommonModule],
  templateUrl: './account-selection.html',
  styleUrl: './account-selection.css',
})
export class AccountSelectionComponent {
  accounts:any[]=[];
  message='';
  constructor(
    private bankService: BankService,
    private router: Router,
    private cdr: ChangeDetectorRef
  ){}
  ngOnInit(){
    console.log("account selection loaded");
    const customerId=localStorage.getItem('customerId');
    if (!customerId){
      this.message="Customer ID not found";
      return;
    }
    this.bankService.getCustomerAccounts(customerId!)
    .subscribe({
      next:(response)=>{
        this.accounts=response;
        if (response.length===0){
          this.message="No accounts found";
        }
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error
        this.cdr.detectChanges(); 
      }
    });
  }

  selectedAccount(accountId:string){
    localStorage.setItem('accountId',accountId);
    this.router.navigate(['/dashboard']
    );
    this.cdr.detectChanges();
  }

}
