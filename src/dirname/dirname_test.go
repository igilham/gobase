package main

import (
	"testing"
)

type dirnameTest struct {
	in, out string
}

var dirnameCases = []dirnameTest{
	dirnameTest{"", "."},
	dirnameTest{".", "."},
	dirnameTest{"..", "."},
	dirnameTest{".." + sep + "a", ".."},
	dirnameTest{".." + sep + "a" + sep + "b", ".." + sep + "a"},
	dirnameTest{"a" + sep + "b" + sep + sep + "c" + sep + sep + sep +
		"d" + sep + sep + sep + sep, "a" + sep + "b" + sep + sep + "c"},
	dirnameTest{sep, sep},
	dirnameTest{sep + "a", sep},
	dirnameTest{sep + "a" + sep + "b", sep + "a"},
}

func TestDirname(t *testing.T) {
	for _, dt := range dirnameCases {
		v := Dirname(dt.in)
		if v != dt.out {
			t.Errorf("dirname(\"%s\") = \"%s\" - expected \"%s\"", dt.in, v, dt.out)
		}
	}
}
