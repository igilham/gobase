package gobase

import (
	"testing"
)

var format = "%g"

type seqTest struct {
	c                int
	start, step, end float64
	exp              []float64
}

// note that testing cannot be exhaustive for floating point numbers and
// rounding errors must be avoided.
var seqCases = []seqTest{
	seqTest{c: 0, start: 1.0, step: 1.0, end: 2.0,
		exp: []float64{1.0, 2.0}},
	seqTest{c: 1, start: 1.0, step: 2.0, end: 3.0,
		exp: []float64{1.0, 3.0}},
	seqTest{c: 2, start: 2.0, step: -1.0, end: 1.0,
		exp: []float64{2.0, 1.0}},
	seqTest{c: 3, start: 2.0, step: -1.0, end: -2.0,
		exp: []float64{2.0, 1.0, 0.0, -1.0, -2.0}},
	seqTest{c: 4, start: -1.0, step: 1.0, end: 1.0,
		exp: []float64{-1.0, 0.0, 1.0}},
	seqTest{c: 5, start: -1.0, step: -1.0, end: -3.0,
		exp: []float64{-1.0, -2.0, -3.0}},
	seqTest{c: 6, start: -0.1, step: 0.1, end: 0.1,
		exp: []float64{-0.1, 0.0, 0.1}},
	seqTest{c: 7, start: 0.1, step: 1.0, end: 1.1,
		exp: []float64{0.1, 1.1}},
}

func TestSeq(t *testing.T) {
	for _, s := range seqCases {
		res := Seq(s.start, s.step, s.end)
		count := len(res)
		expCount := len(s.exp)
		if len(res) != len(s.exp) {
			t.Errorf("c: %d, length was %d, expected %d", s.c, count, expCount)
		} else {
			for i := 0; i < count; i++ {
				if res[i] != s.exp[i] {
					t.Errorf("c: %d, res[%d] = %g, expected %g",
						s.c, i, res[i], s.exp[i])
				}
			}
		}
	}
}
