/*                                                                          
|---------------------------------------------------------------            
| Query builder
|---------------------------------------------------------------            
|
| Experimental build
| 
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/       
package query

import (
	//"fmt"
	"gozen/system/rendering"
	//"gozen/system/formutils"
	//"gozen/system/validation"
	"net/http"
)

// index page
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "query", nil)
}

