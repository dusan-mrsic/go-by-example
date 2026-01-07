package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{height: 5, width: 10}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	pr := &r
	fmt.Println("area: ", pr.area())
	fmt.Println("perim: ", pr.perim())
}
