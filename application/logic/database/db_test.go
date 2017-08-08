package database

import (
	"database/sql"
	"testing"
)

func TestOpenDatabaseConnection(t *testing.T) {
	mydb := MyDb{db: nil}

	mydb.openDatabaseConnection()

	if mydb.db.Ping() != nil {
		t.Error("Could not connect to the database")
	}
}

func TestCloseDatabaseConnection(t *testing.T) {
	mydb := MyDb{db: nil}
	var err error
	mydb.db, err = sql.Open("mysql", "dev:password@/localdb")

	mydb.closeDatabaseConnection()
	err = mydb.db.Ping()

	if err == nil {
		t.Error("Database connection didn't close")
	}
}
