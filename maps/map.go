package maps

type Dictionary map[string]string

var (
	ErrSearchKeyMissing  = DictionaryErr("could not find the key you were looking for")
	ErrUpdateKeyNotFound = DictionaryErr("cannot update word because it doesn't exist in dictionary")
	ErrAddKeyExists      = DictionaryErr("cannot add word because it already exists")
	ErrDeleteKeyMissing  = DictionaryErr("cannot delete word because it is not in dict")
)

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

func (dictionary Dictionary) Search(key string) (string, error) {
	value, ok := dictionary[key]

	if !ok {
		return "", ErrSearchKeyMissing
	}

	return value, nil
}

func (dictionary Dictionary) Add(key, value string) error {
	_, err := dictionary.Search(key)

	switch err {
	case ErrSearchKeyMissing:
		dictionary[key] = value
		return nil
	case nil:
		return ErrAddKeyExists
	default:
		return err
	}

}

func (dictionary Dictionary) Update(key, value string) error {
	err := dictionary.Add(key, value)

	switch err {
	case ErrAddKeyExists:
		dictionary[key] = value
		return nil
	case nil:
		return ErrUpdateKeyNotFound
	default:
		return err
	}

}

func (dictionary Dictionary) Delete(key string) error {
	_, err := dictionary.Search(key)
	switch err {
	case ErrSearchKeyMissing:
		return ErrDeleteKeyMissing
	case nil:
		delete(dictionary, key)
		return nil
	default:
		return err
	}
}
