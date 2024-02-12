package maps

import "testing"

func TestSearch(t *testing.T) {

	key := "test"
	value := "this is just a test"

	dictionary := Dictionary{key: value}

	t.Run("basic search", func(t *testing.T) {

		got, err := dictionary.Search(key)

		assertStrings(t, got, value)
		assertError(t, err, nil)
	})

	t.Run("missing key", func(t *testing.T) {
		_, err := dictionary.Search("nothere")
		want := ErrSearchKeyMissing

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	key := "a_key"
	value := "a_value"
	t.Run("testing basic add", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(key, value)
		assertError(t, err, nil)

		got, err := dictionary.Search(key)

		assertError(t, err, nil)
		assertStrings(t, got, value)
	})

	t.Run("testing no overwrite", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, ErrAddKeyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("test basic update", func(t *testing.T) {
		key := "key"
		dictionary := Dictionary{key: "value"}
		updatedValue := "updated"
		err := dictionary.Update(key, updatedValue)

		assertError(t, err, nil)

		fetchedValue, err := dictionary.Search(key)
		assertError(t, err, nil)
		assertStrings(t, fetchedValue, updatedValue)
	})

	t.Run("test update with no existing key", func(t *testing.T) {
		key := "key"
		dictionary := Dictionary{}

		err := dictionary.Update(key, "thisshouldntwork")

		assertError(t, err, ErrUpdateKeyNotFound)

	})
}

func TestDelete(t *testing.T) {
	t.Run("basic delete", func(t *testing.T) {
		dictionary := Dictionary{"key": "value"}
		err := dictionary.Delete("key")
		assertError(t, err, nil)

		_, err = dictionary.Search("key")
		assertError(t, err, ErrSearchKeyMissing)
	})

	t.Run("delete with no key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("key")

		assertError(t, err, ErrDeleteKeyMissing)
	})
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
