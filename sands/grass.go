package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/reactions"
)

var GRASS = engine.Sand{
	Name:    "Grass",
	Color:   color.RGBA{R: 31, G: 193, B: 5, A: 0},
	Density: 1,

	Behaviors: []engine.SandUpdateFunc{
		grassSpreadBehavior,
		behaviors.Gravity(0, 1),
	},

	Reactions: []engine.SandReactionFunc{
		reactions.Flammable(engine.TagPyro, &FIRE),
	},
}

func grassSpreadBehavior(grain *engine.GrainWithMetadata) {
	for _, neighbor := range []engine.GrainWithMetadata{
		grain.Neighbor(-1, 0),
		grain.Neighbor(1, 0),
		grain.Neighbor(-1, 1),
		grain.Neighbor(1, 1),
		grain.Neighbor(-1, -1),
		grain.Neighbor(1, 1),
	} {
		if !neighbor.IsActionable() {
			continue
		}

		// Only grow on dirt.
		if !neighbor.Kind().HasTag(engine.TagDirt) {
			continue
		}

		// Only grow if the dirt is below air.
		aboveNeighbor := neighbor.Neighbor(0, -1)
		if aboveNeighbor.KindID() == 0 || aboveNeighbor.Kind().HasTag(engine.TagAir) {
			neighbor.SetKind(grain.KindID())
			return
		}
	}
}
