import { Routes } from '@angular/router';
import { CustomerComponent } from './components/customer/customer';
import { AccountComponent } from './components/account/account';
import { DepositComponent } from './components/deposit/deposit';
import { WithdrawComponent } from './components/withdraw/withdraw';
import { TransferComponent } from './components/transfer/transfer';
import { TransactionsComponent } from './components/transactions/transactions';
import { DashboardComponent } from './components/dashboard/dashboard';
import { HomeComponent } from './components/home/home';
import { CustomerSummaryComponent } from './components/customer-summary/customer-summary';

export const routes: Routes = [{path:'', redirectTo:'home', pathMatch:'full'},
    {path:'customer', component:CustomerComponent},
    {path:'account', component:AccountComponent},
    {path:'deposit', component:DepositComponent},
    {path:'withdraw',component:WithdrawComponent},
    {path:'transfer',component:TransferComponent},
    {path:'transactions',component:TransactionsComponent},
    {path:'dashboard',component:DashboardComponent},
    {path:'home', component:HomeComponent},
    {path:'customer-summary', component:CustomerSummaryComponent}
];
