package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{}, 0},
		{Circle{}, 0},
		{Triangle{}, 0},
	}

	for _, pt := range perimeterTests {
		got := pt.shape.Perimeter()

		if got != pt.want {

		}
	}

	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("Perimter got %.2f and want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Area got %.2f and want %.2f", got, tt.want)
		}
	}

}
