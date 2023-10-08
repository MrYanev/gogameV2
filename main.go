package main

import (
	"gogameV2/internal"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := internal.NewGame()
	ebiten.SetWindowTitle("Testing")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
