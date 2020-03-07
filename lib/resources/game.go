package resources

import (
	"arkanoid/lib/ecs"

	"github.com/ByteArena/box2d"
)

// B2PixelRatio is the number of pixels representing 1 meter in box2D world
const B2PixelRatio = 50

// BlockCollisionEvent is triggered when a block collision occurs
type BlockCollisionEvent struct {
	Entity ecs.Entity
}

// LifeEvent is triggered when the player lose a life
type LifeEvent struct{}

// ScoreEvent is triggered when the score changes
type ScoreEvent struct {
	Score int
}

// Events contains game events for communication between game systems
type Events struct {
	BlockCollisionEvents []BlockCollisionEvent
	LifeEvents           []LifeEvent
	ScoreEvents          []ScoreEvent
}

// StateEvent is an event for game progression
type StateEvent int

// List of game progression events
const (
	StateEventNone StateEvent = iota
	StateEventGameOver
	StateEventLevelComplete
)

// Game contains game resources
type Game struct {
	CollisionWorld *box2d.B2World
	Events         Events
	StateEvent     StateEvent
	Lifes          int
	Score          int
}

// NewGame creates a new game
func NewGame() *Game {
	return &Game{Lifes: 5}
}