package maps

import "testing"

func TestSearch(t *testing.T) {

	key := "test"
	value := "this is just a test"

	dictionary := Dictionary{key: value}

	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got %q, but wanted %q", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("Got error %q, wanted %q", got, want)
		}
	}

	t.Run("basic search", func(t *testing.T) {

		got, _ := dictionary.Search(key)

		assertStrings(t, got, value)
		assertNoError(_)
	})

	t.Run("missing key", func(t *testing.T) {
		_, err := dictionary.Search("nothere")
		want := ErrMissingKey

		assertError(t, err, want)
	})
}
