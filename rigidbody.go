package main

import "image/color"

/*
	A RigidBody knows its current velocity and position. See the RigidBody methods getPosition(), getAngle(), getVelocity().
	A RigidBody can calculate its translational and rotational energy and momentum.
	See translationalEnergy(), rotationalEnergy(), momentum().

	@see https://www.myphysicslab.com/develop/docs/Engine2D.html#themathbehindthephysicsengine
*/

type RigidBody struct {
	Velocity, Position Vector2
	Shape              Shape
	Color              color.Color

	UseGravity bool
	IsFrozen   bool
}

func (body RigidBody) GetPosition() Vector2 {
	return body.Position
}

func (body RigidBody) GetVelocity() Vector2 {
	return body.Velocity
}

func (body RigidBody) CollidesWith(other *RigidBody) bool {
	switch shape := body.Shape.(type) {
	case Circle:
		switch otherShape := other.Shape.(type) {
		case Circle:
			distance := body.Position.Distance(other.Position)
			return detectCircleCircleCollision(shape, otherShape, distance)
		case Rectangle:
			// TODO
			return false
		}
	case Rectangle:
		// switch otherShape := other.Shape.(type) {
		// case Circle:
		// 	// TODO
		// 	return false
		// case Rectangle:
		// 	// TODO
		// 	return false
		// }
	}
	return false
}

func detectCircleCircleCollision(a Circle, b Circle, dist float64) bool {
	return a.Radius+b.Radius >= dist
}
