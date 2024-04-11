/*
|---------------------------------------------------------------
| Session helpers
|---------------------------------------------------------------
|
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/
package sessionstore

import "github.com/gorilla/sessions"
import "net/http"
import "log"

//import "fibs/system/hash"

// Fix this! Get from env file
// var t, _ = hash.GenerateKey()

var Store = sessions.NewCookieStore([]byte("secret"))

// Simply set a session variable
func SetSession(w http.ResponseWriter, r *http.Request, n string, t string) {

	session, _ := Store.Get(r, "session-name")
	session.Values[n] = t

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve the session
func GetSession(r *http.Request, t string) string {
	session, _ := Store.Get(r, "session-name")
	b := session.Values[t].(string)
	return b
}

// Destroy all sessions
func DestroySession(w http.ResponseWriter, r *http.Request) {

	session, _ := Store.Get(r, "session-name")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		log.Print("failed to delete session", err)
	}
}

/*
|---------------------------------------------------------------
| Flash data for forms
|---------------------------------------------------------------
*/
func SetOldPostData(w http.ResponseWriter, r *http.Request) {

	//we need to skip the CSRF token
	//and add an arbitary prefix to
	//avoid collision
	for key, values := range r.PostForm {
		for _, value := range values {
			//fmt.Printf("Field: %s, Value: %s\n", key, value)
			SetSession(w, r, key, value)
		}
	}
}

// eg
//
//	PostData := map[string]interface{}{
//	 "name":  "John",
//	 "email": "john@example.com",
//	}
func GetOldPostData(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	//init empty map interface
	PostData := map[string]interface{}{}

	for key, values := range r.PostForm {
		for _ = range values {
			//fmt.Printf("Field: %s, Value: %s\n", key, value)
			tmp := GetSession(r, key)
			PostData[key] = tmp

		}
	}

	return PostData
}
