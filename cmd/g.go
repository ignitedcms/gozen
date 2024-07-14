/*
|---------------------------------------------------------------
| Gozen command line scaffolding utility
|---------------------------------------------------------------
|
| Generates crud
|
| Notes, still needs work to support PostgreSQL and MsSQL
| PostgreSQL and MsSQL do NOT work yet!!!
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/
// Example of usage
// go run % students student name:string,email:string
// omit the id,created_at,updated_at as these are generated automatically
package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gozen/db"
	"io/ioutil"
	"log"
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

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Usage: go run your_program.go table struct fields")
		return
	}

	table := os.Args[1]
	structName := os.Args[2]
	fieldsStr := os.Args[3]

	//Here we will write to the routes file

	writeRoutes(table)

	//End

	fields := parseFields(fieldsStr)
	fields2 := parseFields2(fieldsStr) //for migrations

	fmt.Printf("Table: %s\nStruct: %s\nFields: %v\n", table, structName, fields)

	//now we need to dynamically add id and created,updated at
	// Create a new allFields slice with the desired order
	allFields := []StructField{
		{Name: "id", Type: "string"},
	}
	allFields = append(allFields, fields...)
	allFields = append(allFields,
		StructField{Name: "created_at", Type: "string"},
		StructField{Name: "updated_at", Type: "string"},
	)

	allFields2 := []StructField{
		{Name: "id", Type: "integer"},
	}
	allFields2 = append(allFields2, fields2...)
	allFields2 = append(allFields2,
		StructField{Name: "created_at", Type: "datetime"},
		StructField{Name: "updated_at", Type: "datetime"},
	)

	//Lets load the .env file to retrieve our db type
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConnection := os.Getenv("DB_CONNECTION")

	//Now let's pass it into our function

	sql := generateCreateTableSQL(table, allFields2, dbConnection)

	migrateSql(table, sql)

	var builder strings.Builder

	// Write HTML template
	builder.WriteString("{{ template \"_header.html\" }}\n")
	builder.WriteString("<div class=\"gap\"></div>\n")
	builder.WriteString("<div class=\"default-container\">\n")
	builder.WriteString("    <div class=\"row\">\n")
	builder.WriteString("        <div class=\"col\">\n")
	builder.WriteString(fmt.Sprintf("            <a href=\"/%s/create\" class=\"text-primary underline\">add %s</a>\n", table, table))
	builder.WriteString(fmt.Sprintf("            <h3>%s</h3>\n", capitalize(table)))
	builder.WriteString("            <table class=\"\">\n")
	builder.WriteString("                <tr>\n")

	// Write table headers
	for _, field := range allFields {
		builder.WriteString(fmt.Sprintf("                    <th>%s</th>\n", field.Name))
	}
	builder.WriteString("                    <th>action</th>\n")
	builder.WriteString("                </tr>\n")
	builder.WriteString("                {{range .Data}}\n")
	builder.WriteString("                <tr>\n")

	// Write table data
	for _, field := range allFields {
		if field.Name == "id" {
			builder.WriteString(fmt.Sprintf("                    <td>\n                        <a href=\"/%s/update/{{.Id}}\" class=\"underline text-primary\">\n                            {{.Id}}\n                        </a>\n                    </td>\n", table))
		} else {
			builder.WriteString(fmt.Sprintf("                    <td>{{.%s}}</td>\n", capitalize(field.Name)))
		}
	}
	builder.WriteString(fmt.Sprintf("                    <td><a href=\"/%s/delete/{{.Id}}\" class=\"underline text-primary\">delete</a></td>\n", table))
	builder.WriteString("                </tr>\n")
	builder.WriteString("                {{end}}\n")
	builder.WriteString("            </table>\n")
	builder.WriteString("        </div>\n")
	builder.WriteString("    </div>\n")
	builder.WriteString("</div>\n")
	builder.WriteString("{{ template \"_footer.html\" }}\n")

	// Convert string builder to string
	generatedHTML := builder.String()

	// Write generated HTML to file
	fileName := "views/" + table + "-all.html"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(generatedHTML)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("HTML template file generated successfully:", fileName)

	// Code from html2.go
	builder = strings.Builder{}

	// Write HTML template
	builder.WriteString("{{ template \"_header.html\" }}\n")
	builder.WriteString("<div class=\"gap\"></div>\n")
	builder.WriteString("<div class=\"default-container\">\n")

	// Flash Data
	builder.WriteString("{{ if .Data.FlashData }}\n")
	builder.WriteString("    <div class=\"toasts\">\n")
	builder.WriteString("        <toast ref=\"toast\">\n")
	builder.WriteString("            <div class=\"p-5\">\n")
	builder.WriteString("                <div class=\"text-danger dark:text-white\">Error</div>\n")
	builder.WriteString("                <div class=\"text-danger small dark:text-white\">\n")
	builder.WriteString("                    {{ .Data.FlashData }}\n")
	builder.WriteString("                </div>\n")
	builder.WriteString("            </div>\n")
	builder.WriteString("        </toast>\n")
	builder.WriteString("    </div>\n")
	builder.WriteString("{{ end }}\n")

	// Form
	builder.WriteString("    <div class=\"row\">\n")
	builder.WriteString("        <div class=\"col\">\n")
	builder.WriteString(fmt.Sprintf("            <a href=\"/%s\" class=\"text-primary underline\">%s</a>\n", table, table))
	builder.WriteString(fmt.Sprintf("            <h3>Create %s</h3>\n", capitalize(table)))
	builder.WriteString(fmt.Sprintf("            <form method=\"POST\" action=\"/%s/create\">\n", table))
	builder.WriteString("                {{.Csrf}}\n")

	// Form fields
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("                <div class=\"form-group\">\n"))
		builder.WriteString(fmt.Sprintf("                    <label for=\"%s\">%s</label>\n", field.Name, capitalize(field.Name)))
		builder.WriteString(fmt.Sprintf("                    <input class=\"form-control\" name=\"%s\" value=\"{{.Data.PostData.%s}}\" placeholder=\"\" />\n", field.Name, field.Name))
		builder.WriteString(fmt.Sprintf("                    <div class=\"small text-danger\">{{ .Data.PostDataErrors.%s }}</div>\n", field.Name))
		builder.WriteString(fmt.Sprintf("                </div>\n"))
	}

	// Submit button
	builder.WriteString("                <div class=\"form-group\">\n")
	builder.WriteString("                    <button>Save</button>\n")
	builder.WriteString("                </div>\n")

	builder.WriteString("            </form>\n")
	builder.WriteString("        </div>\n")
	builder.WriteString("    </div>\n")

	builder.WriteString("</div>\n")
	builder.WriteString("{{ template \"_footer.html\" }}\n")

	// Convert string builder to string
	generatedHTML = builder.String()

	// Write generated HTML to file
	fileName = "views/" + table + "-create.html"
	file, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(generatedHTML)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("HTML template file generated successfully:", fileName)

	// Code from html3.go
	builder = strings.Builder{}

	// Write HTML template
	builder.WriteString("{{ template \"_header.html\" }}\n")
	builder.WriteString("<div class=\"gap\"></div>\n")
	builder.WriteString("<div class=\"default-container\">\n")

	// Form
	builder.WriteString("    <div class=\"row\">\n")
	builder.WriteString("        <div class=\"col\">\n")
	builder.WriteString(fmt.Sprintf("            <a href=\"/%s\" class=\"text-primary underline\">%s</a>\n", table, table))
	builder.WriteString(fmt.Sprintf("            <h3>Update %s</h3>\n", capitalize(table)))
	builder.WriteString(fmt.Sprintf("            <form method=\"POST\" action=\"/%s/update/{{.Data.Id}}\">\n", table))
	builder.WriteString("                {{.Csrf}}\n")

	// Form fields
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("                <div class=\"form-group\">\n"))
		builder.WriteString(fmt.Sprintf("                    <label for=\"%s\">%s</label>\n", field.Name, capitalize(field.Name)))
		builder.WriteString(fmt.Sprintf("                    <input class=\"form-control\" name=\"%s\" value=\"{{.Data.%s.%s}}\" placeholder=\"\" />\n", field.Name, capitalize(structName), capitalize(field.Name)))
		builder.WriteString(fmt.Sprintf("                    <div class=\"small text-danger\"></div>\n"))
		builder.WriteString(fmt.Sprintf("                </div>\n"))
	}

	// Submit button
	builder.WriteString("                <div class=\"form-group\">\n")
	builder.WriteString("                    <button>Update</button>\n")
	builder.WriteString("                </div>\n")

	builder.WriteString("            </form>\n")
	builder.WriteString("        </div>\n")
	builder.WriteString("    </div>\n")

	builder.WriteString("</div>\n")
	builder.WriteString("{{ template \"_footer.html\" }}\n")

	// Convert string builder to string
	generatedHTML = builder.String()

	// Write generated HTML to file
	fileName = "views/" + table + "-update.html"
	file, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(generatedHTML)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("HTML template file generated successfully:", fileName)

	//finish html

	// Generate the CRUD operations code
	// We need to pass in dbtype obtained from .env file
	generatedCode := GenerateCRUD(structName, table, allFields, dbConnection)

	// Write the generated code to a file
	fileName = "models/" + table + "/" + table + ".go"
	dir := filepath.Dir(fileName)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	file, err = os.Create(fileName)
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

	// Write the controller code
	var builder2 strings.Builder

	builder2.WriteString("package " + table + "\n\n")
	builder2.WriteString("import (\n")
	builder2.WriteString("\t\"gozen/models/" + table + "\"\n")
	builder2.WriteString("\t\"gozen/system/templates\"\n")
	builder2.WriteString("\t\"gozen/system/validation\"\n")
	builder2.WriteString("\t\"github.com/go-chi/chi/v5\"\n")
	builder2.WriteString("\t\"net/http\"\n")
	builder2.WriteString("\t\"fmt\"\n")
	builder2.WriteString(")\n\n")

	// Write Index function
	builder2.WriteString("func Index(w http.ResponseWriter, r *http.Request) {\n")
	builder2.WriteString("\t" + table + ", err := " + table + ".All()\n")
	builder2.WriteString("\tif err != nil {\n")
	builder2.WriteString("\t\t// Handle error\n")
	builder2.WriteString("\t\treturn\n")
	builder2.WriteString("\t}\n")
	builder2.WriteString("\ttemplates.Render(w, r, \"" + table + "-all\", " + table + ")\n")
	builder2.WriteString("}\n\n")

	// Write CreateView function
	builder2.WriteString("func CreateView(w http.ResponseWriter, r *http.Request) {\n")
	builder2.WriteString("\tdata := templates.TemplateData{\n")
	builder2.WriteString("\t\t// You can set data here as needed\n")
	builder2.WriteString("\t}\n")
	builder2.WriteString("\ttemplates.Render(w, r, \"" + table + "-create\", data)\n")
	builder2.WriteString("}\n\n")

	// Write Create function
	builder2.WriteString("func Create(w http.ResponseWriter, r *http.Request) {\n")
	builder2.WriteString("\tv := &validation.Validator{}\n")

	// Write form value assignments dynamically
	for _, field := range fields {
		builder2.WriteString(fmt.Sprintf("\t%s := r.FormValue(\"%s\")\n", field.Name, field.Name))
	}

	// Write validation rules dynamically
	for _, field := range fields {
		if field.Type == "string" {
			builder2.WriteString(fmt.Sprintf("\tv.Required(\"%s\", %s)\n", field.Name, field.Name))
		}
	}

	builder2.WriteString("\tpostData := templates.PostData(w, r)\n")
	builder2.WriteString("\tif v.HasErrors() {\n")
	builder2.WriteString("\t\ttemplates.Errors(w, r, v, postData, \"" + table + "-create\")\n")
	builder2.WriteString("\t\treturn\n")
	builder2.WriteString("\t}\n")
	builder2.WriteString("\t// Else, no validation errors, proceed with creation\n")

	// Write creation of user dynamically
	builder2.WriteString("\tt, _ := " + table + ".Create(")
	for i, field := range fields {
		builder2.WriteString(field.Name)
		if i < len(fields)-1 {
			builder2.WriteString(", ")
		}
	}
	builder2.WriteString(")\n")
	builder2.WriteString("\tfmt.Print(t)\n")
	builder2.WriteString("\thttp.Redirect(w, r, \"/" + table + "\", http.StatusFound)\n")
	builder2.WriteString("}\n\n")

	// Write UpdateView function
	builder2.WriteString("func UpdateView(w http.ResponseWriter, r *http.Request) {\n")
	builder2.WriteString("\tid := chi.URLParam(r, \"id\")\n")
	builder2.WriteString("\t" + structName + ", err := " + table + ".Read(id)\n")
	builder2.WriteString("\tif err != nil {\n")
	builder2.WriteString("\t\t// Handle error\n")
	builder2.WriteString("\t\treturn\n")
	builder2.WriteString("\t}\n")
	builder2.WriteString("\tdata := map[interface{}]interface{}{\n")
	builder2.WriteString("\t\t\"Id\":   id,\n")
	builder2.WriteString("\t\t\"" + capitalize(structName) + "\": " + structName + ",\n")
	builder2.WriteString("\t}\n")
	builder2.WriteString("\ttemplates.Render(w, r, \"" + table + "-update\", data)\n")
	builder2.WriteString("}\n\n")

	// Write Update function
	builder2.WriteString("func Update(w http.ResponseWriter, r *http.Request) {\n")
	builder2.WriteString("\tid := chi.URLParam(r, \"id\")\n")

	// Write form value assignments dynamically
	for _, field := range fields {
		builder2.WriteString(fmt.Sprintf("\t%s := r.FormValue(\"%s\")\n", field.Name, field.Name))
	}

	builder2.WriteString("\tif err := " + table + ".Update(id, ")
	// Write fields dynamically for Update function
	for i, field := range fields {
		builder2.WriteString(field.Name)
		if i < len(fields)-1 {
			builder2.WriteString(", ")
		}
	}
	builder2.WriteString("); err != nil {\n")
	builder2.WriteString("\t\t// Handle error\n")
	builder2.WriteString("\t\treturn\n")
	builder2.WriteString("\t}\n")
	builder2.WriteString("\thttp.Redirect(w, r, \"/" + table + "\", http.StatusFound)\n")
	builder2.WriteString("}\n\n")

	// Write Destroy function
	builder2.WriteString("func Destroy(w http.ResponseWriter, r *http.Request) {\n")
	builder2.WriteString("\tid := chi.URLParam(r, \"id\")\n")
	builder2.WriteString("\tif err := " + table + ".Delete(id); err != nil {\n")
	builder2.WriteString("\t\t// Handle error\n")
	builder2.WriteString("\t\treturn\n")
	builder2.WriteString("\t}\n")
	builder2.WriteString("\thttp.Redirect(w, r, \"/" + table + "\", http.StatusFound)\n")
	builder2.WriteString("}\n\n")

	generatedCode = builder2.String()

	fileName = "controllers/" + table + "/" + table + ".go"
	dir = filepath.Dir(fileName)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	file, err = os.Create(fileName)
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

	fmt.Println("Controller file generated successfully:", fileName)
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
			//Type: parts[1],
			Type: "string",
		})
	}

	return fields
}

func parseFields2(fieldsStr string) []StructField {
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

type DataType struct {
	GoType     string
	SQLite     string
	MySQL      string
	PostgreSQL string
	SQLServer  string
}

// Let's handle the different db grammars
var dataTypes = []DataType{
	{"integer", "INTEGER", "INT", "INT", "INT"},
	{"float", "REAL", "FLOAT", "FLOAT", "FLOAT"},
	{"string", "TEXT", "VARCHAR(255)", "VARCHAR(255)", "NVARCHAR(255)"},
	{"text", "TEXT", "TEXT", "TEXT", "NVARCHAR(MAX)"},
	{"boolean", "INTEGER", "TINYINT(1)", "BOOLEAN", "BIT"},
	{"date", "TEXT", "DATE", "DATE", "DATE"},
	{"datetime", "TEXT", "DATETIME", "TIMESTAMP", "DATETIME"},
	{"time", "TEXT", "TIME", "TIME", "TIME"},
	{"timestamp", "TEXT", "TIMESTAMP", "TIMESTAMP", "DATETIME"},
}

func generateCreateTableSQL(tableName string, fields []StructField, dbConnection string) string {

	//Do a db check to build grammar specific sql
	idString := ""

	switch dbConnection {
	case "sqlite":
		idString = " PRIMARY KEY"
	case "mysql":
		idString = " PRIMARY KEY AUTO_INCREMENT"
	case "pgsql":
		idString = " SERIAL PRIMARY KEY"
	case "sqlsvr":
		idString = " IDENTITY(1,1) PRIMARY KEY"
	default:
		return "Invalid database type"
	}

	/*
	   |---------------------------------------------------------------
	   | Bug with pgsql
	   |---------------------------------------------------------------
	   |
	   | We need to omit the int/ integer type when creating a primary
	   | key for postgres
	   |
	*/

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", tableName))

	for i, field := range fields {
		fieldType := getFieldType(field.Type, dbConnection)

		//fix pgsql bug
		if dbConnection == "pgsql" && field.Name == "id" {
			sb.WriteString(fmt.Sprintf("    %s ", field.Name))
		} else {
			sb.WriteString(fmt.Sprintf("    %s %s", field.Name, fieldType))
		}

		if field.Name == "id" {
			//sb.WriteString(" PRIMARY KEY AUTO_INCREMENT")
			sb.WriteString(idString)
		}
		if i < len(fields)-1 {
			sb.WriteString(",\n")
		}
	}

	sb.WriteString("\n)")
	return sb.String()
}

