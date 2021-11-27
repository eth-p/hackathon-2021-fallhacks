package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

var FallSticky engine.SandUpdateFunc = func(sand engine.GrainWithMetadata) {
	var below engine.GrainWithMetadata
	if below = sand.Neighbor(0, 1); !below.IsActionable() {
		return
	}

	// Calculate the density difference.
	densityDifference := sand.Kind().Density - below.Kind().Density
	if densityDifference <= 0 {
		return
	}

	// If the current grain is denser than the sand below it by at least 1, immediately drop through.
	if densityDifference >= 1 {
		below.MoveSwap(sand)
		return
	}

	// If the density timer is greater than (1/∆density) frames, drop through.
	timer := sand.Grain.DensityTimer
	timerNeeded := uint8(1 / densityDifference)
	if timer >= timerNeeded {
		below.MoveSwap(sand)
		return
	}

	// Otherwise, increase the density timer.
	sand.Grain.DensityTimer += 1
}
