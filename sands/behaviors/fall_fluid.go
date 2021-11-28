package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

var FallFluid engine.SandUpdateFunc = func(sand engine.GrainWithMetadata) {

	// Try below.
	if tryFall(sand, 0, 1) || tryFall(sand, -1, 1) || tryFall(sand, 1, 1) {
		return
	}

	// Try a grain beside itself.
	if sand.Timer%2 == 0 {
		if tryFall(sand, -1, 0) {
			return
		}
	} else {
		if tryFall(sand, 1, 0) {
			return
		}
	}

	// Swap the direction if it couldn't do it.
	sand.Timer += 1

	// Otherwise, increase the density timer.
	sand.Grain.DensityTimer += 1
}
