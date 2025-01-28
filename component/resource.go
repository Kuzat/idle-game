package component

import "github.com/yohamta/donburi"

type ResourceData struct {
	Type string
}

var Resource = donburi.NewComponentType[ResourceData]()
