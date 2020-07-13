package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shubhamgupta2956/zeronet/server/pkg/sessions"
)

// StartHuman starts session for human player vs AI
func StartHuman(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	session := sessions.Create(db, "human", 4, 0)
	res, _ := json.Marshal(session)
	w.Write([]byte(res))
}

// StartAI starts session for AI vs AI
func StartAI(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	session := sessions.Create(db, "ai", 4, 4)
	res, _ := json.Marshal(session)
	w.Write([]byte(res))
}
