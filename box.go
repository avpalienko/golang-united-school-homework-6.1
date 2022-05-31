package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorExceededCapacity = errors.New("capacity exceeded")
	errorOutOfRange       = errors.New("index out of range")
	errorNoCircle         = errors.New("circles are not exist")
)

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return errorExceededCapacity
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errorOutOfRange
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errorOutOfRange
	}
	ret := b.shapes[i]

	newLen := 0
	for ind, s := range b.shapes {
		if ind != i {
			b.shapes[newLen] = s
			newLen++
		}
	}
	b.shapes = b.shapes[:newLen]
	return ret, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= len(b.shapes) {
		return nil, errorOutOfRange
	}
	ret := b.shapes[i]
	b.shapes[i] = shape
	return ret, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() (ret float64) {
	ret = 0.0
	for _, s := range b.shapes {
		ret += s.CalcPerimeter()
	}
	return ret
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() (ret float64) {
	ret = 0.0
	for _, s := range b.shapes {
		ret += s.CalcArea()
	}
	return ret
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	no := true
	newLen := 0
	for _, s := range b.shapes {
		_, ok1 := s.(Circle)
		_, ok2 := s.(*Circle)
		if ok1 || ok2 {
			no = false
		} else {
			b.shapes[newLen] = s
			newLen++
		}
	}
	if no {
		return errorNoCircle
	} else {
		b.shapes = b.shapes[:newLen]
		return nil
	}
}
