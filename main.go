package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// setup physics engine
	engine := &PhysicsEngine{}

	engine.StartSimulation(BallSimulation{})

	// setup game engine
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Physics Engine")
	if err := ebiten.RunGame(
		&Game{
			Engine: engine,
		},
	); err != nil {
		log.Fatal(err)
	}
}
