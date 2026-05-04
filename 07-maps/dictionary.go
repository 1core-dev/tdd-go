package dictionary

const (
	ErrNotFound  = DictionaryError("not found word")
	ErrWordExist = DictionaryError("not found word")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return res, nil
}

func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = value
		return nil
	case nil:
		return ErrWordExist
	default:
		return err
	}
}
