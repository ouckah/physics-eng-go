package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// setup physics engine
	engine := &PhysicsEngine{
		Entities: map[int]*RigidBody{
			0: {
				Position: Vector2{X: 990, Y: 50},
				Velocity: Vector2{X: -2, Y: -2},
				Shape:    Circle{Radius: 20},
				Color:    color.White,

				UseGravity: true,
				IsFrozen:   false,
			},
			1: {
				Position: Vector2{X: 660, Y: 50},
				Velocity: Vector2{X: 2, Y: -2},
				Shape:    Circle{Radius: 20},
				Color:    color.White,

				UseGravity: true,
				IsFrozen:   false,
			},

			2: {
				Position: Vector2{X: 750, Y: 400},
				Shape:    Circle{Radius: 40},
				Color:    color.White,

				UseGravity: false,
				IsFrozen:   true,
			},
			3: {
				Position: Vector2{X: 840, Y: 400},
				Shape:    Circle{Radius: 40},
				Color:    color.White,

				UseGravity: false,
				IsFrozen:   true,
			},
			4: {
				Position: Vector2{X: 930, Y: 400},
				Shape:    Circle{Radius: 40},
				Color:    color.White,

				UseGravity: false,
				IsFrozen:   true,
			},
			5: {
				Position: Vector2{X: 1020, Y: 400},
				Shape:    Circle{Radius: 40},
				Color:    color.White,

				UseGravity: false,
				IsFrozen:   true,
			},
			6: {
				Position: Vector2{X: 1110, Y: 400},
				Shape:    Circle{Radius: 40},
				Color:    color.White,

				UseGravity: false,
				IsFrozen:   true,
			},
		},
	}

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
