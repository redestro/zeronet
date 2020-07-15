package players

import (
	"fmt"
)

// Player struct
type Player struct {
	token  string `json:"token"`
	kind   string `json: "kind"`
	level  int    `json:"level"`
	symbol string `json:"symbol"`
}

// Play interface for player
type Play interface {
	Play()
}

//Play implementation for human
func (player *Player) Play(board [9]string) (int, string) {
	fmt.Println(player.token) // PLay as AI or give recommendation
	return 2, player.symbol
}

// Init creates a new player
func Init(token string, kind string, level int, symbol string) *Player {
	if kind == "AI" && level == 0 {
		panic("Wrong level for AI")
	}
	return &Player{
		token,
		kind,
		level,
		symbol,
	}
}
