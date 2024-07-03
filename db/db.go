/*
|---------------------------------------------------------------
| Database setup
|---------------------------------------------------------------
|
| Added support for the main four db drivers
| MySQL, SQLite, MsSQl, Postgres
| If necessary load db setup
|
| @author: IgnitedCMS
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

	fmt.Println("Using mysql")
}

func loadPostgres() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

   connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

    // Open the connection
	 var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    //defer db.Close()

    fmt.Println("Using Postgres")
}

func loadSqlsvr() {

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

   connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",dbHost,dbUser,dbPassword,dbPort,dbName)

   // Open the connection
	var err error
	DB, err = sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}


    fmt.Println("Using Sqlsvr")
}

