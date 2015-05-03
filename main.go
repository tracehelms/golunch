package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tracehelms/golunch/yelp"
)

func main() {
	results := yelp.New().Search("hamburger", "80304", 10)
	count := len(results.Businesses)
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
