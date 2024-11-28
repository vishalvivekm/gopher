package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("very-secret-must-not-be-leaked"))

func handler(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session-name")
	session.Options = &sessions.Options{
		Path: "/",
		MaxAge: 3600,
		HttpOnly: true,
	}
	if err != nil {
		log.Println("error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if flashes := session.Flashes(); len(flashes) > 0 {
		for f := range flashes {
			fmt.Println(flashes[f])
		}
	} else {
		session.AddFlash("flash! There are only four rules you need to remember: make the plan, execute the plan, expect the plan to go off the rails, throw away the plan")
	}
	// set values
	session.Values["name"] = "vivek"
	session.Values[124] = 125
	session.Save(r,w)

}

func getSessionData(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "session-name")
    if err != nil {
        log.Println("Error retrieving session:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // retrieve the session values
    name := session.Values["name"]
    number := session.Values[124]
    
    // display em
    fmt.Fprintf(w, "Name: %v, Number: %v", name, number)
}

func main() {

	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/", handler)
	router.HandleFunc("/get-session", getSessionData)
	http.ListenAndServe(":9090", router)
}
