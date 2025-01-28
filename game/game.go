package game

import (
	"github.com/Kuzat/idle-game/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type Game struct {
	scene        Scene
	screenWidth  int
	screenHeight int
}

func NewGame(width, height int) *Game {

	// makes sure to load all assets

	// Set the start Scenes
	g := &Game{
		screenWidth:  width,
		screenHeight: height,
	}

	g.scene = scene.NewOutside(g)
	return g
}

func (g *Game) Width() int {
	return g.screenWidth
}

func (g *Game) Height() int {
	return g.screenHeight
}

func (g *Game) Update() error {
	return g.scene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	if g.screenWidth == 0 || g.screenHeight == 0 {
		return width, height
	}

	return g.screenWidth, g.screenHeight
}
