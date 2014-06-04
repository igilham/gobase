package main

import (
	"os"
	"testing"
)

type wcTest struct {
	path                       string
	bytes, chars, words, lines uint64
}

var wcCases = []wcTest{
	wcTest{"../../resources/test_001.txt", 5, 5, 0, 0},
	wcTest{"../../resources/test_002.txt", 5, 5, 0, 0},
	wcTest{"../../resources/test_003.txt", 44, 44, 8, 6},
}

func TestWc(t *testing.T) {
	for _, tc := range wcCases {
		fd, er := os.Open(tc.path)
		if er != nil {
			t.Fatal(er)
		}
		defer fd.Close()
		actual, ew := Wc(fd)
		if ew != nil {
			t.Fatal(ew)
		}
		if tc.bytes != actual.Bytes {
			t.Errorf("Bytes - expected %d but was %d\n", tc.bytes, actual.Bytes)
		}
		if tc.chars != actual.Chars {
			t.Errorf("Chars - expected %d but was %d\n", tc.chars, actual.Chars)
		}
		if tc.words != actual.Words {
			t.Errorf("Words - expected %d but was %d\n", tc.words, actual.Words)
		}
		if tc.lines != actual.Lines {
			t.Errorf("Lines - expected %d but was %d\n", tc.lines, actual.Lines)
		}
	}
}
