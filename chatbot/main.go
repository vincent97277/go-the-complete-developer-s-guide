package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		base:   10,
		height: 5.5,
	}
	squ := square{
		sideLength: 10,
	}

	printArea(tri)
	printArea(squ)
}

func (t triangle) getArea() float64 {
	return t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
