package sands

import (
	"image/color"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands/behaviors"
)

// AIR is the default type of grain.
var AIR = engine.Sand{
	Color:   color.RGBA{R: 255, G: 0, B: 0, A: 0},
	Density: 0,

	Update: behaviors.Nothing,
}
