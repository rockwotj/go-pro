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
// START SOLUTION
func (m meat) getPrice() float64 {
	return m.pounds * m.pricePerPound
}

func (m milk) getPrice() float64 {
	return m.gallons * m.pricePerGallon
}

type stat struct {
	result float64
	statType string
}

func groceryListStats(r chan result) {
	stats := result{}
	results := make(chan stat)

	go getAverage(groceryList, results)
	go getMax(groceryList, results)
	go getTotal(groceryList, results)
	
	fmt.Println("Started All Calcs")
	for i := 0; i < 3; i++ {
		s := <-results
		switch s.statType {
		case "average":
			stats.average = s.result
		case "max":
			stats.max = s.result
		case "total":
			stats.total = s.result
		}
	}
	r <-stats
}

func getAverage (groceryList []groceryItem, results chan stat) {
		ave := 0.0
		for _, i := range groceryList {
			ave += i.getPrice()
		}
		ave = ave / float64(len(groceryList))
		results<-stat{result:ave, statType:"average"}
	}

func getMax (groceryList []groceryItem, results chan stat) {
		max := 0.0
		for _, i := range groceryList {
			if i.getPrice() > max {
				max = i.getPrice()
			}
		}
		results<-stat{result:max, statType:"max"}
	}

func getTotal (groceryList []groceryItem, results chan stat) {
		total := 0.0
		for _, i := range groceryList {
			total += i.getPrice()
		}
		results<-stat{result:total, statType:"total"}
	}

// END SOLUTION
	
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