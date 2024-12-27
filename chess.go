package main

import (
	"fmt"
	"strings"
)

// Piece represents a chess piece
type Piece struct {
	Type  string // P, R, N, B, Q, K
	Color string // W (White), B (Black)
}

// Board represents the chessboard
type Board [8][8]*Piece

// NewBoard initializes the chessboard with pieces
func NewBoard() *Board {
	board := &Board{}
	setupRow := func(row int, color string) {
		pieces := []string{"R", "N", "B", "Q", "K", "B", "N", "R"}
		for i, p := range pieces {
			board[row][i] = &Piece{Type: p, Color: color}
		}
	}
	for i := 0; i < 8; i++ {
		board[1][i] = &Piece{Type: "P", Color: "B"}
		board[6][i] = &Piece{Type: "P", Color: "W"}
	}
	setupRow(0, "B")
	setupRow(7, "W")
	return board
}

// Display prints the current state of the board
func (b *Board) Display() {
	fmt.Println("  a b c d e f g h")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", 8-i)
		for j := 0; j < 8; j++ {
			if b[i][j] == nil {
				fmt.Print(". ")
			} else {
				color := b[i][j].Color
				if color == "W" {
					fmt.Printf(strings.ToUpper(b[i][j].Type) + " ")
				} else {
					fmt.Printf(strings.ToLower(b[i][j].Type) + " ")
				}
			}
		}
		fmt.Printf("%d\n", 8-i)
	}
	fmt.Println("  a b c d e f g h")
}

// Move attempts to move a piece
func (b *Board) Move(from, to string) bool {
	fx, fy := parsePosition(from)
	tx, ty := parsePosition(to)
	if fx == -1 || fy == -1 || tx == -1 || ty == -1 {
		fmt.Println("Invalid move format")
		return false
	}

	piece := b[fx][fy]
	if piece == nil {
		fmt.Println("No piece at the source position")
		return false
	}

	// Basic move validation (ignoring specific rules)
	if b[tx][ty] != nil && b[tx][ty].Color == piece.Color {
		fmt.Println("Cannot capture your own piece")
		return false
	}

	b[tx][ty] = piece
	b[fx][fy] = nil
	return true
}

// parsePosition converts a position like "e2" into board indices
func parsePosition(pos string) (int, int) {
	if len(pos) != 2 {
		return -1, -1
	}
	file := pos[0] - 'a'
	rank := '8' - pos[1]
	if file < 0 || file > 7 || rank < 0 || rank > 7 {
		return -1, -1
	}
	return int(rank), int(file)
}

func main() {
	board := NewBoard()
	board.Display()

	for {
		var from, to string
		fmt.Print("Enter your move (e.g., e2 e4): ")
		fmt.Scan(&from, &to)
		if from == "quit" || to == "quit" {
			break
		}
		if board.Move(from, to) {
			board.Display()
		}
	}
}
