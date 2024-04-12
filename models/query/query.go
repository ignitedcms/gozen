package query

import (
	"gozen/db"
	//"time"
	"fmt"
)

type TableInfo struct {
	Name    string
	Columns []string
}

func GetAll() ([]TableInfo, error) {
	// Get list of all tables in the database
	var tables []string

	rows, err := db.DB.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	// For each table, get the list of column names
	var tableInfos []TableInfo
	for _, table := range tables {
		columns := make([]string, 0)

		rows, err := db.DB.Query(fmt.Sprintf("SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s' AND TABLE_SCHEMA = DATABASE()", table))
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var column string
			if err := rows.Scan(&column); err != nil {
				return nil, err
			}
			columns = append(columns, column)
		}

		tableInfo := TableInfo{
			Name:    table,
			Columns: columns,
		}
		tableInfos = append(tableInfos, tableInfo)
	}

	return tableInfos, nil
}
