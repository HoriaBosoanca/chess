import { Component, OnInit } from '@angular/core';
import { WebsocketService } from '../websocket.service';
import { Message } from '../app.models';
import { from, toArray } from 'rxjs';

@Component({
  selector: 'app-match',
  imports: [],
  templateUrl: './match.component.html',
  styleUrl: './match.component.css'
})
export class MatchComponent implements OnInit {
  constructor(private wsService: WebsocketService) {}

  ngOnInit(): void {
    this.wsService.unblurer?.next()
    this.renderPieces()
  }

  renderPieces(): void {
    const board = new Board()
    board.create('h5', 'white', 'pawn')
    board.move('h5', 'a1')
    board.move('a1', 'h8')
    board.move('h8', 'f1')
  }
}

class Board {
  private board: (Piece | null)[][] = new Array(8).fill(null).map(() => new Array(8).fill(null));

  public create(position: string, color: string, piece: string): void {
    const newPiece: Piece = new Piece(color, piece)
    newPiece.setPosition(position)
    this.board[newPiece.x][newPiece.y] = newPiece
  }

  public move(fromPos: string, toPos: string): void {
    const fromArr = fromPos.split("")
    const xFrom = xAlgebraicToIndex(fromArr[0])
    const yFrom = yAlgebraicToIndex(fromArr[1])
    const movingPiece = this.board[xFrom][yFrom]
    this.board[xFrom][yFrom] = null
    
    const toArr = toPos.split("")
    const xTo = xAlgebraicToIndex(toArr[0])
    const yTo = yAlgebraicToIndex(toArr[1])
    movingPiece!.setPosition(toPos)
    this.board[xTo][yTo] = movingPiece
  }
}

class Piece {

  private image!: HTMLImageElement 
  public type: string = ""

  public x: number = -1
  public y: number = -1
  
  constructor(color: string, name: string) {
    if(color != 'black' && color != 'white') {
      console.error('wrong color string')
    }
    if(name != 'king' && name != 'queen' && name != 'rook' && name != 'knight' && name != 'bishop' && name != 'pawn') {
      console.error('wrong piece name string')
    }
    this.type = name

    const piece = document.createElement('img')
    piece.src = `pieces/${color}/${name}.png`
    piece.style.position = 'absolute'
    document.querySelector('#originElement')!.appendChild(piece)
    piece.style.height = '100px'
    piece.style.width = '100px'

    this.image = piece
  }

  public setPosition(position: string) {
    if(! /^[a-h][1-8]$/.test(position)) {
      console.error('wrong file or rank')
    }
    const arr = position.split("")
    const x = xAlgebraicToIndex(arr[0])
    const y = yAlgebraicToIndex(arr[1])
    this.image.style.left = x * 100 + 'px'
    this.image.style.top = 800 - (y+1) * 100 + 'px'
    this.x = x
    this.y = y
  }
}

function xAlgebraicToIndex(xAlgebraic: string): number {
  if(! /^[a-h]$/.test(xAlgebraic)) {
    console.error('wrong file input')
  }
  return xAlgebraic.charCodeAt(0) - 97
}

function yAlgebraicToIndex(yAlgebraic: string): number {
  if(! /^[1-8]$/.test(yAlgebraic)) {
    console.error('wrong rank input')
  }
  return Number(yAlgebraic) - 1
}