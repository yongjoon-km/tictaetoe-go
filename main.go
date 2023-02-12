package main

import "fmt"

var board [3][3]int

func main() {
	InitializeGame()
	GameLoop()
}

func InitializeGame() {
	i := 0
	for i < 3 {
		j := 0
		for j < 3 {
			board[i][j] = 0
			j++
		}
		i++
	}
}

func GameLoop() {
	for {
		PrintGame()
	}
}

func PrintGame() {
	i := 0
	for i < 3 {
		j := 0
		for j < 3 {
			c := ConvertToChar(board[i][j])
			fmt.Printf("%s|", c)
			j++
		}
		fmt.Println()
		i++
	}
}

func ConvertToChar(num int) string {
	if num == 0 {
		return " "
	}
	if num == 1 {
		return "O"
	} else {
		return "X"
	}
}
