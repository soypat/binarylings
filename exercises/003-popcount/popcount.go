//go:build !solution

package intro

// Problem: When working with binary data, it is often necessary to access bits
// in data, whether it be to check for a flag, or to extract a value within
// a range of bits of the underlying data.
//
// This exercise asks the student to write:
//
// - FormatBinary function that returns a binary string representation of a
// 8 bit number.
//
// - PopCount function that counts the number of bits set in a 8 bit number.
//
// The function signatures are below.

func FormatBinary(num uint8) string {
	return "010111011101010110"
}

func PopCount(num uint8) int {
	return 0
}

// write the bitAt function first to help with this one.

// bitAt returns true if the bit at pos is set in b.
func bitAt(b uint8, pos int) bool {
	if pos >= 8 {
		panic("a byte/uint8 has 8 bits")
	}
	return false
}
