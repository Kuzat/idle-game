package component

import (
	"github.com/yohamta/donburi"
	"math/rand"
)

type DropsData struct {
	Items map[string]float64 // Map item type to probability to be dropped
}

// GetRandomDrop Get a random dropped item from the Items list using the probabilities in the map
func (d *DropsData) GetRandomDrop() string {
	if len(d.Items) == 0 {
		return "" // Return empty if no items are available
	}

	// Calculate the total weight (sum of all probabilities)
	totalWeight := 0.0
	for _, weight := range d.Items {
		totalWeight += weight
	}

	// Generate a random value between 0 and totalWeight
	randomValue := rand.Float64() * totalWeight

	// Determine which item corresponds to the random value
	cumulativeWeight := 0.0
	for item, weight := range d.Items {
		cumulativeWeight += weight
		if randomValue <= cumulativeWeight {
			return item
		}
	}

	return "" // Return empty if no item matched (this shouldn't happen)
}

var Drops = donburi.NewComponentType[DropsData]()
