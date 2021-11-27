package engine

// SandboxContents is a struct that contains all the Grain grains in a sandbox.
type SandboxContents struct {
	Width   uint
	Height  uint
	Grains  []Grain
	FrameID uint8
}

// At returns the value of the grain at the x and y coordinate pair.
// If the coordinate pair is outside the bounds of the SandboxContents, a nil pointer will be returned.
func (contents *SandboxContents) At(x, y int) *Grain {
	if x < 0 || x >= int(contents.Width) || y < 0 || y >= int(contents.Height) {
		return nil
	}

	index := y*int(contents.Height) + x
	return &contents.Grains[index]
}

// Sandbox is the game field and config of the falling sands game.
type Sandbox struct {
	Contents *SandboxContents
	Sands    []Sand
}

// Size returns the (width, height) of the Sandbox.
func (sandbox *Sandbox) Size() (uint, uint) {
	return sandbox.Width(), sandbox.Height()
}

// Width returns the width of the sandbox.
func (sandbox *Sandbox) Width() uint {
	return sandbox.Contents.Width
}

// Height returns the height of the sandbox.
func (sandbox *Sandbox) Height() uint {
	return sandbox.Contents.Height
}

// GrainAt returns the SandboxGrain at a specific location.
func (sandbox *Sandbox) GrainAt(x, y int) GrainWithMetadata {
	return GrainWithMetadata{
		Grain:   sandbox.Contents.At(x, y),
		X:       x,
		Y:       y,
		sandbox: sandbox,
	}
}

// NewSandbox creates a new Sandbox.
func NewSandbox(width, height uint, sands []Sand) Sandbox {
	return Sandbox{
		Sands: sands,
		Contents: &SandboxContents{
			Width:  width,
			Height: height,
			Grains: make([]Grain, width*height),
		},
	}
}
