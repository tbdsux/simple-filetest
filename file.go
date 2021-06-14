package simplefiletest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// SimilarContents attempts to open and compare the contents of `filename` to `compare`.
// It returns true and nil if it is similar while false and an error otherwise.
func SimilarContents(filename string, compare string) (bool, error) {
	// filename should exist
	if Exists(filename) {
		f, err := ioutil.ReadFile(filename)
		if err != nil {
			return false, err
		}

		contents := string(f)

		if contents == compare {
			return true, nil
		}

		return false, errors.New("both strings are not similar")
	}

	return false, fmt.Errorf("%s does not exist", filename)
}

// Exists checks if `f` exists in the path or not.
// Details: https://stackoverflow.com/a/66405130
func Exists(f string) bool {
	if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