func getFieldType(fieldType string, dbType string) string {
	for _, dt := range dataTypes {
		if dt.GoType == fieldType {
			switch dbType {
			case "sqlite":
				return dt.SQLite
			case "mysql":
				return dt.MySQL
			case "pgsql":
				return dt.PostgreSQL
			case "sqlsvr":
				return dt.SQLServer
			}
		}
	}
	return fieldType // Default case if no match is found
}

func migrateSql(table string, sql string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitDB()

	_, err = db.DB.Exec(sql)

	if err != nil {
		fmt.Println("Error creating table:", err)

	}
}

func buildRoutes(resource string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("\tr.Get(\"/%s\", %s.Index)\n", resource, resource))
	sb.WriteString(fmt.Sprintf("\tr.Get(\"/%s/create\", %s.CreateView)\n", resource, resource))
	sb.WriteString(fmt.Sprintf("\tr.Post(\"/%s/create\", %s.Create)\n", resource, resource))
	sb.WriteString(fmt.Sprintf("\tr.Get(\"/%s/update/{id}\", %s.UpdateView)\n", resource, resource))
	sb.WriteString(fmt.Sprintf("\tr.Post(\"/%s/update/{id}\", %s.Update)\n", resource, resource))
	sb.WriteString(fmt.Sprintf("\tr.Get(\"/%s/delete/{id}\", %s.Destroy)\n", resource, resource))

	return sb.String()
}

