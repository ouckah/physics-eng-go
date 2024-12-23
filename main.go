package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// setup physics engine
	engine := &PhysicsEngine{
		Entities: map[int]*RigidBody{
			0: {Position: Vector2{X: 50, Y: 50}, Shape: Circle{Radius: 20}},
			1: {Position: Vector2{X: 100, Y: 100}, Shape: Circle{Radius: 20}},
			2: {Position: Vector2{X: 200, Y: 200}, Shape: Circle{Radius: 20}},
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
