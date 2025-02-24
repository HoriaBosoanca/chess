package main

import (
	"fmt"
	"mychess/chess"
)

func main() {
	var board chess.Board
	board.Init()

	fmt.Println(board.GetLegalMoves("h2"))
	
	board.Print()
}