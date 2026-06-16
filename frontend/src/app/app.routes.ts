import { mapToCanActivate, Routes } from '@angular/router';
import { CustomerComponent } from './components/customer/customer';
import { AccountComponent } from './components/account/account';
import { DepositComponent } from './components/deposit/deposit';
import { WithdrawComponent } from './components/withdraw/withdraw';
import { TransferComponent } from './components/transfer/transfer';
import { TransactionsComponent } from './components/transactions/transactions';
import { DashboardComponent } from './components/dashboard/dashboard';
import { HomeComponent } from './components/home/home';
import { CustomerSummaryComponent } from './components/customer-summary/customer-summary';
import { LoginComponent } from './components/login/login';
import { authGuard } from './guards/auth-guard';
import { adminGuard } from './guards/admin-guard';
import { AccountSelectionComponent } from './components/account-selection/account-selection';
import { ChangePasswordComponent } from './components/change-password/change-password';

export const routes: Routes = [{path:'', redirectTo:'login', pathMatch:'full'},
    {path:'customer', component:CustomerComponent, canActivate:[authGuard, adminGuard]},
    {path:'account', component:AccountComponent, canActivate:[authGuard, adminGuard]},
    {path:'deposit', component:DepositComponent, canActivate:[authGuard]},
    {path:'withdraw',component:WithdrawComponent, canActivate:[authGuard]},
    {path:'transfer',component:TransferComponent, canActivate:[authGuard]},
    {path:'transactions',component:TransactionsComponent, canActivate:[authGuard]},
    {path:'dashboard',component:DashboardComponent, canActivate:[authGuard]},
    {path:'home', component:HomeComponent, canActivate:[authGuard]},
    {path:'customer-summary', component:CustomerSummaryComponent, canActivate:[authGuard, adminGuard]},
    {path:'login', component:LoginComponent},
    {path:'account-selection', component:AccountSelectionComponent, canActivate:[authGuard]},
    {path:'change-password', component:ChangePasswordComponent, canActivate:[authGuard]}
];
