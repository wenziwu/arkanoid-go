package gamesystem

import (
	c "arkanoid/lib/components"
	"arkanoid/lib/ecs"
	w "arkanoid/lib/ecs/world"

	"github.com/hajimehoshi/ebiten"
)

// MoveBallSystem moves ball
func MoveBallSystem(world w.World) {
	ecs.Join(world.Components.Ball, world.Components.StickyBall.Not(), world.Components.Transform).Visit(ecs.Visit(func(index int) {
		ball := world.Components.Ball.Get(index).(*c.Ball)
		ballTransform := world.Components.Transform.Get(index).(*c.Transform)

		ballTransform.Translation.X += ball.Velocity * ball.Direction.X / ebiten.DefaultTPS
		ballTransform.Translation.Y += ball.Velocity * ball.Direction.Y / ebiten.DefaultTPS
	}))
}