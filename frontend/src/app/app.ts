import{Component} from '@angular/core';
import { RouterOutlet, RouterLink, RouterLinkActive } from '@angular/router';

@Component({
  selector:'app-root',
  standalone:true,
  imports:[RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl:`./app.html`,
  styleUrl:'./app.css'
})
export class App{
  darkMode =false;
  constructor(){
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
}