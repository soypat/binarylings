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

package crc

import "encoding/binary"

func CRCModbus(data []byte) (crc uint16) {
	const (
		initialValue = 0xFFFF
		polynomial   = 0xA001
	)
	crc = initialValue
	for i := 0; i < len(data); i++ {
		crc ^= uint16(data[i])
		for n := 0; n < 8; n++ {
			if crc&1 != 0 { // Examine LSB.
				crc >>= 1
				crc ^= polynomial
			} else {
				crc >>= 1
			}
		}
	}
	return crc
}

func CRCTCPIP(data []byte) (crc uint16) {
	var sum uint32
	N := len(data) / 2
	for i := 0; i < N; i++ {
		sum += uint32(binary.BigEndian.Uint16(data[i*2:]))
	}
	if len(data)%2 != 0 {
		// If odd number of bytes, pad with zero.
		sum += uint32(data[len(data)-1]) << 8
	}
	for sum > 0xffff {
		sum = sum&0xffff + sum>>16
	}
	return ^uint16(sum)
}

func CRCOpenCyphal16(data []byte) (crc16 uint16) {
	const (
		initialValue = 0xffff
		polynomial   = 0x1021
	)
	crc16 = initialValue
	for _, b := range data {
		crc16 ^= uint16(b) << 8
		for i := 0; i < 8; i++ {
			if crc16&0x8000 != 0 {
				crc16 = crc16<<1 ^ polynomial
			} else {
				crc16 <<= 1
			}
		}
	}
	return crc16
}

func CRCOpenCyphal32(data []byte) (crc32 uint32) {
	const (
		initialValue        = 0xffff_ffff
		reflectedPolynomial = 0x1edc_6f41
		// residue             = 0xb798_b438
	)
	crc32 = initialValue
	for _, b := range data {
		crc32 ^= uint32(b) << 8
		for i := 0; i < 8; i++ {
			if crc32&1 != 0 {
				crc32 = crc32>>1 ^ reflectedPolynomial
			} else {
				crc32 >>= 1
			}
		}
	}
	return ^crc32
}
