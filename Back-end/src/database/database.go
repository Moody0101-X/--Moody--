package database

import (
	// "fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	
	// "strings"
)

var DATABASE *sql.DB

func GetDb() *sql.DB {
	return DATABASE;
}

// db initializer: Opens the db, then evluates a global conn variable.
func InitializeDb(dbPath string) (error, string) {	
	var err error;

	DATABASE, err = sql.Open("sqlite3", dbPath); if err != nil {
		return err, ""
	}

	return nil, dbPath
}

func get_id(Table string) int {

	var id int;
	
	row, err := DATABASE.Query("select MAX(ID) from " + Table);
	
	defer row.Close()

	if err != nil { return 0 }

	for row.Next() {
		row.Scan(&id);
	}

	return id;
}