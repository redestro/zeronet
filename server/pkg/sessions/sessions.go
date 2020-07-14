package sessions

import (
	"time"

	"github.com/shubhamgupta2956/zeronet/server/pkg/players"
	"github.com/shubhamgupta2956/zeronet/server/pkg/utils"
)

// Session object
type Session struct {
	Token   string    `json:"token"`
	Moves   [9]int    `json:"moves"`
	Board   [9]string `json:"board"`
	Player1 *players.Player
	Player2 *players.Player
}

// Create makes a new Session entry in the db
func Create(db *Database, mode string, level1 int, level2 int) *Session {
	token := utils.Token(time.Now().Format(time.RFC850))
	player1Symbol := "O"
	player2Symbol := "X"
	emptyMoves := [9]int{0}
	emptyBoard := [9]string{""}
	switch mode {
	case "human":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "Human", level1, player1Symbol),
			players.Init(token, "AI", level2, player2Symbol),
		}
		db.AddSession(token, session)
		return session
	case "ai":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "AI", level1, player1Symbol),
			players.Init(token, "AI", level2, player2Symbol),
		}
		db.AddSession(token, session)
		return session
	case "humanRev":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "AI", level1, player1Symbol),
			players.Init(token, "Human", level2, player2Symbol),
		}
		db.AddSession(token, session)
		return session
	case "human2":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "Human", level1, player1Symbol),
			players.Init(token, "Human", level2, player2Symbol),
		}
		db.AddSession(token, session)
		return session
	default:
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "Human", 4, player1Symbol),
			players.Init(token, "AI", level1, player2Symbol),
		}
		db.AddSession(token, session)
		return session
	}
}
