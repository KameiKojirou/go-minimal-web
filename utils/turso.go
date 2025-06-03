package utils

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/tursodatabase/go-libsql"
)

func TursoDB() (*sql.DB, error) {
	dbName := "file:./local.db"
	db, err := sql.Open("libsql", dbName)
	if err != nil {
    	fmt.Fprintf(os.Stderr, "failed to open db %s", err)
    	os.Exit(1)
  	}
	return db, nil
}

func InitDB() {
	db, err := TursoDB()
	if err != nil {
		os.Exit(1)
	}
	db.Exec("CREATE TABLE IF NOT EXISTS hello (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	db.Exec("CREATE TABLE IF NOT EXISTS example (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	db.Exec("CREATE TABLE IF NOT EXISTS pokemon (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	db.Exec("INSERT INTO hello (name) VALUES ('world')")
	db.Exec("INSERT INTO example (name) VALUES ('world')")
	db.Exec("INSERT INTO pokemon (name) VALUES ('bulbasaur')")
	db.Exec("INSERT INTO pokemon (name) VALUES ('charmander')")
	db.Exec("INSERT INTO pokemon (name) VALUES ('squirtle')")
	defer db.Close()
}