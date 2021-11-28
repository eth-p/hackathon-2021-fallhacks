package ui

import (
	"strings"

	"github.com/blizzy78/ebitenui/widget"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/resources"
)

type SelectSandButton struct {
	*widget.LabeledCheckbox

	Engine *engine.Sandgine
	Sand   *engine.Sand
}

func NewSelectSandButton(engine *engine.Sandgine, sand *engine.Sand) *SelectSandButton {
	box := widget.NewLabeledCheckbox(
		widget.LabeledCheckboxOpts.LabelOpts(
			widget.LabelOpts.Text(strings.ToUpper(sand.Name), resources.Font, resources.SelectSandLabel)),
		widget.LabeledCheckboxOpts.CheckboxOpts(
			widget.CheckboxOpts.ButtonOpts(widget.ButtonOpts.Image(resources.SelectSandCheckboxBtn)),
			widget.CheckboxOpts.Image(resources.SelectSandCheckbox)),
	)

	w := &SelectSandButton{
		LabeledCheckbox: box,
		Engine:          engine,
		Sand:            sand,
	}

	return w
}
