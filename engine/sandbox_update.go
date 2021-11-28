package engine

func (sandbox *Sandbox) Update() {
	// Increment the update counter value.
	sandbox.Contents.FrameID += 1
	currentFrame := sandbox.Contents.FrameID

	// Update all grains of sand.
	xMax := int(sandbox.Width())
	yMax := int(sandbox.Height())

	i := 0
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			grainPtr := &sandbox.Contents.Grains[i]
			kindPtr := &sandbox.sands[grainPtr.Kind]
			i += 1

			sandbox.updateGrain(currentFrame, kindPtr, &GrainWithMetadata{
				X:       x,
				Y:       y,
				Grain:   grainPtr,
				kind:    kindPtr,
				sandbox: sandbox,
			})
		}
	}
}

func (sandbox *Sandbox) updateGrain(currentFrame uint8, kind *Sand, grain *GrainWithMetadata) {
	grain.GenericTimer += 1
	for _, behavior := range kind.Behaviors {
		if grain.updateCounter == currentFrame {
			return
		}

		behavior(grain)
	}
}
