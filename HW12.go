package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {

	var (
		mshape     MultiShape
		c1, c2, c3 Circle
		r1, r2, r3 Rectangle
	)

	c1 = Circle{0, 0, 5}
	c2 = Circle{0, 0, 3}
	c3 = Circle{0, 0, 4}
	r1 = Rectangle{0, 0, 10, 10}
	r2 = Rectangle{5, 6, 2, 1}
	r3 = Rectangle{0, 0, 9, 7}
	mshape.shapes = []Shape{&c1, &c2, &c3, &r1, &r2, &r3}
	fmt.Println(mshape.area())

}
