package main

import "testing"

type Shape interface{
	Area() float64
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6},want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func (t *testing.T) {
			t.Helper()
			got := tt.shape.Area()
			
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})

	}
}