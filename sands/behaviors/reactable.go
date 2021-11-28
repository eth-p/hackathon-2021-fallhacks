package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// Reactable checks if a neighbor causes the grain to react.
func Reactable() engine.SandUpdateFunc {
	return func(grain *engine.GrainWithMetadata) {
		for _, target := range []engine.GrainWithMetadata{
			grain.Neighbor(0, 1),
			grain.Neighbor(0, -1),
			grain.Neighbor(-1, 0),
			grain.Neighbor(1, 0),
		} {
			for _, reaction := range grain.Kind().Reactions {
				if reaction(grain, &target) {
					return
				}
			}
		}
	}
}
