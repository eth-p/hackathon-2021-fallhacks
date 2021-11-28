package behaviors

import (
	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// Generator generates a grain every few frames.
func Generator(frames uint16, dx, dy int, spawns *engine.Sand) engine.SandUpdateFunc {
	return func(grain *engine.GrainWithMetadata) {
		if (grain.GenericTimer % frames) == 0 {
			target := grain.Neighbor(dx, dy)
			if target.IsActionable() && target.Kind().HasTag(engine.TagAir) {
				target.SetKindByRef(spawns)
			}
		}
	}
}
