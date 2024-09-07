package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	tests := []struct {
		name   string
		shape  shape
		expect float64
	}{
		{name: "Square", shape: Square{Side: 3.5}, expect: 12.25},
		{name: "Rectangle", shape: Rectangle{Length: 4, Width: 2}, expect: 8},
		{name: "Circle", shape: Circle{Radius: 3}, expect: 28.274333882308138},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.shape.Area()
			assert.Equal(t, test.expect, actual)
		})
	}
}

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name   string
		shape  shape
		expect float64
	}{
		{name: "Square", shape: Square{Side: 3.5}, expect: 24.5},
		{name: "Rectangle", shape: Rectangle{Length: 4, Width: 2}, expect: 12},
		{name: "Circle", shape: Circle{Radius: 3}, expect: 18.84955592153876},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.shape.Perimeter()
			assert.Equal(t, test.expect, actual)
		})
	}
}
