import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {

  constructor() { }

  baseUrl: string = "ws://localhost:8080/ws"

  websocket: WebSocket | null = null
  connect(gameID: string) {
    this.websocket = new WebSocket(this.baseUrl + '?gameID=' + gameID)
    
    this.websocket.onopen = () => {
      console.log('conn established')
    }

    this.websocket.onclose = () => {
      console.log('conn closed')
    }

    this.websocket.onerror = (error) => {
      console.log('conn error:', error)
    }

    this.websocket.onmessage = (data: MessageEvent) => {
      const message = JSON.parse(data.data)

      if(message.gameID) {
        console.log(message.gameID)
        sessionStorage.setItem("gameID", message.gameID)
      }
      if(message.yourColor) {
        sessionStorage.setItem("myColor", message.yourColor)
      }
      if(message == "it's your turn") {
        console.log('my turn')
      }
    }

  }

  sendMove(move: string) {
    this.websocket!.send(JSON.stringify({
      move: move
    }))
  }
}
