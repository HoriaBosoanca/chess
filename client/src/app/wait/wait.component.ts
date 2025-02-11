import { Component, OnInit } from '@angular/core';
import { WebsocketService } from '../websocket.service';
import { Message } from '../app.models';

@Component({
  selector: 'app-wait',
  imports: [],
  templateUrl: './wait.component.html',
  styleUrl: './wait.component.css'
})
export class WaitComponent implements OnInit {
  
  constructor(private wsService: WebsocketService) {}

  gameID: string = ""

  ngOnInit(): void {
      this.wsService.getWs().subscribe({
        next: (message: Message) => {
          if(message.type == 'gameID') {
            this.gameID = message.message
          } 
        }
      }) 
  }
}
