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
  selectedAccount='';
  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}
  ngOnInit(){
    this.selectedAccount=localStorage.getItem('accountId') || '';
  }

  transferMoney(){
    const accountId=localStorage.getItem('accountId');
    if(!accountId){
      this.message="No account selected";
      return;
    }
    this.transfer.from_account=accountId;
    this.bankService.transferMoney(this.transfer)
    .subscribe({
      next: (response)=>{
        this.message=`${response.message}`;
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
