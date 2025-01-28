package component

import "github.com/yohamta/donburi"

type InventoryData struct {
	Items map[string]int // Map from item name to quantity
}

var Inventory = donburi.NewComponentType[InventoryData]()
