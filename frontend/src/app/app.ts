import{Component} from '@angular/core';
import { Router, RouterOutlet, RouterLink, RouterLinkActive, NavigationEnd } from '@angular/router';
import { CommonModule } from '@angular/common';
import { filter } from 'rxjs';

@Component({
  selector:'app-root',
  standalone:true,
  imports:[RouterOutlet, RouterLink, RouterLinkActive, CommonModule],
  templateUrl:`./app.html`,
  styleUrl:'./app.css'
})
export class App{
  darkMode =false;
  constructor(public router:Router){
    const savedTheme=localStorage.getItem('theme');
    if (savedTheme==='dark'){
      this.darkMode=true;
      document.body.classList.add('dark-theme');
    }
  }
  toogleTheme(){
    this.darkMode= !this.darkMode;
    localStorage.setItem('theme', this.darkMode ? 'dark': 'light');
    document.body.classList.toggle('dark-theme', this.darkMode);
  }

  logout(){
    localStorage.clear();
    this.router.navigate(['/login']);
  }

  get role(): string | null{
    return localStorage.getItem('role');
  }
}