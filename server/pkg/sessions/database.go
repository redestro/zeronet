package sessions

// Database object to store sessions in memory
type Database struct {
	Store map[string]*Session
}

// DB in memory
var DB Database

// InitDB initializes database
func InitDB() *Database {
	DB.Store = make(map[string]*Session)
	return &DB
}

// GetDB returns database
func GetDB() *Database {
	return &DB
}

// AddSession in store
func (db *Database) AddSession(token string, session *Session) {
	db.Store[token] = session
}
