import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Message } from './app.models';

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {

  constructor() { }

  baseUrl: string = "ws://localhost:8080/ws"

  websocket: WebSocket | null = null
  connect(gameID: string): Observable<Message> {
    this.websocket = new WebSocket(this.baseUrl + '?gameID=' + gameID)
    return this.getWs()
  }

  getWs(): Observable<Message> {
    return new Observable<Message>((observer) => {
      this.websocket!.onopen = () => {
        console.log('conn established')
      }
  
      this.websocket!.onclose = () => {
        console.log('conn closed')
      }
  
      this.websocket!.onerror = (error) => {
        console.log(error)
      }
  
      this.websocket!.onmessage = (data: MessageEvent) => {
        const messageObj: Message = JSON.parse(data.data)
        
        console.log(messageObj)

        switch (messageObj.type) {
          case 'gameID':
            observer.next(messageObj)
            break
          case 'error':
            console.log('Error:', messageObj.message)
            observer.error(messageObj)
            break
          case 'color':
            sessionStorage.setItem('color', messageObj.message)
            break
          case 'start':
            observer.next(messageObj)
            break
          case 'turn':
            break
          default:
            console.log('unknown message:', messageObj)
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
