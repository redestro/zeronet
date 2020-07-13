package sessions

import (
	"context"
	"time"

	"github.com/shubhamgupta2956/zeronet/server/pkg/players"

	"github.com/shubhamgupta2956/zeronet/server/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Session object
type Session struct {
	token string
	moves [9]int
	board [9]string
}

// HumanSession object
type HumanSession struct {
	Session
	player1 *players.Human
	player2 *players.AI
}

// AISession object
type AISession struct {
	Session
	player1 *players.AI
	player2 *players.AI
}

// RevHumanSession reverses AI and Human roles
type RevHumanSession struct {
	Session
	player1 *players.AI
	player2 *players.Human
}

// Human2Session is a 2 Human session
type Human2Session struct {
	Session
	player1 *players.AI
	player2 *players.Human
}

// Schema to add data in DB
type Schema struct {
	token   string
	moves   [9]int
	board   [9]string
	player1 string
	player2 string
}

// Create makes a new session entry in the db
func Create(db *mongo.Database, player string) Schema {
	collection := db.Collection("session")

	token := utils.Token(time.Now().Format(time.RFC850))
	emptyMoves := [9]int{0}
	emptyBoard := [9]string{""}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{
		"token":   token,
		"moves":   emptyMoves,
		"board":   emptyBoard,
		"player1": player,
		"player2": "ai",
	})
	if err != nil {
		panic(err)
	}
	var result Schema
	collection.FindOne(ctx, bson.M{
		"token": token,
	}).Decode(&result)
	return result
}

// NewHuman creates a new session instance
func NewHuman(token string) *HumanSession {
	session := &HumanSession{
		Session{
			token,
			moves,                        // Moves Array from db
			board,                        // Board Array from db
			players.InitAI(token, level), // AI level from db
		},
		players.InitHuman(token),
	}
	return session
}
