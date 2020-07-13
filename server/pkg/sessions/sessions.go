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
	player1 *players.Player
	player2 *players.Player
}

// Create makes a new Session entry in the db
func Create(db *Database, mode string, level1 int, level2 int) *Session {
	token := utils.Token(time.Now().Format(time.RFC850))
	emptyMoves := [9]int{0}
	emptyBoard := [9]string{""}
	switch mode {
	case "human":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "Human", 4),
			players.Init(token, "AI", level1),
		}
		db.AddSession(token, session)
		return session
	case "ai":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "AI", level1),
			players.Init(token, "AI", level2),
		}
		db.AddSession(token, session)
		return session
	case "humanRev":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "AI", level1),
			players.Init(token, "Human", 4),
		}
		db.AddSession(token, session)
		return session
	case "human2":
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "Human", 4),
			players.Init(token, "Human", 4),
		}
		db.AddSession(token, session)
		return session
	default:
		session := &Session{
			token,
			emptyMoves,
			emptyBoard,
			players.Init(token, "Human", 4),
			players.Init(token, "AI", level1),
		}
		db.AddSession(token, session)
		return session
	}
}
