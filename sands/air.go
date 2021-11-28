package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// AIR is the default type of grain.
var AIR = engine.Sand{
	Name:    "None",
	Color:   color.RGBA{R: 0, G: 0, B: 0, A: 0},
	Density: 0,
}
