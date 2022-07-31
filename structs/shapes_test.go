package structs

import (
	"testing"
)

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name      string
		rectangle Rectangle
		want      float64
	}{
		// TODO: Add test cases.
		{"pre1", Rectangle{5.0, 5.0}, 20.0},
		{"pre2", Rectangle{3.0, 5.0}, 16.0},
		{"pre3", Rectangle{9.0, 5.0}, 28.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rectangle.Perimeter(); got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Width: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}

}
