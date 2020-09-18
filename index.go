package main

import (
	"fmt"
	"math"
)

// Init Rectangle struct
type Rectangle struct {
	width, length float64
}

// init Circle struct
type Circle struct {
	radius float64
}

// Init Rect area method
func (r *Rectangle) area() float64 {
	return r.width * r.length
}

// Init Circ area method
func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Init Shape interface
type Shape interface {
	area() float64
}

// MultiShape struct
type MultiShape struct {
	shapes []Shape
}

// TotalArea calculate
func totalArea(shapes ...Shape) (area float64) {
	for _, shape := range shapes {
		area += shape.area()
	} 
	return 
}

func main() {
	rect := Rectangle{width: 3, length: 2}
	circ := Circle{radius: 1}

	fmt.Println(totalArea(&rect, &circ))
}