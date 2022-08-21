package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) GetMagnitude() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func (p *Point) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func Sub(from, what Point) *Point {
	return NewPoint(from.x-what.x, from.y-what.y)
}

func GetDistance(from, to Point) float64 {
	return Sub(from, to).GetMagnitude()
}

func main() {
	firstPoint := NewPoint(2, 6)
	secondPoint := NewPoint(1, 2)

	result := GetDistance(*firstPoint, *secondPoint)
	fmt.Println("First point:", firstPoint, "|", "Second point:", secondPoint, "|", "The distance between these points:", result)
}
