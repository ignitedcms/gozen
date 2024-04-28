package dbconvertor

import (
	"fmt"
	"strings"
)

// DataType represents a data type for a column in a database table
type DataType struct {
	Name          string
	SQLiteType    string
	MySQLType     string
	PostgresType  string
	SQLServerType string
}

// Define the mapping of data types for different databases
var dataTypes = []DataType{
	{"integer", "INTEGER", "INT", "INT", "INT"},
	{"float", "REAL", "FLOAT", "FLOAT", "FLOAT"},
	{"string", "TEXT", "VARCHAR(255)", "VARCHAR(255)", "NVARCHAR(255)"},
	{"text", "TEXT", "TEXT", "TEXT", "NVARCHAR(MAX)"}, // Added 'text' data type
	{"boolean", "INTEGER", "TINYINT(1)", "BOOLEAN", "BIT"},
	{"date", "TEXT", "DATE", "DATE", "DATE"},
	{"datetime", "TEXT", "DATETIME", "TIMESTAMP", "DATETIME"},
	{"time", "TEXT", "TIME", "TIME", "TIME"},
	{"timestamp", "TEXT", "TIMESTAMP", "TIMESTAMP", "DATETIME"},
}

// CreateTableSyntax generates the CREATE TABLE syntax for the given database type
func CreateTableSyntax(dbType string, tableName string, columns []string) string {
	var createTableSyntax strings.Builder
	dbType = strings.ToLower(dbType)

	createTableSyntax.WriteString(fmt.Sprintf("CREATE TABLE %s (", tableName))

	// Add default primary key column
	switch dbType {
	case "sqlite":
		createTableSyntax.WriteString("id INTEGER PRIMARY KEY AUTOINCREMENT, ")
	case "mysql":
		createTableSyntax.WriteString("id INT AUTO_INCREMENT PRIMARY KEY, ")
	case "postgres", "pgsql":
		createTableSyntax.WriteString("id SERIAL PRIMARY KEY, ")
	case "sqlserver", "sqlsvr":
		createTableSyntax.WriteString("id INT IDENTITY(1,1) PRIMARY KEY, ")
	default:
		return "Invalid database type"
	}

	for i, column := range columns {
		parts := strings.Split(column, ":")
		columnName := parts[0]
		dataTypeName := parts[1]

		var columnType string
		switch dbType {
		case "sqlite":
			columnType = getDataType(dataTypeName, "SQLiteType")
		case "mysql":
			columnType = getDataType(dataTypeName, "MySQLType")
		case "postgres", "pgsql":
			columnType = getDataType(dataTypeName, "PostgresType")
		case "sqlserver", "sqlsvr":
			columnType = getDataType(dataTypeName, "SQLServerType")
		default:
			return "Invalid database type"
		}

		createTableSyntax.WriteString(fmt.Sprintf("%s %s", columnName, columnType))

		if i < len(columns)-1 {
			createTableSyntax.WriteString(", ")
		}
	}

	createTableSyntax.WriteString(");")

	return createTableSyntax.String()
}

// getDataType retrieves the data type for the given name and database type
func getDataType(name, dbType string) string {
	for _, dt := range dataTypes {
		if strings.ToLower(dt.Name) == strings.ToLower(name) {
			switch dbType {
			case "SQLiteType":
				return dt.SQLiteType
			case "MySQLType":
				return dt.MySQLType
			case "PostgresType":
				return dt.PostgresType
			case "SQLServerType":
				return dt.SQLServerType
			}
		}
	}
	return "UNKNOWN"
}
