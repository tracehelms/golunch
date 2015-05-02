package main

import (
	"fmt"
	"math/rand"

	"github.com/tracehelms/golunch/yelp"
)

func main() {
	restaurants := results()
	fmt.Println("your restaurant is: ", restaurants[rand.Intn(3)])
}
