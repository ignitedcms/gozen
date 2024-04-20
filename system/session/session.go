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
	// Check if the session value exists
	b, ok := session.Values[t].(string)
	if !ok || b == "" {
		// Session value doesn't exist or is an empty string
		return ""
	}

	// Session value exists
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

func SetOldPostData(w http.ResponseWriter, r *http.Request) {

	//we need to skip the CSRF token
	//and add an arbitary prefix to
	//avoid collision
	for key, values := range r.PostForm {
		for _, value := range values {
			//fmt.Printf("Field: %s, Value: %s\n", key, value)
			Set(w, r, key, value)
		}
	}
}

func GetOldPostData(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	//init empty map interface
	PostData := map[string]interface{}{}

	for key, values := range r.PostForm {
		for _ = range values {
			//fmt.Printf("Field: %s, Value: %s\n", key, value)
			tmp := Get(r, key)
			PostData[key] = tmp

		}
	}

	return PostData
}
