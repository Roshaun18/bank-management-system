import {Injectable} from '@angular/core';
import{HttpClient} from '@angular/common/http';
import{Observable} from 'rxjs';

@Injectable({providedIn:'root'})
export class BankService{
    private apiUrl='http://localhost:8080';
    constructor(private http:HttpClient){}

    createCustomer(customer:any):Observable<any>{
        return this.http.post(`${this.apiUrl}/customer`,customer);

    }

    createAccount(account:any):Observable<any>{
        return this.http.post(`${this.apiUrl}/account`,account);
    }

    depositMoney(data:any):Observable<any>{
        return this.http.post(`${this.apiUrl}/deposit`,data);
    }

    withdrawMoney(data:any):Observable<any>{
        return this.http.post(`${this.apiUrl}/withdraw`,data);
    }

    transferMoney(data:any):Observable<any>{
        return this.http.post(`${this.apiUrl}/transfer`,data);
    }

    getTransactions(accountId:string):Observable<any>{
        return this.http.get(`${this.apiUrl}/transactions?account_id=${accountId}`);
    }

    getDashboard(accountId:string):Observable<any>{
        return this.http.get(`${this.apiUrl}/dashboard?account_id=${accountId}`);
    }

    getCustomerSummary():Observable<any>{
        return this.http.get(`${this.apiUrl}/customer-summary`);
    }

    login(data:any):Observable<any>{
        return this.http.post(`${this.apiUrl}/login`,data);
    }

    getCustomerAccounts(customerId:string){
        return this.http.get<any[]>(`${this.apiUrl}/customer/accounts?customer_id=${customerId}`);
    }

    changePassword(data:any):Observable<any>{
        return this.http.post(`${this.apiUrl}/change-password`,data);
    }
}