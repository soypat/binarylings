//go:build !solution

package crc

// Problem: CRC refers to checksums that are used to detect errors in data
// transmission, similarily to LRC. Different protocols often require different
// CRC algorithms.
// This exercise tasks the student with implementing several CRC algorithms
// commonly used in industrial protocols.
//
// - Modbus Serial Line 16bit CRC
// - OpenCyphal 16 and 32bit CRC. See https://opencyphal.org/specification/Cyphal_Specification.pdf
// - IP/TCP RFC 791 16bit CRC

// # Modbus Serial Line 16bit CRC
//
// The CRC is started by first preloading a 16–bit register to all 1’s.
// Then a process begins of applying successive 8–bit bytes of the
// message to the current contents of the register. Only the eight bits of data
// in each character are used for generating the CRC. Start
// and stop bits and the parity bit do not apply to the CRC
//
// During generation of the CRC, each 8–bit character is exclusive ORed with the
// register contents. Then the result is shifted in the direction of the least
// significant bit (LSB), with a zero filled into the most significant bit (MSB) position.
// The LSB is extracted and examined. If the LSB was a 1, the register is
// then exclusive ORed with a preset, fixed value. If the LSB was a 0, no exclusive OR takes place.
//
// This process is repeated until eight shifts have been performed. After the
// last (eighth) shift, the next 8–bit character is exclusive ORed with the
// register’s current value, and the process repeats for eight more shifts as described above.
// The final content of the register, after all the characters of the message have been applied, is the CRC value.

/*
A procedure for generating a Modbus CRC is:
1. Load a 16–bit register with FFFF hex (all 1’s). Call this the CRC register.

2. Exclusive OR the first 8–bit byte of the message with the low–order byte of the 16–bit CRC register, putting the result in the
CRC register.

3. Shift the CRC register one bit to the right (toward the LSB), zero–filling the MSB. Extract and examine the LSB.

4. (If the LSB was 0): Repeat Step 3 (another shift).
(If the LSB was 1): Exclusive OR the CRC register with the polynomial value 0xA001 (1010 0000 0000 0001).

5. Repeat Steps 3 and 4 until 8 shifts have been performed. When this is done, a complete 8–bit byte will have been
processed.

6. Repeat Steps 2 through 5 for the next 8–bit byte of the message. Continue doing this until all bytes have been processed.

7. The final content of the CRC register is the CRC value.

8. When the CRC is placed into the message, its upper and lower bytes must be swapped as described below.
*/
func CRCModbus(data []byte) (crc uint16) {
	return 0
}

func CRCTCPIP(data []byte) (crc uint16) {
	return 0
}

func CRCOpenCyphal16(data []byte) (crc16 uint16) {
	return 0
}

func CRCOpenCyphal32(data []byte) (crc32 uint32) {
	return 0
}
