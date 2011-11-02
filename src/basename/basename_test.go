package main

import "testing"

type basenameTest struct {
	in, out string
}

var cases = []basenameTest{
	//basenameTest{"", ""},
	basenameTest{".", "."},
	basenameTest{"..", ".."},
	basenameTest{"../a", "a"},
	basenameTest{"../a/b", "b"},
	basenameTest{"a/b//c///d////", "d"},
	basenameTest{"/", "/"},
	basenameTest{"//", "/"},
	basenameTest{"/a", "a"},
	basenameTest{"/a/b", "b"},
}

func TestBasename(t *testing.T) {
	for _, dt := range cases {
		v := basename(dt.in)
		if v != dt.out {
			t.Errorf("basename(%s) = %s - expected %s", dt.in, v, dt.out)
		}
	}
}
