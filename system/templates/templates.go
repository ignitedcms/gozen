// templates.go
package templates

import (
	"html/template"
	"strings"
)

var Template *template.Template

// Define your template functions here
func upperCase(str string) string {
	return strings.ToUpper(str)
}

func showChecked(a string, b string) string {
	if strings.Compare(a, b) == 0 {
		return "checked"
	} else {
		return ""
	}
}

func showSelected(a string, b string) string {
	if strings.Compare(a, b) == 0 {
		return "selected"
	} else {
		return ""
	}
}

// LoadTemplates loads all the templates from the views directory
func LoadTemplates() error {
	funcMap := template.FuncMap{
		"upperCase":    upperCase,
		"showChecked":  showChecked,
		"showSelected": showSelected,
	}
	t, err := template.New("").Funcs(funcMap).ParseGlob("views/*.html")
	if err != nil {
		return err
	}
	Template = t
	return nil
}

// GetTemplate returns the loaded template
func GetTemplate() *template.Template {
	return Template
}
