package handlers

import (
	"net/http"

	"github.com/shubhamgupta2956/zeronet/server/pkg/handlers/helpers"
	"github.com/shubhamgupta2956/zeronet/server/pkg/sessions"
)

// DeleteSession removes a completed session
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	db := sessions.GetDB()
	token := helpers.GetKeyFromRequestParam(r, "token")
	_, ok := db.Store[token]
	if ok {
		delete(db.Store, token)
		w.Write([]byte("Session removal successful"))
	} else {
		w.Write([]byte("Session does not exist"))
	}
}
