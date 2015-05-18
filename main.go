package main

import (
	"errors"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/jingweno/negroni-gorelic"
	"github.com/tracehelms/golunch/yelp"
)

type InputQuery struct {
	Query    string
	Location string
}

var cachedResults map[InputQuery]yelp.SearchResult

func main() {
	cachedResults = make(map[InputQuery]yelp.SearchResult)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", SearchHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	// Configure New Relic if license key available
	newRelicLicenseKey := os.Getenv("NewRelicLicenseKey")
	if newRelicLicenseKey != "" {
		n.Use(negronigorelic.New(newRelicLicenseKey, "gruuub", true))
	}

	n.Run(":" + port)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	defer LogTime(time.Now())
	log.Println("In searchhandler")
	tmpl, err := template.ParseFiles("./views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := r.URL.Query().Get("query")
	location := r.URL.Query().Get("location")
	if location == "" {
		location = "80304"
	}
	b := GetRandomBusiness(query, location)

	err = tmpl.Execute(w, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	log.Printf("Added result set for (query: %s, location: %s) to cache.", query, location)
}

func SearchInCache(query string, location string) (yelp.SearchResult, error) {
	input := InputQuery{Query: query, Location: location}
	val, ok := cachedResults[input]
	if ok {
		return val, nil
	}
	return yelp.SearchResult{}, errors.New("not found in cache")

}
