package intro

import (
	"testing"
)

func TestAddU8_saturate(t *testing.T) {
	for _, test := range [][3]uint8{
		{255, 1, 255},
		{255, 0, 255},
		{254, 1, 255},
	} {
		if addU8_saturate(test[0], test[1]) != test[2] {
			t.Error("addU8_saturate(", test[0], ",", test[1], ") != ", test[2])
		}
	}
}
func TestAddU8_wrap(t *testing.T) {
	for _, test := range [][3]uint8{
		{255, 255, 254},
		{255, 1, 0},
		{255, 2, 1},
		{254, 1, 255},
	} {
		if addU8_wrap(test[0], test[1]) != test[2] {
			t.Error("addU8_wrap(", test[0], ",", test[1], ") != ", test[2])
		}
	}
}

func TestAddI8_saturate(t *testing.T) {
	for _, test := range [][3]int8{
		{127, 1, 127},
		{127, 0, 127},
		{126, 1, 127},
		{-127, -2, -128},
		{-128, -0, -128},
	} {
		if addI8_saturate(test[0], test[1]) != test[2] {
			t.Error("addI8_saturate(", test[0], ",", test[1], ") != ", test[2])
		}
	}
}
func TestAddI8_wrap(t *testing.T) {
	for _, test := range [][3]int8{
		{127, 1, -128},
		{127, 3, -126},
		{127, 0, 127},
		{126, 1, 127},
		{-127, -2, 127},
		{-128, -0, -128},
	} {
		if addI8_wrap(test[0], test[1]) != test[2] {
			t.Error("addI8_wrap(", test[0], ",", test[1], ") != ", test[2])
		}
	}
}
