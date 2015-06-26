package main

// https://projecteuler.net/problem=2

func problem2() int {
	fib := nextFib()
	max := 4000000
	sum := 0

	for i := fib(); i < max; i = fib() {
		if i%2 == 0 {
			sum += i
		}
	}

	return sum
}
