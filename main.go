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
	turn = 1 // O first
}

func GameLoop() {
	for {

		x, y, val, hasNextMove := GetNextMove()
		if !hasNextMove {
			fmt.Println("Invalid position to place stone.")
			continue
		}
		ok := PlaceStone(x, y, val)
		if !ok {
			fmt.Println("Invalid position to place stone.")
			continue
		}
		turn = turn * -1 // Get Next Turn

		PrintGame()
		winner := GetWinner(board)
		if winner != 0 {
			fmt.Printf("Game end winner is %d\n", winner)
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

func PlaceStone(x, y, val int) bool {

	if x < 0 || x >= 3 || y < 0 || y >= 3 {
		return false
	}

	if board[x][y] != 0 {
		return false
	}

	board[x][y] = val
	return true
}
