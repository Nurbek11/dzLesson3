package dz3Lesson

func Fibonacci() func() int {
	d := 0
	a := 0
	b := 1
	c := a + b
	return func() int {
		var ret int
		switch {
		case d == 0:
			d++
			ret = 0
		case d == 1:
			d++
			ret = 1
		default:
			ret = c
			a = b
			b = c
			c = a + b
		}
		return ret
	}
}
