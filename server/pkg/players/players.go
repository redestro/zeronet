package players

import (
	"fmt"
)

// Player struct
type Player struct {
	token string
	kind  string
	level int
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

//Play implementation for human
func (human *Human) Play(move int, symbol string) int {
	fmt.Println(human.token) // Update board and moves in db and give recommendation
	return move
}

//Play implementation for human
func (ai *AI) Play(board [9]string, symbol string) int {
	fmt.Println(ai.token) // Update board and moves in db and compute next move and return
	return ai.level
}

// Init creates a new player
func Init(token string, kind string, level int) *Player {
	if kind == "AI" && level == 0 {
		panic("Wrong level for AI")
	}
	return &Player{
		token,
		kind,
		level,
	}
}
