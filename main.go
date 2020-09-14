package main

import (
	"flag"
	"fmt"
	"github.com/chhabra/2048/game"
	log "github.com/sirupsen/logrus"
)

const playInstructionDelay = 2

func main() {
	debug := flag.Bool("debug", false, "debugging flag")
	flag.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	fmt.Printf("Use {W A S D} or {h j k l} or Arrow keys to move the board\n")
	fmt.Printf("Press any key to start\n")
	_, err := game.GetCharKeystroke()
	if err != nil {
		log.Fatal("error while taking input to start the game: %v", err)
	}
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
	}
	fmt.Printf("**** Game Over **** \n")
	max, total := g.CountScore()
	fmt.Printf("Score: Max Tile Value:    %v \n", max)
	fmt.Printf("Score: Total Tiles Value: %v \n", total)
}
