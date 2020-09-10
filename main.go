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
		g.AddElement()
		g.Display()
		g.TakeInput()
		//g.Display()
		fmt.Println("new game")
	}
}
