import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {

  constructor() { }

  baseUrl: string = "ws://localhost:8080/ws"

  websocket: WebSocket | null = null
  connect(gameID: string): Observable<MessageEvent> {
    return new Observable((observer) => {
      this.websocket = new WebSocket(this.baseUrl + '?gameID=' + gameID)
      
      this.websocket.onopen = () => {
        console.log('conn established')
      }
  
      this.websocket.onclose = () => {
        console.log('conn closed')
      }
  
      this.websocket.onerror = (error) => {
        console.log(error)
      }
  
      this.websocket.onmessage = (data: MessageEvent) => {
        const message = JSON.parse(data.data)

        console.log(message)
  
        if(message.gameID) {
          console.log(message.gameID)
          sessionStorage.setItem("gameID", message.gameID)
        }
        if(message.error) {
          console.log(message.error)
          switch(message.error) {
            case "Invalid game ID":
              observer.error(message.error)
              break
            default:
              console.log('unknown error')
              break
          }
        }
        if(message.yourColor) {
          sessionStorage.setItem("myColor", message.yourColor)
        }
        if(message == "it's your turn") {
          // wait for move from player
        }
      }
  
      return () => this.websocket!.close()
    })
  }

  sendMove(move: string) {
    this.websocket!.send(JSON.stringify({
      move: move
    }))
  }
}
