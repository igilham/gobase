package gobase

import (
	"testing"
)

type basenameTest struct {
	in, out string
}

var	basenameCases = []basenameTest{
	//basenameTest{"", ""},
	basenameTest{".", "."},
	basenameTest{"..", ".."},
	basenameTest{".." + sep + "a", "a"},
	basenameTest{".." + sep + "a" + sep + "b", "b"},
	basenameTest{"a" + sep + "b" + sep + sep + "c" + sep + sep + sep +
		"d" + sep + sep + sep + sep, "d"},
	basenameTest{sep, sep},
	basenameTest{sep + sep, sep},
	basenameTest{sep + "a", "a"},
	basenameTest{ sep + "a" + sep + "b", "b"},
}

func TestBasename(t *testing.T) {
	for _, dt := range basenameCases {
		v := Basename(dt.in)
		if v != dt.out {
			t.Errorf("basename(\"%s\") = \"%s\" - expected \"%s\"", dt.in, v, dt.out)
		}
	}
}
