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
package rendering

import (
	"gozen/system/templates"
	"github.com/gorilla/csrf"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {

	tpl := templates.GetTemplate()

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
