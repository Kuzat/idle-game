package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type Game struct {
	resources map[string]int

	// Button position /size
	gatherBushBtnX int
	gatherBushBtnY int
	gatherBushBtnW int
	gatherBushBtnH int
}

func NewGame() *Game {
	return &Game{
		resources: map[string]int{
			"Berries": 0,
			"Fibers":  0,
		},

		gatherBushBtnX: 50,
		gatherBushBtnY: 50,
		gatherBushBtnW: 140,
		gatherBushBtnH: 40,
	}
}

func (g *Game) Update() error {
	// Check for mouse click
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// If the click is inside the 'Gather Bush' button
		if x >= g.gatherBushBtnX && x <= g.gatherBushBtnX+g.gatherBushBtnW &&
			y >= g.gatherBushBtnY && y <= g.gatherBushBtnY+g.gatherBushBtnH {
			// Increase resource counts
			g.resources["Berries"] += 1
			g.resources["Fibers"] += 1
		}
	}

	// In real idle game, we might do more logic here
	// - Passive resource generation
	// Checking for crafting
	// Timers, etc.
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear or draw background color (optional)
	screen.Fill(color.RGBA{R: 0xCC, G: 0xEE, B: 0xFF, A: 0xFF})

	// Draw 'Gatcher Bush' buttons (just a colored rectangle)
	vector.DrawFilledRect(
		screen,
		float32(g.gatherBushBtnX),
		float32(g.gatherBushBtnY),
		float32(g.gatherBushBtnW),
		float32(g.gatherBushBtnH),
		color.RGBA{
			R: 0x55,
			G: 0x88,
			B: 0x55,
			A: 0xFF,
		},
		false,
	)

	// Draw button text
	msg := "Gather Bushes"
	ebitenutil.DebugPrintAt(screen, msg, g.gatherBushBtnX+10, g.gatherBushBtnY+12)

	// show resource counts
	resourceText := fmt.Sprintf("Berries: %d\nFibers: %d", g.resources["Berries"], g.resources["Fibers"])
	ebitenutil.DebugPrintAt(screen, resourceText, 50, 120)

	// in a larger game, you would draw other UI elements here:
	// - Additional buttons for stone, wood, etc.
	// - Crafting grids
	// - Inventory or resource panels
}

// Layout takes the outside window size and returns the logical screen size
// You can keep it simple for a 640x480 game.
func (g *Game) Layout(outsideWith, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := NewGame()

	// set up the window
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Idle Crafting Game")

	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
