import { Routes } from '@angular/router';
import { MainMenuComponent } from './main-menu/main-menu.component';
import { WaitComponent } from './wait/wait.component';

export const routes: Routes = [
    {path: '', component: MainMenuComponent},
    {path: 'waiting', component: WaitComponent},

    // redirects to main menu:
    { path: '', redirectTo: '', pathMatch: 'full' },
    { path: '**', redirectTo: '' }
];
