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
	"gozen/models/query"
	"gozen/system/rendering"
	//"gozen/system/formutils"
	//"gozen/system/validation"
	"net/http"
)

// index page
func Index(w http.ResponseWriter, r *http.Request) {

	TableInfos, err := query.GetAll()
	if err != nil {
		// Handle error
		return
	}

	//for _, tableInfo := range TableInfos {
		//fmt.Printf("Table: %s\nColumns: %v\n", tableInfo.Name, tableInfo.Columns)
	//}

	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "query", TableInfos)
}
