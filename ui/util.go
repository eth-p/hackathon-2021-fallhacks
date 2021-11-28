package ui

import (
	"github.com/blizzy78/ebitenui/widget"

	"github.com/eth-p/hackathon-2021-fallhacks/resources"
)

func createHeader(text string) *widget.Label {
	return widget.NewLabel(widget.LabelOpts.Text(text, resources.HeaderTextFont, resources.HeaderTextColor))
}
