package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shubhamgupta2956/zeronet/server/pkg/handlers/helpers"
	"github.com/shubhamgupta2956/zeronet/server/pkg/sessions"

	"github.com/shubhamgupta2956/zeronet/server/pkg/players"
)

// PlayResponse object for sending move responses
type PlayResponse struct {
	Move       int       `json:"move"`
	Symbol     string    `json:"symbol`
	Board      [9]string `json:"board"`
	RecoMove   int       `json:"recoMove"`
	RecoSymbol string    `json:"recoSymbol`
}

// PlayHuman handler for returning moves from human
func PlayHuman(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	token := r.Context().Value("token").(string)
	session, ok := db.Store[token]
	if ok == false {
		panic("Session error")
	}
	playerPos := helpers.ConvertStringToInt(r.PostFormValue("player"))
	var player, alterPlayer *players.Player
	if playerPos == 1 {
		player = session.Player2 // If we receive response from player1 then player2 should play
		alterPlayer = session.Player1
	} else if playerPos == 2 {
		player = session.Player1 // If we receive response from player2 then player1 should play
		alterPlayer = session.Player2
	} else {
		panic("Invalid player")
	}
	replyMove, replySymbol := player.Play(session.Board)
	session.Update(replyMove, replySymbol)
	recoMove, recoSymbol := alterPlayer.Play(session.Board)

	reponse := PlayResponse{
		replyMove,
		replySymbol,
		session.Board,
		recoMove,
		recoSymbol,
	}
	res, _ := json.Marshal(reponse)
	w.Write([]byte(res))
}

// PlayAI handler for returning first move from AI
func PlayAI(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	token := r.Context().Value("token").(string)
	session, ok := db.Store[token]
	if ok == false {
		panic("Session error")
	}
	player := session.Player1 // AI will always be first player in this case
	alterPlayer := session.Player2
	replyMove, replySymbol := player.Play(session.Board)
	session.Update(replyMove, replySymbol)
	recoMove, recoSymbol := alterPlayer.Play(session.Board)

	reponse := PlayResponse{
		replyMove,
		replySymbol,
		session.Board,
		recoMove,
		recoSymbol,
	}
	res, _ := json.Marshal(reponse)
	w.Write([]byte(res))
}
