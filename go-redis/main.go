package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Info("starting server..")
	http.HandleFunc("/stories", Handler)

	log.Fatalln(http.ListenAndServe(":10000", nil))
}

func Handler(w http.ResponseWriter, request *http.Request) {
	response, err := http.Get("http://hn.algolia.com/api/v1/search?query=foo&tags=story")
	if err != nil {
		log.Fatalln(err)
	}
     
}
