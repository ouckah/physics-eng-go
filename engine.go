package main

type PhysicsEngine struct {
	Entities map[int]*RigidBody
}

func (engine *PhysicsEngine) Update(deltaTime float64) {
	for _, entity := range engine.Entities {
		engine.updateEntityPosition(entity, deltaTime)
	}
}

func (engine *PhysicsEngine) Gravity(velocity Vector2) {
	for _, entity := range engine.Entities {
		entity.Velocity = entity.Velocity.Add(velocity)
	}
}

func (engine *PhysicsEngine) updateEntityPosition(entity *RigidBody, deltaTime float64) {
	var velocityDeltaTime = entity.Velocity.MultiplyByScalar(deltaTime)

	entity.Position = entity.Position.Add(velocityDeltaTime)
}
