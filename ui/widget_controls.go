package ui

import (
	"image/color"

	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/resources"
)

// ControlsWidget is a widget.Container that contains all the game controls.
type ControlsWidget struct {
	*widget.Container

	Engine    *engine.Sandgine
	container *widget.Container
}

func NewControlsWidget(engine *engine.Sandgine) *ControlsWidget {
	Container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(resources.UiBackgroundColor)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Stretch([]bool{true, true}, []bool{false, true}),
				widget.GridLayoutOpts.Spacing(1, 1))))

	w := &ControlsWidget{
		Engine:    engine,
		Container: Container,
	}

	Container.AddChild(widget.NewLabel(widget.LabelOpts.Text("Hi", resources.Font, &widget.LabelColor{
		Idle:     color.RGBA{R: 255, A: 255},
		Disabled: color.RGBA{},
	})))

	Container.AddChild(NewSelectSandPanel(engine))

	return w
}

func (widget *ControlsWidget) PreferredSize() (int, int) {
	return 150, 0
}
