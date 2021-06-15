package simplefiletest

import (
	"testing"
)

var testSimilarContentValue = `module github.com/TheBoringDude/simple-filetest

go 1.16
`

func TestSimilarContents(t *testing.T) {
	ok, err := SimilarContents("go.mod", testSimilarContentValue)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("file's content should be similar")
	}
}

func TestHasContents(t *testing.T) {
	ok, err := HasContents("go.mod", "module github.com/TheBoringDude/simple-filetest")
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("file should contain the string")
	}
}

func TestExists(t *testing.T) {
	if !Exists("go.mod") {
		t.Fatal("exists wrong return")
	}
}

func TestFileExists(t *testing.T) {
	if !FileExists("file.go") {
		t.Fatal("file should exist")
	}
}

func TestFolderExists(t *testing.T) {
	if DirExists("non-existstent folder") {
		t.Fatal("folder doesn't exist but function return that it exists")
	}
}
