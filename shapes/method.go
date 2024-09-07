package shapes

import (
	"math"
)

// Area takes the side for a given square to calculate its area.
func (s Square) Area() float64 {
	return math.Pow(s.Side, 2)
}

// Perimeter takes the side for a given square to calculate its perimeter.
func (s Square) Perimeter() float64 {
	return math.Pow(s.Side, 2) * 2
}

// Area takes the length and width for a given rectangle to calculate its area.
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Perimeter takes the length and width for a given rectangle to calculate its
// perimeter.
func (r Rectangle) Perimeter() float64 {
	return (r.Length + r.Width) * 2
}

// Area takes the radius for a given circle to calculate its area.
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

// Perimeter takes the radius for a given circle to calculate its perimeter.
func (c Circle) Perimeter() float64 {
	return (math.Pi * 2) * c.Radius
}
