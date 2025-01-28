package scene

import (
	"github.com/Kuzat/idle-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/yohamta/donburi"
	"image/color"
)

var (
	BackgroundColor = color.RGBA{
		R: 20,
		G: 146,
		B: 16,
		A: 1,
	}
)

type GameSize interface {
	Width() int
	Height() int
}

type System interface {
	Update() error
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type Game struct {
	size      GameSize
	world     donburi.World
	systems   []System
	drawables []Drawable
}

func NewOutside(size GameSize) *Game {
	return &Game{
		size: size,
	}
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BackgroundColor)
	text.Draw(
		screen,
		"Hello, world",
		assets.NormalFont,
		g.size.Width()/2,
		g.size.Height()/2,
		color.White,
	)
}
