package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat a string 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		verifyStringMatch(got, want, t)
	})

	t.Run("Repeat 0 times, get an empty string", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""
		verifyStringMatch(got, want, t)
	})
}

func verifyStringMatch(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("Got %q and wanted %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
