package main

import (
	"image/color"
	"time"

	"golang.org/x/exp/rand"
)

type Simulation interface {
	Setup(engine *PhysicsEngine)
	Update(engine *PhysicsEngine)
}

type BallSimulation struct{}

func (sim BallSimulation) Setup(engine *PhysicsEngine) {
	engine.Entities = map[int]*RigidBody{
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
	}
}

func (sim BallSimulation) Update(engine *PhysicsEngine) {}

type FountainSimulation struct{}

func (sim FountainSimulation) Setup(engine *PhysicsEngine) {}

func (sim FountainSimulation) Update(engine *PhysicsEngine) {
	rand.Seed(uint64(time.Now().UnixNano()))

	var newID int
	for {
		newID = rand.Int()
		if _, exists := engine.Entities[newID]; !exists {
			break
		}
	}

	// Generate random velocity
	velocityX := rand.Float64()*4 - 2
	velocityY := rand.Float64() * -5

	engine.Entities[newID] = &RigidBody{
		Position:   Vector2{X: 500, Y: 300},
		Velocity:   Vector2{X: velocityX, Y: velocityY},
		Shape:      Circle{Radius: 20},
		Color:      color.White,
		UseGravity: true,
		IsFrozen:   false,
	}
}
