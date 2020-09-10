package main

import (
	"fmt"
	"github.com/chhabra/2048/game"
)

func main() {
	fmt.Printf("Getting started\n")
	g := game.New()
	g.AddElement()
	g.AddElement()

	for true {
		if g.IsOver() {
			break
		}
		g.AddElement()
		g.Display()
		g.TakeInput()
		//g.Display()
		fmt.Println("new game")
	}
	fmt.Printf("**** Game Over **** \n")
}
