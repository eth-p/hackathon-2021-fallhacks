package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
)

var TORCH = engine.Sand{
	Name:    "Torch",
	Color:   color.RGBA{R: 236, G: 60, B: 0, A: 0},
	Density: 0,

	Behaviors: []engine.SandUpdateFunc{
		behaviors.Generator(5, 0, -1, &FIRE),
	},

	Init: func(grain *engine.GrainWithMetadata) {
		grain.SetFlag(engine.GrainLocked, true)
	},
}
