package simplefiletest

import (
	"fmt"
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

	fmt.Println(ok)
}

func TestExists(t *testing.T) {
	if !Exists("go.mod") {
		t.Fatal("exists wrong return")
	}
}
