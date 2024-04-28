/*
|---------------------------------------------------------------
| Database setup
|---------------------------------------------------------------
|
| Added support for the main four db drivers
| MySQL, SQLite, MsSQl, Postgres
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
	_ "github.com/denisenkom/go-mssqldb" // Import the microsoft sqlsvr driver
	_ "github.com/go-sql-driver/mysql"   // MySQL driver
	_ "github.com/lib/pq"                // Import the postgres driver
	_ "github.com/mattn/go-sqlite3"      // Import the SQLite3 driver
	"log"
	"os"
	//"time"
)

var DB *sql.DB // Exported database connection that can be used globally

func InitDB() {

	//First let's set the sql driver as specified from the .env file
	dbConnection := os.Getenv("DB_CONNECTION")

	switch dbConnection {
	case "sqlite":
      loadSqlite()
	case "mysql":
      loadMysql()
	case "pgsql":
      loadPostgres()
	case "sqlsvr":
      loadSqlsvr()
	default:
		fmt.Print("Error database driver not recognised")
	}

}

func loadSqlite() {

	var err error
	DB, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	//defer DB.Close()

	// Create the "user" table
	createTableQuery := `

   CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name VARCHAR(255),
      email VARCHAR(255),
      password VARCHAR(512),
      token VARCHAR(512),
      created_at DATETIME,
      updated_at DATETIME  
   )
    `
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Using sqlite")

}

func loadMysql() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// MySQL connection string
	//connString := "user:password@tcp(localhost:3306)/database_name"
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

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

	createTable := `
   CREATE TABLE IF NOT EXISTS users (
      id INTEGER PRIMARY KEY AUTO_INCREMENT,
      name VARCHAR(255),
      email VARCHAR(255),
      password VARCHAR(512),
      token VARCHAR(512),
      created_at DATETIME,
      updated_at DATETIME  
   )
   `
	_, err = DB.Exec(createTable)

	if err != nil {
		log.Println("Error creating table:", err)
		return

	}
	fmt.Println("Using mysql")
}

func loadPostgres() {

}

func loadSqlsvr() {

}

/*
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

func insertMigrations() {
	// Check if the row exists
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM migrations WHERE table_name = ?", "users").Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	// If the row doesn't exist, execute the INSERT statement
	if count == 0 {
		// Prepare the INSERT statement
		stmt, err := DB.Prepare("INSERT INTO migrations(table_name, created_at, updated_at) VALUES(?, ?, ?)")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()

		// Execute the INSERT statement
		result, err := stmt.Exec("users", time.Now(), time.Now())
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Row inserted successfully:", result)
	} else {
		fmt.Println("Row already exists!")
	}
}
*/
