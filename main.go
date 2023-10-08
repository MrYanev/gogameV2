package main

import (
	"log"

	"github.com/MrYanev/gogameV2/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := internal.NewGame()
	ebiten.SetWindowTitle("Testing")
	ebiten.SetWindowSize(internal.ScreenWidth*3, internal.ScreenHeight*2)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
