package SpecialPythagoreanTriplet

func sq(n int) int {
	return n * n
}

func TripletFor(n int) (int, int, int) {
	for a := 1; a <= n; a++ {
		for b := 2; b <= n-a; b++ {
			c := n - a - b
			if a < b && b < c && sq(a)+sq(b) == sq(c) {
				return a, b, c
			}
		}
	}
	return 0, 0, 0
}
