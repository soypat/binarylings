package intro

import (
	"testing"
)

func TestGetLittleEndian32(t *testing.T) {
	buf := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	if getLittleEndian32(buf) != 0x04030201 {
		t.Fail()
	}
}

func TestGetBigEndian32(t *testing.T) {
	buf := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	if getBigEndian32(buf) != 0x01020304 {
		t.Fail()
	}
}

func TestPutLittleEndian32(t *testing.T) {
	buf := make([]byte, 5)
	putLittleEndian32(buf, 0x04030201)
	if buf[0] != 0x01 || buf[1] != 0x02 || buf[2] != 0x03 || buf[3] != 0x04 {
		t.Fail()
	}
}
func TestPutBigEndian32(t *testing.T) {
	buf := make([]byte, 5)
	putBigEndian32(buf, 0x01020304)
	if buf[0] != 0x01 || buf[1] != 0x02 || buf[2] != 0x03 || buf[3] != 0x04 {
		t.Fail()
	}
}
