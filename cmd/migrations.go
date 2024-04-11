package main

import (
	"gozen/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	//we must re-establish a db connection
	//when running this file by itself
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	createTables()
}

func createTables() {

	createTable := `
   CREATE TABLE IF NOT EXISTS someothertable (
      id INTEGER PRIMARY KEY AUTO_INCREMENT,
      name VARCHAR(255),
      email VARCHAR(255),
      password VARCHAR(512),
      created_at DATETIME,
      updated_at DATETIME  
   )
   `
	_, err := db.DB.Exec(createTable)

	if err != nil {
		log.Println("Error creating table:", err)

	}
}

/*
import "strings""
statements := strings.Split(createTable, ";")

for _, statement := range statements {
    if strings.TrimSpace(statement) != "" {
        _, err := db.DB.Exec(statement)
        if err != nil {
            fmt.Println("Error creating table:", err)
            return
        }
    }
}*/
