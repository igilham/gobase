package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetLines_1Line(t *testing.T) {
	fd, er := os.Open(
		filepath.Join("..", "..", "resources", "test_004.txt"))
	if er != nil {
		t.Fatal(er)
	}
	defer fd.Close()
	var expected = []string{
		"001",
	}
	actual, _ := GetLines(fd, 1)
	verifyGetLinesResults(t, expected, actual)
}

func TestGetLines_AllLinesInOneLineFile(t *testing.T) {
	fd, er := os.Open(
		filepath.Join("..", "..", "resources", "test_001.txt"))
	if er != nil {
		t.Fatal(er)
	}
	defer fd.Close()
	var expected = []string{
		"hello",
	}
	actual, _ := GetLines(fd, 0)
	verifyGetLinesResults(t, expected, actual)
}

func TestGetLines_10Lines(t *testing.T) {
	fd, er := os.Open(filepath.Join("..", "..", "resources", "test_004.txt"))
	if er != nil {
		t.Fatal(er)
	}
	defer fd.Close()
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
	actual, _ := GetLines(fd, 10)
	verifyGetLinesResults(t, expected, actual)
}

func TestGetLines_AllLines(t *testing.T) {
	fd, er := os.Open(
		filepath.Join("..", "..", "resources", "test_004.txt"))
	if er != nil {
		t.Fatal(er)
	}
	defer fd.Close()
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
	actual, _ := GetLines(fd, 0)
	verifyGetLinesResults(t, expected, actual)
}

func TestGetLines_19Lines(t *testing.T) {
	fd, er := os.Open(
		filepath.Join("..", "..", "resources", "test_004.txt"))
	if er != nil {
		t.Fatal(er)
	}
	defer fd.Close()
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
	actual, _ := GetLines(fd, 19)
	verifyGetLinesResults(t, expected, actual)
}

func verifyGetLinesResults(t *testing.T, expected, actual []string) {
	if len(expected) != len(actual) {
		t.Errorf("line count: expected %d but was %d\n", len(expected), len(actual))
	}
	for i, val := range expected {
		if val != actual[i] {
			t.Errorf("value mis-match: expected \"%s\" but found \"%s\"\n", val, actual[i])
		}
	}
}
