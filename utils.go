package main

func nextFib() func() int {
	a := 1
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}
