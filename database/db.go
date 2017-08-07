package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type MyDb struct {
	db *sql.DB
}

func New() MyDb {
	m := MyDb{db: nil}
	m.openDatabaseConnection()

	return m
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

// Creates user in the mysql database
func (s *MyDb) SetUser(username, password string) (string, bool) {

	var user string
	err := s.db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "Server error, unable to create your account.", true
		}

		_, err = s.db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			return "Server error, unable to create your account.", true
		}

		return "User created!", false
	case err != nil:
		return "Server error, unable to create your account.", true
	default:
		return "", false
	}
}

// Attemps to log user into the mysql database
func (s *MyDb) LoginUser(username, password string) bool {

	var databaseUsername string
	var databasePassword string

	err := s.db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
	if err != nil {
		return true
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		return true
	}
	return false
}
