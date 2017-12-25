package main

import (
	"fmt"
	"github.com/major24/mypkg1"
	"github.com/major24/mypkg2"
	"github.com/major24/math"
	"github.com/major24/geometry"
	"github.com/major24/shape"
	"github.com/major24/finance"
	"github.com/major24/creditcard"
	)

func main() {
	fmt.Printf("Hello, world from gocode src in hello.\n")
	fmt.Println(mypkg1.Test1("a"))
	fmt.Println(mypkg2.Test1("a"))
	// test math avg
	calcAverage()

	calcAreaFromRect()

	calcCreditCardTotal()
}

func calcAverage(){
	xs := []float64{1,2,3,4}
	avg := math.Average(xs)
	fmt.Println(avg)
}

// use from interface
func calcAreaFromRect(){
	r1 := shape.Rect{Width: 5, Height: 4}
	r2 := shape.Rect{Width: 2, Height: 3}
	c1 := shape.Circle{Radius: 3}
	c2 := shape.Circle{Radius: 6}

	geos := []geometry.Geometry{r1,r2,c1,c2}

	for _, geo := range geos{
		fmt.Println("Area=", geo, geo.Area())
		fmt.Println("Perim=", geo, geo.Perim())
	}
}


func calcCreditCardTotal() {
	v1 := creditcard.Visa{Rate: 0.1, Total: 1000.0}

	fin := []finance.Finance{v1}

	fmt.Println("Visa=", fin[0], fin[0].TotalInterest)

}