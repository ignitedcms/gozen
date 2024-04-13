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
	"fmt"
	"gozen/models/query"
	//"gozen/system/rendering"
	//"gozen/system/formutils"
	//"gozen/system/validation"
	"net/http"
)

// index page
func Index(w http.ResponseWriter, r *http.Request) {

	tableInfos, err := query.GetAll()
	if err != nil {
		// Handle error
		return
	}

	for _, tableInfo := range tableInfos {
		fmt.Printf("Table: %s (UUID: %s)\n", tableInfo.Table.Name, tableInfo.Table.UUID)
		fmt.Println("Columns:")
		for _, column := range tableInfo.Columns {
			fmt.Printf("- %s (UUID: %s)\n", column.Name, column.UUID)
		}
		fmt.Println()
	}

	// Render the template and write it to the response
	//rendering.RenderTemplate(w, r, "query", TableInfos)
}
