package main

import (
	"fmt"
	"math"
)

//no matter which struct is chosen both are a form of geometry and therefore the both have area and perimeter
type geometry interface {
	area() float64
	perim() float64
}

//define rectangles
type rect struct {
	width, height float64
}

//define circles
type circle struct {
	radius float64
}

//define area and perim for rectangles
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//define area and perim for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//defines measure funtion that takes geometrys as input
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 69, height: 420}
	c := circle{radius: 42069}

	measure(r)
	measure(c)
}
