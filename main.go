package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var board [3][3]string
var currentPlayer string

func main() {
	clearScreen()
	initializeBoard()
	currentPlayer = "X"

	for {
		printBoard()
		fmt.Printf("Player %s's turn\n", currentPlayer)
		fmt.Print("Enter row (0-2): ")
		row := getUserInput()
		fmt.Print("Enter column (0-2): ")
		col := getUserInput()

		if isValidMove(row, col) {
			makeMove(row, col)
			if isWin() {
				printBoard()
				fmt.Printf("Player %s wins!\n", currentPlayer)
				break
			} else if isDraw() {
				printBoard()
				fmt.Println("It's a draw!")
				break
			}
			switchPlayer()
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

func printBoard() {
	fmt.Println("---------")
	for i := 0; i < 3; i++ {
		fmt.Printf("| %s | %s | %s |\n", board[i][0], board[i][1], board[i][2])
		fmt.Println("---------")
	}
}

func getUserInput() int {
	var input int
	fmt.Scanln(&input)
	return input
}

func isValidMove(row, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	if board[row][col] != " " {
		return false
	}
	return true
}

func makeMove(row, col int) {
	board[row][col] = currentPlayer
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func isWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}
	return false
}

func isDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}
