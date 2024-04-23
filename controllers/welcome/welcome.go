package welcome

import (
	"gozen/system/templates"
	"net/http"
)

// welcome page
func Index(w http.ResponseWriter, r *http.Request) {

	// Render the template and write it to the response
	templates.Render(w, r, "welcome", nil)
}
