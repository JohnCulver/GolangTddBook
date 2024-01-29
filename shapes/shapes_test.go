package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("Perimter got %.2f and want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("Area got %.2f and want %.2f", got, want)
		}
	}

	t.Run("Area of rectangle", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, rect, want)
	})

	t.Run("Area of circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})

}
