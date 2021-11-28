package engine

func (sandbox *Sandbox) Update() {
	// Increment the update counter value.
	sandbox.Contents.FrameID += 1
	currentFrame := sandbox.Contents.FrameID
	println(currentFrame)
	// Update all grains of sand.
	xMax := int(sandbox.Width())
	yMax := int(sandbox.Height())

	i := 0
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			grainPtr := &sandbox.Contents.Grains[i]
			kindPtr := &sandbox.Sands[grainPtr.Kind]
			i += 1

			// If the grain was already updated this frame, skip processing it.
			if grainPtr.updateCounter == currentFrame {
				continue
			}

			// Update the grain.
			kindPtr.Update(GrainWithMetadata{
				X:       x,
				Y:       y,
				Grain:   grainPtr,
				kind:    kindPtr,
				sandbox: sandbox,
			})

		}
	}
}
