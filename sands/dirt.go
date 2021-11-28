package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
)

var DIRT = engine.Sand{
	Name:    "Dirt",
	Color:   color.RGBA{R: 72, G: 38, B: 5, A: 0},
	Density: 2.30,

	Tags: engine.TagDirt,

	Behaviors: []engine.SandUpdateFunc{
		behaviors.Gravity(0, 1),
		behaviors.Gravity(-1, 1),
		behaviors.Gravity(1, 1),
	},
}
