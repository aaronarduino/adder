package main

import (
	"fmt"

	"github.com/aaronarduino/adder/api"
)

func main() {
	a := api.Bin4(1, 0, 1, 0)
	b := api.Bin4(0, 1, 1, 1)
	fmt.Println(adder(a[3], b[3], a[2], b[2], a[1], b[1], a[0], b[0]))
}

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

func adder(a0, b0, a1, b1, a2, b2, a3, b3 bool) (int, int, int, int, int) {
	sum0, carry0 := add(a0, b0, false)
	sum1, carry1 := add(a1, b1, carry0)
	sum2, carry2 := add(a2, b2, carry1)
	sum3, cout := add(a3, b3, carry2)
	return api.Bool2int(sum0), api.Bool2int(sum1), api.Bool2int(sum2), api.Bool2int(sum3), api.Bool2int(cout)
}
