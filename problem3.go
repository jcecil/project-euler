package main

import "math"

// https://projecteuler.net/problem=2

func problem3() int {
	number := 600851475143
	start := number / 2
	for ; start > 0; start -= 1 {
		if number%start == 0 && isPrime(start) {
			return start
		}
	}
	return 0
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
