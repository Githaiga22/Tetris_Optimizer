package main

import (
	"fmt"
	"os"

	"allan/tetris"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] != "" {
		result, err := Start(os.Args[1])
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Error: command is not correct")
		fmt.Println("Example: go run . tetris.txt")
	}
}

// Start function takes the name of the input file as a parameter,
// opens the file, reads the tetrominoes from it, and attempts to solve the puzzle.
// It returns the solution as a string or an error if the process fails.
func Start(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()

	// Read the input file to get the tetrominoes as an array.
	myArray, err := tetris.ReadInputFile(file)
	if err != nil {
		return "File Error", err
	} else {
		//solve the puzzle using the read tetrominoes.
		tetris.Solve(myArray)
		// Return the solution as a string.
		return tetris.PrintSolution(), nil
	}
}