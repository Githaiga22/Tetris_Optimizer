package tetris

import (
	"math"
)

// variable to hold the board where tetrominoes will be placed.
var board [][]string

// BacktrackSolver attempts to place all tetrominoes on the board using a recursive backtracking approach.
// It returns true if all tetrominoes are successfully placed, otherwise false.
func BacktrackSolver(tetrominoes [][4][4]string, n int) bool {
	if n == len(tetrominoes) { // base condition when all tetrominoes are placed, board is solved
		return true
	}
	// Iterate through each cell of the board to find a position for the current tetromino
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if CheckPosition(i, j, tetrominoes[n]) { // check if we can place current tetrominoe on the board anywhere
				Insert(i, j, tetrominoes[n])           // if we can place it at this location, check if we can place another piece
				if BacktrackSolver(tetrominoes, n+1) { // Recursively attempt to place the next tetromino.
					return true
				}
				Remove(i, j, tetrominoes[n]) // if the next piece can't be placed, backtrack
			}
		}
	} // if we can't place tetro anywhere, return false
	return false
}

// Insert places a tetromino on the board at the specified position (i, j).
// It updates the board by marking the positions occupied by the tetromino.
func Insert(i, j int, tetro [4][4]string) { // insert piece and when all 4 piecec "#" are placed, no need to place '.'
	a, b, c := 0, 0, 0
	for a < 4 {
		for b < 4 {
			if tetro[a][b] != "." {
				c++
				board[i+a][j+b] = tetro[a][b]
				if c == 4 {
					break
				}
			}
			b++
		}
		b = 0
		a++
	}
}

// Remove takes a tetromino off the board from the specified position (i, j).
// It updates the board to mark the positions as empty (".").
func Remove(i, j int, tetro [4][4]string) { // remove piece at current location
	a, b, c := 0, 0, 0 // Initialize counters for row, column, and block count
	for a < 4 {
		for b < 4 {
			// If the current position in the tetromino is not empty (not ".").
			if tetro[a][b] != "." {
				if c == 4 {
					break
				}
				board[i+a][j+b] = "."
			}
			b++
		}
		b = 0
		a++
	}
}

// Solve initializes the board and attempts to solve the Tetris puzzle using the provided tetrominoes.
// It dynamically adjusts the board size until a solution is found and returns the completed board.
func Solve(tetrominoes [][4][4]string) [][]string {
	// initial board starts with dimmension 4*4, if we can't place all tetrominoes
	// increase size by 1 and initialize board
	l := int(math.Ceil(math.Sqrt(float64(4 * len(tetrominoes)))))
	
	board = InitSquare(l)
	for !BacktrackSolver(tetrominoes, 0) {
		l++
		board = InitSquare(l)
	}
	// BacktrackSolver(tetrominoes, 0)
	return board
}

// PrintSolution formats the board into a string representation for output.
// It concatenates each row of the board into a single string with line breaks.
func PrintSolution() string {
	result := ""
	for i := range board {
		if i > 0 {
			result += "\n"
		}
		for j := range board {
			result += board[i][j] // Append each cell to the result string.
		}
	}
	return result
}

// CheckPosition checks if a tetromino can be placed at the specified position (i, j) on the board.
// It ensures that the tetromino fits within the board boundaries and does not overlap with existing blocks.
func CheckPosition(i, j int, tetro [4][4]string) bool {
	for a := 0; a < 4; a++ {
		for b := 0; b < 4; b++ {
			if tetro[a][b] != "." {
				if i+a == len(board) || j+b == len(board) || board[i+a][j+b] != "." {
					return false
				}
			}
		}
	}
	return true
}

// OptimizeTetromino removes unnecessary empty rows and columns from the tetromino.
// It shifts the tetromino upwards and leftwards to minimize its size.
func OptimizeTetromino(tetromino [4][4]string) [4][4]string {
	// optimzes tetromino
	i := 0
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[i][j] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { // if row is all zeroes, shift by 1 row to top
			tetromino = ShiftVertical(tetromino)
			continue
		}
		break
	}
	for {
		zeroes := 0
		for j := 0; j < 4; j++ {
			if tetromino[j][i] == "." {
				zeroes++
			}
		}
		if zeroes == 4 { // if col is all zeroes, shift by 1 col to left
			tetromino = ShiftHorizontal(tetromino)
			continue
		}
		break
	}
	return tetromino
}

// ShiftHorizontal shifts the tetromino left by one column.
// It effectively removes the first column and moves the remaining columns left.
func ShiftVertical(tetromino [4][4]string) [4][4]string {
	// shifts tetromino row by 1
	temp := tetromino[0]
	tetromino[0] = tetromino[1]
	tetromino[1] = tetromino[2]
	tetromino[2] = tetromino[3]
	tetromino[3] = temp
	return tetromino
}

// ShiftHorizontal shifts the tetromino left by one column.
// It effectively removes the first column and moves the remaining columns left.
func ShiftHorizontal(tetromino [4][4]string) [4][4]string {
	// shifts tetromino col by 1
	tetromino = Transpose(tetromino)
	tetromino = ShiftVertical(tetromino)
	tetromino = Transpose(tetromino)
	return tetromino
}

// Transpose switches the rows and columns of the tetromino.
// It creates a new 4x4 array where the value at (i, j) is moved to (j, i).
func Transpose(slice [4][4]string) [4][4]string {
	// transpose tetromino
	xl := len(slice[0])
	yl := len(slice)
	var result [4][4]string
	for i := range result {
		result[i] = [4]string{}
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
