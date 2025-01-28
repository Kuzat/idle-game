package component

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type PositionData math.Vec2

var Position = donburi.NewComponentType[PositionData]()