func writeRoutes(table string) {

	existingContent, err := ioutil.ReadFile("routes/routes.go")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string
	contentStr := string(existingContent)

	// Add a new import statement
	newImport := "\n\t\"gozen/controllers/" + table + "\""
	contentStr = strings.Replace(contentStr, "import (", "import ("+newImport, 1)

	// Add a new string before the closing brace
	newString := buildRoutes(table)
	contentStr = strings.Replace(contentStr, "} //end", newString+"} //end", 1)

	// Write the updated content back to the file
	err = ioutil.WriteFile("routes/routes.go", []byte(contentStr), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("File overwritten successfully!")

}

// GenerateCRUD generates CRUD operations for a given struct
func GenerateCRUD(structName string, table string, fields []StructField, dbConnection string) string {

	// Here we need to use the dbConnection type to dynamically change our
	// SQL grammar so it supports the four databases. In general, sqlite and
	// mysql are largely the same, postgres and mssql have the main differences
	//
	// Postgres uses $1,$2 for prepared statements
	// Mssql uses @name, @email for prepared statements

	if dbConnection == "hey" {
		fmt.Print("")
	} else {
		fmt.Print("")
	}

	var builder strings.Builder

	// Generate package and imports
	builder.WriteString(fmt.Sprintf("package %s\n\n", table))
	builder.WriteString("import (\n")
	builder.WriteString("\t\"gozen/db\"\n")
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
