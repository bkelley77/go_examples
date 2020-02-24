package main

// ******************************************************************
// imports
// ******************************************************************

import (
	"fmt"
	"math"
)

// ******************************************************************
// interfaces
// ******************************************************************

type Shape interface {
	perimeter() float64
}

// ******************************************************************
// structs
// ******************************************************************

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// ******************************************************************
// functions
// ******************************************************************

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := float64(5)
	w := float64(7)
	return 2 * (l + w)
}

func totalArea(shapes ...Shape) float64 {
	var per float64
	for _, s := range shapes {
		per += s.perimeter()
	}
	return per
}

// ******************************************************************
// main
// ******************************************************************

func main() {

	c := Circle {5,4,3}
	r := Rectangle{0, 0, 10, 10}
	t := totalArea( &c, &r)
	fmt.Println("total area = ", t)

}