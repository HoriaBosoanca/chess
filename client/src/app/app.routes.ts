import { Routes } from '@angular/router';
import { MainMenuComponent } from './main-menu/main-menu.component';
import { WaitComponent } from './wait/wait.component';
import { MatchComponent } from './match/match.component';

export const routes: Routes = [
    {path: '', component: MainMenuComponent},
    {path: 'waiting', component: WaitComponent},
    {path: 'game', component: MatchComponent},

    // redirects to main menu:
    { path: '', redirectTo: '', pathMatch: 'full' },
    { path: '**', redirectTo: '' }
];
