package gobase

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
	dirnameTest{"../a", ".."},
	dirnameTest{"../a/b", "../a"},
	dirnameTest{"a/b//c///d////", "a/b//c"},
	dirnameTest{"/", "/"},
	dirnameTest{"/a", "/"},
	dirnameTest{"/a/b", "/a"},
}

func TestDirname(t *testing.T) {
	for _, dt := range dirnameCases {
		v := Dirname(dt.in)
		if v != dt.out {
			t.Errorf("dirname(%s) = %s - expected %s", dt.in, v, dt.out)
		}
	}
}
