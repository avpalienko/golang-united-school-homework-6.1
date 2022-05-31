package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (s Rectangle) CalcPerimeter() float64 {
	return (s.Height + s.Weight) * 2
}

func (s Rectangle) CalcArea() float64 {
	return s.Height * s.Weight
}
