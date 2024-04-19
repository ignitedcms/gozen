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
package session

import "github.com/gorilla/sessions"
import "net/http"
import "log"
import "os"

// Simply set a session variable
func Set(w http.ResponseWriter, r *http.Request, n string, t string) {

	key := os.Getenv("APP_KEY")
	Store := sessions.NewCookieStore([]byte(key))

	session, _ := Store.Get(r, "session-name")
	session.Values[n] = t

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Retrieve the session
func Get(r *http.Request, t string) string {
	key := os.Getenv("APP_KEY")
	Store := sessions.NewCookieStore([]byte(key))
	session, _ := Store.Get(r, "session-name")
	b := session.Values[t].(string)
	return b
}

// Destroy all sessions
func Destroy(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("APP_KEY")
	Store := sessions.NewCookieStore([]byte(key))

	session, _ := Store.Get(r, "session-name")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		log.Print("failed to delete session", err)
	}
}

