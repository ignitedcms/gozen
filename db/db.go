/*
|---------------------------------------------------------------
| Database setup
|---------------------------------------------------------------
|
| Test database connection (mySQL only)
| If necessary load db setup
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"log"
	"os"
	"time"
)

var DB *sql.DB // Exported database connection that can be used globally

func InitDB() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// MySQL connection string
	//connString := "user:password@tcp(localhost:3306)/database_name"
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	// Open a connection
	var err error
	DB, err = sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	createTables()
	createTables2()
}

// Auto increment needs underscore!
// This probably needs to be refactored
// to somewhere else
func createTables() {

	createTable := `
   CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTO_INCREMENT,
      name VARCHAR(255),
      email VARCHAR(255),
      password VARCHAR(512),
      created_at DATETIME,
      updated_at DATETIME  
   )
   `
	_, err := DB.Exec(createTable)

	if err != nil {
		log.Println("Error creating table:", err)

	}
}

func createTables2() {

	createTable := `
   CREATE TABLE IF NOT EXISTS migrations (
      id INTEGER PRIMARY KEY AUTO_INCREMENT,
      table_name VARCHAR(255),
      created_at DATETIME,
      updated_at DATETIME  
   )
   `
	_, err := DB.Exec(createTable)

	if err != nil {
		log.Println("Error creating table:", err)

	}
}

func insertMigrations(){
   
	stmt, err := DB.Prepare("INSERT INTO migrations(table_name, created_at, updated_at) VALUES(?, ?, ?)")
	if err != nil {
		//return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec( "users", time.Now(), time.Now())
	if err != nil {
		//return 0, err
	}
   fmt.Print(result)
}
