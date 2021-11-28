package engine

import "image/color"

// SandID is a unique identifier for a kind of sand.
type SandID uint16

// SandUpdateFunc is the function called to update a grain of sand on each frame.
type SandUpdateFunc func(grain *GrainWithMetadata)

// SandReactionFunc is the function called to check and update a grain of sand when it comes into contact with another
// grain.
type SandReactionFunc func(grain *GrainWithMetadata, contact *GrainWithMetadata) bool

// SandInitializeFunc is the function called when creating a grain of sand.
type SandInitializeFunc func(grain *GrainWithMetadata)

// Sand is an archetype for a grain of sand.
// This describes what a kind of sand is and how it behaves.
type Sand struct {
	Name    string
	Color   color.RGBA
	Density float64
	Tags    uint64

	Reactions []SandReactionFunc
	Behaviors []SandUpdateFunc
	Init      SandInitializeFunc
}

// HasTag returns true if the Sand has a tag.
func (sand *Sand) HasTag(tag uint64) bool {
	return sand.Tags&tag != 0
}

// SandWithSandID is a pointer to a game's Sand and its associated SandID.
type SandWithSandID struct {
	*Sand
	ID SandID
}
