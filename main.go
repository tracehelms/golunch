package main

import (
	"fmt"
	"math/rand"
)

func main() {
	restaurants := results()
	fmt.Println("your restaurant is: ", restaurants[rand.Intn(3)])
}
