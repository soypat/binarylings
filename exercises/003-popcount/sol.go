//go:build solution

// Binary numbers: https://www.interviewcake.com/concept/java/binary-numbers
// Hint 1: https://www.interviewcake.com/concept/java/bit-shift
// Hint 2: https://www.interviewcake.com/concept/java/and
// Hint 3: https://en.wikichip.org/wiki/population_count

/**/

/**/

/**/

/**/

/**/

/**/

/**/

/**/

/**/

/*
///////////////////////////////////////////

///////////////////////////////////////////

WARNING: SOLUTION BELOW!

Continue at your own risk to your learning experience! Solutions ahead!

///////////////////////////////////////////

///////////////////////////////////////////
*/

package intro

import (
	"fmt"
	"math/bits"
)

func FormatBinary(num uint8) (s string) {
	for i := 0; i < 8; i++ {
		if bitAt(num, i) {
			s += "1"
		} else {
			s += "0"
		}
	}
	return s
}

func PopCount(num uint8) (count int) {
	for i := 0; i < 8; i++ {
		if bitAt(num, i) {
			count++
		}
	}
	return count
}

// bitAt returns true if the bit at pos is set in b.
func bitAt(b uint8, pos int) bool {
	return b&(1<<pos) != 0
}

// ======================
// Alternative solutions.
// ======================

func PopCount_KerniganAndRitchie(buf uint8) (count int) {
	// Iterate until buf is zero.
	for buf != 0 {
		count += int(buf & 1) // Zero out all bits except the LSB and add (this will always add 1 or 0).
		buf >>= 1             // Shift buf right 1 bit.
	}
	return count
}

func PopCount_Stdlib(buf uint8) (count int) {
	// The standard library implements popcount with lookup table for
	// max performance.
	return bits.OnesCount8(buf)
}

func FormatBinary_Stdlib(num uint8) string {
	return fmt.Sprintf("%b", num)
}

func FormatBinary_efficient(num uint8) string {
	// We know our string has exactly 8 bytes so we allocate memory only once.
	// Setting a value in a slice is MUCH faster than concatenating strings.
	var buf [8]byte
	for i := 0; i < 8; i++ {
		if bitAt(num, i) {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}
	return string(buf[:])
}
