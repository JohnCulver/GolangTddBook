package maps

import "errors"

type Dictionary map[string]string

var ErrMissingKey = errors.New("Key not in dict.")

func (dictionary Dictionary) Search(key string) (string, error) {
	value, ok := dictionary[key]

	if !ok {
		return "", ErrMissingKey
	}

	return value, nil
}
