//go:build solution

/* Hint 1: The modbus protocol gives us a C program that calculates the LRC:

static unsigned char LRC(auchMsg, usDataLen) // the function returns the LRC as a type unsigned char
unsigned char *auchMsg ;                     // message to calculate LRC upon
unsigned short usDataLen ;                   // quantity of bytes in message
{
	unsigned char uchLRC = 0 ;                   // LRC char initialized
	while (usDataLen--)                          // pass through message buffer
	uchLRC += *auchMsg++ ;                       // add buffer byte without carry
	return ((unsigned char)(-((char)uchLRC))) ;  // return twos complement
}
*/

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

func LRC(data []byte) (lrc uint8) {
	for i := range data {
		lrc += data[i] // Go's overflow behaviour same as C: wraps discarding carries.
	}
	lrc = 0xFF - lrc // Step 2: Subtract final field value from FF to get ones complement.
	return lrc + 1   // Step 3: Add 1 to the result to produce two's complement
}

// ======================
// Alternative solutions.
// ======================

func LRC_twoscomplementDirect(data []byte) (lrc uint8) {
	for i := range data {
		lrc += data[i] // Go's overflow behaviour same as C: wraps discarding carries.
	}
	char := -int8(lrc) // Produce two's complement directly.
	return uint8(char)
}
