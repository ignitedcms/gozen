package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// Define the flags
	tablePtr := flag.String("table", "comments", "the name of the table")
	structNamePtr := flag.String("struct", "comment", "the name of the struct")

	// Define a slice of StructField pointers
	allFieldsPtr := flag.String("fields", "", "the fields of the struct, comma-separated")

	// Parse the command-line arguments
	flag.Parse()

	// Convert the allFields string to a slice of StructField
	table := *tablePtr
	structName := *structNamePtr
	allFields := parseFields(*allFieldsPtr)

	// Print the values
	fmt.Printf("Table: %s\nStruct: %s\nFields: %v\n", *tablePtr, *structNamePtr, allFields)
}

func parseFields(fieldsStr string) []StructField {
	var fields []StructField

	// Split the string by comma to get each field
	for _, field := range strings.Split(fieldsStr, ",") {
		// Split each field by colon to get the name and type
		parts := strings.Split(field, ":")
		if len(parts) != 2 {
			continue
		}
		fields = append(fields, StructField{
			Name: parts[0],
			Type: parts[1],
		})
	}

	return fields
}

type StructField struct {
	Name string
	Type string
}
