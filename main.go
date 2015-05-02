package main

import (
	"fmt"
	"math/rand"

	"github.com/tracehelms/golunch/yelp"
)

func main() {
	yc := yelp.New()
	results := yc.Search("hamburger")
	fmt.Println(results)
	// restaurants := yelp.Results()
	// fmt.Println("your restaurant is: ", restaurants[rand.Intn(3)])
}
