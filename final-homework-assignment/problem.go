package main

// Summary of points:
// 1 point for coding style
// 1 point for defining a struct for milk that 'implements' groceryItem
// 1 point for a function to average a grocery list
// 1 point for a function to total a grocery list
// 1 point for a function to find the maximal priced item in a grocery list
// 1 point for making a channel for each of the above functions
// 4 points for running the above 3 functions all at once 
//   (for 2 points calculate the values sequentially 
//      - you'll have to comment out the select statement for this to work)


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

func (f fruit) getPrice() float64 {
	if f.fresh {
		return f.price
	} else {
		return 0
	}
}

type meat struct {
	pounds float64
	pricePerPound float64
}

func (m meat) getPrice() float64 {
	return m.pounds * m.pricePerPound
}

// START PROBLEM

// define a struct for milk with fields: gallons and pricePerGallon as both float64 types (1 point)

// Define a function to calculate the price of milk via: gallons * pricePerGallon
// Do this in a way so that milk 'implements' groceryItem

func groceryListStats(groceryList []groceryItem, r chan result) {
	stats := result{}
	// Open up a float 64 channel for each of the goroutines (1 point)
	totalResult := nil
	averageResult := nil
	maxResult := nil

	// Start each of the go routines (4 points)
	
	fmt.Println("Started All Calcs")
	for i := 0; i < 3; i++ {
		select {
			case stats.total = <- totalResult:
				fmt.Println("Recieved Total")
			case stats.average = <- averageResult:
				fmt.Println("Recieved Average")
			case stats.max = <- maxResult:
				fmt.Println("Recieved Max")
		}
	}
	r <-stats
}

// Write a function that totals the prices of all items in a grocery list (1 point)

// Write a function that finds the maximal price of an item in a grocery list (1 point)

// Write a function that finds the average of an item in a grocery list (1 point)
// You might find the methods len() [to find the length of a slice]
// and float64() [to convert an int to a float64] to be useful.

// END PROBLEM
	
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
		milk{gallons:.5, pricePerGallon: 4.33}, // MAKE SURE THIS MATCHES YOUR STRUCT
	}
	channel := make(chan result)
	go groceryListStats(groceryList, channel)
	results := <-channel
	results.print()
}
