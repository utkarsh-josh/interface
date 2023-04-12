package rectangle

import (
	"fmt"
)

type Rectangle interface {
	Area()
	Perimeter()
}

type rect struct {
	length  int
	breadth int
}

func NewRectangle(l, b int) Rectangle {
	return &rect{
		length:  l,
		breadth: b,
	}
}

func (r *rect) Area() {
	fmt.Println("Area of rectangle : ", r.length*r.breadth)
}

func (r *rect) Perimeter() {
	fmt.Println("Perimiter of rectangle : ", 2*(r.length+r.breadth))
}
