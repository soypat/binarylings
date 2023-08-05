//go:build !solution

package intro

// Problem: CRCs and LRCs are used in many protocols to detect errors in data transmission.
// They are a simple checksum that is easy to compute and verify.
//
// We are asked to implement the LRC checksum as per the Modbus over Serial Line
// Protocol specification v1.02. From the Spec:

/*
# Appendix: Section 6.2.1 LRC Generation

The Longitudinal Redundancy Checking (LRC) field is one byte, containing an 8–bit
binary value. The LRC value is calculated by the transmitting device, which appends
the LRC to the message. The device that receives recalculates an LRC during receipt of the
message, and compares the calculated value to the actual value it received in the
LRC field. If the two values are not equal, an error results.

The LRC is calculated by adding together successive 8–bit bytes in the message,
discarding any carries, and then two’s complementing the result.
The LRC is an 8–bit field, therefore each new addition of a character that would
result in a value higher than 255 decimal simply ‘rolls over’ the field’s value
through zero. Because there is no ninth bit, the carry is discarded automatically.

A procedure for generating an LRC is:

1. Add all bytes in the message, excluding the starting ‘colon’ and ending CRLF.
Add them into an 8–bit field, so that carries will be discarded.

2. Subtract the final field value from FF hex (all 1’s), to produce the ones–complement.

3. Add 1 to produce the twos–complement.
*/
func LRC(data []byte) (lrc uint8) {
	return lrc
}
