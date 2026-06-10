import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BankService } from '../../services/bank';

@Component({
  selector: 'app-customer-summary',
  standalone:true,
  imports: [CommonModule],
  templateUrl: './customer-summary.html',
  styleUrl: './customer-summary.css',
})
export class CustomerSummaryComponent implements OnInit {
  customers:any[]=[];
  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}

  ngOnInit(): void {
    this.loadCustomers();
  }

  loadCustomers(){
    this.bankService.getCustomerSummary()
    .subscribe({
      next:(response)=>{
        this.customers=response;
        console.log(response);
        this.cdr.detectChanges();
      },
      error: (err)=>{
        console.error(err);
        this.cdr.detectChanges();
      }
    });
  }
}
