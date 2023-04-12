package main

import (
	"github.com/utkarsh-josh/interface/circle"
	"github.com/utkarsh-josh/interface/rectangle"
)

func main() {
	rect := rectangle.NewRectangle(5, 10)
	circ := circle.NewCircle(14)
	rect.Area()
	rect.Perimeter()
	circ.Area()
	circ.Perimeter()
}
