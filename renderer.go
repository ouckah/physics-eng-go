package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Engine *PhysicsEngine
}

func (g *Game) Update() error {
	g.Engine.Gravity(Vector2{X: 0, Y: 0.04})
	g.Engine.Update(1)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entity := range g.Engine.Entities {
		switch shape := entity.Shape.(type) {
		case Circle:
			vector.DrawFilledCircle(screen, float32(entity.Position.X), float32(entity.Position.Y), float32(shape.Radius), entity.Color, true)
		case Rectangle:
			vector.DrawFilledRect(screen, float32(entity.Position.X), float32(entity.Position.Y), float32(shape.Width), float32(shape.Height), entity.Color, true)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
