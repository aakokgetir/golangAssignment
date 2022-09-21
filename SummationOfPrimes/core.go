package SummationOfPrimes

import (
	"fmt"
	"multiplesOfThreeOrFive/SomePrime"
	"time"
)

func SummationOfPrimes(n int) int {
	sum := 0
	dt := time.Now()
	for i, nthPrime := 1, SomePrime.NthPrime(1); nthPrime <= n; i, nthPrime = i+1, SomePrime.NthPrime(i+1) {
		sum += nthPrime
	}
	fmt.Println("Took", time.Now().Sub(dt))
	return sum
}
