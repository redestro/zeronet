package players

import (
	"github.com/redestro/zeronet/server/pkg/games"
)

// Player struct
type Player struct {
	token  string `json:"token"`
	kind   string `json: "kind"`
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
func Init(token string, kind string, symbol string) *Player {
	return &Player{
		token,
		kind,
		symbol,
	}
}
