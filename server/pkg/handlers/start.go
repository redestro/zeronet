package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shubhamgupta2956/zeronet/server/pkg/handlers/helpers"

	"github.com/shubhamgupta2956/zeronet/server/pkg/sessions"
)

type response struct {
	Token string    `json:"token"`
	Moves [9]int    `json:"moves"`
	Board [9]string `json:"board"`
}

// StartHuman starts session for human player vs AI
func StartHuman(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	level1 := helpers.GetintKeyFromRequestParam(r, "level1")
	level2 := helpers.GetintKeyFromRequestParam(r, "level2")
	session := sessions.Create(db, "human", level1, level2)
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
	level1 := helpers.GetintKeyFromRequestParam(r, "level1")
	level2 := helpers.GetintKeyFromRequestParam(r, "level2")
	session := sessions.Create(db, "ai", level1, level2)
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
	level1 := helpers.GetintKeyFromRequestParam(r, "level1")
	level2 := helpers.GetintKeyFromRequestParam(r, "level2")
	session := sessions.Create(db, "human", level1, level2)
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
	level1 := helpers.GetintKeyFromRequestParam(r, "level1")
	level2 := helpers.GetintKeyFromRequestParam(r, "level2")
	session := sessions.Create(db, "human", level1, level2)
	response := response{
		session.Token,
		session.Moves,
		session.Board,
	}
	res, _ := json.Marshal(response)
	w.Write([]byte(res))
}