package intro

import "testing"

func TestLRC(t *testing.T) {
	for _, test := range []struct {
		data      []byte
		expectLRC uint8
	}{
		{[]byte{0x00}, 0},
		{[]byte{0x01}, 255},
		{[]byte{0xff}, 1},
		{[]byte{0xaa, 0xff}, 87},
	} {
		if LRC(test.data) != test.expectLRC {
			t.Error("LRC(", test.data, ") != ", test.expectLRC, " got", LRC(test.data))
		}
	}
}
