package players

import (
	"fmt"
)

// Player struct
type Player struct {
	token string
}

// Human struct
type Human struct {
	Player
}

// AI struct
type AI struct {
	Player
	level int
}

// Play interface for player
type Play interface {
	Play()
}

// Init interface for player
type Init interface {
	Init()
}

//Play implementation for human
func (human *Human) Play(move int, symbol string) int {
	fmt.Println(human.token) // Update board and moves in db and give recommendation
	return suggestion
}

//Play implementation for human
func (ai *AI) Play(board [9]string, symbol string) int {
	fmt.Println(ai.token) // Update board and moves in db and compute next move and return
	return move
}

// InitHuman implementation for human
func InitHuman(token string) *Human {
	return &Human{
		Player{
			token,
		},
	}
}

// InitAI implementation for AI
func InitAI(token string, level int) *AI {
	return &AI{
		Player{
			token,
		},
		level,
	}
}
