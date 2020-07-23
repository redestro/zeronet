package sessions

import (
	"time"

	"github.com/redestro/zeronet/server/pkg/players"
	"github.com/redestro/zeronet/server/pkg/utils"
)

// Session object
type Session struct {
	Token     string    `json:"token"`
	Moves     [9]int    `json:"moves"`
	Board     [9]string `json:"board"`
	MoveCount int       `json:"moveCount"`
	Player1   *players.Player
	Player2   *players.Player
}

// Create makes a new Session entry in the db
func Create(db *Database, mode string) *Session {
	token := utils.Token(time.Now().Format(time.RFC850))
	player1Symbol := "O"
	player2Symbol := "X"
	emptyMoves := [9]int{0}
	emptyBoard := [9]string{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
	var player1 *players.Player
	var player2 *players.Player
	var session *Session
	switch mode {
	case "human":
		player1 = players.Init(token, "Human", player1Symbol)
		player2 = players.Init(token, "AI", player2Symbol)
	case "ai":
		player1 = players.Init(token, "AI", player1Symbol)
		player2 = players.Init(token, "AI", player2Symbol)
	case "humanRev":
		player1 = players.Init(token, "AI", player1Symbol)
		player2 = players.Init(token, "Human", player2Symbol)
	case "human2":
		player1 = players.Init(token, "Human", player1Symbol)
		player2 = players.Init(token, "Human", player2Symbol)
	default:
		player1 = players.Init(token, "Human", player1Symbol)
		player2 = players.Init(token, "AI", player2Symbol)
	}
	session = &Session{
		token,
		emptyMoves,
		emptyBoard,
		0,
		player1,
		player2,
	}
	db.AddSession(token, session)
	return session
}

// Update updates session status
func (session *Session) Update(move int, symbol string) {
	moveCount := session.MoveCount
	session.Board[move] = symbol
	session.Moves[moveCount] = move
	session.MoveCount++
}
