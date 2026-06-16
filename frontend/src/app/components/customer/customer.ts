import {Component} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {BankService } from '../../services/bank';
import { ChangeDetectorRef } from '@angular/core';

@Component({
  selector:'app-customer',
  standalone:true,
  imports:[FormsModule],
  templateUrl:'./customer.html',
  styleUrl:'./customer.css'
})
export class CustomerComponent{
  customer={
    name:'',
    email:'',
    username:'',
    password:''
  };

  message='';
  constructor(
    private bankService:BankService,
  private cdr:ChangeDetectorRef) {}

  createCustomer(){
    const payload={
      name:this.customer.name,
      email:this.customer.email,
      username:this.customer.username,
      password:this.customer.password
    };

    this.bankService.createCustomer(payload)
    .subscribe({
      next:(response)=>{
        console.log("Response:",response);
        this.message=`${response.message}(ID: ${response.id})`;
        this.customer.name='';
        this.customer.email='';
        this.customer.username='';
        this.customer.password='';
        this.cdr.detectChanges();
      },
      error:(err)=>{
        this.message=err.error;
        this.cdr.detectChanges();
      }
    });
  }
}