package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// StructField represents a field in a struct
type StructField struct {
	Name string
	Type string
}

// GenerateCRUD generates CRUD operations for a given struct
func GenerateCRUD(structName string, table string, fields []StructField) string {
	var builder strings.Builder

	// Generate package and imports
	builder.WriteString(fmt.Sprintf("package %s\n\n", table))
	builder.WriteString("import (\n")
	builder.WriteString("\t\"fibs/db\"\n")
	builder.WriteString("\t\"time\"\n")
	builder.WriteString(")\n\n")

	// Generate struct type
	builder.WriteString(fmt.Sprintf("// %s represents a %s in the system\n", structName, structName))
	builder.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("\t%s %s\n", capitalize(field.Name), field.Type))
	}
	builder.WriteString("}\n\n")

	// Generate Insert function
	builder.WriteString(fmt.Sprintf("// Insert inserts a new %s into the database\n", structName))
	builder.WriteString(fmt.Sprintf("func Create("))
	firstField := true
	for _, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			if !firstField {
				builder.WriteString(", ")
			} else {
				firstField = false
			}
			builder.WriteString(fmt.Sprintf("%s %s", strings.ToLower(field.Name), field.Type))
		}
	}
	builder.WriteString(") (int64, error) {\n")
	builder.WriteString("\tstmt, err := db.DB.Prepare(\"INSERT INTO " + table + "(")
	var insertFields []string
	for _, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			insertFields = append(insertFields, field.Name)
		}
	}
	builder.WriteString(strings.Join(insertFields, ", "))
	builder.WriteString(", created_at, updated_at) VALUES(")
	for i := range insertFields {
		builder.WriteString("?")
		if i < len(insertFields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(", ?, ?)\")") // Add created_at and updated_at
	builder.WriteString("\n\tif err != nil {\n")
	builder.WriteString("\t\treturn 0, err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\tdefer stmt.Close()\n")
	builder.WriteString("\tresult, err := stmt.Exec(")
	for i, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			builder.WriteString(strings.ToLower(field.Name))
			if i < len(fields)-1 && (fields[i+1].Name != "created_at" && fields[i+1].Name != "updated_at") {
				builder.WriteString(", ")
			}
		}
	}
	builder.WriteString(", time.Now(), time.Now())") // Add created_at and updated_at
	builder.WriteString("\n\tif err != nil {\n")
	builder.WriteString("\t\treturn 0, err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\tlastInsertID, err := result.LastInsertId()\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn 0, err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn lastInsertID, nil\n")
	builder.WriteString("}\n\n")

	// Generate Update function
	builder.WriteString(fmt.Sprintf("// Update updates an existing %s in the database\n", structName))
	builder.WriteString(fmt.Sprintf("func Update(id string, "))
	firstField = true
	for _, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			if !firstField {
				builder.WriteString(", ")
			} else {
				firstField = false
			}
			builder.WriteString(fmt.Sprintf("%s %s", strings.ToLower(field.Name), field.Type))
		}
	}
	builder.WriteString(") error {\n")
	builder.WriteString("\tstmt, err := db.DB.Prepare(\"UPDATE " + table + " SET ")
	var updateFields []string
	for _, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			updateFields = append(updateFields, fmt.Sprintf("%s = ?", field.Name))
		}
	}
	builder.WriteString(strings.Join(updateFields, ", "))
	builder.WriteString(", updated_at = ? WHERE id = ?\")\n") // Add updated_at
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\tdefer stmt.Close()\n")
	builder.WriteString("\t_, err = stmt.Exec(")
	for i, field := range fields {
		if field.Name != "id" && field.Name != "created_at" && field.Name != "updated_at" {
			builder.WriteString(strings.ToLower(field.Name))
			if i < len(fields)-1 && (fields[i+1].Name != "created_at" && fields[i+1].Name != "updated_at") {
				builder.WriteString(", ")
			}
		}
	}
	builder.WriteString(", time.Now(), id)\n") // Add updated_at
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn nil\n")
	builder.WriteString("}\n\n")

	// Generate Delete function
	builder.WriteString(fmt.Sprintf("// Delete deletes an existing %s from the database\n", structName))
	builder.WriteString(fmt.Sprintf("func Delete( id string) error {\n"))
	builder.WriteString("\tstmt, err := db.DB.Prepare(\"DELETE FROM " + table + " WHERE id = ?\")\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\tdefer stmt.Close()\n")
	builder.WriteString("\t_, err = stmt.Exec(id)\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn nil\n")
	builder.WriteString("}\n\n")

	// Generate All function
	builder.WriteString(fmt.Sprintf("// All returns all %ss from the database\n", structName))
	builder.WriteString(fmt.Sprintf("func All() ([]%s, error) {\n", structName))
	builder.WriteString("\trows, err := db.DB.Query(\"SELECT ")
	for i, field := range fields {
		builder.WriteString(field.Name)
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(" FROM " + table + "\")\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn nil, err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\tdefer rows.Close()\n")
	builder.WriteString("\tvar result []")
	builder.WriteString(structName)
	builder.WriteString("\n")
	builder.WriteString("\tfor rows.Next() {\n")
	builder.WriteString("\t\tvar u ")
	builder.WriteString(structName)
	builder.WriteString("\n")
	builder.WriteString("\t\terr := rows.Scan(")
	for i, field := range fields {
		builder.WriteString("&u.")
		builder.WriteString(capitalize(field.Name))
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")\n")
	builder.WriteString("\t\tif err != nil {\n")
	builder.WriteString("\t\t\treturn nil, err\n")
	builder.WriteString("\t\t}\n")
	builder.WriteString("\t\tresult = append(result, u)\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn result, nil\n")
	builder.WriteString("}\n\n")

	// Generate Read function
	builder.WriteString(fmt.Sprintf("// Read%s reads a single %s from the database by its ID\n", structName, structName))
	builder.WriteString(fmt.Sprintf("func Read(id string) (*%s, error) {\n", structName))
	builder.WriteString("\tvar result " + structName + "\n")
	builder.WriteString("\terr := db.DB.QueryRow(\"SELECT ")
	for i, field := range fields {
		builder.WriteString(field.Name)
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(" FROM " + table + " WHERE id = ?\", id).Scan(")
	for i, field := range fields {
		builder.WriteString("&result." + capitalize(field.Name))
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\treturn nil, err\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\treturn &result, nil\n")
	builder.WriteString("}\n\n")

	return builder.String()
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func main() {
	// Define the struct name and fields
	table := "users"
	structName := "User"
	fields := []StructField{
		{Name: "id", Type: "string"},
		{Name: "name", Type: "string"},
		{Name: "email", Type: "string"},
		{Name: "password", Type: "string"},
		{Name: "created_at", Type: "string"},
		{Name: "updated_at", Type: "string"},
	}

	// Generate the CRUD operations code
	generatedCode := GenerateCRUD(structName, table, fields)

	// Write the generated code to a file
	fileName := "models/" + table + "/" + table + ".go"
	dir := filepath.Dir(fileName)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(generatedCode)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("CRUD operations code generated successfully in", fileName)
}
