package main

import "fmt"

func main() {
	GameLoop()
}

func GameLoop() {
	for {
		PrintGame()
	}
}

func PrintGame() {
	fmt.Println("hello game")
}
