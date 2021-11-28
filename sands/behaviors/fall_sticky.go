package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

var FallSticky engine.SandUpdateFunc = func(sand engine.GrainWithMetadata) {
	if tryFall(sand, 0, 1) {
		return
	}

	// Otherwise, increase the density timer.
	sand.Grain.DensityTimer += 1
}
