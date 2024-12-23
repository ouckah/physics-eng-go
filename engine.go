package main

import "image/color"

type PhysicsEngine struct {
	Entities map[int]*RigidBody
}

func (engine *PhysicsEngine) Update(deltaTime float64) {
	for _, entity := range engine.Entities {
		engine.updateEntityPosition(entity, deltaTime)
		engine.simulateCollisions()
	}
}

func (engine *PhysicsEngine) Gravity(velocity Vector2) {
	for _, entity := range engine.Entities {
		if entity.UseGravity {
			entity.Velocity = entity.Velocity.Add(velocity)
		}
	}
}

func (engine *PhysicsEngine) updateEntityPosition(entity *RigidBody, deltaTime float64) {
	var velocityDeltaTime = entity.Velocity.MultiplyByScalar(deltaTime)

	entity.Position = entity.Position.Add(velocityDeltaTime)
}

func (engine *PhysicsEngine) simulateCollisions() {
	for i, a := range engine.Entities {
		for j := i + 1; j < len(engine.Entities); j++ {
			b := engine.Entities[j]

			// check for collision and handle
			if a.CollidesWith(b) {
				engine.handleCollision(a, b)
			}
		}
	}
}

var (
	Red  = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	Blue = color.RGBA{R: 0, G: 0, B: 255, A: 255}
)

func (engine *PhysicsEngine) handleCollision(a *RigidBody, b *RigidBody) {

	// color changes for visual effects
	a.Color = Red
	b.Color = Blue

	// edge cases: a or b are frozen rigidbodies
	// for now, simply reverse the velocity of the moving object
	if a.IsFrozen {
		b.Velocity = b.Velocity.MultiplyByScalar(-0.8)
		return
	}
	if b.IsFrozen {
		a.Velocity = a.Velocity.MultiplyByScalar(-0.8)
		return
	}

	/*
		Perfectly Elastic Collision Formula:
		-----------------------------------
		For two rigid bodies, A and B:

		v1' = v1 - (2 * m2 / (m1 + m2)) * ((v1 - v2) ⋅ n) * n
		v2' = v2 - (2 * m1 / (m1 + m2)) * ((v2 - v1) ⋅ n) * n

		Where:
		- v1, v2: Initial velocities of A and B
		- v1', v2': Final velocities of A and B
		- m1, m2: Masses of A and B
		- n: Normalized collision normal (vector pointing from A to B)
		- ⋅ : Dot product operator

		Since, for now, masses are always equal, the formula reduces to:

		v1' = v1 - ((v1 - v2) ⋅ n) * n
		v2' = v2 - ((v2 - v1) ⋅ n) * n

		v1' = v2
		v2' = v1

		Therefore, the bodies exchange velocities in this case.

		This formula ensures conservation of momentum and kinetic energy in a perfectly elastic collision.
	*/

	vA, vB := a.Velocity, b.Velocity
	a.Velocity = vB.MultiplyByScalar(0.8)
	b.Velocity = vA.MultiplyByScalar(0.8)
}
