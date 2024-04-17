package mail

import (
	"fmt"
	"gozen/models/users"
	//"gozen/system/mail"
	"gozen/system/rendering"
	"net/http"
)

// index page for mail
func Index(w http.ResponseWriter, r *http.Request) {
	// Render the template and write it to the response
	rendering.RenderTemplate(w, r, "mail/index", nil)
}

func MailView(w http.ResponseWriter, r *http.Request) {
	rendering.RenderTemplate(w, r, "mail", nil)
}

func SendMail(w http.ResponseWriter, r *http.Request) {

	query, err := users.GetHash("foo@mail.com")
   //no record found
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(query.Email)

      to := query.Email
      //mail.Test(to)
	}
}
