package gobase

// Print a sequence of numbers.
// step > 0; end > start
// step < 0; end < start
// format: formatting string
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