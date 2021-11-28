package ui

import (
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/resources"
)

// SelectSandPanel is a widget.Container that contains all the sand selection buttons.
type SelectSandPanel struct {
	*widget.Container

	Engine  *engine.Sandgine
	buttons []*SelectSandButton
}

func (w *SelectSandPanel) SelectSand(id engine.SandID) {
	w.Engine.Config.SelectedSand = id

	for _, btn := range w.buttons {
		btn.GetWidget().Disabled = btn.Kind.ID == id
	}
}

func NewSelectSandPanel(eng *engine.Sandgine) *SelectSandPanel {
	Container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(resources.UiBackgroundColor)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Spacing(1, 1))))

	w := &SelectSandPanel{
		Engine:    eng,
		Container: Container,
	}

	// Add a button for each sand type.
	btns := make([]*SelectSandButton, 0)
	for _, kind := range eng.Sandbox.Sands() {
		btn := newSelectSandButton(eng, kind)
		btn.parent = w

		btns = append(btns, btn)
		Container.AddChild(btn)
	}

	// Turn it into a radio group.
	w.buttons = btns
	return w
}
