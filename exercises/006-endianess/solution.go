//go:build solution

// Hint 1: The binary package from the standard library has functions for
// converting between integers and byte slices.

/**/

/**/

// Hint 2: These functions can be written by calling the appropriate binary
// package functions, no code is required.

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

func getLittleEndian32(b []byte) uint32 {
	// binary.LittleEndian.Uint32(b)
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func getBigEndian32(b []byte) uint32 {
	// binary.BigEndian.Uint32(b)
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func putLittleEndian32(b []byte, v uint32) {
	// binary.LittleEndian.PutUint32(b, v)
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}

func putBigEndian32(b []byte, v uint32) {
	// binary.BigEndian.PutUint32(b, v)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}
