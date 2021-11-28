package main

import (
	"fmt"
	"time"

	"github.com/blizzy78/ebitenui"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/ui"
)

type App struct {
	ui     ebitenui.UI
	engine *engine.Sandgine
}

func (game *App) Draw(screen *ebiten.Image) {
	game.ui.Draw(screen)
}

func (game *App) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (game *App) Update() error {
	game.ui.Update()

	updateStart := time.Now()
	err := game.engine.Update()
	updateEnd := time.Now()

	updateTime := updateEnd.Sub(updateStart)
	fmt.Printf("Frame took %v\n", updateTime)
	return err
}

func NewApp(engine *engine.Sandgine) *App {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{true, false}),
			widget.GridLayoutOpts.Spacing(1, 1))))
	container.AddChild(ui.NewGameWidget(engine))

	//container.AddChild(ui.NewControlsWidget(engine))

	return &App{
		engine: engine,
		ui: ebitenui.UI{
			Container: container,
		},
	}
}
