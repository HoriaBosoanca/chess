import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WebsocketService } from './websocket.service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit {
  title = 'Chess';

  constructor(private wsService: WebsocketService) {}

  ngOnInit(): void {
    this.wsService.unblurBoard().subscribe({
      next: () => {
        const img = document.querySelector('img')
        img?.classList.remove('blur-sm')
      }
    })
  }
}