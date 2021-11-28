package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/reactions"
)

var ICE = engine.Sand{
	Name:    "Ice",
	Color:   color.RGBA{R: 100, G: 116, B: 189, A: 0},
	Density: 1,

	Behaviors: []engine.SandUpdateFunc{
		behaviors.Reactable(),
	},

	Init: func(grain *engine.GrainWithMetadata) {
		grain.SetFlag(engine.GrainLocked, true)
	},

	Reactions: []engine.SandReactionFunc{
		reactions.Flammable(engine.TagPyro, &WATER),
	},
}
