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
	fmt.Println(results.Businesses[randNumber])
}
