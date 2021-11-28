package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/reactions"
)

var WOOD = engine.Sand{
	Name:    "WOOD",
	Color:   color.RGBA{R: 78, G: 78, B: 0, A: 0},
	Density: 0,

	Behaviors: []engine.SandUpdateFunc{
		behaviors.Reactable(),
	},

	Init: func(grain *engine.GrainWithMetadata) {
		grain.SetFlag(engine.GrainLocked, true)
	},

	Reactions: []engine.SandReactionFunc{
		reactions.Flammable(engine.TagPyro, &FIRE),
	},
}
