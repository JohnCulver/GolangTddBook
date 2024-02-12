package maps

import "testing"

func TestSearch(t *testing.T) {

	key := "test"
	value := "this is just a test"

	dictionary := Dictionary{key: value}

	t.Run("basic search", func(t *testing.T) {

		got, err := dictionary.Search(key)

		assertStrings(t, got, value)
		assertNoError(t, err)
	})

	t.Run("missing key", func(t *testing.T) {
		_, err := dictionary.Search("nothere")
		want := ErrMissingKey

		assertError(t, err, want)
	})
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal("Expected no error, but we got one.")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Got error %q, wanted %q", got, want)
	}
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, but wanted %q", got, want)
	}
}
