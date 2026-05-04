package dictionary

type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}

func Search(dict map[string]string, word string) string {
	return dict[word]
}
