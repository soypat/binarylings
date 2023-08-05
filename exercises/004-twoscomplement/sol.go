//go:build solution

// Hint 1: Make use of Go's built-in overflow behavior.

/**/

/**/

// Hint 2: Go's underflow behaviour wraps.

/**/

/**/

/**/

// Hint 3: To check if value overflows, try casting arguments to a larger type.
// i.e. if adding two uint32s, cast them to uint64 and add them as uint64s and check
// if the result is larger than what a uint32 can store. The standard library math package
// had MaxUint32 and MaxUint64 constants.

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

func addU8_wrap(a, b uint8) uint8 {
	return a + b
}

func addU8_saturate(a, b uint8) uint8 {
	if uint16(a)+uint16(b) > 0xff {
		return 255
	}
	return a + b
}

func addI8_wrap(a, b int8) int8 {
	return a + b
}

func addI8_saturate(a, b int8) int8 {
	res := int16(a) + int16(b)
	if res < -128 {
		return -128
	}
	if res > 127 {
		return 127
	}
	return a + b
}
