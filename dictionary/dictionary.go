package dictionary

import (
	"errors"
)

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

func Search(m map[string]string, key string) string {
	return m[key]
}

func (m Dictionary) Search(key string) (string, error) {
	definition, ok := m[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (m Dictionary) Add(word, definition string) error {
	_, err := m.Search(word)
	switch err {
	case ErrNotFound:
		m[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (m Dictionary) Update(word, definition string) error {
	_, err := m.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		m[word] = definition
	default:
		return err
	}
	return nil
}

func (m Dictionary) Delete(word string) {
	delete(m, word)
}
func (e DictionaryErr) Error() string {
	return string(e)
}
