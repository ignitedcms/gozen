package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// StructField represents a field in a struct
type StructField struct {
	Name string
	Type string
}

func main() {
	table := "users"

	fields := []StructField{
		{Name: "name", Type: "string"},
		{Name: "email", Type: "string"},
		{Name: "password", Type: "string"},
	}

	var builder strings.Builder

	// Write package declaration and imports
	builder.WriteString("package " + table + "\n\n")
	builder.WriteString("import (\n")
	builder.WriteString("\t\"fibs/models/" + table + "\"\n")
	builder.WriteString("\t\"fibs/system/rendering\"\n")
	builder.WriteString("\t\"fibs/system/formutils\"\n")
	builder.WriteString("\t\"fibs/system/validation\"\n")
	builder.WriteString("\t\"github.com/go-chi/chi/v5\"\n")
	builder.WriteString("\t\"net/http\"\n")
	builder.WriteString("\t\"fmt\"\n")
	builder.WriteString(")\n\n")

	// Write Index function
	builder.WriteString("func Index(w http.ResponseWriter, r *http.Request) {\n")
	builder.WriteString("\t" + table + ", err := " + table + ".All()\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\t// Handle error\n")
	builder.WriteString("\t\treturn\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\trendering.RenderTemplate(w, r, \"" + table + "-all\", " + table + ")\n")
	builder.WriteString("}\n\n")

	// Write CreateView function
	builder.WriteString("func CreateView(w http.ResponseWriter, r *http.Request) {\n")
	builder.WriteString("\tdata := formutils.TemplateData{\n")
	builder.WriteString("\t\t// You can set data here as needed\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\trendering.RenderTemplate(w, r, \"" + table + "-create\", data)\n")
	builder.WriteString("}\n\n")

	// Write Create function
	builder.WriteString("func Create(w http.ResponseWriter, r *http.Request) {\n")
	builder.WriteString("\tvalidator := &validation.Validator{}\n")

	// Write form value assignments dynamically
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("\t%s := r.FormValue(\"%s\")\n", field.Name, field.Name))
	}

	// Write validation rules dynamically
	for _, field := range fields {
		if field.Type == "string" {
			builder.WriteString(fmt.Sprintf("\tvalidator.Required(\"%s\", %s)\n", field.Name, field.Name))
		}
	}

	builder.WriteString("\tpostData := formutils.SetAndGetPostData(w, r)\n")
	builder.WriteString("\tif validator.HasErrors() {\n")
	builder.WriteString("\t\tformutils.HandleValidationErrors(w, r, validator, postData, \"" + table + "-create\")\n")
	builder.WriteString("\t\treturn\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\t// Else, no validation errors, proceed with creation\n")

	// Write creation of user dynamically
	builder.WriteString("\tt, _ := " + table + ".Create(")
	for i, field := range fields {
		builder.WriteString(field.Name)
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(")\n")
	builder.WriteString("\tfmt.Print(t)\n")
	builder.WriteString("\thttp.Redirect(w, r, \"/" + table + "\", http.StatusFound)\n")
	builder.WriteString("}\n\n")

	// Write UpdateView function
	builder.WriteString("func UpdateView(w http.ResponseWriter, r *http.Request) {\n")
	builder.WriteString("\tid := chi.URLParam(r, \"id\")\n")
	builder.WriteString("\tuser, err := " + table + ".Read(id)\n")
	builder.WriteString("\tif err != nil {\n")
	builder.WriteString("\t\t// Handle error\n")
	builder.WriteString("\t\treturn\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\tdata := map[interface{}]interface{}{\n")
	builder.WriteString("\t\t\"Id\":   id,\n")
	builder.WriteString("\t\t\"User\": user,\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\trendering.RenderTemplate(w, r, \"" + table + "-update\", data)\n")
	builder.WriteString("}\n\n")

	// Write Update function
	builder.WriteString("func Update(w http.ResponseWriter, r *http.Request) {\n")
	builder.WriteString("\tid := chi.URLParam(r, \"id\")\n")

	// Write form value assignments dynamically
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("\t%s := r.FormValue(\"%s\")\n", field.Name, field.Name))
	}

	builder.WriteString("\tif err := " + table + ".Update(id, ")
	// Write fields dynamically for Update function
	for i, field := range fields {
		builder.WriteString(field.Name)
		if i < len(fields)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("); err != nil {\n")
	builder.WriteString("\t\t// Handle error\n")
	builder.WriteString("\t\treturn\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\thttp.Redirect(w, r, \"/" + table + "\", http.StatusFound)\n")
	builder.WriteString("}\n\n")

	// Write Destroy function
	builder.WriteString("func Destroy(w http.ResponseWriter, r *http.Request) {\n")
	builder.WriteString("\tid := chi.URLParam(r, \"id\")\n")
	builder.WriteString("\tif err := " + table + ".Delete(id); err != nil {\n")
	builder.WriteString("\t\t// Handle error\n")
	builder.WriteString("\t\treturn\n")
	builder.WriteString("\t}\n")
	builder.WriteString("\thttp.Redirect(w, r, \"/" + table + "\", http.StatusFound)\n")
	builder.WriteString("}\n\n")

	// Convert string builder to string
	generatedCode := builder.String()

	fileName := "controllers/" + table + "/" + table + ".go"
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

	fmt.Println("Controller file generated successfully:", fileName)
}
