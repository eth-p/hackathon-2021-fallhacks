package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

var Fall engine.SandUpdateFunc = func(sand engine.GrainWithMetadata) {
	if tryFall(sand, 0, 1) || tryFall(sand, -1, 1) || tryFall(sand, 1, 1) {
		return
	}

	// Otherwise, increase the density timer.
	sand.Grain.DensityTimer += 1
}

// tryFall attempts to make a grain of sand fall.
func tryFall(sand engine.GrainWithMetadata, dx, dy int) bool {
	var below engine.GrainWithMetadata
	if below = sand.Neighbor(dx, dy); !below.IsActionable() {
		return false
	}

	// If the grain below is the same kind, skip it.
	if below.Kind() == sand.Kind() {
		return false
	}

	// Calculate the density difference.
	densityDifference := sand.Kind().Density - below.Kind().Density
	if densityDifference <= 0 {
		return false
	}

	// If the current grain is denser than the sand below it by at least 1, immediately drop through.
	if densityDifference >= 1 {
		below.MoveSwap(sand)
		return true
	}

	// If the density timer is greater than (1/∆density) frames, drop through.
	//timer := sand.Grain.DensityTimer
	//timerNeeded := uint8(1 / densityDifference)
	//if timer >= timerNeeded {
	below.MoveSwap(sand)
	return true
	//}

	//return false
}
