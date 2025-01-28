package main

import (
	"github.com/Kuzat/idle-game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	// Load some config?

	mw, mh := ebiten.Monitor().Size()

	w, h := int(float64(mw)/1.5), int(float64(mh)/1.5)
	ebiten.SetWindowSize(w, h)

	err := ebiten.RunGame(game.NewGame(w, h))
	if err != nil {
		log.Fatal(err)
	}
}
