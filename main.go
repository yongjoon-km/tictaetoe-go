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
		winner := GetWinner(board)
		if winner != 0 {
			fmt.Printf("Game end winner is %d\n", winner)
			break
		}
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

func GetWinner(board [3][3]int) int {

	// horizontal match
	i := 0
	for i < 3 {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
		i++
	}

	// vertical match

	j := 0
	for j < 3 {
		if board[0][j] == board[1][j] && board[0][j] == board[2][j] {
			return board[0][j]
		}
		j++
	}

	// diagonal match
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return board[0][2]
	}
	return 0
}
