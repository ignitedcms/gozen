package query

import (
	"gozen/db"
	//"time"
	"fmt"
)

func GetAll() {
	// Get list of all tables in the database
   fmt.Print("testing")
	var tables []string
	rows, err := db.DB.Query("SHOW TABLES")
	if err != nil {
		// Handle error
		return
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			// Handle error
			return
		}
		tables = append(tables, table)
	}

	// For each table, get the list of column names
	for _, table := range tables {
		columns := make([]string, 0)
      rows, err := db.DB.Query(fmt.Sprintf("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s' AND TABLE_SCHEMA = DATABASE()", table))
		if err != nil {
			// Handle error
			return
		}
		defer rows.Close()

		for rows.Next() {
			var column string
			if err := rows.Scan(&column); err != nil {
				// Handle error
				return
			}
			columns = append(columns, column)
		}

		fmt.Printf("Table: %s\nColumns: %v\n", table, columns)
	}
}
