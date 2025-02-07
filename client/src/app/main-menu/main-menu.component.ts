import { Component } from '@angular/core';
import { WebsocketService } from '../websocket.service';
import { FormsModule } from '@angular/forms'
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-main-menu',
  imports: [FormsModule, CommonModule],
  templateUrl: './main-menu.component.html',
  styleUrl: './main-menu.component.css'
})
export class MainMenuComponent {
  constructor(private wsService: WebsocketService) {}

  startGame() {
    this.wsService.connect('create').subscribe()
  }

  gameID: string = ''
  inputError: string = ''
  joinGame() {
    this.wsService.connect(this.gameID).subscribe({
      error: (error) => {
        this.inputError = error
        this.gameID = ''
      }
    })
  }
}
