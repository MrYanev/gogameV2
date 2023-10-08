package main

import (
	"log"

	"github.com/MrYanev/gogameV2/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := internal.NewGame()
	ebiten.SetWindowTitle("Testing")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
