package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type SpriteData ebiten.Image

var Sprite = donburi.NewComponentType[SpriteData]()
