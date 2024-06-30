package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbFilePath string = "database/task_db.db"
var setupDBScript string = "database/scripts/setup.sql"

// InitDB initializes the database
func InitDB() *sql.DB {
	// check if the db file exists
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		fmt.Println("Initializing New Database, no old database was found.")
		return loadDB(true)
	} else {
		fmt.Println("Database already present. Skipping DB initialization.")
		return loadDB(false)
	}
}

func loadDB(setupDB bool) *sql.DB {
	// load the database file
	database, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatalf("Error Occurred while loading the database %s\n", err)
	} else {
		if setupDB {
			fmt.Println("Setting up Database ...")
			// load the setup script
			setupScript, err := os.ReadFile(setupDBScript)
			if err != nil {
				log.Fatalf("Failed to setup database: %s\n", err)
			}
			// execute the sql script
			if _, err := database.Exec(string(setupScript)); err != nil {
				log.Fatalf("Failed while setting up database: %s\n", err)
			}
		}
	}
	defer database.Close()
	return database
}

// GetDbConnection returns a connection to the database
func GetDbConnection() *sql.DB {
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatalf("Error Occurred while loading the database %s\n", err)
	}
	return db
}
