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
		{"  ", "BP", "  ", "  ", "  ", "  ", "BP", "  "},
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
		{"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
		{"  ", "WP", "  ", "  ", "  ", "  ", "  ", "  "},
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
	fmt.Println(file, rank)
	piece := board.Board[rank][file]

	moves = []string{}

	switch string(piece[1]) {
	case "P":
		if string(piece[0]) == "W" && rank != 7 {
			// move
			if board.Board[rank+1][file] == "  " {
				moves = append(moves, FromIndexes(file, rank+1))
				if rank == 1 && board.Board[rank+2][file] == "  " {
					moves = append(moves, FromIndexes(file, rank+2))
				}
			}
			// capture
			if file != 0 && string(board.Board[rank+1][file-1][0]) == "B" {
				moves = append(moves, FromIndexes(file-1, rank+1))
			}
			if file != 7 && string(board.Board[rank+1][file+1][0]) == "B" {
				moves = append(moves, FromIndexes(file+1, rank+1))
			}
		}
		if string(piece[0]) == "B" && rank != 0 {
			// move
			if board.Board[rank-1][file] == "  " {
				moves = append(moves, FromIndexes(file, rank-1))
				if rank == 6 && board.Board[rank-2][file] == "  " {
					moves = append(moves, FromIndexes(file, rank-2))
				}
			}
			// capture
			if file != 0 && string(board.Board[rank-1][file-1][0]) == "W" {
				moves = append(moves, FromIndexes(file-1, rank-1))
			}
			if file != 7 && string(board.Board[rank-1][file+1][0]) == "W" {
				moves = append(moves, FromIndexes(file+1, rank-1))
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
	return string(rune(file + 'a')) + string(rune(rank + '1'))
}