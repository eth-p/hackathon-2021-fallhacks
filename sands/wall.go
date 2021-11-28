package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

var WALL = engine.Sand{
	Name:    "WALL",
	Color:   color.RGBA{R: 78, G: 0, B: 0, A: 0},
	Density: 0,

	Behaviors: []engine.SandUpdateFunc{},
	Init: func(grain *engine.GrainWithMetadata) {
		grain.SetFlag(engine.GrainLocked, true)
	},
}
