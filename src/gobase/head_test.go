package gobase

import (
	"os"
	"testing"
)

func TestHead(t *testing.T) {
	fd, er := os.Open("test_004.txt")
	if er != nil {
		t.Fatal(er)
	}
	var expected = []string{
		"001",
		"002",
		"003",
		"004",
		"005",
		"006",
		"007",
		"008",
		"009",
		"010",
	}
	actual := Head(fd, 10)
	if len(expected) != len(actual) {
		t.Errorf("line count: expected %d but was %d\n", len(expected), len(actual))
	}
	for i, val := range expected {
		if val != actual[i] {
			t.Errorf("value mis-match: expected \"%s\" but found \"%s\"\n", val, actual[i])
		}
	}
}