package behaviors

import (
	"math"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// Gravity creates a gravity behavior.
func Gravity(dx, dy int) engine.SandUpdateFunc {
	return func(grain *engine.GrainWithMetadata) {
		doGravity(grain, dx, dy)
	}
}

// doGravity attempts to apply gravity on a grain of sand.
// If the grain of sand is unable to move, `false` will be returned.
func doGravity(sand *engine.GrainWithMetadata, dx, dy int) bool {
	sandKind := sand.Kind()

	var target engine.GrainWithMetadata
	if target = sand.Neighbor(dx, dy); !target.IsActionable() {
		return false
	}

	if target.GetFlag(engine.GrainLocked) {
		sand.SetUpdated()
		return true
	}

	// If the target grain is the same kind, skip it.
	if target.Kind() == sandKind {
		return false
	}

	// Check if there's a reaction first.
	// If one occurs, cancel the movement.
	for _, reaction := range sandKind.Reactions {
		if reaction(sand, &target) {
			return true
		}
	}

	// Calculate the density difference.
	densityDifference := sand.Kind().Density - target.Kind().Density

	// If the density difference is too great, it can't fall through the target.
	if sandKind.Density > 0 && densityDifference <= 0 || sandKind.Density < 0 && densityDifference >= 0 {
		return false
	}

	// If the current grain is denser than target grain by at least 1, immediately drop through.
	densityDifferenceAbs := math.Abs(densityDifference)
	if densityDifferenceAbs >= 1 {
		target.MoveSwap(*sand)
		return true
	}

	// If the density timer is greater than (1/abs(∆density)) frames, drop through.
	timer := sand.Grain.DensityTimer
	timerNeeded := uint8(1 / densityDifferenceAbs)
	if timer >= timerNeeded {
		target.MoveSwap(*sand)
		return true
	}

	return false
}
