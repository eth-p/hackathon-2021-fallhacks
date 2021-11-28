package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// Fluid creates a fluid behavior.
func Fluid() engine.SandUpdateFunc {
	return func(grain *engine.GrainWithMetadata) {
		if !doGravity(grain, 1, 0) {
			doGravity(grain, -1, 0)
		}
	}
}
