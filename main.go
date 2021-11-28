package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/eth-p/hackathon-2021-fallhacks/engine"
	"github.com/eth-p/hackathon-2021-fallhacks/sands"
)

const SANDBOX_WIDTH = 100
const SANDBOX_HEIGHT = 100

func main() {
	sandbox := engine.NewSandbox(SANDBOX_WIDTH, SANDBOX_HEIGHT, []engine.Sand{
		sands.AIR,
		sands.SAND,
		sands.WATER,
	})

	eng := engine.NewSandgine(&sandbox, engine.Config{
		UpdateInterval: 0,
	})

	ebiten.SetWindowSize(900, 600)
	ebiten.SetWindowTitle("Fall Sands")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(NewApp(eng)); err != nil {
		log.Fatal(err)
	}
}
