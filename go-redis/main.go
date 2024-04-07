package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.Info("starting server..")
	http.HandleFunc("/api", Handler)

	log.Fatalln(http.ListenAndServe(":9090", nil))
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	log.Info("from server")

	resp := APIReponse{
		Cache: false,
		Data:  make([]NominatimResponse, 0),
	}
	err := json.NewEncoder(writer).Encode(resp)
	if err != nil {
		log.Fatalf("error encoding the response: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

type APIReponse struct {
	Cache bool                `json:"cache"`
	Data  []NominatimResponse `json:"data"`
}
type NominatimResponse struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	PlaceRank   int      `json:"place_rank"`
	Importance  float64  `json:"importance"`
	Addresstype string   `json:"addresstype"`
	Name        string   `json:"name"`
	DisplayName string   `json:"display_name"`
	Boundingbox []string `json:"boundingbox"`
}

