import { Component, ChangeDetectorRef} from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BankService } from '../../services/bank';

@Component({
  selector: 'app-dashboard',
  standalone:true,
  imports: [FormsModule],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.css',
})
export class DashboardComponent {
  accountId='';
  message='';
  dashboard:any=null;

  constructor(
    private bankService: BankService,
    private cdr: ChangeDetectorRef
  ){}

  getDashboard(){
    if(!this.accountId){
      this.message="Please enter AccountID";
      this.cdr.detectChanges();
      return;
    }
    this.bankService.getDashboard(this.accountId)
    .subscribe({
      next:(response)=>{
        this.dashboard=response;
        this.message='';
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.dashboard=null;
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });

  }
}
