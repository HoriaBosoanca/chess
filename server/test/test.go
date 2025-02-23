package main

import (
	"fmt"
	"game/chess"
)

func main() {
	var board chess.Board
	board.Init()

	fmt.Println(board.GetLegalMoves("e4"))

	board.SetPiece("e2", "WP")
	board.Move("e2", "e4")
	
	board.Print()
}