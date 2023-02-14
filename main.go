package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var board [3][3]int
var turn int

const (
	NONE = iota
	ZERO
	CROSS
)

func main() {
	InitializeGame()
	GameLoop()
}

func InitializeGame() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = 0
		}
	}
	turn = 1 // O first
}

func GameLoop() {
	for {

		x, y, val, hasNextMove := GetNextMove()
		if !hasNextMove {
			fmt.Println("Invalid position to place stone.")
			continue
		}
		if ok := PlaceStone(x, y, val); !ok {
			fmt.Println("Invalid position to place stone.")
			continue
		}
		turn = turn * -1 // Get Next Turn

		PrintGame()
		if winner := GetWinner(board); winner != 0 {
			fmt.Printf("Game end winner is %d\n", winner)
		if winner := GetWinner(board); winner != NONE {
			fmt.Printf("Game end winner is %s\n", ConvertToChar(winner))
			break
		}
	}
}

func GetNextMove() (int, int, int, bool) {
	fmt.Print("x y > ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	if len(s) == 0 {
		return 0, 0, 0, false
	}
	coordinates := strings.Fields(s)
	if len(coordinates) != 2 {
		return 0, 0, 0, false
	}
	x, _ := strconv.ParseInt(coordinates[0], 10, 0)
	y, _ := strconv.ParseInt(coordinates[1], 10, 0)
	return int(x), int(y), turn, true
}

func PrintGame() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c := ConvertToChar(board[i][j])
			fmt.Printf("%s|", c)
		}
		fmt.Println()
	}
}

func ConvertToChar(num int) string {
	if num == NONE {
		return " "
	}
	if num == ZERO {
		return "O"
	} else {
		return "X"
	}
}

func GetWinner(board [3][3]int) int {

	// horizontal match
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return board[i][0]
		}
	}

	// vertical match
	for j := 0; j < 3; j++ {
		if board[0][j] == board[1][j] && board[0][j] == board[2][j] {
			return board[0][j]
		}
	}

	// diagonal match
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return board[0][2]
	}
	return NONE
}

func PlaceStone(x, y, val int) bool {

	if x < 0 || x >= 3 || y < 0 || y >= 3 {
		return false
	}

	if board[x][y] != NONE {
		return false
	}

	board[x][y] = val
	return true
}
