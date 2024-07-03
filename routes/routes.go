/*
|---------------------------------------------------------------
| Routes
|---------------------------------------------------------------
|
| We define all routes here that the main app uses
| We must import the controllers we need
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/
package routes

import (
	"github.com/go-chi/chi/v5"
	"gozen/controllers/welcome"
)

func LoadRoutes(r *chi.Mux) {

	r.Get("/", welcome.Index)

} //end
