/*
|---------------------------------------------------------------
| Templating
|---------------------------------------------------------------
|
| A helper for rendering server side templates
|
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package templates

import (
	"github.com/gorilla/csrf"
	"html/template"
	"net/http"
	"strings"
   "gozen/system/session"
	"gozen/system/validation"

)

type TemplateData struct {
	PostData       map[string]interface{}
	PostDataErrors map[string]interface{}
	FlashData      string
	Foo            string
}


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

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {

	tpl := GetTemplate()

	tmp := tmpl + ".html"

	// Execute the template and write it to the response
	err := tpl.ExecuteTemplate(w, tmp, struct {
		Data interface{}
		Csrf template.HTML
	}{
		Data: data,
		Csrf: csrf.TemplateField(r), // Get CSRF token from the request
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func SetAndGetPostData(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	session.SetOldPostData(w, r)
	return session.GetOldPostData(w, r)
}

func Errors(w http.ResponseWriter,
	r *http.Request,
	validator *validation.Validator,
	postData map[string]interface{}, templatePath string) {
	postDataErrors := make(map[string]interface{})
	for _, err := range validator.GetErrors() {
		postDataErrors[err.Field] = err.Message
	}

	data := TemplateData{
		PostData:       postData,
		PostDataErrors: postDataErrors,
		FlashData:      "Failed, error occurred",
	}
	RenderTemplate(w, r, templatePath, data)
}

