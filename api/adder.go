package api

func xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}

func and(a, b bool) bool {
	return a && b
}

func or(a, b bool) bool {
	return a || b
}

// add outputs in format (sum, carry)
func add(a, b, c bool) (bool, bool) {
	tmpSum := xor(a, b)
	return xor(tmpSum, c), or(and(tmpSum, c), and(a, b))
}

func adder(a0, b0, a1, b1, a2, b2, a3, b3 bool) Bin8 {
	sum0, carry0 := add(a0, b0, false)
	sum1, carry1 := add(a1, b1, carry0)
	sum2, carry2 := add(a2, b2, carry1)
	sum3, cout := add(a3, b3, carry2)
	return Bin8{false, false, false, cout, sum3, sum2, sum1, sum0}
}
