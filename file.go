package simplefiletest

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

// file reader wrapper
func fileReader(filename string) (string, error) {
	if FileExists(filename) {
		f, err := ioutil.ReadFile(filename)
		if err != nil {
			return "", err
		}

		return string(f), nil

	}

	return "", fmt.Errorf("%s does not exist", filename)
}

// SimilarContents attempts to open and compare the contents of `filename` to `compare`.
// It returns true and nil if it is similar while false and an error otherwise.
func SimilarContents(filename string, compare string) (bool, error) {
	contents, err := fileReader(filename)
	if err != nil {
		return false, err
	}

	if contents == compare {
		return true, nil
	}

	return false, nil
}

// HasContents attempts to open and check the contents of `filename` if
// it includes `compare`.
func HasContents(filename string, compare string) (bool, error) {
	contents, err := fileReader(filename)
	if err != nil {
		return false, err
	}

	if strings.Contains(contents, compare) {
		return true, nil
	}

	return false, nil
}

// details: https://stackoverflow.com/a/66405130
func pathExists(path string) (fs.FileInfo, bool) {
	info, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		return nil, false
	}

	return info, true
}

// Exists checks if `f` exists in the path or not.
func Exists(f string) bool {
	_, ok := pathExists(f)
	return ok
}

// FileExists checks if file exists in path.
func FileExists(file string) bool {
	info, ok := pathExists(file)
	if !ok {
		return false
	}

	return !info.IsDir()
}

// DirExists checks if the dir exists in path.
func DirExists(dir string) bool {
	info, ok := pathExists(dir)
	if !ok {
		return false
	}

	return info.IsDir()
}
