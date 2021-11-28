package reactions

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// Flammable reactions will cause the grain to turn into another grain upon contact with the given tag.
func Flammable(tag uint64, into *engine.Sand) engine.SandReactionFunc {
	return func(grain *engine.GrainWithMetadata, target *engine.GrainWithMetadata) bool {
		if target.Kind().HasTag(tag) {
			grain.SetKindByRef(into)
			return true
		}

		return false
	}
}
