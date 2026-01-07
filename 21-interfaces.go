package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius: ", c)
	}
}

func main() {
	r := rectangle{4, 3}
	c := circle{5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
