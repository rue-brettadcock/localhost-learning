package logic

import (
	"log"

	"github.com/rue-brettadcock/localhost-learning/logic/database"
	"golang.org/x/crypto/bcrypt"
)

type Logic struct {
	mydb *database.MyDb
}

func New() *Logic {
	l := Logic{mydb: database.New()}
	return &l
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (l *Logic) Register(username, password string) (string, bool) {
	var hashedPassword []byte
	var err error

	if len(username) <= 1 {
		log.Fatal("Invalid username")
	}

	hasUsername := l.mydb.UsernameExists(username)
	if hasUsername == true {
		return "Server error, unable to create your account.", true
	}

	if hasUsername == false {
		hashedPassword, err = hashPassword(password)
		if err != nil {
			return "Server error, unable to create your account.", true
		}
		err = l.mydb.CreateUser(username, hashedPassword)
		if err != nil {
			return "Server error, unable to create your account.", true
		}
		return "User Created", false
	}
	return "", true
}

func (l *Logic) SignIn(username, password string) bool {

	databasePassword := l.mydb.GetStoredPassword(username)

	err := bcrypt.CompareHashAndPassword(databasePassword, []byte(password))
	if err != nil {
		return true
	}
	return false
}
