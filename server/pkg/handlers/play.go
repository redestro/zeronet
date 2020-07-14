package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shubhamgupta2956/zeronet/server/pkg/players"

	"github.com/shubhamgupta2956/zeronet/server/pkg/handlers/helpers"
	"github.com/shubhamgupta2956/zeronet/server/pkg/sessions"
)

// PlayResponse object for sending move responses
type PlayResponse struct {
	Move   int       `json:"move"`
	Symbol string    `json:"symbol`
	Board  [9]string `json:"board"`
}

// Play handler for returning moves
func Play(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	token := helpers.GetKeyFromRequestParam(r, "token")
	session := db.Store[token]
	playerPos := r.PostFormValue("player")
	var player *players.Player
	if playerPos == "1" {
		player = session.Player2 // If we receive response from player1 then player2 should play
	} else if playerPos == "2" {
		player = session.Player1 // If we receive response from player2 then player1 should play
	} else {
		panic("Invalid player")
	}
	symbol := r.PostFormValue("symbol")
	move64, err := strconv.ParseInt(r.PostFormValue("move"), 10, 64)
	move := int(move64)
	if err != nil {
		panic(err)
	}
	replyMove, replySymbol := player.Play(move, symbol)
	reponse := PlayResponse{
		replyMove,
		replySymbol,
		session.Board,
	}
	res, _ := json.Marshal(reponse)
	w.Write([]byte(res))
}
