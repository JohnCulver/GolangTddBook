package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 10}, 40},
		{Circle{10}, 62},
		{Triangle{10, 10}, -1},
	}

	for _, pt := range perimeterTests {
		got := pt.shape.Perimeter()
		if got != pt.want {
			t.Errorf("Perimter of %q got %.2f and want %.2f", pt.shape.Name(), got, pt.want)
		}
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
			t.Errorf("Area of %q got %.2f and want %.2f", tt.shape.Name(), got, tt.want)
		}
	}

}
