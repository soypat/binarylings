//go:build !solution

package intro

// Problem: Primitive integer types in Go have a fixed width. For example, the
// int type is 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
// The same is true for uint.
//
// uint8, uint16, uint32 and uint64 are, intuitively, 8, 16, 32 and 64 bits wide.
// Same is true for int8, int16, int32 and int64, though these are signed, which
// is to say they can represent negative numbers at the cost of a smaller magnitudes.
//
// i.e: uint8 can represent 0..255, int8 can represent -128..127. Worth noting
// both can represent a total of 256 distinct values, which is 2^8.
//
// If 8 bit integers can't represent greater than 255, what happens when we add 1
// to 255? Different programming languages have different answers to this question,
// and some even let the user choose. This is a case of "integer overflow".
//
// The two most common ways of handling integer overflow are "wrap around" and
// "saturate".
//
// Wrap around means that the value "wraps around" to the smallest
// value of the type. For example, if we add 1 to 255, we get 0. This is the
// most common way of handling since it finds plenty of use in binary protocols.
//
// Saturate means that the value "saturates" at the limit of the type, whether
// it be the maximum or minimum value. For example, if we add 1 to 255, we get 255.
//
// Complete the following functions to add two integers and return the result
// according to the above rules.
// https://en.wikipedia.org/wiki/Integer_overflow

func addU8_wrap(a, b uint8) uint8 {
	return 0
}

func addU8_saturate(a, b uint8) uint8 {
	return 0
}

func addI8_wrap(a, b int8) int8 {
	return 0
}

func addI8_saturate(a, b int8) int8 {
	return 0
}
