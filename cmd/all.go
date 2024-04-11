package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// StructField represents a field in a struct
type StructField struct {
	Name string
	Type string
}

func main() {
	table := "users"
	allFields := []StructField{
		{Name: "id", Type: "string"},
		{Name: "name", Type: "string"},
		{Name: "email", Type: "string"},
		{Name: "password", Type: "string"},
		{Name: "created_at", Type: "string"},
		{Name: "updated_at", Type: "string"},
	}

	// Fields for create and update HTML templates
	fields := []StructField{
		{Name: "name", Type: "string"},
		{Name: "email", Type: "string"},
		{Name: "password", Type: "string"},
	}

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
	builder.WriteString("                    <button-component variant=\"primary\">Save</button-component>\n")
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
		builder.WriteString(fmt.Sprintf("                    <input class=\"form-control\" name=\"%s\" value=\"{{.Data.%s.%s}}\" placeholder=\"\" />\n", field.Name, capitalize(table), capitalize(field.Name)))
		builder.WriteString(fmt.Sprintf("                    <div class=\"small text-danger\"></div>\n"))
		builder.WriteString(fmt.Sprintf("                </div>\n"))
	}

	// Submit button
	builder.WriteString("                <div class=\"form-group\">\n")
	builder.WriteString("                    <button-component variant=\"primary\">Update</button-component>\n")
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
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
