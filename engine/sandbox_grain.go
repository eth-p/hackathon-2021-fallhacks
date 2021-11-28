package engine

// Grain is a grain of sand in a Sandbox.
// It is kept to a machine word size (i.e. 64 bits) for performance reasons.
type Grain struct {
	Kind          SandID
	flags         uint8
	updateCounter uint8

	DensityTimer uint8
	GenericTimer uint16

	State uint8
}

// GrainWithMetadata is a Grain with additional metadata relating to its location in a Sandbox.
type GrainWithMetadata struct {
	*Grain
	X       int
	Y       int
	sandbox *Sandbox
	kind    *Sand
}

// ---------------------------------------------------------------------------------------------------------------------

// GrainFlag is a flag applied to a Grain.
type GrainFlag uint8

const (
	GrainLocked = 0b00000001
)

// GetFlag gets a GrainFlag.
func (grain *GrainWithMetadata) GetFlag(flag GrainFlag) bool {
	if grain.Grain == nil {
		return false
	}

	return (grain.flags & uint8(flag)) != 0
}

// SetFlag sets a GrainFlag.
func (grain *GrainWithMetadata) SetFlag(flag GrainFlag, value bool) {
	grain.flags &= ^uint8(flag) // Remove the flag.
	if value {
		grain.flags |= uint8(flag)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Kind returns the Sand kind of this Grain.
func (grain *GrainWithMetadata) Kind() *Sand {
	if grain.kind == nil {
		grain.kind = &grain.sandbox.sands[grain.Grain.Kind]
	}

	return grain.kind
}

// SetKind sets the Grain to a specific SandID, using its default values.
func (grain *GrainWithMetadata) SetKind(id SandID) {
	grain.Grain.Kind = id
	grain.kind = &grain.sandbox.sands[id]
	grain.setMoved()

	// Initialize the grain (if applicable).
	if grain.kind.Init != nil {
		grain.kind.Init(grain)
	}
}

// SetKindByRef sets the Grain to a specific Sand, using its default values.
func (grain *GrainWithMetadata) SetKindByRef(kind *Sand) {
	grain.SetKind(grain.sandbox.sandsByName[kind.Name])
}

// Clear clears the Grain to the default value.
func (grain *GrainWithMetadata) Clear() {
	*grain.Grain = Grain{}
	grain.SetUpdated()
}

// ---------------------------------------------------------------------------------------------------------------------

// HasUpdated returns true if the Grain was already updated in the current frame.
// Grains that have already been updated should not be changed until the next frame.
func (grain *GrainWithMetadata) HasUpdated() bool {
	return grain.sandbox.Contents.FrameID == grain.updateCounter
}

// IsActionable returns true if the Grain has not updated this frame, and is inside the bounds of the Sandbox.
func (grain *GrainWithMetadata) IsActionable() bool {
	return grain.Grain != nil && !grain.HasUpdated()
}

// SetUpdated sets the Grain to be considered updated.
func (grain *GrainWithMetadata) SetUpdated() {
	if grain.Grain.Kind != 0 {
		grain.updateCounter = grain.sandbox.Contents.FrameID
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// Neighbor returns a neighboring GrainWithMetadata.
// dX and dY refer to the relative delta of the target's X and Y coordinates.
//
// Example:
//   Neighbor(0, 0) -> Current grain.
//   Neighbor(0, 1) -> The grain below.
func (grain *GrainWithMetadata) Neighbor(dX, dY int) GrainWithMetadata {
	return grain.sandbox.GrainAt(grain.X+dX, grain.Y+dY)
}

// ---------------------------------------------------------------------------------------------------------------------

// MoveSwap swaps the Grain with another Grain.
// dX and dY refer to the relative delta of the target's X and Y coordinates.
//
// Example:
//   Neighbor(0, 0) -> Current grain.
//   Neighbor(0, 1) -> The grain below.
func (grain *GrainWithMetadata) MoveSwap(other GrainWithMetadata) {
	// Use copy-by-value semantics to swap the grain values.
	otherCopy := *other.Grain

	*other.Grain = *grain.Grain
	*grain.Grain = otherCopy

	// Update
	grain.setMoved()
	other.setMoved()
}

// setMoved sets the Grain to be considered moved.
// This will set the grain as updated and clear the density timer.
func (grain *GrainWithMetadata) setMoved() {
	grain.SetUpdated()
	grain.DensityTimer = 0
}
