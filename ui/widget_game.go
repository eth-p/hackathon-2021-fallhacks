package ui

import (
	"image"
	"math"

	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
)

// GameWidget is a widget.Widget that renders the game to the screen.
type GameWidget struct {
	widget.Widget

	Engine *engine.Sandgine
}

func (g GameWidget) GetWidget() *widget.Widget {
	return &g.Widget
}

func (g GameWidget) PreferredSize() (int, int) {
	return int(g.Engine.Sandbox.Width()), int(g.Engine.Sandbox.Height())
}

func (g GameWidget) SetLocation(rect image.Rectangle) {
}

func (g GameWidget) Render(screen *ebiten.Image, def widget.DeferredRenderFunc) {
	img := g.Engine.Render()

	// Calculate the scaling factor.
	screenSize := screen.Bounds().Size()
	imgSize := img.Bounds().Size()

	maxScaleX := float64(screenSize.X) / float64(imgSize.X)
	maxScaleY := float64(screenSize.Y) / float64(imgSize.Y)
	scaleFactor := math.Min(maxScaleX, maxScaleY)

	// Calculate the X and Y offsets.
	renderWidth := int(scaleFactor * float64(imgSize.X))
	renderHeight := int(scaleFactor * float64(imgSize.X))
	renderX := float64(screenSize.X-renderWidth) / 2.0
	renderY := float64(screenSize.Y-renderHeight) / 2.0

	// Draw the game screen to the UI.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleFactor, scaleFactor)
	op.GeoM.Translate(renderX, renderY)
	screen.DrawImage(img, op)
}
