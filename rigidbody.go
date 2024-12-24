package main

import (
	"image/color"
	"math"
)

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
			return detectCircleCircleCollision(body.Position.X, body.Position.Y, other.Position.X, other.Position.Y, shape, otherShape)
		case Rect:
			return detectCircleRectCollision(body.Position.X, body.Position.Y, other.Position.X, other.Position.Y, shape, otherShape)
		}
	case Rect:
		switch otherShape := other.Shape.(type) {
		case Circle:
			return detectCircleRectCollision(body.Position.X, body.Position.Y, other.Position.X, other.Position.Y, otherShape, shape)
		case Rect:
			return detectRectRectCollision(body.Position.X, body.Position.Y, other.Position.X, other.Position.Y, shape, otherShape)
		}
	}
	return false
}

func detectCircleCircleCollision(ax, ay, bx, by float64, a, b Circle) bool {
	dist := math.Sqrt(math.Pow(ax-bx, 2) + math.Pow(ay-by, 2))
	return a.Radius+b.Radius >= dist
}

func detectCircleRectCollision(ax, ay, bx, by float64, a Circle, b Rect) bool {
	return false
}

func detectRectRectCollision(ax, ay, bx, by float64, a, b Rect) bool {
	return false
}
