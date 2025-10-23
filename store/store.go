package store

import "errors"

var (
	store = map[string]string{}
)

func Insert(key, val string) {
	// i need a store to put stuff in
	// i need to be able to retrieve the key
	store[key] = val
}

func Fetch(key string) (string, error) {
	val, ok := store[key]
	if !ok {
		return "", errors.New("key not found")
	}

	return val, nil
}
