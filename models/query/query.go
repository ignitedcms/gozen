package query

import (
	"gozen/db"
	//"time"
	"fmt"
	"github.com/google/uuid"
)

func randomFilename() string {
	return uuid.New().String()
}

type TableInfo struct {
	Table struct {
		Name string
		UUID string
	}
	Columns []struct {
		Name string
		UUID string
	}
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
		tableInfo := TableInfo{}
		tableInfo.Table.Name = table
		tableInfo.Table.UUID = uuid.New().String()

		columns := make([]struct {
			Name string
			UUID string
		}, 0)

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
			columnInfo := struct {
				Name string
				UUID string
			}{
				Name: column,
				UUID: uuid.New().String(),
			}
			columns = append(columns, columnInfo)
		}

		tableInfo.Columns = columns
		tableInfos = append(tableInfos, tableInfo)
	}

	return tableInfos, nil
}
