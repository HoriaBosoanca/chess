import { Component, OnInit } from '@angular/core';
import { WebsocketService } from '../websocket.service';
import { Message } from '../app.models';

@Component({
  selector: 'app-match',
  imports: [],
  templateUrl: './match.component.html',
  styleUrl: './match.component.css'
})
export class MatchComponent implements OnInit {
  constructor(private wsService: WebsocketService) {}

  ngOnInit(): void {
    this.renderPieces()
  }

  renderPieces(): void {
    this.wsService.unblurer?.next()
  }
}
