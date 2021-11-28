package engine

import "github.com/hajimehoshi/ebiten/v2"

type Sandgine struct {
	Sandbox *Sandbox
	Config  Config

	currentUpdateIntervalCounter uint

	pxBuf []byte
	pxImg *ebiten.Image
}

func NewSandgine(sandbox *Sandbox, config Config) *Sandgine {
	return &Sandgine{
		Sandbox: sandbox,
		Config:  config,

		pxBuf: make([]byte, 4*sandbox.Width()*sandbox.Height()),
		pxImg: ebiten.NewImage(int(sandbox.Width()), int(sandbox.Height())),
	}
}

func (engine *Sandgine) Update() error {
	if engine.Config.Paused {
		return nil
	}

	if engine.currentUpdateIntervalCounter >= engine.Config.UpdateInterval {
		engine.currentUpdateIntervalCounter = 0
		engine.Sandbox.Update()
	} else {
		engine.currentUpdateIntervalCounter += 1
	}

	return nil
}

func (engine *Sandgine) Render() *ebiten.Image {
	pxBuf := engine.pxBuf

	// Render each individual sand.
	for i, grain := range engine.Sandbox.Contents.Grains {
		archetype := engine.Sandbox.sands[grain.Kind]

		iS := 4 * i
		pxBuf[iS+0] = archetype.Color.R
		pxBuf[iS+1] = archetype.Color.G
		pxBuf[iS+2] = archetype.Color.B
		pxBuf[iS+3] = 255
	}

	// Update the texture.
	engine.pxImg.ReplacePixels(pxBuf)
	return engine.pxImg
}
