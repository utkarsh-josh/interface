package circle

import (
	"fmt"
)

type Circle interface {
	Area()
	Perimeter()
}

type circ struct {
	radius int
}

func NewCircle(r int) Circle {
	return &circ{
		radius: r,
	}
}

func (c *circ) Area() {
	fmt.Println("Area of Circle : ", (22/7) * c.radius * c.radius)
}

func (c *circ) Perimeter() {
	fmt.Println("Perimeter of Circle : ", 2 * (22/7) * c.radius)
}