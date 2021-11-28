package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
)

var WATER = engine.Sand{
	Name:    "Water",
	Color:   color.RGBA{R: 0, G: 16, B: 89, A: 0},
	Density: 1,

	Update: behaviors.FallFluid,
}
