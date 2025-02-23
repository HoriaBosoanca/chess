package chess

import (
	"fmt"
)

type Board struct {
	Board [][]string
}

func (board *Board) Init() {
	board.Board = [][]string{
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
		{"WP", "WP", "WP", "WP", "WP", "WP", "WP", "WP"},
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
		{"BP", "BP", "BP", "BP", "BP", "BP", "BP", "BP"},
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
	}
}

func (board *Board) Print() {
	for i := 7; i >= 0; i-- {
		rank := ""
		for _, v := range board.Board[i] {
			rank += v + " "
		}
		fmt.Println(rank)
	}
}

func (board *Board) SetPiece(position, piece string) {
	file, rank := ToIndexes(position)
	board.Board[rank][file] = piece
}

func (board *Board) GetPiece(position string) (string) {
	file, rank := ToIndexes(position)
	return board.Board[rank][file]
}

func (board *Board) Clear(position string) {
	file, rank := ToIndexes(position)
	board.Board[rank][file] = "  "
}

func (board *Board) Move(from string, to string) {
	piece := board.GetPiece(from)
	board.Clear(from)
	board.SetPiece(to, piece)
}

func (board *Board) GetLegalMoves(position string) (moves []string) {
	file, rank := ToIndexes(position)
	piece := board.Board[rank][file]
	fmt.Println(piece)

	moves = []string{}

	switch string(piece[1]) {
	case "P":
		if rank < 7 {
			if board.Board[rank+1][file] != "XX" {
				moves = append(moves, FromIndexes(rank+1, file))
			}
		}
	case "N":
	case "B":
	case "R":
	case "Q":
	case "K":
	}

	return moves
}

func ToIndexes(position string) (file, rank int) {
	file = int(position[0] - 'a')
	rank = int(position[1] - '1')
	return file, rank
}

func FromIndexes(file, rank int) (position string) {
	return string(file + 65)
}