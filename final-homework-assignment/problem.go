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
		return f.price
	} else {
		return 0
	}
}
func (m meat) getPrice() float64 {
	return m.pounds * m.pricePerPound
}
// START SOLUTION

// END SOLUTION
	
type result struct {
	average float64
	max float64
	total float64
}

func (r result) print() {
	fmt.Println("Printing Results")
	fmt.Println("Average:", r.average)
	fmt.Println("Max:", r.max)
	fmt.Println("Total:", r.total)	
}

func main() {
	groceryList := []groceryItem {
		fruit{fresh:false, price:12.33},
		fruit{fresh:true, price:2.33},
		meat{pounds:3.3, pricePerPound:1.61},
		meat{pounds:1.2, pricePerPound:24.1},
		milk{gallons:.5, pricePerGallon: 4.33},
	}
	channel := make(chan result)
	go groceryListStats(groceryList, channel)
	results := <-channel
	results.print()
}