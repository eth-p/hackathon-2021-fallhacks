package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
)

var FIRE = engine.Sand{
	Name:    "Fire",
	Color:   color.RGBA{R: 236, G: 82, B: 0, A: 0},
	Density: 0,

	Tags: engine.TagPyro,

	Behaviors: []engine.SandUpdateFunc{
		behaviors.Expires(40),
		behaviors.Gravity(0, -1),
	},

	Init: func(grain *engine.GrainWithMetadata) {
		grain.GenericTimer = 0
	},
}
