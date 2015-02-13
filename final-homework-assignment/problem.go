package main

import(
	"fmt"
)

type groceryItem interface {
	getPrice() float64
}

type fruit struct {
	fresh bool
	price float64
}

type meat struct {
	pounds float64
	pricePerPound float64
}

type milk struct {
	gallons float64
	pricePerGallon float64	
}

func (f fruit) getPrice() float64 {
	if f.fresh {
		return price
	} else {
		0
	}
}

// YOUR SOLUTION HERE

type result struct {
	average float64
	max float64
	total float64
}

var groceryList = []groceryItem {
	fruit{fresh:false, price:12.33},
	fruit{fresh:true, price:2.33},
	meat{pounds:3.3, pricePerPound:1.61},
	meat{pounds:1.2, pricePerPound:24.1},
	milk{gallons:.5, pricePerGallon: 4.33},
}

func main() {
	// We write some tests too
	channel := make(chan result)
	go groceryListStats(channel)
	results := <-channel
	fmt.Println(results)
}