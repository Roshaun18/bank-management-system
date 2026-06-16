import { Component, ChangeDetectorRef } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { BankService } from '../../services/bank';

@Component({
  selector: 'app-login',
  standalone:true,
  imports: [FormsModule],
  templateUrl: './login.html',
  styleUrl: './login.css',
})
export class LoginComponent {
  loginData={
    username:'',
    password:'',
    role:'admin'
  };

  message='';

  constructor(
    private bankService: BankService,
    private router: Router,
    private cdr: ChangeDetectorRef
  ){}

  login(){
    this.bankService.login(this.loginData)
    .subscribe({
      next:(response)=>{
        console.log(response);
        localStorage.setItem(
          'loggedIn',
          'true'
        );
        localStorage.setItem(
          'role',
          response.role
        );
        console.log("Role =", response.role);

        if(response.role==='customer'){
          localStorage.setItem(
            'customerId',
            response.customer_id);
            localStorage.setItem(
              'accountId', response.account_id
            );
        }

        if (response.role==='customer'){
          this.router.navigate(['/account-selection']);
        }
        else{
          this.router.navigate(['/dashboard']);
        }

        

        this.cdr.detectChanges();

      },

      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }
}
