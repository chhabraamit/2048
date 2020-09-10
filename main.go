package main

import (
	"fmt"
	"github.com/chhabra/2048/game"
)

func main() {
	fmt.Printf("Getting started\n")
	g := game.New()
	for i := 0; i < 10; i++ {
		g.Display()
		g.AddElement()
		g.TakeInput()
		fmt.Println("new game")
	}
}
