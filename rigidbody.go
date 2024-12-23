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
}

func (body RigidBody) GetPosition() Vector2 {
	return body.Position
}

func (body RigidBody) GetVelocity() Vector2 {
	return body.Velocity
}
