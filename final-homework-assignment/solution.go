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

func (m milk) getPrice() float64 {
	return m.gallons * m.pricePerGallon
}

func groceryListStats(groceryList []groceryItem, r chan result) {
	stats := result{}
	averageResult := make(chan float64)
	maxResult := make(chan float64)
	totalResult := make(chan float64)

	go getAverage(groceryList, averageResult)
	go getMax(groceryList, maxResult)
	go getTotal(groceryList, totalResult)
	
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

func getAverage (groceryList []groceryItem, results chan float64) {
		ave := 0.0
		for _, i := range groceryList {
			ave += i.getPrice()
		}
		ave = ave / float64(len(groceryList))
		results<-ave
	}

func getMax (groceryList []groceryItem, results chan float64) {
		max := 0.0
		for _, i := range groceryList {
			if i.getPrice() > max {
				max = i.getPrice()
			}
		}
		results<-max
	}

func getTotal (groceryList []groceryItem, results chan float64) {
		total := 0.0
		for _, i := range groceryList {
			total += i.getPrice()
		}
		results<-total
	}

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