package src

import "testing"

func TestShortEncode(t *testing.T) {
	if Encode(1) != "1" {
		t.Error("encode 1 error")
	}

	if Encode(408245137) != "hello" {
		t.Error("encode hello error")
	}
}

func TestDecode(t *testing.T) {
	val, err := Decode("hello")
	if val != 408245137 || err != nil {
		t.Error("decode hello error")
	}
}
