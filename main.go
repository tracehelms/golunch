package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tracehelms/golunch/yelp"
)

type InputQuery struct {
	Query    string
	Location string
}

var cachedResults map[InputQuery]yelp.SearchResult

func main() {
	cachedResults = make(map[InputQuery]yelp.SearchResult)
	r := mux.NewRouter()
	r.HandleFunc("/api/search", SearchHandler).Methods("GET")
	http.Handle("/api/", r)
	http.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	defer LogTime(time.Now())
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("query")
	location := r.URL.Query().Get("location")
	if location == "" {
		location = "80304"
	}
	b := GetRandomBusiness(query, location)

	jsonResp, _ := json.Marshal(b)
	w.Write(jsonResp)
}

func GetRandomBusiness(query string, location string) yelp.Business {
	results, err := SearchInCache(query, location)
	if err != nil {
		results = yelp.New().Search(query, location)
		AddToCache(query, location, results)
	}
	count := len(results.Businesses)
	if count <= 0 {
		return yelp.Business{}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNumber := r.Intn(count)
	return results.Businesses[randNumber]
}

func LogTime(t time.Time) {
	elapsed := time.Since(t)
	log.Println("Response time: ", elapsed)
}

func AddToCache(query string, location string, results yelp.SearchResult) {
	input := InputQuery{Query: query, Location: location}
	cachedResults[input] = results
	log.Println("Added result set to cache...")
	log.Println("query: ", query)
	log.Println("location: ", location)
}

func SearchInCache(query string, location string) (yelp.SearchResult, error) {
	input := InputQuery{Query: query, Location: location}
	val, ok := cachedResults[input]
	if ok {
		return val, nil
	}
	return yelp.SearchResult{}, errors.New("not found in cache")

}
