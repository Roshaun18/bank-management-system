import { Component, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BankService } from '../../services/bank'; 

@Component({
  selector: 'app-transfer',
  standalone:true,
  imports: [FormsModule],
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
  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}

  transferMoney(){
    this.bankService.transferMoney(this.transfer)
    .subscribe({
      next: (response)=>{
        this.message=`${response.message}`;
        this.transfer.from_account='';
        this.transfer.to_account='';
        this.transfer.amount=0;
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }

}
