package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

func main() {
	log.Info("starting server..")
	http.HandleFunc("/api", Handler)

	log.Fatalln(http.ListenAndServe(":9090", nil))
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	log.Info("in the handler")

	q := request.URL.Query().Get("q")
	data, err := getData(q)
	if err != nil {
		fmt.Printf("error calling data source: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
	resp := APIReponse{
		Cache: false,
		Data:  data,
	}
	err = json.NewEncoder(writer).Encode(resp)
	if err != nil {
		fmt.Printf("error encoding the response: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}
func getData(q string) ([]NominatimResponse, error) {
	// San Francisco , san%20francisco

	escapedQ := url.PathEscape(q)
	address := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", escapedQ)
	resp, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	data := make([]NominatimResponse, 0)
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
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
