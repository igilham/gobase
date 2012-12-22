package gobase

import (
	"os"
	"testing"
)

func TestHead_1Line(t *testing.T) {
	fd, er := os.Open("test_004.txt")
	if er != nil {
		t.Fatal(er)
	}
	var expected = []string{
		"001",
	}
	actual := Head(fd, 1)
	verifyHeadResults(t, expected, actual)
}

func TestHead_10Lines(t *testing.T) {
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
	verifyHeadResults(t, expected, actual)
}

func TestHead_AllLines(t *testing.T) {
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
		"011",
		"012",
		"013",
		"014",
		"015",
		"016",
		"017",
		"018",
		"019",
		"020",
	}
	actual := Head(fd, 0)
	verifyHeadResults(t, expected, actual)
}

func TestHead_19Lines(t *testing.T) {
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
		"011",
		"012",
		"013",
		"014",
		"015",
		"016",
		"017",
		"018",
		"019",
	}
	actual := Head(fd, 19)
	verifyHeadResults(t, expected, actual)
}

func verifyHeadResults(t *testing.T, expected, actual []string) {
	if len(expected) != len(actual) {
		t.Errorf("line count: expected %d but was %d\n", len(expected), len(actual))
	}
	for i, val := range expected {
		if val != actual[i] {
			t.Errorf("value mis-match: expected \"%s\" but found \"%s\"\n", val, actual[i])
		}
	}
}
