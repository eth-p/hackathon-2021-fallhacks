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

	Engine *engine.Sandgine
}

func NewSelectSandPanel(engine *engine.Sandgine) *SelectSandPanel {
	Container := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(resources.UiBackgroundColor)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(1),
				widget.GridLayoutOpts.Spacing(1, 1))))

	w := &SelectSandPanel{
		Engine:    engine,
		Container: Container,
	}

	// Add a button for each sand type.
	btns := make([]*widget.Checkbox, 0)
	for _, kind := range engine.Sandbox.Sands {
		btn := NewSelectSandButton(engine, &kind)
		btns = append(btns, btn.LabeledCheckbox.Checkbox())
		Container.AddChild(btn)
	}

	// Turn it into a radio group.
	//_ = widget.NewRadioGroup(widget.RadioGroupOpts.Checkboxes(btns...))

	return w
}
