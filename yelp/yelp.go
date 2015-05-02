package yelp

type Restaurant struct {
	name   string
	zip    string
	rating uint32
}

func Results() []Restaurant {
	resp := make([]Restaurant, 3)
	resp[0] = Restaurant{name: "Bob's Burgers", zip: "80304", rating: 3}
	resp[1] = Restaurant{name: "Jane's Joint", zip: "80304", rating: 4}
	resp[2] = Restaurant{name: "Dick's Diner", zip: "80304", rating: 5}
	return resp
}
