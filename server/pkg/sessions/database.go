package sessions

// Database object to store sessions in memory
type Database struct {
	Store map[string]*Session
}

// db in memory
var db Database

// InitDB initializes database
func InitDB() *Database {
	db.Store = make(map[string]*Session)
	return &db
}

// GetDB returns database
func GetDB() *Database {
	return &db
}

// AddSession in store
func (db *Database) AddSession(token string, session *Session) {
	db.Store[token] = session
}
