package main

import "image/color"

type PhysicsEngine struct {
	Entities   map[int]*RigidBody
	Simulation Simulation
}

func (engine *PhysicsEngine) Update(deltaTime float64) {
	if engine.Simulation != nil {
		engine.Simulation.Update(engine)
	}

	for _, entity := range engine.Entities {
		engine.updateEntityPosition(entity, deltaTime)
		engine.simulateCollisions()
	}
}

func (engine *PhysicsEngine) StartSimulation(sim Simulation) {
	engine.Entities = map[int]*RigidBody{}
	engine.Simulation = sim

	sim.Setup(engine)
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

	// handle penetration
	handlePenetration(a, b)

	// edge cases: a or b are frozen rigidbodies
	// for now, simply reverse the velocity of the moving object
	if a.IsFrozen {
		n := a.Position.Sub(b.Position).Normalize()
		newBVel := n.MultiplyByScalar(-b.Velocity.Length() * ELASTICITY)
		b.Velocity = newBVel
		return
	}
	if b.IsFrozen {
		n := b.Position.Sub(a.Position).Normalize()
		newAVel := n.MultiplyByScalar(-a.Velocity.Length() * ELASTICITY)
		a.Velocity = newAVel
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
	a.Velocity = vB.MultiplyByScalar(ELASTICITY)
	b.Velocity = vA.MultiplyByScalar(ELASTICITY)
}

func handlePenetration(a, b *RigidBody) {
	switch shape := a.Shape.(type) {
	case Circle:
		switch otherShape := b.Shape.(type) {
		case Circle:
			handleCircleCirclePenetration(a, b, shape, otherShape)
		case Rect:
			handleCircleRectPenetration(a, b, shape, otherShape)
		}
	case Rect:
		switch otherShape := b.Shape.(type) {
		case Circle:
			handleCircleRectPenetration(b, a, otherShape, shape)
		case Rect:
			handleRectRectPenetration(a, b, shape, otherShape)
		}
	}
}

func handleCircleCirclePenetration(a, b *RigidBody, circleA, circleB Circle) {
	n := a.Position.Sub(b.Position)

	penetrationDistance := (circleA.Radius + circleB.Radius) - n.Length() + BUFFER

	if penetrationDistance <= 0 {
		return
	}

	penetrationVector := n.Normalize().MultiplyByScalar(penetrationDistance / 2)

	newAPos, newBPos := a.Position, b.Position
	if !a.IsFrozen && !b.IsFrozen {
		newAPos = a.Position.Add(penetrationVector)
		newBPos = b.Position.Sub(penetrationVector)
	} else if a.IsFrozen {
		newBPos = b.Position.Sub(penetrationVector.MultiplyByScalar(2))
	} else if b.IsFrozen {
		newAPos = a.Position.Add(penetrationVector.MultiplyByScalar(2))
	}

	a.Position, b.Position = newAPos, newBPos
}

func handleCircleRectPenetration(a, b *RigidBody, circleA Circle, rectB Rect) {
	// TODO
	return
}

func handleRectRectPenetration(a, b *RigidBody, rectA, rectB Rect) {
	// TODO
	return
}
