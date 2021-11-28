package ui

import (
	"strings"

	"github.com/blizzy78/ebitenui/widget"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/resources"
)

type SelectSandButton struct {
	*widget.Button
	parent *SelectSandPanel

	Engine *engine.Sandgine
	Kind   engine.SandWithSandID
}

func newSelectSandButton(engine *engine.Sandgine, kind engine.SandWithSandID) *SelectSandButton {
	w := &SelectSandButton{
		Engine: engine,
		Kind:   kind,
	}

	w.Button = widget.NewButton(
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			w.parent.SelectSand(kind.ID)
		}),

		widget.ButtonOpts.Text(strings.ToUpper(kind.Name), resources.Font, resources.SelectSandButtonText),
		widget.ButtonOpts.Image(resources.SelectSandButtonImage),
	)

	return w
}
