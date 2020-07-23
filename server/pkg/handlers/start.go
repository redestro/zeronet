package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redestro/zeronet/server/pkg/sessions"
)

type response struct {
	Token string    `json:"token"`
	Moves [9]int    `json:"moves"`
	Board [9]string `json:"board"`
}

// StartHuman starts session for human player vs AI
func StartHuman(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	session := sessions.Create(db, "human")
	response := response{
		session.Token,
		session.Moves,
		session.Board,
	}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}

// StartAI starts session for AI vs AI
func StartAI(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	session := sessions.Create(db, "ai")
	response := response{
		session.Token,
		session.Moves,
		session.Board,
	}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}

// StartRevHuman starts session for AI vs human player
func StartRevHuman(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	session := sessions.Create(db, "humanRev")
	response := response{
		session.Token,
		session.Moves,
		session.Board,
	}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}

// StartHuman2 starts session for human player vs human player
func StartHuman2(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	session := sessions.Create(db, "human2")
	response := response{
		session.Token,
		session.Moves,
		session.Board,
	}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}
