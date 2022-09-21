package SomePrime

import "multiplesOfThreeOrFive/LargestPrimeFactor"

var nthPrimes map[int]int = map[int]int{1: 2} // nth->prime number itself

func NthPrime(n int) int {
	val, ok := nthPrimes[n]
	if ok {
		return val
	}
	for primeCandidate := NthPrime(n-1) + 1; true; primeCandidate++ {
		if LargestPrimeFactor.IsPrime(primeCandidate) {
			nthPrimes[n] = primeCandidate
			return primeCandidate
		}
	}
	return -1
}

func NthPrimeWithoutMemo(n int) int {
	nth := 1
	prime := 2
	nthPrimes[1] = 2
	for {
		if nth == n {
			nthPrimes[nth] = prime
			return prime
		}
		i := prime + 1
		for {
			if LargestPrimeFactor.IsPrime(i) {
				prime = i
				break
			}
			i += 1
		}
		nth += 1
	}
}
