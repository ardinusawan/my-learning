package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")
var ErrWordExist = errors.New("word already exist")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}
