package dictionary

import "errors"

var ErrNotFound = errors.New("not found word")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return res, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
