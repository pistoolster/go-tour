package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0] = append(board[0], "test")

	board[1][2] = "X"
	board[0][2] = "O"
	board[2][0] = "X"
	board[2][2] = "O"
	board[0][0] = "X"
	board[1][0] = "O"

	//for i := 0; i < len(board); i ++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}

	for i := range board {
		fmt.Printf("%d %s\n", i, strings.Join(board[i], " "))
	}
}
