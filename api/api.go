package api

import "fmt"

type Bin4 [4]bool

type Bin8 [8]bool

// NewBin4 converts a 4 bit binary number to bools
func NewBin4(a, b, c, d int) Bin4 {
	return Bin4{int2bool(a), int2bool(b), int2bool(c), int2bool(d)}
}

// NewBin4 converts a 4 bit binary number to bools
func NewBin8(a, b, c, d, a1, b1, c1, d1 int) Bin8 {
	return Bin8{
		int2bool(a),
		int2bool(b),
		int2bool(c),
		int2bool(d),
		int2bool(a1),
		int2bool(b1),
		int2bool(c1),
		int2bool(d1),
	}
}

func (b Bin4) Add(addend Bin4) Bin8 {
	return adder(addend[0], b[0], addend[1], b[1], addend[2], b[2], addend[3], b[3])
}

func (b *Bin4) String() string {
	return fmt.Sprintf("%v%v%v%v", Bool2int(b[0]), Bool2int(b[1]), Bool2int(b[2]), Bool2int(b[3]))
}

func (b *Bin8) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v", b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7])
}

func int2bool(a int) bool {
	if a == 0 {
		return false
	}
	return true
}

// Bool2int converts a bool to either a 1 or 0
func Bool2int(a bool) int {
	if a == false {
		return 0
	}
	return 1
}