import { Component } from '@angular/core';
import { WebsocketService } from '../websocket.service';
import { FormsModule } from '@angular/forms'
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { Message } from '../app.models'

@Component({
  selector: 'app-main-menu',
  imports: [FormsModule, CommonModule],
  templateUrl: './main-menu.component.html',
  styleUrl: './main-menu.component.css'
})
export class MainMenuComponent {
  constructor(private wsService: WebsocketService, private router: Router) {}

  startGame() {
    this.wsService.connect('create').subscribe()
    this.router.navigate(['waiting'])
  }

  gameID: string = ''
  inputError: string = ''
  joinGame() {
    this.wsService.connect(this.gameID).subscribe({
      error: (error: Message) => {
        this.inputError = error.message
        this.gameID = ''
      },
      next: (message: Message) => {
        if(message.type == 'start') {
          this.router.navigate(['game'])
        }
      }
    })
  }
}
