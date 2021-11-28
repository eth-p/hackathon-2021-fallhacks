package ui

import (
	"image"
	"math"

	"github.com/blizzy78/ebitenui/input"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/resources"
)

// GameWidget is a widget.Widget that renders the game to the screen and handles user input.
type GameWidget struct {
	*widget.Widget

	Engine *engine.Sandgine

	cursorEntered  bool
	cursorPainting bool
	cursorLast     image.Point

	renderScaleFactor float64
	renderX           float64
	renderY           float64
}

func NewGameWidget(engine *engine.Sandgine) *GameWidget {
	w := &GameWidget{
		Engine: engine,
	}

	w.Widget = widget.NewWidget(
		widget.WidgetOpts.CursorEnterHandler(func(args *widget.WidgetCursorEnterEventArgs) {
			x, y := input.CursorPosition()
			w.cursorEntered = true
			w.cursorLast = image.Point{X: x, Y: y}
		}),

		widget.WidgetOpts.CursorExitHandler(func(args *widget.WidgetCursorExitEventArgs) {
			w.cursorEntered = false
		}),

		widget.WidgetOpts.MouseButtonPressedHandler(func(args *widget.WidgetMouseButtonPressedEventArgs) {
			if args.Button == ebiten.MouseButtonLeft {
				w.cursorPainting = true
			}
		}),

		widget.WidgetOpts.MouseButtonReleasedHandler(func(args *widget.WidgetMouseButtonReleasedEventArgs) {
			if args.Button == ebiten.MouseButtonLeft {
				w.cursorPainting = false
			}
		}),
	)

	return w
}

func (g *GameWidget) GetWidget() *widget.Widget {
	return g.Widget
}

func (g *GameWidget) PreferredSize() (int, int) {
	return int(g.Engine.Sandbox.Width()), int(g.Engine.Sandbox.Height())
}

func (g *GameWidget) Render(screen *ebiten.Image, def widget.DeferredRenderFunc) {
	g.Widget.Render(screen, def)
	g.handleInput()

	screen.Fill(resources.GameBackgroundColor)
	img := g.Engine.Render()

	// Calculate the scaling factor.
	screenSize := g.Widget.Rect.Size()
	imgSize := img.Bounds().Size()

	maxScaleX := float64(screenSize.X) / float64(imgSize.X)
	maxScaleY := float64(screenSize.Y) / float64(imgSize.Y)
	g.renderScaleFactor = math.Min(maxScaleX, maxScaleY)

	// Calculate the X and Y offsets.
	renderWidth := int(g.renderScaleFactor * float64(imgSize.X))
	renderHeight := int(g.renderScaleFactor * float64(imgSize.X))
	g.renderX = float64(screenSize.X-renderWidth) / 2.0
	g.renderY = float64(screenSize.Y-renderHeight) / 2.0

	// Draw the game screen to the UI.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(g.renderScaleFactor, g.renderScaleFactor)
	op.GeoM.Translate(g.renderX, g.renderY)
	screen.DrawImage(img, op)
}

// handleInput paints a grain of sand under the cursor.
func (g *GameWidget) handleInput() {
	if g.cursorEntered && g.cursorPainting {
		x, y := input.CursorPosition()
		paintAtX, paintAtY := g.cursorToCoords(x, y)

		// Set the sand grain at the mouse coordinates.
		grain := g.Engine.Sandbox.GrainAt(paintAtX, paintAtY)
		if grain.Grain != nil {
			grain.SetKind(g.Engine.Config.SelectedSand)
		}
	}
}

// cursorToCoords converts cursor coordinates to game coordinates.
func (g *GameWidget) cursorToCoords(x, y int) (int, int) {
	return int((float64(x-g.Rect.Min.X) - g.renderX) / g.renderScaleFactor),
		int((float64(y-g.Rect.Min.Y) - g.renderY) / g.renderScaleFactor)
}
