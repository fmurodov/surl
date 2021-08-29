package storage

import "testing"

// test RandString return string length is 10
func TestRandString(t *testing.T) {
	s := RandString()
	if len(s) != 10 {
		t.Error("RandString return string length is 10")
	}
}
