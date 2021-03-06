package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
)

var SAND = engine.Sand{
	Name:    "Sand",
	Color:   color.RGBA{R: 176, G: 152, B: 0, A: 0},
	Density: 2,

	Behaviors: []engine.SandUpdateFunc{
		behaviors.Gravity(0, 1),
		behaviors.Gravity(-1, 1),
		behaviors.Gravity(1, 1),
	},
}
