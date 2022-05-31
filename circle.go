package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (s Circle) CalcPerimeter() float64 {
	return 2 * math.Pi * s.Radius
}

func (s Circle) CalcArea() float64 {
	return math.Pi * s.Radius * s.Radius
}
