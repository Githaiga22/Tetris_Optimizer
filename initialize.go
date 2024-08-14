package tetris

import (
	"bufio"
	"errors"
	"io"
)
// ReadInputFile reads the input file containing tetromino definitions.
// It parses the file and returns an array of tetrominoes or an error if the format is invalid.
func ReadInputFile(file io.Reader) ([][4][4]string, error) {
	fileError := errors.New("File Error")
	var tetrominoArray [][4][4]string // initialize slice for the pieces
	var tetromino [4][4]string
	scanner := bufio.NewScanner(file) // to read the file line by line
	index := 0
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // String to map tetrominoes to letters.
	for scanner.Scan() {
		for i := 0; i < 4; i++ {
			if i > 0 && !scanner.Scan() {
				return nil, fileError
			}
			str := scanner.Text()
			if str == "" {
				return nil, fileError
			} else {
				var arr [4]string
				if len(str) != 4 { // check that the piece has 4 lines which is the correct format
					return nil, fileError
				}
				for ind := range arr {
					if rune(str[ind]) == '.' {
						arr[ind] = "."
					} else if rune(str[ind]) == '#' {
						arr[ind] = string(alpha[index])
					} else {
						return nil, fileError
					}
				}
				tetromino[i] = arr
			}
		}
		index++
		if !CheckPiece(tetromino) {
			return nil, fileError
		}
		tetromino = OptimizeTetromino(tetromino)
		tetrominoArray = append(tetrominoArray, tetromino)
		if scanner.Scan() && scanner.Text() != "" {
			return nil, fileError
		}
	}
	if len(tetrominoArray) == 0 {
		return nil, fileError
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return tetrominoArray, nil
}

// InitSquare initializes a square board of size n x n filled with empty spaces (".").
// It returns a 2D slice representing the empty board.
func InitSquare(n int) [][]string {
	//initializes a square
	var Square [][]string
	var row []string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			row = append(row, ".")
		}
		Square = append(Square, row)
		row = []string{}
	}
	return Square
}
// CheckPiece validates a tetromino to ensure it has exactly four blocks and appropriate connections.
// It checks that the tetromino does not have more or less than four filled blocks and that they are connected.
func CheckPiece(tetromino [4][4]string) bool {
	c := 0  // Counter for filled blocks.
	d := 0  // Counter for adjacent connections.

	for a, elem := range tetromino {
		for b, elem2 := range elem {
			if elem2 != "." {
				d++  // Increment filled block count.
				// Check for adjacent blocks in all four directions.
				if a+1 < 4 && tetromino[a+1][b] != "." {
					c++
				}
				if a-1 >= 0 && tetromino[a-1][b] != "." {
					c++
				}
				if b+1 < 4 && tetromino[a][b+1] != "." {
					c++
				}
				if b-1 >= 0 && tetromino[a][b-1] != "." {
					c++
				}
			}
		}
	}
	if d != 4 {
		return false // Return false if the tetromino does not have exactly four blocks.
	}
	// Return true if the tetromino has valid connections.
	if c == 6 || c == 8 {
		return true
	}
	return false
}