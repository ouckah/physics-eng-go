package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// setup physics engine
	engine := &PhysicsEngine{
		Entities: map[int]*RigidBody{
			0: {Position: Vector2{X: 50, Y: 50}},
			1: {Position: Vector2{X: 100, Y: 100}},
			2: {Position: Vector2{X: 200, Y: 200}},
		},
	}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Physics Engine")
	if err := ebiten.RunGame(
		&Game{
			Engine: engine,
		},
	); err != nil {
		log.Fatal(err)
	}
}
