package main

import (
	"fmt"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("The Area of the Square is", s.side*s.side)
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("I'm visiting circle and calculating area it's area")
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	fmt.Println("I'm visiting rectangle and calculating area it's area")
}
