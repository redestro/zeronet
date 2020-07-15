package games

import (
	"fmt"
)

type move struct {
	row int
	col int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func configureBoard(board [9]string) [3][3]string {
	newBoard := [3][3]string{
		[3]string{"_"},
		[3]string{"_"},
		[3]string{"_"},
	}
	for i, v := range board {
		var row int = i / 3
		var col int = i % 3
		newBoard[row][col] = v
	}
	fmt.Println(newBoard)
	return newBoard
}

func isMovesLeft(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "_" {
				return true
			}
		}
	}
	return false
}

func evaluate(b [3][3]string, player string, opponent string) int {
	// Checking for Rows for X or O victory.
	for row := 0; row < 3; row++ {
		if b[row][0] == b[row][1] && b[row][1] == b[row][2] {
			if b[row][0] == player {
				return +10
			} else if b[row][0] == opponent {
				return -10
			}
		}
	}

	// Checking for Columns for X or O victory.
	for col := 0; col < 3; col++ {
		if b[0][col] == b[1][col] && b[1][col] == b[2][col] {
			if b[0][col] == player {
				return +10
			} else if b[0][col] == opponent {
				return -10
			}
		}
	}

	// Checking for Diagonals for X or O victory.
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		if b[0][0] == player {
			return +10
		} else if b[0][0] == opponent {
			return -10
		}
	}

	if b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		if b[0][2] == player {
			return +10
		} else if b[0][2] == opponent {
			return -10
		}
	}

	// Else if none of them have won then return 0
	return 0
}

// This is the minimax function. It considers all
// the possible ways the game can go and returns
// the value of the board
func minimax(board [3][3]string, depth int, isMax bool, player string, opponent string) int {
	score := evaluate(board, player, opponent)

	// If Maximizer has won the game return his/her
	// evaluated score
	if score == 10 {
		return score
	}

	// If Minimizer has won the game return his/her
	// evaluated score
	if score == -10 {
		return score
	}

	// If there are no more moves and no winner then
	// it is a tie
	if isMovesLeft(board) == false {
		return 0
	}

	// If this maximizer's move
	if isMax {
		best := -1000

		// Traverse all cells
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// Check if cell is empty
				if board[i][j] == "_" {
					// Make the move
					board[i][j] = player

					// Call minimax recursively and choose
					// the maximum value
					best = max(best, minimax(board, depth+1, !isMax, player, opponent))

					// Undo the move
					board[i][j] = "_"
				}
			}
		}
		return best
	} else { // If this minimizer's move
		best := 1000

		// Traverse all cells
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// Check if cell is empty
				if board[i][j] == "_" {
					// Make the move
					board[i][j] = opponent

					// Call minimax recursively and choose
					// the minimum value
					best = min(best, minimax(board, depth+1, !isMax, player, opponent))

					// Undo the move
					board[i][j] = "_"
				}
			}
		}
		return best
	}
}

// ComputeMove returns the best possible move
func ComputeMove(board [9]string, playerSymbol string, opponentSymbol string) int {
	newBoard := configureBoard(board)
	bestVal := -1000
	var bestMove move
	bestMove.row = -1
	bestMove.col = -1

	// Traverse all cells, evaluate minimax function for
	// all empty cells. And return the cell with optimal
	// value.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// Check if cell is empty
			if newBoard[i][j] == "_" {
				// Make the move
				newBoard[i][j] = playerSymbol
				// compute evaluation function for this
				// move.
				moveVal := minimax(newBoard, 0, false, playerSymbol, opponentSymbol)
				// Undo the move
				newBoard[i][j] = "_"

				// If the value of the current move is
				// more than the best value, then update
				// best/
				if moveVal > bestVal {
					bestMove.row = i
					bestMove.col = j
					bestVal = moveVal
				}
			}
		}
	}
	move := bestMove.row*3 + bestMove.col
	return move
}
