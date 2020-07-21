package middlewares

import (
	"context"
	"net/http"

	"github.com/redestro/zeronet/server/pkg/handlers/helpers"
	"github.com/redestro/zeronet/server/pkg/sessions"
)

// Update the current state of the board
func Update(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := sessions.GetDB()
		token := helpers.GetKeyFromRequestParam(r, "token")
		session := db.Store[token]
		symbol := r.PostFormValue("symbol")
		move := helpers.ConvertStringToInt(r.PostFormValue("move"))
		player := helpers.ConvertStringToInt(r.PostFormValue("player"))
		if player != 0 {
			session.Update(move, symbol)
		}
		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
