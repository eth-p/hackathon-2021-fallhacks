package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// Expires causes a grain to expire after a given duration.
func Expires(frames uint16) engine.SandUpdateFunc {
	return func(grain *engine.GrainWithMetadata) {
		if grain.GenericTimer > frames {
			grain.Clear()
			return
		}

		grain.GenericTimer += 1
	}
}
