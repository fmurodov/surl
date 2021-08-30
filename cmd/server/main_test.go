package main

import (
	"testing"
)

// TODO: Test the server.

// test getEnv function
func TestGetEnv(t *testing.T) {
	var env = "TEST"
	var expected = "defautval"
	var actual = getEnv(env, expected)
	if actual != expected {
		t.Errorf("getEnv(%s) = %s; want %s", env, actual, expected)
	}
}
