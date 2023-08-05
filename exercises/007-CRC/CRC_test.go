package crc

import (
	"testing"
)

func TestCRCModbus(t *testing.T) {
	for _, test := range []struct {
		data      []byte
		expectCRC uint16
	}{
		{[]byte("hello"), 13558},
		{[]byte("hello!"), 24180},
		{[]byte(""), 0xffff},
		{[]byte("m"), 44414},
	} {
		if CRCModbus(test.data) != test.expectCRC {
			t.Error("CRCModbus(", test.data, ") != ", test.expectCRC, "got", CRCModbus(test.data))
		}
	}
}

func TestCRCTCPIP(t *testing.T) {
	for _, test := range []struct {
		data      []byte
		expectCRC uint16
	}{
		{[]byte("hello"), 48173},
		{[]byte("hello!"), 48140},
		{[]byte(""), 0xffff},
		{[]byte("m"), 37631},
		{[]byte("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"), 25565},
	} {
		if CRCTCPIP(test.data) != test.expectCRC {
			t.Error("CRCModbus(", test.data, ") != ", test.expectCRC, "got", CRCTCPIP(test.data))
		}
	}
}

func TestCRCOpenCyphal16(t *testing.T) {
	for _, test := range []struct {
		data      []byte
		expectCRC uint16
	}{
		{[]byte("hello"), 53870},
		{[]byte("hello!"), 45436},
		{[]byte(""), 0xffff},
		{[]byte("m"), 23803},
		{[]byte("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"), 24827},
	} {
		if CRCOpenCyphal16(test.data) != test.expectCRC {
			t.Error("CRCModbus(", test.data, ") != ", test.expectCRC, "got", CRCOpenCyphal16(test.data))
		}
	}
}

func TestCRCOpenCyphal32(t *testing.T) {
	for _, test := range []struct {
		data      []byte
		expectCRC uint32
	}{
		{[]byte("hello"), 4248433346},
		{[]byte("hello!"), 4136573492},
		{[]byte(""), 0},
		{[]byte("m"), 3986853568},
		{[]byte("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz"), 4064057660},
	} {
		if CRCOpenCyphal32(test.data) != test.expectCRC {
			t.Error("CRCModbus(", test.data, ") != ", test.expectCRC, "got", CRCOpenCyphal32(test.data))
		}
	}
}
