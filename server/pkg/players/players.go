package players

import (
	"github.com/shubhamgupta2956/zeronet/server/pkg/games"
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

func opponentSymbol(symbol string) string {
	if symbol == "X" {
		return "O"
	} else {
		return "X"
	}
}

//Play implementation for human
func (player *Player) Play(board [9]string) (int, string) {
	opponentSymbol := opponentSymbol(player.symbol)
	move := games.ComputeMove(board, player.symbol, opponentSymbol)
	return move, player.symbol
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
