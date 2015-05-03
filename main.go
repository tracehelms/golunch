package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/tracehelms/golunch/yelp"
)

func main() {
	location := flag.String("location", "80304", "the location you want to search in")
	flag.Parse()

	query := flag.Args()[0]

	results := yelp.New().Search(query, *location)
	count := len(results.Businesses)
	if count <= 0 {
		fmt.Println("No results were found!")
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNumber := r.Intn(count)

	printResult(results.Businesses[randNumber])
}

func printResult(business yelp.Business) {
	fmt.Println("Time for grub! Your destination:")
	fmt.Println("Name:     ", business.Name)
	fmt.Println("Rating:   ", business.Rating)
	fmt.Println("Location: ", business.Location.Address)
	fmt.Println("URL:      ", business.Url)
}
