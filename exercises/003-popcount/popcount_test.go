package intro

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	for _, test := range [][2]uint8{
		{255, 8},
		{1, 1},
		{0, 0},
		{2, 1},
		{4, 1},
		{64, 1},
		{64 | 4 | 2, 3},
	} {
		if PopCount(test[0]) != int(test[1]) {
			t.Error("PopCount(", test[0], ") != ", test[1], "got", PopCount(test[0]))
		}
	}
}
