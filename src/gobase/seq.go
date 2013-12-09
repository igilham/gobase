package gobase

// Seq returns a slice containing a sequence of numbers from start to end in
// step increments. If the following conditions do not hold, Seq returns the
// empty slice.
// step != 0
// if step > 0; end > start
// if step < 0; end < start
func Seq(start, step, end float64) []float64 {
	// are we counting up or down?
	var out []float64
	var direction float64
	switch {
	case step > 0:
		direction = 1.0
	case step < 0:
		direction = -1.0
	default:
		return out
	}
	if (start * direction) > (end * direction) {
		return out
	}
	for num := start; (num * direction) <= (end * direction); num += step {
		out = append(out, num)
	}
	return out
}
