import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WebsocketService } from './websocket.service';
import { FormsModule } from '@angular/forms'

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, FormsModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'angularClient';

  constructor(private wsService: WebsocketService) {}

  startGame() {
    this.wsService.connect('create')
  }

  gameID: string = 'Enter a game code here'
  joinGame() {
    console.log(this.gameID)
    this.wsService.connect(this.gameID)
  }
}
