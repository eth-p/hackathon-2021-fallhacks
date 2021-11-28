package engine

import "image/color"

// SandID is a unique identifier for a kind of sand.
type SandID uint16

// SandUpdateFunc is the function called to update a grain of sand on each frame.
type SandUpdateFunc func(grain GrainWithMetadata)

// Sand is an archetype for a grain of sand.
// This describes what a kind of sand is and how it behaves.
type Sand struct {
	Name    string
	Color   color.RGBA
	Density float64

	Update SandUpdateFunc
}
