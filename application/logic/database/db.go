package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MyDb struct {
	db *sql.DB
}

func New() *MyDb {
	m := MyDb{db: nil}
	m.openDatabaseConnection()

	return &m
}

// Opens database connection
func (s *MyDb) openDatabaseConnection() {
	var err error

	s.db, err = sql.Open("mysql", "dev:password@/localdb")
	if err != nil {
		log.Fatal(err)
	}

	err = s.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// closes database connection
func (s *MyDb) closeDatabaseConnection() {
	s.db.Close()
}

func (s *MyDb) UsernameExists(username string) bool {
	var user string
	err := s.db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func (s *MyDb) GetStoredPassword(username string) []byte {
	var databasePassword string
	s.db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&databasePassword)
	return []byte(databasePassword)
}

func (s *MyDb) CreateUser(username string, hashedPassword []byte) error {
	_, err := s.db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
	return err
}
