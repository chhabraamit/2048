package main

import (
	"flag"
	"fmt"
	"github.com/chhabra/2048/game"
	log "github.com/sirupsen/logrus"
)

func main() {
	debug := flag.Bool("debug", false, "debugging flag")
	flag.Parse()
	if *debug {
		log.SetLevel(log.DebugLevel)
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
