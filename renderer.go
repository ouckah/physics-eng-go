package main

import (
	"image/color"

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
		vector.DrawFilledCircle(screen, float32(entity.Position.X), float32(entity.Position.Y), 20, color.White, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}
