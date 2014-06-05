package main

import (
	"path/filepath"
	"testing"
)

type dirnameTest struct {
	in, out string
}

var dirnameCases = []dirnameTest{
	dirnameTest{"", "."},
	dirnameTest{".", "."},
	dirnameTest{"..", "."},
	dirnameTest{filepath.Join("..", "a"), ".."},
	dirnameTest{filepath.Join("..", "a", "b"), filepath.Join("..", "a")},
	dirnameTest{"a" + separator + "b" + separator + separator + "c" +
		separator + separator + separator +
		"d" + separator + separator + separator + separator,
		"a" + separator + "b" + separator + separator + "c"},
	dirnameTest{separator, separator},
	dirnameTest{separator + "a", separator},
	dirnameTest{separator + "a" + separator + "b", separator + "a"},
}

func TestDirname(t *testing.T) {
	for _, dt := range dirnameCases {
		v := Dirname(dt.in)
		if v != dt.out {
			t.Errorf("dirname(\"%s\") = \"%s\" - expected \"%s\"", dt.in, v, dt.out)
		}
	}
}
