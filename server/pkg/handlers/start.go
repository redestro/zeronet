package handlers

import (
	"net/http"
	"time"

	"github.com/shubhamgupta2956/zeronet/server/pkg/players"

	"github.com/shubhamgupta2956/zeronet/server/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

// StartHuman starts session for human player vs AI
func StartHuman(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	// session := sessions.Create(db, "human")
	// res, _ := json.Marshal(session)
	token := utils.Token(time.Now().Format(time.RFC850))
	human := players.InitHuman(token)
	human.Play()
	w.Write([]byte("human"))
}

// StartAI starts session for AI vs AI
func StartAI(w http.ResponseWriter, r *http.Request, db *mongo.Database) {
	w.Write([]byte("AI"))
}
